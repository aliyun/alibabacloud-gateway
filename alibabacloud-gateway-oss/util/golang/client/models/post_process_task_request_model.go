// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostProcessTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PostProcessTaskRequest
	GetBody() *string
}

type PostProcessTaskRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostProcessTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PostProcessTaskRequest) GoString() string {
	return s.String()
}

func (s *PostProcessTaskRequest) GetBody() *string {
	return s.Body
}

func (s *PostProcessTaskRequest) SetBody(v string) *PostProcessTaskRequest {
	s.Body = &v
	return s
}

func (s *PostProcessTaskRequest) Validate() error {
	return dara.Validate(s)
}
