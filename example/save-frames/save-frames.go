package main

// Code based on a tutorial at http://dranger.com/ffmpeg/tutorial01.html

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/tetsu-koba/goav/avcodec"
	"github.com/tetsu-koba/goav/avformat"
	"github.com/tetsu-koba/goav/avutil"
	"github.com/tetsu-koba/goav/swscale"
)

// SaveFrame writes a single frame to disk as a PPM file
func SaveFrame(frame *avutil.Frame, width, height, frameNumber int) {
	// Open file
	fileName := fmt.Sprintf("frame%d.ppm", frameNumber)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error Reading")
	}
	defer file.Close()

	// Write header
	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	file.Write([]byte(header))

	// Write pixel data
	for y := 0; y < height; y++ {
		data0 := frame.Data()[0]
		buf := make([]byte, width*3)
		startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(frame.Linesize()[0])
		for i := 0; i < width*3; i++ {
			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
			buf[i] = element
		}
		file.Write(buf)
	}
}

func SaveFrames(inputFileName string) int {
	// Open video file
	pFormatContext := avformat.AvformatAllocContext()
	if avformat.AvformatOpenInput(&pFormatContext, inputFileName, nil, nil) != 0 {
		fmt.Printf("Unable to open file %s\n", inputFileName)
		return 1
	}
	defer avformat.AvformatCloseInput(pFormatContext)

	// Retrieve stream information
	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		return 1
	}

	// Dump information about file onto standard error
	pFormatContext.AvDumpFormat(0, inputFileName, 0)

	// Find the first video stream
	videoStreamIndex := -1
	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		if pFormatContext.Streams()[i].CodecParameters().CodecType() == avcodec.AVMEDIA_TYPE_VIDEO {
			videoStreamIndex = i
			break
		}
	}
	if videoStreamIndex == -1 {
		fmt.Println("no video stream")
		return 1
	}

	// Get a pointer to the codec context for the video stream
	pCodecCtx := pFormatContext.Streams()[videoStreamIndex].Codec()
	// Find the decoder for the video stream
	pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(pCodecCtx.CodecId()))
	if pCodec == nil {
		fmt.Println("Unsupported codec!")
		return 1
	}

	// Open codec
	if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
		fmt.Println("Could not open codec")
		return 1
	}
	defer pCodecCtx.AvcodecClose()

	// Allocate video frame
	pFrame := avutil.AvFrameAlloc()
	defer avutil.AvFrameFree(pFrame)

	// Allocate an AVFrame structure
	pFrameRGB := avutil.AvFrameAlloc()
	if pFrameRGB == nil {
		fmt.Println("Unable to allocate RGB Frame")
		return 1
	}
	defer avutil.AvFrameFree(pFrameRGB)
	pFrameRGB.SetFormat(avutil.AV_PIX_FMT_RGB24)
	pFrameRGB.SetWidth(pCodecCtx.Width())
	pFrameRGB.SetHeight(pCodecCtx.Height())
	avutil.AvFrameGetBuffer(pFrameRGB, 0)

	// initialize SWS context for software scaling
	swsCtx := swscale.SwsGetcontext(
		pCodecCtx.Width(),
		pCodecCtx.Height(),
		pCodecCtx.PixFmt(),
		pCodecCtx.Width(),
		pCodecCtx.Height(),
		avutil.AV_PIX_FMT_RGB24,
		swscale.SWS_BILINEAR,
		nil,
		nil,
		nil,
	)

	// Read frames and save first five frames to disk
	frameNumber := 1
	packet := avcodec.AvPacketAlloc()
	defer avcodec.AvPacketFree(packet)
	for pFormatContext.AvReadFrame(packet) >= 0 {
		// Is this a packet from the video stream?
		if packet.StreamIndex() != videoStreamIndex {
			packet.AvPacketUnref()
			continue
		}
		// Decode video frame
		response := avcodec.AvcodecSendPacket(pCodecCtx, packet)
		if response < 0 {
			fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.AvStrerr(response))
		}
		for response >= 0 {
			response = avcodec.AvcodecReceiveFrame(pCodecCtx, pFrame)
			if response == avutil.AVERROR_EOF {
				break
			} else if response == avutil.AVERROR_EAGAIN {
				continue
			} else if response < 0 {
				fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.AvStrerr(response))
				return 1
			}

			if frameNumber <= 5 {
				// Convert the image from its native format to RGB
				swscale.SwsScale(swsCtx,
					pFrame.Data(),
					pFrame.Linesize(),
					0,
					pCodecCtx.Height(),
					pFrameRGB.Data(),
					pFrameRGB.Linesize())

				// Save the frame to disk
				fmt.Printf("Writing frame %d\n", frameNumber)
				SaveFrame(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), frameNumber)
			} else {
				return 0
			}
			frameNumber++
		}
		// Free the packet that was allocated by av_read_frame
		packet.AvPacketUnref()
	}
	return 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie file")
		os.Exit(1)
	}
	os.Exit(SaveFrames(os.Args[1]))
}
