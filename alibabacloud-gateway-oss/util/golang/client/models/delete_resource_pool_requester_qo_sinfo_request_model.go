// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourcePoolRequesterQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQosRequester(v string) *DeleteResourcePoolRequesterQoSInfoRequest
	GetQosRequester() *string
	SetResourcePool(v string) *DeleteResourcePoolRequesterQoSInfoRequest
	GetResourcePool() *string
}

type DeleteResourcePoolRequesterQoSInfoRequest struct {
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
	// rp-01
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
}

func (s DeleteResourcePoolRequesterQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourcePoolRequesterQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourcePoolRequesterQoSInfoRequest) GetQosRequester() *string {
	return s.QosRequester
}

func (s *DeleteResourcePoolRequesterQoSInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *DeleteResourcePoolRequesterQoSInfoRequest) SetQosRequester(v string) *DeleteResourcePoolRequesterQoSInfoRequest {
	s.QosRequester = &v
	return s
}

func (s *DeleteResourcePoolRequesterQoSInfoRequest) SetResourcePool(v string) *DeleteResourcePoolRequesterQoSInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *DeleteResourcePoolRequesterQoSInfoRequest) Validate() error {
	return dara.Validate(s)
}
