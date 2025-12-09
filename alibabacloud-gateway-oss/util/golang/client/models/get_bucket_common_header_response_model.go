// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCommonHeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketCommonHeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketCommonHeaderResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketCommonHeaderResponseBody) *GetBucketCommonHeaderResponse
	GetBody() *GetBucketCommonHeaderResponseBody
}

type GetBucketCommonHeaderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketCommonHeaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketCommonHeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCommonHeaderResponse) GoString() string {
	return s.String()
}

func (s *GetBucketCommonHeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketCommonHeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketCommonHeaderResponse) GetBody() *GetBucketCommonHeaderResponseBody {
	return s.Body
}

func (s *GetBucketCommonHeaderResponse) SetHeaders(v map[string]*string) *GetBucketCommonHeaderResponse {
	s.Headers = v
	return s
}

func (s *GetBucketCommonHeaderResponse) SetStatusCode(v int32) *GetBucketCommonHeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketCommonHeaderResponse) SetBody(v *GetBucketCommonHeaderResponseBody) *GetBucketCommonHeaderResponse {
	s.Body = v
	return s
}

func (s *GetBucketCommonHeaderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
