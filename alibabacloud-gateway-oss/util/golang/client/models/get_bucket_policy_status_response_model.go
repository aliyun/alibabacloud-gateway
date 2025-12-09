// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketPolicyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketPolicyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketPolicyStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketPolicyStatusResponseBody) *GetBucketPolicyStatusResponse
	GetBody() *GetBucketPolicyStatusResponseBody
}

type GetBucketPolicyStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketPolicyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketPolicyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketPolicyStatusResponse) GoString() string {
	return s.String()
}

func (s *GetBucketPolicyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketPolicyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketPolicyStatusResponse) GetBody() *GetBucketPolicyStatusResponseBody {
	return s.Body
}

func (s *GetBucketPolicyStatusResponse) SetHeaders(v map[string]*string) *GetBucketPolicyStatusResponse {
	s.Headers = v
	return s
}

func (s *GetBucketPolicyStatusResponse) SetStatusCode(v int32) *GetBucketPolicyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketPolicyStatusResponse) SetBody(v *GetBucketPolicyStatusResponseBody) *GetBucketPolicyStatusResponse {
	s.Body = v
	return s
}

func (s *GetBucketPolicyStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
