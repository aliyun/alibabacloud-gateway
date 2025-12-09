// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectGroupIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileGroup(v *FileGroupInfo) *GetObjectGroupIndexResponseBody
	GetFileGroup() *FileGroupInfo
}

type GetObjectGroupIndexResponseBody struct {
	FileGroup *FileGroupInfo `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
}

func (s GetObjectGroupIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectGroupIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectGroupIndexResponseBody) GetFileGroup() *FileGroupInfo {
	return s.FileGroup
}

func (s *GetObjectGroupIndexResponseBody) SetFileGroup(v *FileGroupInfo) *GetObjectGroupIndexResponseBody {
	s.FileGroup = v
	return s
}

func (s *GetObjectGroupIndexResponseBody) Validate() error {
	if s.FileGroup != nil {
		if err := s.FileGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}
