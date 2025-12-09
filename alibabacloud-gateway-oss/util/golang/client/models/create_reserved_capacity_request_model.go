// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReservedCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReservedCapacityCreateConfiguration(v *ReservedCapacityCreateConfiguration) *CreateReservedCapacityRequest
	GetReservedCapacityCreateConfiguration() *ReservedCapacityCreateConfiguration
}

type CreateReservedCapacityRequest struct {
	ReservedCapacityCreateConfiguration *ReservedCapacityCreateConfiguration `json:"ReservedCapacityCreateConfiguration,omitempty" xml:"ReservedCapacityCreateConfiguration,omitempty"`
}

func (s CreateReservedCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReservedCapacityRequest) GoString() string {
	return s.String()
}

func (s *CreateReservedCapacityRequest) GetReservedCapacityCreateConfiguration() *ReservedCapacityCreateConfiguration {
	return s.ReservedCapacityCreateConfiguration
}

func (s *CreateReservedCapacityRequest) SetReservedCapacityCreateConfiguration(v *ReservedCapacityCreateConfiguration) *CreateReservedCapacityRequest {
	s.ReservedCapacityCreateConfiguration = v
	return s
}

func (s *CreateReservedCapacityRequest) Validate() error {
	if s.ReservedCapacityCreateConfiguration != nil {
		if err := s.ReservedCapacityCreateConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
