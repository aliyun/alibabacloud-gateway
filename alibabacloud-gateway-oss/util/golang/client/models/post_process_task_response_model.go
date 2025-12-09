// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostProcessTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostProcessTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostProcessTaskResponse
	GetStatusCode() *int32
}

type PostProcessTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PostProcessTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PostProcessTaskResponse) GoString() string {
	return s.String()
}

func (s *PostProcessTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostProcessTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostProcessTaskResponse) SetHeaders(v map[string]*string) *PostProcessTaskResponse {
	s.Headers = v
	return s
}

func (s *PostProcessTaskResponse) SetStatusCode(v int32) *PostProcessTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PostProcessTaskResponse) Validate() error {
	return dara.Validate(s)
}
