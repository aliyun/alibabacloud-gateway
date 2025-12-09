// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolBucketsRequest
	GetContinuationToken() *string
	SetMaxKeys(v int64) *ListResourcePoolBucketsRequest
	GetMaxKeys() *int64
	SetResourcePool(v string) *ListResourcePoolBucketsRequest
	GetResourcePool() *string
}

type ListResourcePoolBucketsRequest struct {
	// example:
	//
	// test-bucket
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rp-01
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
}

func (s ListResourcePoolBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolBucketsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListResourcePoolBucketsRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolBucketsRequest) SetContinuationToken(v string) *ListResourcePoolBucketsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketsRequest) SetMaxKeys(v int64) *ListResourcePoolBucketsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListResourcePoolBucketsRequest) SetResourcePool(v string) *ListResourcePoolBucketsRequest {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolBucketsRequest) Validate() error {
	return dara.Validate(s)
}
