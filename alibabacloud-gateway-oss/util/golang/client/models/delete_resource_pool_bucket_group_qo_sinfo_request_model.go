// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourcePoolBucketGroupQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourcePool(v string) *DeleteResourcePoolBucketGroupQoSInfoRequest
	GetResourcePool() *string
	SetResourcePoolBucketGroup(v string) *DeleteResourcePoolBucketGroupQoSInfoRequest
	GetResourcePoolBucketGroup() *string
}

type DeleteResourcePoolBucketGroupQoSInfoRequest struct {
	// This parameter is required.
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
	// This parameter is required.
	ResourcePoolBucketGroup *string `json:"resourcePoolBucketGroup,omitempty" xml:"resourcePoolBucketGroup,omitempty"`
}

func (s DeleteResourcePoolBucketGroupQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourcePoolBucketGroupQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourcePoolBucketGroupQoSInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *DeleteResourcePoolBucketGroupQoSInfoRequest) GetResourcePoolBucketGroup() *string {
	return s.ResourcePoolBucketGroup
}

func (s *DeleteResourcePoolBucketGroupQoSInfoRequest) SetResourcePool(v string) *DeleteResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *DeleteResourcePoolBucketGroupQoSInfoRequest) SetResourcePoolBucketGroup(v string) *DeleteResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePoolBucketGroup = &v
	return s
}

func (s *DeleteResourcePoolBucketGroupQoSInfoRequest) Validate() error {
	return dara.Validate(s)
}
