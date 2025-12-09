// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAccessPointPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAccessPointPublicAccessBlockResponse
	GetStatusCode() *int32
}

type PutAccessPointPublicAccessBlockResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutAccessPointPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *PutAccessPointPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAccessPointPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAccessPointPublicAccessBlockResponse) SetHeaders(v map[string]*string) *PutAccessPointPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *PutAccessPointPublicAccessBlockResponse) SetStatusCode(v int32) *PutAccessPointPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAccessPointPublicAccessBlockResponse) Validate() error {
	return dara.Validate(s)
}
