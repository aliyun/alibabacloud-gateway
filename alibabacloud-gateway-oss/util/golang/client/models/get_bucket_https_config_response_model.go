// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketHttpsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketHttpsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketHttpsConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketHttpsConfigResponseBody) *GetBucketHttpsConfigResponse
	GetBody() *GetBucketHttpsConfigResponseBody
}

type GetBucketHttpsConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketHttpsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketHttpsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketHttpsConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBucketHttpsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketHttpsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketHttpsConfigResponse) GetBody() *GetBucketHttpsConfigResponseBody {
	return s.Body
}

func (s *GetBucketHttpsConfigResponse) SetHeaders(v map[string]*string) *GetBucketHttpsConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBucketHttpsConfigResponse) SetStatusCode(v int32) *GetBucketHttpsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketHttpsConfigResponse) SetBody(v *GetBucketHttpsConfigResponseBody) *GetBucketHttpsConfigResponse {
	s.Body = v
	return s
}

func (s *GetBucketHttpsConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
