// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketPolicyResponse
	GetStatusCode() *int32
	SetBody(v string) *GetBucketPolicyResponse
	GetBody() *string
}

type GetBucketPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetBucketPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketPolicyResponse) GetBody() *string {
	return s.Body
}

func (s *GetBucketPolicyResponse) SetHeaders(v map[string]*string) *GetBucketPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetBucketPolicyResponse) SetStatusCode(v int32) *GetBucketPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketPolicyResponse) SetBody(v string) *GetBucketPolicyResponse {
	s.Body = &v
	return s
}

func (s *GetBucketPolicyResponse) Validate() error {
	return dara.Validate(s)
}
