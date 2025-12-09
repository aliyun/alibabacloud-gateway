// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketEventNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationConfiguration(v *EventNotificationConfiguration) *GetBucketEventNotificationResponseBody
	GetNotificationConfiguration() *EventNotificationConfiguration
}

type GetBucketEventNotificationResponseBody struct {
	// Function Compute service configuration for a bucket.
	NotificationConfiguration *EventNotificationConfiguration `json:"NotificationConfiguration,omitempty" xml:"NotificationConfiguration,omitempty"`
}

func (s GetBucketEventNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketEventNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketEventNotificationResponseBody) GetNotificationConfiguration() *EventNotificationConfiguration {
	return s.NotificationConfiguration
}

func (s *GetBucketEventNotificationResponseBody) SetNotificationConfiguration(v *EventNotificationConfiguration) *GetBucketEventNotificationResponseBody {
	s.NotificationConfiguration = v
	return s
}

func (s *GetBucketEventNotificationResponseBody) Validate() error {
	if s.NotificationConfiguration != nil {
		if err := s.NotificationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
