// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourcePoolRequesterQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQoSConfiguration(v *QoSConfiguration) *PutResourcePoolRequesterQoSInfoRequest
	GetQoSConfiguration() *QoSConfiguration
	SetQosRequester(v string) *PutResourcePoolRequesterQoSInfoRequest
	GetQosRequester() *string
	SetResourcePool(v string) *PutResourcePoolRequesterQoSInfoRequest
	GetResourcePool() *string
}

type PutResourcePoolRequesterQoSInfoRequest struct {
	QoSConfiguration *QoSConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26753xxxxxxxx14340
	QosRequester *string `json:"qosRequester,omitempty" xml:"qosRequester,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rp-test
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
}

func (s PutResourcePoolRequesterQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourcePoolRequesterQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *PutResourcePoolRequesterQoSInfoRequest) GetQoSConfiguration() *QoSConfiguration {
	return s.QoSConfiguration
}

func (s *PutResourcePoolRequesterQoSInfoRequest) GetQosRequester() *string {
	return s.QosRequester
}

func (s *PutResourcePoolRequesterQoSInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *PutResourcePoolRequesterQoSInfoRequest) SetQoSConfiguration(v *QoSConfiguration) *PutResourcePoolRequesterQoSInfoRequest {
	s.QoSConfiguration = v
	return s
}

func (s *PutResourcePoolRequesterQoSInfoRequest) SetQosRequester(v string) *PutResourcePoolRequesterQoSInfoRequest {
	s.QosRequester = &v
	return s
}

func (s *PutResourcePoolRequesterQoSInfoRequest) SetResourcePool(v string) *PutResourcePoolRequesterQoSInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *PutResourcePoolRequesterQoSInfoRequest) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
