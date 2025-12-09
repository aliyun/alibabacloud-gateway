// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryRespSubtitle interface {
	dara.Model
	String() string
	GoString() string
	SetCodecName(v string) *MetaQueryRespSubtitle
	GetCodecName() *string
	SetDuration(v float64) *MetaQueryRespSubtitle
	GetDuration() *float64
	SetLanguage(v string) *MetaQueryRespSubtitle
	GetLanguage() *string
	SetStartTime(v float64) *MetaQueryRespSubtitle
	GetStartTime() *float64
}

type MetaQueryRespSubtitle struct {
	// example:
	//
	// mov_text
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// 71.378
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 0.000000
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s MetaQueryRespSubtitle) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryRespSubtitle) GoString() string {
	return s.String()
}

func (s *MetaQueryRespSubtitle) GetCodecName() *string {
	return s.CodecName
}

func (s *MetaQueryRespSubtitle) GetDuration() *float64 {
	return s.Duration
}

func (s *MetaQueryRespSubtitle) GetLanguage() *string {
	return s.Language
}

func (s *MetaQueryRespSubtitle) GetStartTime() *float64 {
	return s.StartTime
}

func (s *MetaQueryRespSubtitle) SetCodecName(v string) *MetaQueryRespSubtitle {
	s.CodecName = &v
	return s
}

func (s *MetaQueryRespSubtitle) SetDuration(v float64) *MetaQueryRespSubtitle {
	s.Duration = &v
	return s
}

func (s *MetaQueryRespSubtitle) SetLanguage(v string) *MetaQueryRespSubtitle {
	s.Language = &v
	return s
}

func (s *MetaQueryRespSubtitle) SetStartTime(v float64) *MetaQueryRespSubtitle {
	s.StartTime = &v
	return s
}

func (s *MetaQueryRespSubtitle) Validate() error {
	return dara.Validate(s)
}
