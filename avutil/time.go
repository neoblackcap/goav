package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/time.h>
import "C"

func AvGetTimeRelative() int64 {
	return int64(C.av_gettime_relative())
}
