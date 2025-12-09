// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlPolicy(v *GetBucketAclResponseBodyAccessControlPolicy) *GetBucketAclResponseBody
	GetAccessControlPolicy() *GetBucketAclResponseBodyAccessControlPolicy
}

type GetBucketAclResponseBody struct {
	// The container that stores the ACL information.
	AccessControlPolicy *GetBucketAclResponseBodyAccessControlPolicy `json:"AccessControlPolicy,omitempty" xml:"AccessControlPolicy,omitempty" type:"Struct"`
}

func (s GetBucketAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponseBody) GetAccessControlPolicy() *GetBucketAclResponseBodyAccessControlPolicy {
	return s.AccessControlPolicy
}

func (s *GetBucketAclResponseBody) SetAccessControlPolicy(v *GetBucketAclResponseBodyAccessControlPolicy) *GetBucketAclResponseBody {
	s.AccessControlPolicy = v
	return s
}

func (s *GetBucketAclResponseBody) Validate() error {
	if s.AccessControlPolicy != nil {
		if err := s.AccessControlPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketAclResponseBodyAccessControlPolicy struct {
	// The class of the container that stores the ACL information.
	AccessControlList *GetBucketAclResponseBodyAccessControlPolicyAccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Struct"`
	// The container that stores the information about the bucket owner.
	Owner *Owner `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s GetBucketAclResponseBodyAccessControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAclResponseBodyAccessControlPolicy) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponseBodyAccessControlPolicy) GetAccessControlList() *GetBucketAclResponseBodyAccessControlPolicyAccessControlList {
	return s.AccessControlList
}

func (s *GetBucketAclResponseBodyAccessControlPolicy) GetOwner() *Owner {
	return s.Owner
}

func (s *GetBucketAclResponseBodyAccessControlPolicy) SetAccessControlList(v *GetBucketAclResponseBodyAccessControlPolicyAccessControlList) *GetBucketAclResponseBodyAccessControlPolicy {
	s.AccessControlList = v
	return s
}

func (s *GetBucketAclResponseBodyAccessControlPolicy) SetOwner(v *Owner) *GetBucketAclResponseBodyAccessControlPolicy {
	s.Owner = v
	return s
}

func (s *GetBucketAclResponseBodyAccessControlPolicy) Validate() error {
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

type GetBucketAclResponseBodyAccessControlPolicyAccessControlList struct {
	// The ACL of the bucket.
	Grant *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s GetBucketAclResponseBodyAccessControlPolicyAccessControlList) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAclResponseBodyAccessControlPolicyAccessControlList) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponseBodyAccessControlPolicyAccessControlList) GetGrant() *string {
	return s.Grant
}

func (s *GetBucketAclResponseBodyAccessControlPolicyAccessControlList) SetGrant(v string) *GetBucketAclResponseBodyAccessControlPolicyAccessControlList {
	s.Grant = &v
	return s
}

func (s *GetBucketAclResponseBodyAccessControlPolicyAccessControlList) Validate() error {
	return dara.Validate(s)
}
