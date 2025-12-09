// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirtualBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVirtualBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVirtualBucketResponse
	GetStatusCode() *int32
	SetBody(v *GetVirtualBucketResponseBody) *GetVirtualBucketResponse
	GetBody() *GetVirtualBucketResponseBody
}

type GetVirtualBucketResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVirtualBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVirtualBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVirtualBucketResponse) GoString() string {
	return s.String()
}

func (s *GetVirtualBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVirtualBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVirtualBucketResponse) GetBody() *GetVirtualBucketResponseBody {
	return s.Body
}

func (s *GetVirtualBucketResponse) SetHeaders(v map[string]*string) *GetVirtualBucketResponse {
	s.Headers = v
	return s
}

func (s *GetVirtualBucketResponse) SetStatusCode(v int32) *GetVirtualBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVirtualBucketResponse) SetBody(v *GetVirtualBucketResponseBody) *GetVirtualBucketResponse {
	s.Body = v
	return s
}

func (s *GetVirtualBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
