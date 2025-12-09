// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResourcePoolBucketGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourcePool(v string) *PutBucketResourcePoolBucketGroupRequest
	GetResourcePool() *string
	SetResourcePoolBucketGroup(v string) *PutBucketResourcePoolBucketGroupRequest
	GetResourcePoolBucketGroup() *string
}

type PutBucketResourcePoolBucketGroupRequest struct {
	// This parameter is required.
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
	// This parameter is required.
	ResourcePoolBucketGroup *string `json:"resourcePoolBucketGroup,omitempty" xml:"resourcePoolBucketGroup,omitempty"`
}

func (s PutBucketResourcePoolBucketGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResourcePoolBucketGroupRequest) GoString() string {
	return s.String()
}

func (s *PutBucketResourcePoolBucketGroupRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *PutBucketResourcePoolBucketGroupRequest) GetResourcePoolBucketGroup() *string {
	return s.ResourcePoolBucketGroup
}

func (s *PutBucketResourcePoolBucketGroupRequest) SetResourcePool(v string) *PutBucketResourcePoolBucketGroupRequest {
	s.ResourcePool = &v
	return s
}

func (s *PutBucketResourcePoolBucketGroupRequest) SetResourcePoolBucketGroup(v string) *PutBucketResourcePoolBucketGroupRequest {
	s.ResourcePoolBucketGroup = &v
	return s
}

func (s *PutBucketResourcePoolBucketGroupRequest) Validate() error {
	return dara.Validate(s)
}
