// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketLocationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketLocationResponseBody) *GetBucketLocationResponse
	GetBody() *GetBucketLocationResponseBody
}

type GetBucketLocationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLocationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketLocationResponse) GetBody() *GetBucketLocationResponseBody {
	return s.Body
}

func (s *GetBucketLocationResponse) SetHeaders(v map[string]*string) *GetBucketLocationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketLocationResponse) SetStatusCode(v int32) *GetBucketLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketLocationResponse) SetBody(v *GetBucketLocationResponseBody) *GetBucketLocationResponse {
	s.Body = v
	return s
}

func (s *GetBucketLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
