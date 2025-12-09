// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketHashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketHashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketHashResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketHashResponseBody) *GetBucketHashResponse
	GetBody() *GetBucketHashResponseBody
}

type GetBucketHashResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketHashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketHashResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketHashResponse) GoString() string {
	return s.String()
}

func (s *GetBucketHashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketHashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketHashResponse) GetBody() *GetBucketHashResponseBody {
	return s.Body
}

func (s *GetBucketHashResponse) SetHeaders(v map[string]*string) *GetBucketHashResponse {
	s.Headers = v
	return s
}

func (s *GetBucketHashResponse) SetStatusCode(v int32) *GetBucketHashResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketHashResponse) SetBody(v *GetBucketHashResponseBody) *GetBucketHashResponse {
	s.Body = v
	return s
}

func (s *GetBucketHashResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
