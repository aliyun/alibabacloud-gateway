// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetObjectLink(v *ObjectLinkInfo) *GetObjectLinkResponseBody
	GetObjectLink() *ObjectLinkInfo
}

type GetObjectLinkResponseBody struct {
	ObjectLink *ObjectLinkInfo `json:"ObjectLink,omitempty" xml:"ObjectLink,omitempty"`
}

func (s GetObjectLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectLinkResponseBody) GetObjectLink() *ObjectLinkInfo {
	return s.ObjectLink
}

func (s *GetObjectLinkResponseBody) SetObjectLink(v *ObjectLinkInfo) *GetObjectLinkResponseBody {
	s.ObjectLink = v
	return s
}

func (s *GetObjectLinkResponseBody) Validate() error {
	if s.ObjectLink != nil {
		if err := s.ObjectLink.Validate(); err != nil {
			return err
		}
	}
	return nil
}
