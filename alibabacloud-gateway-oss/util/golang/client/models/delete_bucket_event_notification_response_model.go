// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketEventNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketEventNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketEventNotificationResponse
	GetStatusCode() *int32
}

type DeleteBucketEventNotificationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketEventNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketEventNotificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketEventNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketEventNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketEventNotificationResponse) SetHeaders(v map[string]*string) *DeleteBucketEventNotificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketEventNotificationResponse) SetStatusCode(v int32) *DeleteBucketEventNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketEventNotificationResponse) Validate() error {
	return dara.Validate(s)
}
