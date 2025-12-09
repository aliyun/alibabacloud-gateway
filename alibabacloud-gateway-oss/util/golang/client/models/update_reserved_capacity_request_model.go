// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReservedCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReservedCapacityUpdateConfiguration(v *ReservedCapacityUpdateConfiguration) *UpdateReservedCapacityRequest
	GetReservedCapacityUpdateConfiguration() *ReservedCapacityUpdateConfiguration
	SetXOssReservedCapacityId(v string) *UpdateReservedCapacityRequest
	GetXOssReservedCapacityId() *string
}

type UpdateReservedCapacityRequest struct {
	ReservedCapacityUpdateConfiguration *ReservedCapacityUpdateConfiguration `json:"ReservedCapacityUpdateConfiguration,omitempty" xml:"ReservedCapacityUpdateConfiguration,omitempty"`
	// This parameter is required.
	XOssReservedCapacityId *string `json:"x-oss-reserved-capacity-id,omitempty" xml:"x-oss-reserved-capacity-id,omitempty"`
}

func (s UpdateReservedCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReservedCapacityRequest) GoString() string {
	return s.String()
}

func (s *UpdateReservedCapacityRequest) GetReservedCapacityUpdateConfiguration() *ReservedCapacityUpdateConfiguration {
	return s.ReservedCapacityUpdateConfiguration
}

func (s *UpdateReservedCapacityRequest) GetXOssReservedCapacityId() *string {
	return s.XOssReservedCapacityId
}

func (s *UpdateReservedCapacityRequest) SetReservedCapacityUpdateConfiguration(v *ReservedCapacityUpdateConfiguration) *UpdateReservedCapacityRequest {
	s.ReservedCapacityUpdateConfiguration = v
	return s
}

func (s *UpdateReservedCapacityRequest) SetXOssReservedCapacityId(v string) *UpdateReservedCapacityRequest {
	s.XOssReservedCapacityId = &v
	return s
}

func (s *UpdateReservedCapacityRequest) Validate() error {
	if s.ReservedCapacityUpdateConfiguration != nil {
		if err := s.ReservedCapacityUpdateConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
