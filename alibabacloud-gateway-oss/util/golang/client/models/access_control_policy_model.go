// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessControlPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlList(v *AccessControlList) *AccessControlPolicy
	GetAccessControlList() *AccessControlList
	SetOwner(v *Owner) *AccessControlPolicy
	GetOwner() *Owner
}

type AccessControlPolicy struct {
	AccessControlList *AccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty"`
	Owner             *Owner             `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s AccessControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s AccessControlPolicy) GoString() string {
	return s.String()
}

func (s *AccessControlPolicy) GetAccessControlList() *AccessControlList {
	return s.AccessControlList
}

func (s *AccessControlPolicy) GetOwner() *Owner {
	return s.Owner
}

func (s *AccessControlPolicy) SetAccessControlList(v *AccessControlList) *AccessControlPolicy {
	s.AccessControlList = v
	return s
}

func (s *AccessControlPolicy) SetOwner(v *Owner) *AccessControlPolicy {
	s.Owner = v
	return s
}

func (s *AccessControlPolicy) Validate() error {
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
