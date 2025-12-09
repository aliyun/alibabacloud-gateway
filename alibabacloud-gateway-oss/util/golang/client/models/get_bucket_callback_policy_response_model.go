// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCallbackPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketCallbackPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketCallbackPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketCallbackPolicyResponseBody) *GetBucketCallbackPolicyResponse
	GetBody() *GetBucketCallbackPolicyResponseBody
}

type GetBucketCallbackPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketCallbackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketCallbackPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCallbackPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetBucketCallbackPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketCallbackPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketCallbackPolicyResponse) GetBody() *GetBucketCallbackPolicyResponseBody {
	return s.Body
}

func (s *GetBucketCallbackPolicyResponse) SetHeaders(v map[string]*string) *GetBucketCallbackPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetBucketCallbackPolicyResponse) SetStatusCode(v int32) *GetBucketCallbackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketCallbackPolicyResponse) SetBody(v *GetBucketCallbackPolicyResponseBody) *GetBucketCallbackPolicyResponse {
	s.Body = v
	return s
}

func (s *GetBucketCallbackPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
