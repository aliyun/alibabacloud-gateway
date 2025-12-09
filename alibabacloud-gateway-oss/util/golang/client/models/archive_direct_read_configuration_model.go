// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveDirectReadConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ArchiveDirectReadConfiguration
	GetEnabled() *bool
}

type ArchiveDirectReadConfiguration struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ArchiveDirectReadConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ArchiveDirectReadConfiguration) GoString() string {
	return s.String()
}

func (s *ArchiveDirectReadConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ArchiveDirectReadConfiguration) SetEnabled(v bool) *ArchiveDirectReadConfiguration {
	s.Enabled = &v
	return s
}

func (s *ArchiveDirectReadConfiguration) Validate() error {
	return dara.Validate(s)
}
