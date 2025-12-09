// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCacheConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketCacheConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketCacheConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketCacheConfigurationResponseBody) *GetBucketCacheConfigurationResponse
	GetBody() *GetBucketCacheConfigurationResponseBody
}

type GetBucketCacheConfigurationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketCacheConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketCacheConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCacheConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketCacheConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketCacheConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketCacheConfigurationResponse) GetBody() *GetBucketCacheConfigurationResponseBody {
	return s.Body
}

func (s *GetBucketCacheConfigurationResponse) SetHeaders(v map[string]*string) *GetBucketCacheConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketCacheConfigurationResponse) SetStatusCode(v int32) *GetBucketCacheConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketCacheConfigurationResponse) SetBody(v *GetBucketCacheConfigurationResponseBody) *GetBucketCacheConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetBucketCacheConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
