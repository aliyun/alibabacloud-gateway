// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVectorBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVectorBucketResponse
	GetStatusCode() *int32
	SetBody(v *GetVectorBucketResponseBody) *GetVectorBucketResponse
	GetBody() *GetVectorBucketResponseBody
}

type GetVectorBucketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVectorBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVectorBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVectorBucketResponse) GoString() string {
	return s.String()
}

func (s *GetVectorBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVectorBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVectorBucketResponse) GetBody() *GetVectorBucketResponseBody {
	return s.Body
}

func (s *GetVectorBucketResponse) SetHeaders(v map[string]*string) *GetVectorBucketResponse {
	s.Headers = v
	return s
}

func (s *GetVectorBucketResponse) SetStatusCode(v int32) *GetVectorBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVectorBucketResponse) SetBody(v *GetVectorBucketResponseBody) *GetVectorBucketResponse {
	s.Body = v
	return s
}

func (s *GetVectorBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
