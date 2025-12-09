// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolBucketGroupQoSInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResourcePoolBucketGroupQoSInfo(v *ResourcePoolBucketGroupQoSInfo) *GetResourcePoolBucketGroupQoSInfoResponseBody
	GetResourcePoolBucketGroupQoSInfo() *ResourcePoolBucketGroupQoSInfo
}

type GetResourcePoolBucketGroupQoSInfoResponseBody struct {
	ResourcePoolBucketGroupQoSInfo *ResourcePoolBucketGroupQoSInfo `json:"ResourcePoolBucketGroupQoSInfo,omitempty" xml:"ResourcePoolBucketGroupQoSInfo,omitempty"`
}

func (s GetResourcePoolBucketGroupQoSInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolBucketGroupQoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePoolBucketGroupQoSInfoResponseBody) GetResourcePoolBucketGroupQoSInfo() *ResourcePoolBucketGroupQoSInfo {
	return s.ResourcePoolBucketGroupQoSInfo
}

func (s *GetResourcePoolBucketGroupQoSInfoResponseBody) SetResourcePoolBucketGroupQoSInfo(v *ResourcePoolBucketGroupQoSInfo) *GetResourcePoolBucketGroupQoSInfoResponseBody {
	s.ResourcePoolBucketGroupQoSInfo = v
	return s
}

func (s *GetResourcePoolBucketGroupQoSInfoResponseBody) Validate() error {
	if s.ResourcePoolBucketGroupQoSInfo != nil {
		if err := s.ResourcePoolBucketGroupQoSInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
