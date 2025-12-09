// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourcePoolBucketGroupQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourcePoolBucketGroupQoSInfo(v *ResourcePoolBucketGroupQoSInfo) *PutResourcePoolBucketGroupQoSInfoRequest
	GetResourcePoolBucketGroupQoSInfo() *ResourcePoolBucketGroupQoSInfo
	SetResourcePool(v string) *PutResourcePoolBucketGroupQoSInfoRequest
	GetResourcePool() *string
	SetResourcePoolBucketGroup(v string) *PutResourcePoolBucketGroupQoSInfoRequest
	GetResourcePoolBucketGroup() *string
}

type PutResourcePoolBucketGroupQoSInfoRequest struct {
	ResourcePoolBucketGroupQoSInfo *ResourcePoolBucketGroupQoSInfo `json:"ResourcePoolBucketGroupQoSInfo,omitempty" xml:"ResourcePoolBucketGroupQoSInfo,omitempty"`
	// This parameter is required.
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
	// This parameter is required.
	ResourcePoolBucketGroup *string `json:"resourcePoolBucketGroup,omitempty" xml:"resourcePoolBucketGroup,omitempty"`
}

func (s PutResourcePoolBucketGroupQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourcePoolBucketGroupQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) GetResourcePoolBucketGroupQoSInfo() *ResourcePoolBucketGroupQoSInfo {
	return s.ResourcePoolBucketGroupQoSInfo
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) GetResourcePoolBucketGroup() *string {
	return s.ResourcePoolBucketGroup
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) SetResourcePoolBucketGroupQoSInfo(v *ResourcePoolBucketGroupQoSInfo) *PutResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePoolBucketGroupQoSInfo = v
	return s
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) SetResourcePool(v string) *PutResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) SetResourcePoolBucketGroup(v string) *PutResourcePoolBucketGroupQoSInfoRequest {
	s.ResourcePoolBucketGroup = &v
	return s
}

func (s *PutResourcePoolBucketGroupQoSInfoRequest) Validate() error {
	if s.ResourcePoolBucketGroupQoSInfo != nil {
		if err := s.ResourcePoolBucketGroupQoSInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
