// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateObjectLink(v *CreateObjectLinkResult) *PutObjectLinkResponseBody
	GetCreateObjectLink() *CreateObjectLinkResult
}

type PutObjectLinkResponseBody struct {
	CreateObjectLink *CreateObjectLinkResult `json:"CreateObjectLink,omitempty" xml:"CreateObjectLink,omitempty"`
}

func (s PutObjectLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLinkResponseBody) GoString() string {
	return s.String()
}

func (s *PutObjectLinkResponseBody) GetCreateObjectLink() *CreateObjectLinkResult {
	return s.CreateObjectLink
}

func (s *PutObjectLinkResponseBody) SetCreateObjectLink(v *CreateObjectLinkResult) *PutObjectLinkResponseBody {
	s.CreateObjectLink = v
	return s
}

func (s *PutObjectLinkResponseBody) Validate() error {
	if s.CreateObjectLink != nil {
		if err := s.CreateObjectLink.Validate(); err != nil {
			return err
		}
	}
	return nil
}
