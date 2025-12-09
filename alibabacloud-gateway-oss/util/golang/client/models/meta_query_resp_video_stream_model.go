// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryRespVideoStream interface {
	dara.Model
	String() string
	GoString() string
	SetBitDepth(v int64) *MetaQueryRespVideoStream
	GetBitDepth() *int64
	SetBitrate(v int64) *MetaQueryRespVideoStream
	GetBitrate() *int64
	SetCodecName(v string) *MetaQueryRespVideoStream
	GetCodecName() *string
	SetColorSpace(v string) *MetaQueryRespVideoStream
	GetColorSpace() *string
	SetDuration(v float64) *MetaQueryRespVideoStream
	GetDuration() *float64
	SetFrameCount(v int64) *MetaQueryRespVideoStream
	GetFrameCount() *int64
	SetFrameRate(v string) *MetaQueryRespVideoStream
	GetFrameRate() *string
	SetHeight(v int64) *MetaQueryRespVideoStream
	GetHeight() *int64
	SetLanguage(v string) *MetaQueryRespVideoStream
	GetLanguage() *string
	SetPixelFormat(v string) *MetaQueryRespVideoStream
	GetPixelFormat() *string
	SetStartTime(v float64) *MetaQueryRespVideoStream
	GetStartTime() *float64
	SetWidth(v int64) *MetaQueryRespVideoStream
	GetWidth() *int64
}

type MetaQueryRespVideoStream struct {
	// example:
	//
	// 8
	BitDepth *int64 `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	// example:
	//
	// 5407765
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// bt709
	ColorSpace *string `json:"ColorSpace,omitempty" xml:"ColorSpace,omitempty"`
	// example:
	//
	// 22.88
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 572
	FrameCount *int64 `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	// example:
	//
	// 25/1
	FrameRate *string `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// 720
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// yuv420p
	PixelFormat *string `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	// example:
	//
	// 0.000000
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1280
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s MetaQueryRespVideoStream) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryRespVideoStream) GoString() string {
	return s.String()
}

func (s *MetaQueryRespVideoStream) GetBitDepth() *int64 {
	return s.BitDepth
}

func (s *MetaQueryRespVideoStream) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *MetaQueryRespVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *MetaQueryRespVideoStream) GetColorSpace() *string {
	return s.ColorSpace
}

func (s *MetaQueryRespVideoStream) GetDuration() *float64 {
	return s.Duration
}

func (s *MetaQueryRespVideoStream) GetFrameCount() *int64 {
	return s.FrameCount
}

func (s *MetaQueryRespVideoStream) GetFrameRate() *string {
	return s.FrameRate
}

func (s *MetaQueryRespVideoStream) GetHeight() *int64 {
	return s.Height
}

func (s *MetaQueryRespVideoStream) GetLanguage() *string {
	return s.Language
}

func (s *MetaQueryRespVideoStream) GetPixelFormat() *string {
	return s.PixelFormat
}

func (s *MetaQueryRespVideoStream) GetStartTime() *float64 {
	return s.StartTime
}

func (s *MetaQueryRespVideoStream) GetWidth() *int64 {
	return s.Width
}

func (s *MetaQueryRespVideoStream) SetBitDepth(v int64) *MetaQueryRespVideoStream {
	s.BitDepth = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetBitrate(v int64) *MetaQueryRespVideoStream {
	s.Bitrate = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetCodecName(v string) *MetaQueryRespVideoStream {
	s.CodecName = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetColorSpace(v string) *MetaQueryRespVideoStream {
	s.ColorSpace = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetDuration(v float64) *MetaQueryRespVideoStream {
	s.Duration = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetFrameCount(v int64) *MetaQueryRespVideoStream {
	s.FrameCount = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetFrameRate(v string) *MetaQueryRespVideoStream {
	s.FrameRate = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetHeight(v int64) *MetaQueryRespVideoStream {
	s.Height = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetLanguage(v string) *MetaQueryRespVideoStream {
	s.Language = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetPixelFormat(v string) *MetaQueryRespVideoStream {
	s.PixelFormat = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetStartTime(v float64) *MetaQueryRespVideoStream {
	s.StartTime = &v
	return s
}

func (s *MetaQueryRespVideoStream) SetWidth(v int64) *MetaQueryRespVideoStream {
	s.Width = &v
	return s
}

func (s *MetaQueryRespVideoStream) Validate() error {
	return dara.Validate(s)
}
