// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPublicAccessBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutAccessPointPublicAccessBlockRequest
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
	SetXOssAccessPointName(v string) *PutAccessPointPublicAccessBlockRequest
	GetXOssAccessPointName() *string
}

type PutAccessPointPublicAccessBlockRequest struct {
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
	// The name of the access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-01
	XOssAccessPointName *string `json:"x-oss-access-point-name,omitempty" xml:"x-oss-access-point-name,omitempty"`
}

func (s PutAccessPointPublicAccessBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPublicAccessBlockRequest) GoString() string {
	return s.String()
}

func (s *PutAccessPointPublicAccessBlockRequest) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *PutAccessPointPublicAccessBlockRequest) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *PutAccessPointPublicAccessBlockRequest) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutAccessPointPublicAccessBlockRequest {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *PutAccessPointPublicAccessBlockRequest) SetXOssAccessPointName(v string) *PutAccessPointPublicAccessBlockRequest {
	s.XOssAccessPointName = &v
	return s
}

func (s *PutAccessPointPublicAccessBlockRequest) Validate() error {
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
