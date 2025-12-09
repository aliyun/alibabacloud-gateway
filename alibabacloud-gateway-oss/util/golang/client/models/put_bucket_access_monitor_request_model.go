// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketAccessMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMonitorConfiguration(v *AccessMonitorConfiguration) *PutBucketAccessMonitorRequest
	GetAccessMonitorConfiguration() *AccessMonitorConfiguration
}

type PutBucketAccessMonitorRequest struct {
	// The access tracking configurations of the bucket.
	AccessMonitorConfiguration *AccessMonitorConfiguration `json:"AccessMonitorConfiguration,omitempty" xml:"AccessMonitorConfiguration,omitempty"`
}

func (s PutBucketAccessMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketAccessMonitorRequest) GoString() string {
	return s.String()
}

func (s *PutBucketAccessMonitorRequest) GetAccessMonitorConfiguration() *AccessMonitorConfiguration {
	return s.AccessMonitorConfiguration
}

func (s *PutBucketAccessMonitorRequest) SetAccessMonitorConfiguration(v *AccessMonitorConfiguration) *PutBucketAccessMonitorRequest {
	s.AccessMonitorConfiguration = v
	return s
}

func (s *PutBucketAccessMonitorRequest) Validate() error {
	if s.AccessMonitorConfiguration != nil {
		if err := s.AccessMonitorConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
