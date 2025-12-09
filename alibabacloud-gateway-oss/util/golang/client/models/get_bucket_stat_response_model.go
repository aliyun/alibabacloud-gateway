// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketStatResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketStatResponseBody) *GetBucketStatResponse
	GetBody() *GetBucketStatResponseBody
}

type GetBucketStatResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketStatResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketStatResponse) GoString() string {
	return s.String()
}

func (s *GetBucketStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketStatResponse) GetBody() *GetBucketStatResponseBody {
	return s.Body
}

func (s *GetBucketStatResponse) SetHeaders(v map[string]*string) *GetBucketStatResponse {
	s.Headers = v
	return s
}

func (s *GetBucketStatResponse) SetStatusCode(v int32) *GetBucketStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketStatResponse) SetBody(v *GetBucketStatResponseBody) *GetBucketStatResponse {
	s.Body = v
	return s
}

func (s *GetBucketStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
