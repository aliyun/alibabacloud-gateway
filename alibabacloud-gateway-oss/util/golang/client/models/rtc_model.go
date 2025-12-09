// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRTC interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *RTC
	GetStatus() *string
}

type RTC struct {
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RTC) String() string {
	return dara.Prettify(s)
}

func (s RTC) GoString() string {
	return s.String()
}

func (s *RTC) GetStatus() *string {
	return s.Status
}

func (s *RTC) SetStatus(v string) *RTC {
	s.Status = &v
	return s
}

func (s *RTC) Validate() error {
	return dara.Validate(s)
}
