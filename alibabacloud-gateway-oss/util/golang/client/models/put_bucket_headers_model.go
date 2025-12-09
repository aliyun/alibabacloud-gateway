// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutBucketHeaders
	GetCommonHeaders() map[string]*string
	SetAcl(v string) *PutBucketHeaders
	GetAcl() *string
	SetXOssBucketTagging(v string) *PutBucketHeaders
	GetXOssBucketTagging() *string
	SetXOssResourceGroupId(v string) *PutBucketHeaders
	GetXOssResourceGroupId() *string
}

type PutBucketHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The access control list (ACL) of the bucket to be created. Valid values:
	//
	// 	- public-read-write
	//
	// 	- public-read
	//
	// 	- private (default)
	//
	// For more information, see [Bucket ACL](https://help.aliyun.com/document_detail/31843.html).
	Acl               *string `json:"x-oss-acl,omitempty" xml:"x-oss-acl,omitempty"`
	XOssBucketTagging *string `json:"x-oss-bucket-tagging,omitempty" xml:"x-oss-bucket-tagging,omitempty"`
	// The ID of the resource group.
	//
	// 	- If you include the header in the request and specify the ID of the resource group, the bucket that you create belongs to the resource group. If the specified resource group ID is rg-default-id, the bucket that you create belongs to the default resource group.
	//
	// 	- If you do not include the header in the request, the bucket that you create belongs to the default resource group.
	//
	// You can obtain the ID of a resource group in the Resource Management console or by calling the ListResourceGroups operation. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html) and [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// >  You cannot configure a resource group for an Anywhere Bucket.
	XOssResourceGroupId *string `json:"x-oss-resource-group-id,omitempty" xml:"x-oss-resource-group-id,omitempty"`
}

func (s PutBucketHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutBucketHeaders) GoString() string {
	return s.String()
}

func (s *PutBucketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutBucketHeaders) GetAcl() *string {
	return s.Acl
}

func (s *PutBucketHeaders) GetXOssBucketTagging() *string {
	return s.XOssBucketTagging
}

func (s *PutBucketHeaders) GetXOssResourceGroupId() *string {
	return s.XOssResourceGroupId
}

func (s *PutBucketHeaders) SetCommonHeaders(v map[string]*string) *PutBucketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketHeaders) SetAcl(v string) *PutBucketHeaders {
	s.Acl = &v
	return s
}

func (s *PutBucketHeaders) SetXOssBucketTagging(v string) *PutBucketHeaders {
	s.XOssBucketTagging = &v
	return s
}

func (s *PutBucketHeaders) SetXOssResourceGroupId(v string) *PutBucketHeaders {
	s.XOssResourceGroupId = &v
	return s
}

func (s *PutBucketHeaders) Validate() error {
	return dara.Validate(s)
}
