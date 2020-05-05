package avcodec

import "C"

func (cp *CodecParameters) CodecType() MediaType {
	return MediaType(cp.codec_type)
}

func (cp *CodecParameters) CodecId() CodecId {
	return CodecId(cp.codec_id)
}

func (cp *CodecParameters) SetCodecTag(tag uint) {
	cp.codec_tag = C.uint(tag)
}
