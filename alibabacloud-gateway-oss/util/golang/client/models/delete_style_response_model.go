// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStyleResponse
	GetStatusCode() *int32
}

type DeleteStyleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStyleResponse) GoString() string {
	return s.String()
}

func (s *DeleteStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStyleResponse) SetHeaders(v map[string]*string) *DeleteStyleResponse {
	s.Headers = v
	return s
}

func (s *DeleteStyleResponse) SetStatusCode(v int32) *DeleteStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStyleResponse) Validate() error {
	return dara.Validate(s)
}
