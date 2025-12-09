// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketNotificationResponse
	GetStatusCode() *int32
}

type PutBucketNotificationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketNotificationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketNotificationResponse) SetHeaders(v map[string]*string) *PutBucketNotificationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketNotificationResponse) SetStatusCode(v int32) *PutBucketNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketNotificationResponse) Validate() error {
	return dara.Validate(s)
}
