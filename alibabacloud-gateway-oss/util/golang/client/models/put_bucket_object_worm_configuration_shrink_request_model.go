// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketObjectWormConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *PutBucketObjectWormConfigurationShrinkRequest
	GetBodyShrink() *string
}

type PutBucketObjectWormConfigurationShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutBucketObjectWormConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketObjectWormConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutBucketObjectWormConfigurationShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *PutBucketObjectWormConfigurationShrinkRequest) SetBodyShrink(v string) *PutBucketObjectWormConfigurationShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *PutBucketObjectWormConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
