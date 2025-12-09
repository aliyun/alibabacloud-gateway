// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketWebsiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketWebsiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketWebsiteResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketWebsiteResponseBody) *GetBucketWebsiteResponse
	GetBody() *GetBucketWebsiteResponseBody
}

type GetBucketWebsiteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketWebsiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketWebsiteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketWebsiteResponse) GoString() string {
	return s.String()
}

func (s *GetBucketWebsiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketWebsiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketWebsiteResponse) GetBody() *GetBucketWebsiteResponseBody {
	return s.Body
}

func (s *GetBucketWebsiteResponse) SetHeaders(v map[string]*string) *GetBucketWebsiteResponse {
	s.Headers = v
	return s
}

func (s *GetBucketWebsiteResponse) SetStatusCode(v int32) *GetBucketWebsiteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketWebsiteResponse) SetBody(v *GetBucketWebsiteResponseBody) *GetBucketWebsiteResponse {
	s.Body = v
	return s
}

func (s *GetBucketWebsiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
