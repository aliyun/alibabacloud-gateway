// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessPoint interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointName(v string) *AccessPoint
	GetAccessPointName() *string
	SetAlias(v string) *AccessPoint
	GetAlias() *string
	SetBucket(v string) *AccessPoint
	GetBucket() *string
	SetNetworkOrigin(v string) *AccessPoint
	GetNetworkOrigin() *string
	SetStatus(v string) *AccessPoint
	GetStatus() *string
	SetVpcConfiguration(v *AccessPointVpcConfiguration) *AccessPoint
	GetVpcConfiguration() *AccessPointVpcConfiguration
}

type AccessPoint struct {
	AccessPointName  *string                      `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	Alias            *string                      `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Bucket           *string                      `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	NetworkOrigin    *string                      `json:"NetworkOrigin,omitempty" xml:"NetworkOrigin,omitempty"`
	Status           *string                      `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcConfiguration *AccessPointVpcConfiguration `json:"VpcConfiguration,omitempty" xml:"VpcConfiguration,omitempty"`
}

func (s AccessPoint) String() string {
	return dara.Prettify(s)
}

func (s AccessPoint) GoString() string {
	return s.String()
}

func (s *AccessPoint) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *AccessPoint) GetAlias() *string {
	return s.Alias
}

func (s *AccessPoint) GetBucket() *string {
	return s.Bucket
}

func (s *AccessPoint) GetNetworkOrigin() *string {
	return s.NetworkOrigin
}

func (s *AccessPoint) GetStatus() *string {
	return s.Status
}

func (s *AccessPoint) GetVpcConfiguration() *AccessPointVpcConfiguration {
	return s.VpcConfiguration
}

func (s *AccessPoint) SetAccessPointName(v string) *AccessPoint {
	s.AccessPointName = &v
	return s
}

func (s *AccessPoint) SetAlias(v string) *AccessPoint {
	s.Alias = &v
	return s
}

func (s *AccessPoint) SetBucket(v string) *AccessPoint {
	s.Bucket = &v
	return s
}

func (s *AccessPoint) SetNetworkOrigin(v string) *AccessPoint {
	s.NetworkOrigin = &v
	return s
}

func (s *AccessPoint) SetStatus(v string) *AccessPoint {
	s.Status = &v
	return s
}

func (s *AccessPoint) SetVpcConfiguration(v *AccessPointVpcConfiguration) *AccessPoint {
	s.VpcConfiguration = v
	return s
}

func (s *AccessPoint) Validate() error {
	if s.VpcConfiguration != nil {
		if err := s.VpcConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
