// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketObjectWormConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketObjectWormConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketObjectWormConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketObjectWormConfigurationResponseBody) *GetBucketObjectWormConfigurationResponse
	GetBody() *GetBucketObjectWormConfigurationResponseBody
}

type GetBucketObjectWormConfigurationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketObjectWormConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketObjectWormConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketObjectWormConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketObjectWormConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketObjectWormConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketObjectWormConfigurationResponse) GetBody() *GetBucketObjectWormConfigurationResponseBody {
	return s.Body
}

func (s *GetBucketObjectWormConfigurationResponse) SetHeaders(v map[string]*string) *GetBucketObjectWormConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketObjectWormConfigurationResponse) SetStatusCode(v int32) *GetBucketObjectWormConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketObjectWormConfigurationResponse) SetBody(v *GetBucketObjectWormConfigurationResponseBody) *GetBucketObjectWormConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetBucketObjectWormConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
