// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePublicAccessBlockResponse
	GetStatusCode() *int32
}

type DeletePublicAccessBlockResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeletePublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePublicAccessBlockResponse) SetHeaders(v map[string]*string) *DeletePublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicAccessBlockResponse) SetStatusCode(v int32) *DeletePublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicAccessBlockResponse) Validate() error {
	return dara.Validate(s)
}
