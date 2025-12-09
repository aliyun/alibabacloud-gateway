// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelAudio interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *LiveChannelAudio
	GetBandwidth() *int64
	SetCodec(v string) *LiveChannelAudio
	GetCodec() *string
	SetSampleRate(v int64) *LiveChannelAudio
	GetSampleRate() *int64
}

type LiveChannelAudio struct {
	Bandwidth  *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	SampleRate *int64  `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s LiveChannelAudio) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelAudio) GoString() string {
	return s.String()
}

func (s *LiveChannelAudio) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *LiveChannelAudio) GetCodec() *string {
	return s.Codec
}

func (s *LiveChannelAudio) GetSampleRate() *int64 {
	return s.SampleRate
}

func (s *LiveChannelAudio) SetBandwidth(v int64) *LiveChannelAudio {
	s.Bandwidth = &v
	return s
}

func (s *LiveChannelAudio) SetCodec(v string) *LiveChannelAudio {
	s.Codec = &v
	return s
}

func (s *LiveChannelAudio) SetSampleRate(v int64) *LiveChannelAudio {
	s.SampleRate = &v
	return s
}

func (s *LiveChannelAudio) Validate() error {
	return dara.Validate(s)
}
