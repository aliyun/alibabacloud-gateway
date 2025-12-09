// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorBucketsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListVectorBucketsHeaders
	GetCommonHeaders() map[string]*string
	SetXOssResourceGroupId(v string) *ListVectorBucketsHeaders
	GetXOssResourceGroupId() *string
}

type ListVectorBucketsHeaders struct {
	CommonHeaders       map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XOssResourceGroupId *string            `json:"x-oss-resource-group-id,omitempty" xml:"x-oss-resource-group-id,omitempty"`
}

func (s ListVectorBucketsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListVectorBucketsHeaders) GoString() string {
	return s.String()
}

func (s *ListVectorBucketsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListVectorBucketsHeaders) GetXOssResourceGroupId() *string {
	return s.XOssResourceGroupId
}

func (s *ListVectorBucketsHeaders) SetCommonHeaders(v map[string]*string) *ListVectorBucketsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListVectorBucketsHeaders) SetXOssResourceGroupId(v string) *ListVectorBucketsHeaders {
	s.XOssResourceGroupId = &v
	return s
}

func (s *ListVectorBucketsHeaders) Validate() error {
	return dara.Validate(s)
}
