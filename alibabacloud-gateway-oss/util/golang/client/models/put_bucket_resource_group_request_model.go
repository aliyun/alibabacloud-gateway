// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketResourceGroupConfiguration(v *BucketResourceGroupConfiguration) *PutBucketResourceGroupRequest
	GetBucketResourceGroupConfiguration() *BucketResourceGroupConfiguration
}

type PutBucketResourceGroupRequest struct {
	// The container that contains the ID of the resource group.
	BucketResourceGroupConfiguration *BucketResourceGroupConfiguration `json:"BucketResourceGroupConfiguration,omitempty" xml:"BucketResourceGroupConfiguration,omitempty"`
}

func (s PutBucketResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *PutBucketResourceGroupRequest) GetBucketResourceGroupConfiguration() *BucketResourceGroupConfiguration {
	return s.BucketResourceGroupConfiguration
}

func (s *PutBucketResourceGroupRequest) SetBucketResourceGroupConfiguration(v *BucketResourceGroupConfiguration) *PutBucketResourceGroupRequest {
	s.BucketResourceGroupConfiguration = v
	return s
}

func (s *PutBucketResourceGroupRequest) Validate() error {
	if s.BucketResourceGroupConfiguration != nil {
		if err := s.BucketResourceGroupConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
