// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupQoSInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolBucketGroupQoSInfosRequest
	GetContinuationToken() *string
	SetMaxKeys(v string) *ListResourcePoolBucketGroupQoSInfosRequest
	GetMaxKeys() *string
	SetResourcePool(v string) *ListResourcePoolBucketGroupQoSInfosRequest
	GetResourcePool() *string
}

type ListResourcePoolBucketGroupQoSInfosRequest struct {
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	MaxKeys           *string `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// This parameter is required.
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
}

func (s ListResourcePoolBucketGroupQoSInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupQoSInfosRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) GetMaxKeys() *string {
	return s.MaxKeys
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) SetContinuationToken(v string) *ListResourcePoolBucketGroupQoSInfosRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) SetMaxKeys(v string) *ListResourcePoolBucketGroupQoSInfosRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) SetResourcePool(v string) *ListResourcePoolBucketGroupQoSInfosRequest {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosRequest) Validate() error {
	return dara.Validate(s)
}
