// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolBucketGroupQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourcePool(v string) *GetResourcePoolBucketGroupQoSInfoRequest
	GetResourcePool() *string
	SetResourcePoolBucketGroup(v string) *GetResourcePoolBucketGroupQoSInfoRequest
	GetResourcePoolBucketGroup() *string
}

type GetResourcePoolBucketGroupQoSInfoRequest struct {
	// This parameter is required.
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
	// This parameter is required.
	ResourcePoolBucketGroup *string `json:"resourcePoolBucketGroup,omitempty" xml:"resourcePoolBucketGroup,omitempty"`
}

func (s GetResourcePoolBucketGroupQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolBucketGroupQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePoolBucketGroupQoSInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *GetResourcePoolBucketGroupQoSInfoRequest) GetResourcePoolBucketGroup() *string {
	return s.ResourcePoolBucketGroup
}

func (s *GetResourcePoolBucketGroupQoSInfoRequest) SetResourcePool(v string) *GetResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *GetResourcePoolBucketGroupQoSInfoRequest) SetResourcePoolBucketGroup(v string) *GetResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePoolBucketGroup = &v
	return s
}

func (s *GetResourcePoolBucketGroupQoSInfoRequest) Validate() error {
	return dara.Validate(s)
}
