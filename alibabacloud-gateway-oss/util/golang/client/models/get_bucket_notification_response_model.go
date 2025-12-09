// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketNotificationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketNotificationResponseBody) *GetBucketNotificationResponse
	GetBody() *GetBucketNotificationResponseBody
}

type GetBucketNotificationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketNotificationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketNotificationResponse) GetBody() *GetBucketNotificationResponseBody {
	return s.Body
}

func (s *GetBucketNotificationResponse) SetHeaders(v map[string]*string) *GetBucketNotificationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketNotificationResponse) SetStatusCode(v int32) *GetBucketNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketNotificationResponse) SetBody(v *GetBucketNotificationResponseBody) *GetBucketNotificationResponse {
	s.Body = v
	return s
}

func (s *GetBucketNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
