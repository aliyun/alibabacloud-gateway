// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessPointPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessPointPublicAccessBlockResponse
	GetStatusCode() *int32
}

type DeleteAccessPointPublicAccessBlockResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAccessPointPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessPointPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessPointPublicAccessBlockResponse) SetHeaders(v map[string]*string) *DeleteAccessPointPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPointPublicAccessBlockResponse) SetStatusCode(v int32) *DeleteAccessPointPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPointPublicAccessBlockResponse) Validate() error {
	return dara.Validate(s)
}
