// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketNotificationResponse
	GetStatusCode() *int32
}

type DeleteBucketNotificationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketNotificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketNotificationResponse) SetHeaders(v map[string]*string) *DeleteBucketNotificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketNotificationResponse) SetStatusCode(v int32) *DeleteBucketNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketNotificationResponse) Validate() error {
	return dara.Validate(s)
}
