// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryRespAudioStream interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int64) *MetaQueryRespAudioStream
	GetBitrate() *int64
	SetChannels(v int64) *MetaQueryRespAudioStream
	GetChannels() *int64
	SetCodecName(v string) *MetaQueryRespAudioStream
	GetCodecName() *string
	SetDuration(v float64) *MetaQueryRespAudioStream
	GetDuration() *float64
	SetLanguage(v string) *MetaQueryRespAudioStream
	GetLanguage() *string
	SetSampleRate(v int64) *MetaQueryRespAudioStream
	GetSampleRate() *int64
	SetStartTime(v float64) *MetaQueryRespAudioStream
	GetStartTime() *float64
}

type MetaQueryRespAudioStream struct {
	// example:
	//
	// 320087
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 2
	Channels *int64 `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// 3.690667
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 48000
	SampleRate *int64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 0.0235
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s MetaQueryRespAudioStream) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryRespAudioStream) GoString() string {
	return s.String()
}

func (s *MetaQueryRespAudioStream) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *MetaQueryRespAudioStream) GetChannels() *int64 {
	return s.Channels
}

func (s *MetaQueryRespAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *MetaQueryRespAudioStream) GetDuration() *float64 {
	return s.Duration
}

func (s *MetaQueryRespAudioStream) GetLanguage() *string {
	return s.Language
}

func (s *MetaQueryRespAudioStream) GetSampleRate() *int64 {
	return s.SampleRate
}

func (s *MetaQueryRespAudioStream) GetStartTime() *float64 {
	return s.StartTime
}

func (s *MetaQueryRespAudioStream) SetBitrate(v int64) *MetaQueryRespAudioStream {
	s.Bitrate = &v
	return s
}

func (s *MetaQueryRespAudioStream) SetChannels(v int64) *MetaQueryRespAudioStream {
	s.Channels = &v
	return s
}

func (s *MetaQueryRespAudioStream) SetCodecName(v string) *MetaQueryRespAudioStream {
	s.CodecName = &v
	return s
}

func (s *MetaQueryRespAudioStream) SetDuration(v float64) *MetaQueryRespAudioStream {
	s.Duration = &v
	return s
}

func (s *MetaQueryRespAudioStream) SetLanguage(v string) *MetaQueryRespAudioStream {
	s.Language = &v
	return s
}

func (s *MetaQueryRespAudioStream) SetSampleRate(v int64) *MetaQueryRespAudioStream {
	s.SampleRate = &v
	return s
}

func (s *MetaQueryRespAudioStream) SetStartTime(v float64) *MetaQueryRespAudioStream {
	s.StartTime = &v
	return s
}

func (s *MetaQueryRespAudioStream) Validate() error {
	return dara.Validate(s)
}
