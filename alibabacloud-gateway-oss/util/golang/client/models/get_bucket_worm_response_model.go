// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketWormResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketWormResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketWormResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketWormResponseBody) *GetBucketWormResponse
	GetBody() *GetBucketWormResponseBody
}

type GetBucketWormResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketWormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketWormResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketWormResponse) GoString() string {
	return s.String()
}

func (s *GetBucketWormResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketWormResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketWormResponse) GetBody() *GetBucketWormResponseBody {
	return s.Body
}

func (s *GetBucketWormResponse) SetHeaders(v map[string]*string) *GetBucketWormResponse {
	s.Headers = v
	return s
}

func (s *GetBucketWormResponse) SetStatusCode(v int32) *GetBucketWormResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketWormResponse) SetBody(v *GetBucketWormResponseBody) *GetBucketWormResponse {
	s.Body = v
	return s
}

func (s *GetBucketWormResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
