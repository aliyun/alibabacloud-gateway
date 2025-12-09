// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCallbackPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketCallbackPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketCallbackPolicyResponse
	GetStatusCode() *int32
}

type PutBucketCallbackPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketCallbackPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCallbackPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutBucketCallbackPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketCallbackPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketCallbackPolicyResponse) SetHeaders(v map[string]*string) *PutBucketCallbackPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutBucketCallbackPolicyResponse) SetStatusCode(v int32) *PutBucketCallbackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketCallbackPolicyResponse) Validate() error {
	return dara.Validate(s)
}
