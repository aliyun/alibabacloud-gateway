// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketPublicAccessBlockResponse
	GetStatusCode() *int32
}

type PutBucketPublicAccessBlockResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *PutBucketPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketPublicAccessBlockResponse) SetHeaders(v map[string]*string) *PutBucketPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *PutBucketPublicAccessBlockResponse) SetStatusCode(v int32) *PutBucketPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketPublicAccessBlockResponse) Validate() error {
	return dara.Validate(s)
}
