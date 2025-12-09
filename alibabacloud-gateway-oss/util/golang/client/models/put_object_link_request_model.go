// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateObjectLink(v *ObjectLinkInfo) *PutObjectLinkRequest
	GetCreateObjectLink() *ObjectLinkInfo
}

type PutObjectLinkRequest struct {
	CreateObjectLink *ObjectLinkInfo `json:"CreateObjectLink,omitempty" xml:"CreateObjectLink,omitempty"`
}

func (s PutObjectLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLinkRequest) GoString() string {
	return s.String()
}

func (s *PutObjectLinkRequest) GetCreateObjectLink() *ObjectLinkInfo {
	return s.CreateObjectLink
}

func (s *PutObjectLinkRequest) SetCreateObjectLink(v *ObjectLinkInfo) *PutObjectLinkRequest {
	s.CreateObjectLink = v
	return s
}

func (s *PutObjectLinkRequest) Validate() error {
	if s.CreateObjectLink != nil {
		if err := s.CreateObjectLink.Validate(); err != nil {
			return err
		}
	}
	return nil
}
