// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketResponseHeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketResponseHeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketResponseHeaderResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketResponseHeaderResponseBody) *GetBucketResponseHeaderResponse
	GetBody() *GetBucketResponseHeaderResponseBody
}

type GetBucketResponseHeaderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketResponseHeaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketResponseHeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketResponseHeaderResponse) GoString() string {
	return s.String()
}

func (s *GetBucketResponseHeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketResponseHeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketResponseHeaderResponse) GetBody() *GetBucketResponseHeaderResponseBody {
	return s.Body
}

func (s *GetBucketResponseHeaderResponse) SetHeaders(v map[string]*string) *GetBucketResponseHeaderResponse {
	s.Headers = v
	return s
}

func (s *GetBucketResponseHeaderResponse) SetStatusCode(v int32) *GetBucketResponseHeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketResponseHeaderResponse) SetBody(v *GetBucketResponseHeaderResponseBody) *GetBucketResponseHeaderResponse {
	s.Body = v
	return s
}

func (s *GetBucketResponseHeaderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
