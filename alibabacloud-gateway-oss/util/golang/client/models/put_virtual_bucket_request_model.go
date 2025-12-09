// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVirtualBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVirtualBucketConfiguration(v *VirtualBucketConfiguration) *PutVirtualBucketRequest
	GetVirtualBucketConfiguration() *VirtualBucketConfiguration
}

type PutVirtualBucketRequest struct {
	VirtualBucketConfiguration *VirtualBucketConfiguration `json:"VirtualBucketConfiguration,omitempty" xml:"VirtualBucketConfiguration,omitempty"`
}

func (s PutVirtualBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s PutVirtualBucketRequest) GoString() string {
	return s.String()
}

func (s *PutVirtualBucketRequest) GetVirtualBucketConfiguration() *VirtualBucketConfiguration {
	return s.VirtualBucketConfiguration
}

func (s *PutVirtualBucketRequest) SetVirtualBucketConfiguration(v *VirtualBucketConfiguration) *PutVirtualBucketRequest {
	s.VirtualBucketConfiguration = v
	return s
}

func (s *PutVirtualBucketRequest) Validate() error {
	if s.VirtualBucketConfiguration != nil {
		if err := s.VirtualBucketConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
