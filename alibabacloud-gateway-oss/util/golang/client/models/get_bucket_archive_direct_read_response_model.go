// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketArchiveDirectReadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketArchiveDirectReadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketArchiveDirectReadResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketArchiveDirectReadResponseBody) *GetBucketArchiveDirectReadResponse
	GetBody() *GetBucketArchiveDirectReadResponseBody
}

type GetBucketArchiveDirectReadResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketArchiveDirectReadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketArchiveDirectReadResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketArchiveDirectReadResponse) GoString() string {
	return s.String()
}

func (s *GetBucketArchiveDirectReadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketArchiveDirectReadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketArchiveDirectReadResponse) GetBody() *GetBucketArchiveDirectReadResponseBody {
	return s.Body
}

func (s *GetBucketArchiveDirectReadResponse) SetHeaders(v map[string]*string) *GetBucketArchiveDirectReadResponse {
	s.Headers = v
	return s
}

func (s *GetBucketArchiveDirectReadResponse) SetStatusCode(v int32) *GetBucketArchiveDirectReadResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketArchiveDirectReadResponse) SetBody(v *GetBucketArchiveDirectReadResponseBody) *GetBucketArchiveDirectReadResponse {
	s.Body = v
	return s
}

func (s *GetBucketArchiveDirectReadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
