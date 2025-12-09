// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListCacheRequest
	GetMarker() *string
	SetMaxKeys(v int64) *ListCacheRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListCacheRequest
	GetPrefix() *string
}

type ListCacheRequest struct {
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxKeys *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix  *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCacheRequest) GoString() string {
	return s.String()
}

func (s *ListCacheRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListCacheRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListCacheRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListCacheRequest) SetMarker(v string) *ListCacheRequest {
	s.Marker = &v
	return s
}

func (s *ListCacheRequest) SetMaxKeys(v int64) *ListCacheRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListCacheRequest) SetPrefix(v string) *ListCacheRequest {
	s.Prefix = &v
	return s
}

func (s *ListCacheRequest) Validate() error {
	return dara.Validate(s)
}
