// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketEventNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketEventNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketEventNotificationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketEventNotificationResponseBody) *GetBucketEventNotificationResponse
	GetBody() *GetBucketEventNotificationResponseBody
}

type GetBucketEventNotificationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketEventNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketEventNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketEventNotificationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketEventNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketEventNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketEventNotificationResponse) GetBody() *GetBucketEventNotificationResponseBody {
	return s.Body
}

func (s *GetBucketEventNotificationResponse) SetHeaders(v map[string]*string) *GetBucketEventNotificationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketEventNotificationResponse) SetStatusCode(v int32) *GetBucketEventNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketEventNotificationResponse) SetBody(v *GetBucketEventNotificationResponseBody) *GetBucketEventNotificationResponse {
	s.Body = v
	return s
}

func (s *GetBucketEventNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
