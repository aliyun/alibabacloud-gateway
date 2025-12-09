// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointName(v string) *CreateAccessPointConfiguration
	GetAccessPointName() *string
	SetNetworkOrigin(v string) *CreateAccessPointConfiguration
	GetNetworkOrigin() *string
	SetVpcConfiguration(v *AccessPointVpcConfiguration) *CreateAccessPointConfiguration
	GetVpcConfiguration() *AccessPointVpcConfiguration
}

type CreateAccessPointConfiguration struct {
	AccessPointName  *string                      `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	NetworkOrigin    *string                      `json:"NetworkOrigin,omitempty" xml:"NetworkOrigin,omitempty"`
	VpcConfiguration *AccessPointVpcConfiguration `json:"VpcConfiguration,omitempty" xml:"VpcConfiguration,omitempty"`
}

func (s CreateAccessPointConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAccessPointConfiguration) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *CreateAccessPointConfiguration) GetNetworkOrigin() *string {
	return s.NetworkOrigin
}

func (s *CreateAccessPointConfiguration) GetVpcConfiguration() *AccessPointVpcConfiguration {
	return s.VpcConfiguration
}

func (s *CreateAccessPointConfiguration) SetAccessPointName(v string) *CreateAccessPointConfiguration {
	s.AccessPointName = &v
	return s
}

func (s *CreateAccessPointConfiguration) SetNetworkOrigin(v string) *CreateAccessPointConfiguration {
	s.NetworkOrigin = &v
	return s
}

func (s *CreateAccessPointConfiguration) SetVpcConfiguration(v *AccessPointVpcConfiguration) *CreateAccessPointConfiguration {
	s.VpcConfiguration = v
	return s
}

func (s *CreateAccessPointConfiguration) Validate() error {
	if s.VpcConfiguration != nil {
		if err := s.VpcConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
