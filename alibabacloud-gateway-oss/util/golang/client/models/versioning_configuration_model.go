// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVersioningConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *VersioningConfiguration
	GetStatus() *string
}

type VersioningConfiguration struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VersioningConfiguration) String() string {
	return dara.Prettify(s)
}

func (s VersioningConfiguration) GoString() string {
	return s.String()
}

func (s *VersioningConfiguration) GetStatus() *string {
	return s.Status
}

func (s *VersioningConfiguration) SetStatus(v string) *VersioningConfiguration {
	s.Status = &v
	return s
}

func (s *VersioningConfiguration) Validate() error {
	return dara.Validate(s)
}
