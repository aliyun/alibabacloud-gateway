// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketCorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketCorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketCorsResponse
	GetStatusCode() *int32
}

type DeleteBucketCorsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketCorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketCorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketCorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketCorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketCorsResponse) SetHeaders(v map[string]*string) *DeleteBucketCorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketCorsResponse) SetStatusCode(v int32) *DeleteBucketCorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketCorsResponse) Validate() error {
	return dara.Validate(s)
}
