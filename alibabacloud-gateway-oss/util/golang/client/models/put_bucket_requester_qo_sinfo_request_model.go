// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRequesterQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQoSConfiguration(v *QoSConfiguration) *PutBucketRequesterQoSInfoRequest
	GetQoSConfiguration() *QoSConfiguration
	SetQosRequester(v string) *PutBucketRequesterQoSInfoRequest
	GetQosRequester() *string
}

type PutBucketRequesterQoSInfoRequest struct {
	QoSConfiguration *QoSConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26753xxxxxxxx14340
	QosRequester *string `json:"qosRequester,omitempty" xml:"qosRequester,omitempty"`
}

func (s PutBucketRequesterQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRequesterQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequesterQoSInfoRequest) GetQoSConfiguration() *QoSConfiguration {
	return s.QoSConfiguration
}

func (s *PutBucketRequesterQoSInfoRequest) GetQosRequester() *string {
	return s.QosRequester
}

func (s *PutBucketRequesterQoSInfoRequest) SetQoSConfiguration(v *QoSConfiguration) *PutBucketRequesterQoSInfoRequest {
	s.QoSConfiguration = v
	return s
}

func (s *PutBucketRequesterQoSInfoRequest) SetQosRequester(v string) *PutBucketRequesterQoSInfoRequest {
	s.QosRequester = &v
	return s
}

func (s *PutBucketRequesterQoSInfoRequest) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
