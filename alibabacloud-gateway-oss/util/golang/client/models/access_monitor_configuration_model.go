// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessMonitorConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCopy(v bool) *AccessMonitorConfiguration
	GetAllowCopy() *bool
	SetStatus(v string) *AccessMonitorConfiguration
	GetStatus() *string
}

type AccessMonitorConfiguration struct {
	// example:
	//
	// true
	AllowCopy *bool   `json:"AllowCopy,omitempty" xml:"AllowCopy,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AccessMonitorConfiguration) String() string {
	return dara.Prettify(s)
}

func (s AccessMonitorConfiguration) GoString() string {
	return s.String()
}

func (s *AccessMonitorConfiguration) GetAllowCopy() *bool {
	return s.AllowCopy
}

func (s *AccessMonitorConfiguration) GetStatus() *string {
	return s.Status
}

func (s *AccessMonitorConfiguration) SetAllowCopy(v bool) *AccessMonitorConfiguration {
	s.AllowCopy = &v
	return s
}

func (s *AccessMonitorConfiguration) SetStatus(v string) *AccessMonitorConfiguration {
	s.Status = &v
	return s
}

func (s *AccessMonitorConfiguration) Validate() error {
	return dara.Validate(s)
}
