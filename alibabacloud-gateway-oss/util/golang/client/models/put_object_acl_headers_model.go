// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectAclHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutObjectAclHeaders
	GetCommonHeaders() map[string]*string
	SetAcl(v string) *PutObjectAclHeaders
	GetAcl() *string
}

type PutObjectAclHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The access control list (ACL) of the object.
	//
	// This parameter is required.
	Acl *string `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
}

func (s PutObjectAclHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutObjectAclHeaders) GoString() string {
	return s.String()
}

func (s *PutObjectAclHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutObjectAclHeaders) GetAcl() *string {
	return s.Acl
}

func (s *PutObjectAclHeaders) SetCommonHeaders(v map[string]*string) *PutObjectAclHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutObjectAclHeaders) SetAcl(v string) *PutObjectAclHeaders {
	s.Acl = &v
	return s
}

func (s *PutObjectAclHeaders) Validate() error {
	return dara.Validate(s)
}
