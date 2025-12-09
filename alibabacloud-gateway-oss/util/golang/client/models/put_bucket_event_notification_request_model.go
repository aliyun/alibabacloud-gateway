// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketEventNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationConfiguration(v *EventNotificationConfiguration) *PutBucketEventNotificationRequest
	GetNotificationConfiguration() *EventNotificationConfiguration
}

type PutBucketEventNotificationRequest struct {
	NotificationConfiguration *EventNotificationConfiguration `json:"NotificationConfiguration,omitempty" xml:"NotificationConfiguration,omitempty"`
}

func (s PutBucketEventNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketEventNotificationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketEventNotificationRequest) GetNotificationConfiguration() *EventNotificationConfiguration {
	return s.NotificationConfiguration
}

func (s *PutBucketEventNotificationRequest) SetNotificationConfiguration(v *EventNotificationConfiguration) *PutBucketEventNotificationRequest {
	s.NotificationConfiguration = v
	return s
}

func (s *PutBucketEventNotificationRequest) Validate() error {
	if s.NotificationConfiguration != nil {
		if err := s.NotificationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
