// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListBucketsHeaders
	GetCommonHeaders() map[string]*string
	SetXOssResourceGroupId(v string) *ListBucketsHeaders
	GetXOssResourceGroupId() *string
}

type ListBucketsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The ID of the resource group to which the bucket belongs.
	XOssResourceGroupId *string `json:"x-oss-resource-group-id,omitempty" xml:"x-oss-resource-group-id,omitempty"`
}

func (s ListBucketsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsHeaders) GoString() string {
	return s.String()
}

func (s *ListBucketsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListBucketsHeaders) GetXOssResourceGroupId() *string {
	return s.XOssResourceGroupId
}

func (s *ListBucketsHeaders) SetCommonHeaders(v map[string]*string) *ListBucketsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListBucketsHeaders) SetXOssResourceGroupId(v string) *ListBucketsHeaders {
	s.XOssResourceGroupId = &v
	return s
}

func (s *ListBucketsHeaders) Validate() error {
	return dara.Validate(s)
}
