// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQoSConfiguration(v *QoSConfiguration) *PutBucketQoSInfoRequest
	GetQoSConfiguration() *QoSConfiguration
}

type PutBucketQoSInfoRequest struct {
	QoSConfiguration *QoSConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
}

func (s PutBucketQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *PutBucketQoSInfoRequest) GetQoSConfiguration() *QoSConfiguration {
	return s.QoSConfiguration
}

func (s *PutBucketQoSInfoRequest) SetQoSConfiguration(v *QoSConfiguration) *PutBucketQoSInfoRequest {
	s.QoSConfiguration = v
	return s
}

func (s *PutBucketQoSInfoRequest) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
