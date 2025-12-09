// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelVideo interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *LiveChannelVideo
	GetBandwidth() *int64
	SetCodec(v string) *LiveChannelVideo
	GetCodec() *string
	SetFrameRate(v int64) *LiveChannelVideo
	GetFrameRate() *int64
	SetHeight(v int64) *LiveChannelVideo
	GetHeight() *int64
	SetWidth(v int64) *LiveChannelVideo
	GetWidth() *int64
}

type LiveChannelVideo struct {
	Bandwidth *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Codec     *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	FrameRate *int64  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Height    *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width     *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s LiveChannelVideo) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelVideo) GoString() string {
	return s.String()
}

func (s *LiveChannelVideo) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *LiveChannelVideo) GetCodec() *string {
	return s.Codec
}

func (s *LiveChannelVideo) GetFrameRate() *int64 {
	return s.FrameRate
}

func (s *LiveChannelVideo) GetHeight() *int64 {
	return s.Height
}

func (s *LiveChannelVideo) GetWidth() *int64 {
	return s.Width
}

func (s *LiveChannelVideo) SetBandwidth(v int64) *LiveChannelVideo {
	s.Bandwidth = &v
	return s
}

func (s *LiveChannelVideo) SetCodec(v string) *LiveChannelVideo {
	s.Codec = &v
	return s
}

func (s *LiveChannelVideo) SetFrameRate(v int64) *LiveChannelVideo {
	s.FrameRate = &v
	return s
}

func (s *LiveChannelVideo) SetHeight(v int64) *LiveChannelVideo {
	s.Height = &v
	return s
}

func (s *LiveChannelVideo) SetWidth(v int64) *LiveChannelVideo {
	s.Width = &v
	return s
}

func (s *LiveChannelVideo) Validate() error {
	return dara.Validate(s)
}
