// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlPolicy(v *GetObjectAclResponseBodyAccessControlPolicy) *GetObjectAclResponseBody
	GetAccessControlPolicy() *GetObjectAclResponseBodyAccessControlPolicy
}

type GetObjectAclResponseBody struct {
	// The container that stores the results of the GetObjectACL request.
	AccessControlPolicy *GetObjectAclResponseBodyAccessControlPolicy `json:"AccessControlPolicy,omitempty" xml:"AccessControlPolicy,omitempty" type:"Struct"`
}

func (s GetObjectAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponseBody) GetAccessControlPolicy() *GetObjectAclResponseBodyAccessControlPolicy {
	return s.AccessControlPolicy
}

func (s *GetObjectAclResponseBody) SetAccessControlPolicy(v *GetObjectAclResponseBodyAccessControlPolicy) *GetObjectAclResponseBody {
	s.AccessControlPolicy = v
	return s
}

func (s *GetObjectAclResponseBody) Validate() error {
	if s.AccessControlPolicy != nil {
		if err := s.AccessControlPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetObjectAclResponseBodyAccessControlPolicy struct {
	// The container that stores the ACL information.
	AccessControlList *GetObjectAclResponseBodyAccessControlPolicyAccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Struct"`
	// The container that stores the information about the bucket owner.
	Owner *Owner `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s GetObjectAclResponseBodyAccessControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetObjectAclResponseBodyAccessControlPolicy) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponseBodyAccessControlPolicy) GetAccessControlList() *GetObjectAclResponseBodyAccessControlPolicyAccessControlList {
	return s.AccessControlList
}

func (s *GetObjectAclResponseBodyAccessControlPolicy) GetOwner() *Owner {
	return s.Owner
}

func (s *GetObjectAclResponseBodyAccessControlPolicy) SetAccessControlList(v *GetObjectAclResponseBodyAccessControlPolicyAccessControlList) *GetObjectAclResponseBodyAccessControlPolicy {
	s.AccessControlList = v
	return s
}

func (s *GetObjectAclResponseBodyAccessControlPolicy) SetOwner(v *Owner) *GetObjectAclResponseBodyAccessControlPolicy {
	s.Owner = v
	return s
}

func (s *GetObjectAclResponseBodyAccessControlPolicy) Validate() error {
	if s.AccessControlList != nil {
		if err := s.AccessControlList.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetObjectAclResponseBodyAccessControlPolicyAccessControlList struct {
	// The ACL of the object. Default value: default.
	ACL *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s GetObjectAclResponseBodyAccessControlPolicyAccessControlList) String() string {
	return dara.Prettify(s)
}

func (s GetObjectAclResponseBodyAccessControlPolicyAccessControlList) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponseBodyAccessControlPolicyAccessControlList) GetACL() *string {
	return s.ACL
}

func (s *GetObjectAclResponseBodyAccessControlPolicyAccessControlList) SetACL(v string) *GetObjectAclResponseBodyAccessControlPolicyAccessControlList {
	s.ACL = &v
	return s
}

func (s *GetObjectAclResponseBodyAccessControlPolicyAccessControlList) Validate() error {
	return dara.Validate(s)
}
