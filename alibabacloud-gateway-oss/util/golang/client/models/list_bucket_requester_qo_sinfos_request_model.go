// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRequesterQoSInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListBucketRequesterQoSInfosRequest
	GetContinuationToken() *string
	SetMaxKeys(v int64) *ListBucketRequesterQoSInfosRequest
	GetMaxKeys() *int64
}

type ListBucketRequesterQoSInfosRequest struct {
	// example:
	//
	// 26753xxxxxxxx14340
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListBucketRequesterQoSInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRequesterQoSInfosRequest) GoString() string {
	return s.String()
}

func (s *ListBucketRequesterQoSInfosRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListBucketRequesterQoSInfosRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListBucketRequesterQoSInfosRequest) SetContinuationToken(v string) *ListBucketRequesterQoSInfosRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListBucketRequesterQoSInfosRequest) SetMaxKeys(v int64) *ListBucketRequesterQoSInfosRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListBucketRequesterQoSInfosRequest) Validate() error {
	return dara.Validate(s)
}
