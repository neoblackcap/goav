package main

// based on remuxing.c
// https://git.ffmpeg.org/gitweb/ffmpeg.git/blob_plain/HEAD:/doc/examples/remuxing.c

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/tetsu-koba/goav/avcodec"
	"github.com/tetsu-koba/goav/avformat"
	"github.com/tetsu-koba/goav/avutil"
)

const SPEED_FACTOR = 20

func Timelapse(inputFile, outputFile string) int {
	ifmtCtx := avformat.AvformatAllocContext()
	if avformat.AvformatOpenInput(&ifmtCtx, inputFile, nil, nil) != 0 {
		fmt.Printf("Unable to open file %s\n", inputFile)
		return 1
	}
	defer avformat.AvformatCloseInput(ifmtCtx)

	if ifmtCtx.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		return 1
	}

	ifmtCtx.AvDumpFormat(0, inputFile, 0)

	ofmtCtx := avformat.AvformatAllocContext()
	if avformat.AvformatAllocOutputContext2(&ofmtCtx, nil, "", outputFile) != 0 {
		fmt.Printf("Unable to alloc output context for %s\n", outputFile)
		return 1
	}
	defer ofmtCtx.AvformatFreeContext()

	videoStreamIndex := -1
	for i, inStream := range ifmtCtx.Streams() {
		inCodecPar := inStream.CodecParameters()
		if inCodecPar.CodecType() == avcodec.AVMEDIA_TYPE_VIDEO {
			videoStreamIndex = i
			break
		}
	}
	if videoStreamIndex == -1 {
		fmt.Printf("No video streamss\n")
		return 1
	}
	inStream := ifmtCtx.Streams()[videoStreamIndex]
	inCodecPar := inStream.CodecParameters()
	if inCodecPar.CodecId() != avcodec.CodecId(avcodec.AV_CODEC_ID_H264) {
		fmt.Printf("Sorry, Only H.264 is supported\n")
		return 1
	}
	outStream := ofmtCtx.AvformatNewStream(nil)
	if outStream == nil {
		fmt.Println("Failed allocating output stream")
		return 1
	}
	ret := avcodec.AvcodecParametersCopy(outStream.CodecParameters(), inCodecPar)
	if ret < 0 {
		fmt.Printf("Failed copy codec parameters: ret=%d\n", ret)
		return 1
	}
	t := inStream.TimeBase()
	outStream.SetTimeBase(avutil.NewRational(t.Num(), t.Den()*SPEED_FACTOR))
	outCodecPar := outStream.CodecParameters()
	outCodecPar.SetCodecTag(0)

	ofmtCtx.AvDumpFormat(0, outputFile, 1)

	pb := (*avformat.AvIOContext)(nil)
	ret = avformat.AvIOOpen(&pb, outputFile, avformat.AVIO_FLAG_WRITE)
	if ret < 0 {
		fmt.Printf("Could not open output file '%s'\n", outputFile)
		return 1
	}
	ofmtCtx.SetPb(pb)
	defer func() {
		pb := ofmtCtx.Pb()
		avformat.AvIOClosep(&pb)
	}()

	ret = ofmtCtx.AvformatWriteHeader(nil)
	if ret < 0 {
		fmt.Printf("Error occurred when opening output file\n")
		return 1
	}
	defer ofmtCtx.AvWriteTrailer()

	return loop(ofmtCtx, ifmtCtx, videoStreamIndex)
}

func isPframe(packet *avcodec.Packet) bool {
	d := avutil.PointerToUint8Slice(unsafe.Pointer(packet.Data()), packet.Size())
	return len(d) < 4 || (d[4]&0x1f) == 1
}

func loop(ofmtCtx *avformat.Context, ifmtCtx *avformat.Context, videoStreamIndex int) int {
	packet := avcodec.AvPacketAlloc()
	defer avcodec.AvPacketFree(packet)
	for ifmtCtx.AvReadFrame(packet) >= 0 {
		if packet.StreamIndex() != videoStreamIndex ||
			isPframe(packet) {
			packet.AvPacketUnref()
			continue
		}
		packet.SetStreamIndex(0)
		ret := ofmtCtx.AvInterleavedWriteFrame(packet)
		packet.AvPacketUnref()
		if ret < 0 {
			fmt.Printf("Error muxing packet\n")
			return 1
		}
	}
	return 0
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s infile outfile\n", os.Args[0])
		os.Exit(1)
	}
	os.Exit(Timelapse(os.Args[1], os.Args[2]))
}
