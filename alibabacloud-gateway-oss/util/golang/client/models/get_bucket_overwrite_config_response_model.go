// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketOverwriteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketOverwriteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketOverwriteConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketOverwriteConfigResponseBody) *GetBucketOverwriteConfigResponse
	GetBody() *GetBucketOverwriteConfigResponseBody
}

type GetBucketOverwriteConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketOverwriteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketOverwriteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketOverwriteConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBucketOverwriteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketOverwriteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketOverwriteConfigResponse) GetBody() *GetBucketOverwriteConfigResponseBody {
	return s.Body
}

func (s *GetBucketOverwriteConfigResponse) SetHeaders(v map[string]*string) *GetBucketOverwriteConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBucketOverwriteConfigResponse) SetStatusCode(v int32) *GetBucketOverwriteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketOverwriteConfigResponse) SetBody(v *GetBucketOverwriteConfigResponseBody) *GetBucketOverwriteConfigResponse {
	s.Body = v
	return s
}

func (s *GetBucketOverwriteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
