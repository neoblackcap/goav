package main

import (
	"log"

	"github.com/neoblackcap/goav/avcodec"
	"github.com/neoblackcap/goav/avdevice"
	"github.com/neoblackcap/goav/avfilter"
	"github.com/neoblackcap/goav/avformat"
	"github.com/neoblackcap/goav/avutil"
	"github.com/neoblackcap/goav/swresample"
	"github.com/neoblackcap/goav/swscale"
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
