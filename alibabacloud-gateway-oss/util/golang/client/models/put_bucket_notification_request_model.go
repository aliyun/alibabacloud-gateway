// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationConfiguration(v *NotificationConfiguration) *PutBucketNotificationRequest
	GetNotificationConfiguration() *NotificationConfiguration
}

type PutBucketNotificationRequest struct {
	NotificationConfiguration *NotificationConfiguration `json:"NotificationConfiguration,omitempty" xml:"NotificationConfiguration,omitempty"`
}

func (s PutBucketNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketNotificationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketNotificationRequest) GetNotificationConfiguration() *NotificationConfiguration {
	return s.NotificationConfiguration
}

func (s *PutBucketNotificationRequest) SetNotificationConfiguration(v *NotificationConfiguration) *PutBucketNotificationRequest {
	s.NotificationConfiguration = v
	return s
}

func (s *PutBucketNotificationRequest) Validate() error {
	if s.NotificationConfiguration != nil {
		if err := s.NotificationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
