// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketTagsResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketTagsResponseBody) *GetBucketTagsResponse
	GetBody() *GetBucketTagsResponseBody
}

type GetBucketTagsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketTagsResponse) GoString() string {
	return s.String()
}

func (s *GetBucketTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketTagsResponse) GetBody() *GetBucketTagsResponseBody {
	return s.Body
}

func (s *GetBucketTagsResponse) SetHeaders(v map[string]*string) *GetBucketTagsResponse {
	s.Headers = v
	return s
}

func (s *GetBucketTagsResponse) SetStatusCode(v int32) *GetBucketTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketTagsResponse) SetBody(v *GetBucketTagsResponseBody) *GetBucketTagsResponse {
	s.Body = v
	return s
}

func (s *GetBucketTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
