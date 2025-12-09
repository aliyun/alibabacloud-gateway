// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketCorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketCorsResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketCorsResponseBody) *GetBucketCorsResponse
	GetBody() *GetBucketCorsResponseBody
}

type GetBucketCorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketCorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketCorsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCorsResponse) GoString() string {
	return s.String()
}

func (s *GetBucketCorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketCorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketCorsResponse) GetBody() *GetBucketCorsResponseBody {
	return s.Body
}

func (s *GetBucketCorsResponse) SetHeaders(v map[string]*string) *GetBucketCorsResponse {
	s.Headers = v
	return s
}

func (s *GetBucketCorsResponse) SetStatusCode(v int32) *GetBucketCorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketCorsResponse) SetBody(v *GetBucketCorsResponseBody) *GetBucketCorsResponse {
	s.Body = v
	return s
}

func (s *GetBucketCorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
