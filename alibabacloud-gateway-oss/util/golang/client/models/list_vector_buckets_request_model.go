// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListVectorBucketsRequest
	GetMarker() *string
	SetMaxKeys(v int64) *ListVectorBucketsRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListVectorBucketsRequest
	GetPrefix() *string
}

type ListVectorBucketsRequest struct {
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxKeys *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix  *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListVectorBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVectorBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListVectorBucketsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListVectorBucketsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListVectorBucketsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListVectorBucketsRequest) SetMarker(v string) *ListVectorBucketsRequest {
	s.Marker = &v
	return s
}

func (s *ListVectorBucketsRequest) SetMaxKeys(v int64) *ListVectorBucketsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListVectorBucketsRequest) SetPrefix(v string) *ListVectorBucketsRequest {
	s.Prefix = &v
	return s
}

func (s *ListVectorBucketsRequest) Validate() error {
	return dara.Validate(s)
}
