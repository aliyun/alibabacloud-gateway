// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketAclHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutBucketAclHeaders
	GetCommonHeaders() map[string]*string
	SetAcl(v string) *PutBucketAclHeaders
	GetAcl() *string
}

type PutBucketAclHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The ACL that you want to configure or modify for the bucket. The x-oss-acl header is included in PutBucketAcl requests to configure or modify the ACL of the bucket. If this header is not included, the ACL configurations do not take effect.\\
	//
	// Valid values:
	//
	// 	- public-read-write: All users can read and write objects in the bucket. Exercise caution when you set the value to public-read-write.
	//
	// 	- public-read: Only the owner and authorized users of the bucket can read and write objects in the bucket. Other users can only read objects in the bucket. Exercise caution when you set the value to public-read.
	//
	// 	- private: Only the owner and authorized users of this bucket can read and write objects in the bucket. Other users cannot access objects in the bucket.
	//
	// This parameter is required.
	Acl *string `json:"x-oss-acl,omitempty" xml:"x-oss-acl,omitempty"`
}

func (s PutBucketAclHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutBucketAclHeaders) GoString() string {
	return s.String()
}

func (s *PutBucketAclHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutBucketAclHeaders) GetAcl() *string {
	return s.Acl
}

func (s *PutBucketAclHeaders) SetCommonHeaders(v map[string]*string) *PutBucketAclHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketAclHeaders) SetAcl(v string) *PutBucketAclHeaders {
	s.Acl = &v
	return s
}

func (s *PutBucketAclHeaders) Validate() error {
	return dara.Validate(s)
}
