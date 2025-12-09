// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketResourceGroupConfiguration(v *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) *GetBucketResourceGroupResponseBody
	GetBucketResourceGroupConfiguration() *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration
}

type GetBucketResourceGroupResponseBody struct {
	// The container that stores the ID of the resource group.
	BucketResourceGroupConfiguration *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration `json:"BucketResourceGroupConfiguration,omitempty" xml:"BucketResourceGroupConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketResourceGroupResponseBody) GetBucketResourceGroupConfiguration() *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration {
	return s.BucketResourceGroupConfiguration
}

func (s *GetBucketResourceGroupResponseBody) SetBucketResourceGroupConfiguration(v *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) *GetBucketResourceGroupResponseBody {
	s.BucketResourceGroupConfiguration = v
	return s
}

func (s *GetBucketResourceGroupResponseBody) Validate() error {
	if s.BucketResourceGroupConfiguration != nil {
		if err := s.BucketResourceGroupConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration struct {
	// The ID of the resource group to which the bucket belongs.
	//
	// If this element is not specified, the bucket is moved to the default resource group.
	//
	// example:
	//
	// rg-asdfklj***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) SetResourceGroupId(v string) *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration) Validate() error {
	return dara.Validate(s)
}
