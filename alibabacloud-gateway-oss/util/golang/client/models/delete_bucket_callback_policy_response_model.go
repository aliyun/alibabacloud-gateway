// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketCallbackPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketCallbackPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketCallbackPolicyResponse
	GetStatusCode() *int32
}

type DeleteBucketCallbackPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketCallbackPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketCallbackPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketCallbackPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketCallbackPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketCallbackPolicyResponse) SetHeaders(v map[string]*string) *DeleteBucketCallbackPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketCallbackPolicyResponse) SetStatusCode(v int32) *DeleteBucketCallbackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketCallbackPolicyResponse) Validate() error {
	return dara.Validate(s)
}
