// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetID(v string) *RtcConfiguration
	GetID() *string
	SetRTC(v *RTC) *RtcConfiguration
	GetRTC() *RTC
}

type RtcConfiguration struct {
	// example:
	//
	// test_replication_rule_1
	ID  *string `json:"ID,omitempty" xml:"ID,omitempty"`
	RTC *RTC    `json:"RTC,omitempty" xml:"RTC,omitempty"`
}

func (s RtcConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RtcConfiguration) GoString() string {
	return s.String()
}

func (s *RtcConfiguration) GetID() *string {
	return s.ID
}

func (s *RtcConfiguration) GetRTC() *RTC {
	return s.RTC
}

func (s *RtcConfiguration) SetID(v string) *RtcConfiguration {
	s.ID = &v
	return s
}

func (s *RtcConfiguration) SetRTC(v *RTC) *RtcConfiguration {
	s.RTC = v
	return s
}

func (s *RtcConfiguration) Validate() error {
	if s.RTC != nil {
		if err := s.RTC.Validate(); err != nil {
			return err
		}
	}
	return nil
}
