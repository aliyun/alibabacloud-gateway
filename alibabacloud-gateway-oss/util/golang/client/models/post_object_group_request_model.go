// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostObjectGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateFileGroup(v *CreateFileGroup) *PostObjectGroupRequest
	GetCreateFileGroup() *CreateFileGroup
}

type PostObjectGroupRequest struct {
	CreateFileGroup *CreateFileGroup `json:"CreateFileGroup,omitempty" xml:"CreateFileGroup,omitempty"`
}

func (s PostObjectGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s PostObjectGroupRequest) GoString() string {
	return s.String()
}

func (s *PostObjectGroupRequest) GetCreateFileGroup() *CreateFileGroup {
	return s.CreateFileGroup
}

func (s *PostObjectGroupRequest) SetCreateFileGroup(v *CreateFileGroup) *PostObjectGroupRequest {
	s.CreateFileGroup = v
	return s
}

func (s *PostObjectGroupRequest) Validate() error {
	if s.CreateFileGroup != nil {
		if err := s.CreateFileGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}
