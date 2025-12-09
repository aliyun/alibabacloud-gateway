// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketQoSInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketQoSInfoResponseBody) *GetBucketQoSInfoResponse
	GetBody() *GetBucketQoSInfoResponseBody
}

type GetBucketQoSInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketQoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBucketQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketQoSInfoResponse) GetBody() *GetBucketQoSInfoResponseBody {
	return s.Body
}

func (s *GetBucketQoSInfoResponse) SetHeaders(v map[string]*string) *GetBucketQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBucketQoSInfoResponse) SetStatusCode(v int32) *GetBucketQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketQoSInfoResponse) SetBody(v *GetBucketQoSInfoResponseBody) *GetBucketQoSInfoResponse {
	s.Body = v
	return s
}

func (s *GetBucketQoSInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
