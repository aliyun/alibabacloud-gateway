// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketAccessMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMonitorConfiguration(v *AccessMonitorConfiguration) *GetBucketAccessMonitorResponseBody
	GetAccessMonitorConfiguration() *AccessMonitorConfiguration
}

type GetBucketAccessMonitorResponseBody struct {
	// The container that stores access monitor configuration.
	AccessMonitorConfiguration *AccessMonitorConfiguration `json:"AccessMonitorConfiguration,omitempty" xml:"AccessMonitorConfiguration,omitempty"`
}

func (s GetBucketAccessMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAccessMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketAccessMonitorResponseBody) GetAccessMonitorConfiguration() *AccessMonitorConfiguration {
	return s.AccessMonitorConfiguration
}

func (s *GetBucketAccessMonitorResponseBody) SetAccessMonitorConfiguration(v *AccessMonitorConfiguration) *GetBucketAccessMonitorResponseBody {
	s.AccessMonitorConfiguration = v
	return s
}

func (s *GetBucketAccessMonitorResponseBody) Validate() error {
	if s.AccessMonitorConfiguration != nil {
		if err := s.AccessMonitorConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
