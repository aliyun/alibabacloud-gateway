// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolsRequest
	GetContinuationToken() *string
	SetMaxKeys(v int64) *ListResourcePoolsRequest
	GetMaxKeys() *int64
}

type ListResourcePoolsRequest struct {
	// example:
	//
	// rp-01
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListResourcePoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolsRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePoolsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListResourcePoolsRequest) SetContinuationToken(v string) *ListResourcePoolsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolsRequest) SetMaxKeys(v int64) *ListResourcePoolsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListResourcePoolsRequest) Validate() error {
	return dara.Validate(s)
}
