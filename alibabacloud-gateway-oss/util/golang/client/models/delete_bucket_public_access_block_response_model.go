// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketPublicAccessBlockResponse
	GetStatusCode() *int32
}

type DeleteBucketPublicAccessBlockResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketPublicAccessBlockResponse) SetHeaders(v map[string]*string) *DeleteBucketPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketPublicAccessBlockResponse) SetStatusCode(v int32) *DeleteBucketPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketPublicAccessBlockResponse) Validate() error {
	return dara.Validate(s)
}
