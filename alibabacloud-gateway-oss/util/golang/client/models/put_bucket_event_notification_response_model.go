// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketEventNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketEventNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketEventNotificationResponse
	GetStatusCode() *int32
}

type PutBucketEventNotificationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketEventNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketEventNotificationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketEventNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketEventNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketEventNotificationResponse) SetHeaders(v map[string]*string) *PutBucketEventNotificationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketEventNotificationResponse) SetStatusCode(v int32) *PutBucketEventNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketEventNotificationResponse) Validate() error {
	return dara.Validate(s)
}
