// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketVersioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketVersioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketVersioningResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketVersioningResponseBody) *GetBucketVersioningResponse
	GetBody() *GetBucketVersioningResponseBody
}

type GetBucketVersioningResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketVersioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketVersioningResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketVersioningResponse) GoString() string {
	return s.String()
}

func (s *GetBucketVersioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketVersioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketVersioningResponse) GetBody() *GetBucketVersioningResponseBody {
	return s.Body
}

func (s *GetBucketVersioningResponse) SetHeaders(v map[string]*string) *GetBucketVersioningResponse {
	s.Headers = v
	return s
}

func (s *GetBucketVersioningResponse) SetStatusCode(v int32) *GetBucketVersioningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketVersioningResponse) SetBody(v *GetBucketVersioningResponseBody) *GetBucketVersioningResponse {
	s.Body = v
	return s
}

func (s *GetBucketVersioningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
