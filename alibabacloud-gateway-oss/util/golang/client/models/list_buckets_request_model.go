// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListBucketsRequest
	GetMarker() *string
	SetMaxKeys(v int64) *ListBucketsRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListBucketsRequest
	GetPrefix() *string
	SetTagKey(v string) *ListBucketsRequest
	GetTagKey() *string
	SetTagValue(v string) *ListBucketsRequest
	GetTagValue() *string
	SetTagging(v string) *ListBucketsRequest
	GetTagging() *string
}

type ListBucketsRequest struct {
	// The name of the bucket from which the buckets start to return. The buckets whose names are alphabetically after the value of marker are returned. If this parameter is not specified, all results are returned. By default, this parameter is left empty.
	//
	// example:
	//
	// mybucket10
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of buckets that can be returned. Valid values: 1 to 1000. Default value: 100
	//
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// The prefix that the names of returned buckets must contain. If this parameter is not specified, prefixes are not used to filter returned buckets. By default, this parameter is left empty.
	//
	// example:
	//
	// my
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// A tag key of target buckets. The listing results will only include Buckets that have been tagged with this key.
	//
	// example:
	//
	// test-key
	TagKey *string `json:"tag-key,omitempty" xml:"tag-key,omitempty"`
	// A tag value for the target buckets. If this parameter is specified in the request, the tag-key must also be specified. The listing results will only include Buckets that have been tagged with this key-value pair.
	//
	// example:
	//
	// test-value
	TagValue *string `json:"tag-value,omitempty" xml:"tag-value,omitempty"`
	// Tag list of target buckets. Only Buckets that match all the key-value pairs in the list will added into the listing results. The tagging parameter cannot be used with the tag-key and tag-value parameters in a request.
	//
	// example:
	//
	// "test-key":"test-value"
	Tagging *string `json:"tagging,omitempty" xml:"tagging,omitempty"`
}

func (s ListBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListBucketsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListBucketsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListBucketsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListBucketsRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListBucketsRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *ListBucketsRequest) GetTagging() *string {
	return s.Tagging
}

func (s *ListBucketsRequest) SetMarker(v string) *ListBucketsRequest {
	s.Marker = &v
	return s
}

func (s *ListBucketsRequest) SetMaxKeys(v int64) *ListBucketsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListBucketsRequest) SetPrefix(v string) *ListBucketsRequest {
	s.Prefix = &v
	return s
}

func (s *ListBucketsRequest) SetTagKey(v string) *ListBucketsRequest {
	s.TagKey = &v
	return s
}

func (s *ListBucketsRequest) SetTagValue(v string) *ListBucketsRequest {
	s.TagValue = &v
	return s
}

func (s *ListBucketsRequest) SetTagging(v string) *ListBucketsRequest {
	s.Tagging = &v
	return s
}

func (s *ListBucketsRequest) Validate() error {
	return dara.Validate(s)
}
