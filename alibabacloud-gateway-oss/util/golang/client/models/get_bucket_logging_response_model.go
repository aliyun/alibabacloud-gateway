// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLoggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketLoggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketLoggingResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketLoggingResponseBody) *GetBucketLoggingResponse
	GetBody() *GetBucketLoggingResponseBody
}

type GetBucketLoggingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketLoggingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLoggingResponse) GoString() string {
	return s.String()
}

func (s *GetBucketLoggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketLoggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketLoggingResponse) GetBody() *GetBucketLoggingResponseBody {
	return s.Body
}

func (s *GetBucketLoggingResponse) SetHeaders(v map[string]*string) *GetBucketLoggingResponse {
	s.Headers = v
	return s
}

func (s *GetBucketLoggingResponse) SetStatusCode(v int32) *GetBucketLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketLoggingResponse) SetBody(v *GetBucketLoggingResponseBody) *GetBucketLoggingResponse {
	s.Body = v
	return s
}

func (s *GetBucketLoggingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
