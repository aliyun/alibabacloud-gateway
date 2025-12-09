// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessControlList interface {
	dara.Model
	String() string
	GoString() string
	SetGrant(v string) *AccessControlList
	GetGrant() *string
}

type AccessControlList struct {
	Grant *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s AccessControlList) String() string {
	return dara.Prettify(s)
}

func (s AccessControlList) GoString() string {
	return s.String()
}

func (s *AccessControlList) GetGrant() *string {
	return s.Grant
}

func (s *AccessControlList) SetGrant(v string) *AccessControlList {
	s.Grant = &v
	return s
}

func (s *AccessControlList) Validate() error {
	return dara.Validate(s)
}
