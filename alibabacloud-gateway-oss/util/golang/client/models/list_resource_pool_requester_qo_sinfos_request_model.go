// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolRequesterQoSInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolRequesterQoSInfosRequest
	GetContinuationToken() *string
	SetMaxKeys(v int64) *ListResourcePoolRequesterQoSInfosRequest
	GetMaxKeys() *int64
	SetResourcePool(v string) *ListResourcePoolRequesterQoSInfosRequest
	GetResourcePool() *string
}

type ListResourcePoolRequesterQoSInfosRequest struct {
	// example:
	//
	// 26753xxxxxxxx14340
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

func (s ListResourcePoolRequesterQoSInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolRequesterQoSInfosRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePoolRequesterQoSInfosRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolRequesterQoSInfosRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListResourcePoolRequesterQoSInfosRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolRequesterQoSInfosRequest) SetContinuationToken(v string) *ListResourcePoolRequesterQoSInfosRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosRequest) SetMaxKeys(v int64) *ListResourcePoolRequesterQoSInfosRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosRequest) SetResourcePool(v string) *ListResourcePoolRequesterQoSInfosRequest {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosRequest) Validate() error {
	return dara.Validate(s)
}
