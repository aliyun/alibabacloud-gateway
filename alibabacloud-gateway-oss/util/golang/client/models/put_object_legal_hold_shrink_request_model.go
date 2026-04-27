// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectLegalHoldShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *PutObjectLegalHoldShrinkRequest
	GetBodyShrink() *string
}

type PutObjectLegalHoldShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutObjectLegalHoldShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLegalHoldShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutObjectLegalHoldShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *PutObjectLegalHoldShrinkRequest) SetBodyShrink(v string) *PutObjectLegalHoldShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *PutObjectLegalHoldShrinkRequest) Validate() error {
	return dara.Validate(s)
}
