// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotificationConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetTopicConfiguration(v []*NotificationConfigurationTopicConfiguration) *NotificationConfiguration
	GetTopicConfiguration() []*NotificationConfigurationTopicConfiguration
}

type NotificationConfiguration struct {
	TopicConfiguration []*NotificationConfigurationTopicConfiguration `json:"TopicConfiguration,omitempty" xml:"TopicConfiguration,omitempty" type:"Repeated"`
}

func (s NotificationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s NotificationConfiguration) GoString() string {
	return s.String()
}

func (s *NotificationConfiguration) GetTopicConfiguration() []*NotificationConfigurationTopicConfiguration {
	return s.TopicConfiguration
}

func (s *NotificationConfiguration) SetTopicConfiguration(v []*NotificationConfigurationTopicConfiguration) *NotificationConfiguration {
	s.TopicConfiguration = v
	return s
}

func (s *NotificationConfiguration) Validate() error {
	if s.TopicConfiguration != nil {
		for _, item := range s.TopicConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NotificationConfigurationTopicConfiguration struct {
	// example:
	//
	// test-topic-1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s NotificationConfigurationTopicConfiguration) String() string {
	return dara.Prettify(s)
}

func (s NotificationConfigurationTopicConfiguration) GoString() string {
	return s.String()
}

func (s *NotificationConfigurationTopicConfiguration) GetId() *string {
	return s.Id
}

func (s *NotificationConfigurationTopicConfiguration) SetId(v string) *NotificationConfigurationTopicConfiguration {
	s.Id = &v
	return s
}

func (s *NotificationConfigurationTopicConfiguration) Validate() error {
	return dara.Validate(s)
}
