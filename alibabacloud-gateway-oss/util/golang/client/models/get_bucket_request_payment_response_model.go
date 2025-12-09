// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRequestPaymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketRequestPaymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketRequestPaymentResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketRequestPaymentResponseBody) *GetBucketRequestPaymentResponse
	GetBody() *GetBucketRequestPaymentResponseBody
}

type GetBucketRequestPaymentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketRequestPaymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketRequestPaymentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRequestPaymentResponse) GoString() string {
	return s.String()
}

func (s *GetBucketRequestPaymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketRequestPaymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketRequestPaymentResponse) GetBody() *GetBucketRequestPaymentResponseBody {
	return s.Body
}

func (s *GetBucketRequestPaymentResponse) SetHeaders(v map[string]*string) *GetBucketRequestPaymentResponse {
	s.Headers = v
	return s
}

func (s *GetBucketRequestPaymentResponse) SetStatusCode(v int32) *GetBucketRequestPaymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketRequestPaymentResponse) SetBody(v *GetBucketRequestPaymentResponseBody) *GetBucketRequestPaymentResponse {
	s.Body = v
	return s
}

func (s *GetBucketRequestPaymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
