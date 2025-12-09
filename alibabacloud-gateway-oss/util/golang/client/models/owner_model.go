// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOwner interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *Owner
	GetDisplayName() *string
	SetID(v string) *Owner
	GetID() *string
}

type Owner struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ID          *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s Owner) String() string {
	return dara.Prettify(s)
}

func (s Owner) GoString() string {
	return s.String()
}

func (s *Owner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *Owner) GetID() *string {
	return s.ID
}

func (s *Owner) SetDisplayName(v string) *Owner {
	s.DisplayName = &v
	return s
}

func (s *Owner) SetID(v string) *Owner {
	s.ID = &v
	return s
}

func (s *Owner) Validate() error {
	return dara.Validate(s)
}
