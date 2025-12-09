// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAccessPointConfiguration(v *CreateAccessPointConfiguration) *CreateAccessPointRequest
	GetCreateAccessPointConfiguration() *CreateAccessPointConfiguration
}

type CreateAccessPointRequest struct {
	// The container that stores the information about the access point.
	CreateAccessPointConfiguration *CreateAccessPointConfiguration `json:"CreateAccessPointConfiguration,omitempty" xml:"CreateAccessPointConfiguration,omitempty"`
}

func (s CreateAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessPointRequest) GetCreateAccessPointConfiguration() *CreateAccessPointConfiguration {
	return s.CreateAccessPointConfiguration
}

func (s *CreateAccessPointRequest) SetCreateAccessPointConfiguration(v *CreateAccessPointConfiguration) *CreateAccessPointRequest {
	s.CreateAccessPointConfiguration = v
	return s
}

func (s *CreateAccessPointRequest) Validate() error {
	if s.CreateAccessPointConfiguration != nil {
		if err := s.CreateAccessPointConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
