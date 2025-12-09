// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketPolicyResponse
	GetStatusCode() *int32
}

type DeleteBucketPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketPolicyResponse) SetHeaders(v map[string]*string) *DeleteBucketPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketPolicyResponse) SetStatusCode(v int32) *DeleteBucketPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketPolicyResponse) Validate() error {
	return dara.Validate(s)
}
