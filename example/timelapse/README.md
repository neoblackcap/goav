# Timelapse

Make time-lapse movie. 10 minutes movie is squeezed to 30 seconds.

[![Watch the sample video](https://img.youtube.com/vi/Et1FjP9R_Uc/hqdefault.jpg)](https://youtu.be/Et1FjP9R_Uc)


(c) copyright 2008, Blender Foundation / www.bigbuckbunny.org

## How to build

Install ffmpeg developer libraries.
If Ubuntu or Debian,

```
sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
```

Resolve dependency

```
go get .
```

Build

```
go build
```

## How to use

```
./timelapse <input_video_file> <output_video_file>
```

Supported Video codec is H.264 only.

For example

```
./timelapse bbb480.mp4 out.mp4
```

Speed factor is hard-coded to 20x.
You can change it.

```
const SPEED_FACTOR = 20
```

Enjoy.
