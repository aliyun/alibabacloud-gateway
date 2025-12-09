// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketResponse
	GetStatusCode() *int32
}

type DeleteBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketResponse) SetHeaders(v map[string]*string) *DeleteBucketResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketResponse) SetStatusCode(v int32) *DeleteBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketResponse) Validate() error {
	return dara.Validate(s)
}
