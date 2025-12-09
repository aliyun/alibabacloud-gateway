// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketAntiDDosInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListBucketAntiDDosInfoRequest
	GetMarker() *string
	SetMaxKeys(v string) *ListBucketAntiDDosInfoRequest
	GetMaxKeys() *string
}

type ListBucketAntiDDosInfoRequest struct {
	// The name of the Anti-DDoS instance from which the list starts. The Anti-DDoS instances whose names are alphabetically after the value of marker are returned.
	//
	// >  You can set marker to an empty string in the first request. If IsTruncated is returned in the response and the value of IsTruncated is true, you must use the value of Marker in the response as the value of marker in the next request.
	//
	// example:
	//
	// nextMarker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of Anti-DDoS instances that can be returned.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxKeys *string `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListBucketAntiDDosInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucketAntiDDosInfoRequest) GoString() string {
	return s.String()
}

func (s *ListBucketAntiDDosInfoRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListBucketAntiDDosInfoRequest) GetMaxKeys() *string {
	return s.MaxKeys
}

func (s *ListBucketAntiDDosInfoRequest) SetMarker(v string) *ListBucketAntiDDosInfoRequest {
	s.Marker = &v
	return s
}

func (s *ListBucketAntiDDosInfoRequest) SetMaxKeys(v string) *ListBucketAntiDDosInfoRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListBucketAntiDDosInfoRequest) Validate() error {
	return dara.Validate(s)
}
