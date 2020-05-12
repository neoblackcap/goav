# Timelapse

Make time-lapse movie. 10 minutes movie is squeezed to 30 seconds.

[![Watch the sample video](https://img.youtube.com/vi/Et1FjP9R_Uc/hqdefault.jpg)](https://youtu.be/Et1FjP9R_Uc)

Click this image to watch the sample video.

(c) copyright 2008, Blender Foundation / www.bigbuckbunny.org

## How to build

Install ffmpeg developer libraries.
If Ubuntu 20.04

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

## How to use

```
./timelapse <input_video_file> <output_video_file>
```

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
