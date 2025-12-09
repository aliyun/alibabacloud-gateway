// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketResourceGroupConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *BucketResourceGroupConfiguration
	GetResourceGroupId() *string
}

type BucketResourceGroupConfiguration struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s BucketResourceGroupConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BucketResourceGroupConfiguration) GoString() string {
	return s.String()
}

func (s *BucketResourceGroupConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BucketResourceGroupConfiguration) SetResourceGroupId(v string) *BucketResourceGroupConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *BucketResourceGroupConfiguration) Validate() error {
	return dara.Validate(s)
}
