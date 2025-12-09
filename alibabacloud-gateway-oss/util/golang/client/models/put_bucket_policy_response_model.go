// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketPolicyResponse
	GetStatusCode() *int32
}

type PutBucketPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutBucketPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketPolicyResponse) SetHeaders(v map[string]*string) *PutBucketPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutBucketPolicyResponse) SetStatusCode(v int32) *PutBucketPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketPolicyResponse) Validate() error {
	return dara.Validate(s)
}
