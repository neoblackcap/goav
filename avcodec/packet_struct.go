// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

func (p *Packet) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}

func (p *Packet) Duration() int64 {
	return int64(p.duration)
}

func (p *Packet) SetDuration(d int64) {
	p.duration = C.int64_t(d)
}

func (p *Packet) Flags() int {
	return int(p.flags)
}

func (p *Packet) SetFlags(f int) {
	p.flags = C.int(f)
}

func (p *Packet) SideDataElems() int {
	return int(p.side_data_elems)
}

func (p *Packet) Size() int {
	return int(p.size)
}

func (p *Packet) SetSize(s int) {
	p.size = C.int(s)
}

func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}

func (p *Packet) SetStreamIndex(i int) {
	p.stream_index = C.int(i)
}

func (p *Packet) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}

func (p *Packet) Dts() int64 {
	return int64(p.dts)
}

func (p *Packet) SetDts(v int64) {
	p.dts = C.int64_t(v)
}

func (p *Packet) Pos() int64 {
	return int64(p.pos)
}

func (p *Packet) SetPos(v int64) {
	p.pos = C.int64_t(v)
}

func (p *Packet) Pts() int64 {
	return int64(p.pts)
}

func (p *Packet) SetPts(v int64) {
	p.pts = C.int64_t(v)
}

func (p *Packet) Data() *uint8 {
	return (*uint8)(p.data)
}

func (p *Packet) SetData(d *uint8) {
	p.data = (*C.uint8_t)(d)
}
