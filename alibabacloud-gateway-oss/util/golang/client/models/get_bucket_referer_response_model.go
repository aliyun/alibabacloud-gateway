// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRefererResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketRefererResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketRefererResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketRefererResponseBody) *GetBucketRefererResponse
	GetBody() *GetBucketRefererResponseBody
}

type GetBucketRefererResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketRefererResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketRefererResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRefererResponse) GoString() string {
	return s.String()
}

func (s *GetBucketRefererResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketRefererResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketRefererResponse) GetBody() *GetBucketRefererResponseBody {
	return s.Body
}

func (s *GetBucketRefererResponse) SetHeaders(v map[string]*string) *GetBucketRefererResponse {
	s.Headers = v
	return s
}

func (s *GetBucketRefererResponse) SetStatusCode(v int32) *GetBucketRefererResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketRefererResponse) SetBody(v *GetBucketRefererResponseBody) *GetBucketRefererResponse {
	s.Body = v
	return s
}

func (s *GetBucketRefererResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
