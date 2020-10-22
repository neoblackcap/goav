package main

import (
	"log"

	"github.com/sxhxliang/goav/avcodec"
	"github.com/sxhxliang/goav/avdevice"
	"github.com/sxhxliang/goav/avfilter"
	"github.com/sxhxliang/goav/avformat"
	"github.com/sxhxliang/goav/avutil"
	"github.com/sxhxliang/goav/swresample"
	"github.com/sxhxliang/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
