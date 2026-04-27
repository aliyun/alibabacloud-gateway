// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectRetentionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *PutObjectRetentionShrinkRequest
	GetBodyShrink() *string
}

type PutObjectRetentionShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutObjectRetentionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectRetentionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutObjectRetentionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *PutObjectRetentionShrinkRequest) SetBodyShrink(v string) *PutObjectRetentionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *PutObjectRetentionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
