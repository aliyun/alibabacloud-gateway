// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutPublicAccessBlockResponse
	GetStatusCode() *int32
}

type PutPublicAccessBlockResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s PutPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *PutPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutPublicAccessBlockResponse) SetHeaders(v map[string]*string) *PutPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *PutPublicAccessBlockResponse) SetStatusCode(v int32) *PutPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *PutPublicAccessBlockResponse) Validate() error {
	return dara.Validate(s)
}
