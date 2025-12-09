// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationConfiguration(v *NotificationConfiguration) *GetBucketNotificationResponseBody
	GetNotificationConfiguration() *NotificationConfiguration
}

type GetBucketNotificationResponseBody struct {
	NotificationConfiguration *NotificationConfiguration `json:"NotificationConfiguration,omitempty" xml:"NotificationConfiguration,omitempty"`
}

func (s GetBucketNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketNotificationResponseBody) GetNotificationConfiguration() *NotificationConfiguration {
	return s.NotificationConfiguration
}

func (s *GetBucketNotificationResponseBody) SetNotificationConfiguration(v *NotificationConfiguration) *GetBucketNotificationResponseBody {
	s.NotificationConfiguration = v
	return s
}

func (s *GetBucketNotificationResponseBody) Validate() error {
	if s.NotificationConfiguration != nil {
		if err := s.NotificationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
