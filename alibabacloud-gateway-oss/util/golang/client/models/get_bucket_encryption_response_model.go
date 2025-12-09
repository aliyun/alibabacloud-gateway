// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketEncryptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketEncryptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketEncryptionResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketEncryptionResponseBody) *GetBucketEncryptionResponse
	GetBody() *GetBucketEncryptionResponseBody
}

type GetBucketEncryptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketEncryptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketEncryptionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketEncryptionResponse) GoString() string {
	return s.String()
}

func (s *GetBucketEncryptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketEncryptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketEncryptionResponse) GetBody() *GetBucketEncryptionResponseBody {
	return s.Body
}

func (s *GetBucketEncryptionResponse) SetHeaders(v map[string]*string) *GetBucketEncryptionResponse {
	s.Headers = v
	return s
}

func (s *GetBucketEncryptionResponse) SetStatusCode(v int32) *GetBucketEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketEncryptionResponse) SetBody(v *GetBucketEncryptionResponseBody) *GetBucketEncryptionResponse {
	s.Body = v
	return s
}

func (s *GetBucketEncryptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
