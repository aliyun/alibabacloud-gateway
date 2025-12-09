// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolBucketGroupsRequest
	GetContinuationToken() *string
	SetMaxKeys(v string) *ListResourcePoolBucketGroupsRequest
	GetMaxKeys() *string
	SetResourcePool(v string) *ListResourcePoolBucketGroupsRequest
	GetResourcePool() *string
}

type ListResourcePoolBucketGroupsRequest struct {
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	MaxKeys           *string `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// This parameter is required.
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
}

func (s ListResourcePoolBucketGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolBucketGroupsRequest) GetMaxKeys() *string {
	return s.MaxKeys
}

func (s *ListResourcePoolBucketGroupsRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolBucketGroupsRequest) SetContinuationToken(v string) *ListResourcePoolBucketGroupsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketGroupsRequest) SetMaxKeys(v string) *ListResourcePoolBucketGroupsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListResourcePoolBucketGroupsRequest) SetResourcePool(v string) *ListResourcePoolBucketGroupsRequest {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolBucketGroupsRequest) Validate() error {
	return dara.Validate(s)
}
