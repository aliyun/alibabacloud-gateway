// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostObjectGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateFileGroup(v *CreateFileGroupResult) *PostObjectGroupResponseBody
	GetCreateFileGroup() *CreateFileGroupResult
}

type PostObjectGroupResponseBody struct {
	CreateFileGroup *CreateFileGroupResult `json:"CreateFileGroup,omitempty" xml:"CreateFileGroup,omitempty"`
}

func (s PostObjectGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostObjectGroupResponseBody) GoString() string {
	return s.String()
}

func (s *PostObjectGroupResponseBody) GetCreateFileGroup() *CreateFileGroupResult {
	return s.CreateFileGroup
}

func (s *PostObjectGroupResponseBody) SetCreateFileGroup(v *CreateFileGroupResult) *PostObjectGroupResponseBody {
	s.CreateFileGroup = v
	return s
}

func (s *PostObjectGroupResponseBody) Validate() error {
	if s.CreateFileGroup != nil {
		if err := s.CreateFileGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}
