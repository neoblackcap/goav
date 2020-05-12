# Transcode audio

derived from transcode_aac.c from ffmpeg example
https://git.ffmpeg.org/gitweb/ffmpeg.git/blob_plain/HEAD:/doc/examples/transcode_aac.c

## How to build

Install ffmpeg developer libraries.
If Ubuntu 20.04,

```
sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
```

Otherwise, you have to build ffmpeg 4.2 or newer.  Build with configure `--enable-shared`.


Resolve dependency

```
go get .
```

Build

```
go build
```

I tested on Ubuntu 18.04 on x86_64, Ubuntu 18.04 on arm64(jetson nano) and Raspbian on Raspberry Pi 2.

## How to use

```
./transcode-audio <input_audio_file> <output_audio_file>
```

For example

```
./transcode-audio test001.opus out.mp4
```

Audio CODEC is hard-coded to AAC.
If you want to transcode other than AAC, Change this line of transcode-audio.go

```
	outputCodec := avcodec.AvcodecFindEncoder(avcodec.CodecId(avcodec.AV_CODEC_ID_AAC))
```

Enjoy.
