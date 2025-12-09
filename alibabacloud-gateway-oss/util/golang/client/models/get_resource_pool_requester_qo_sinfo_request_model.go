// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolRequesterQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQosRequester(v string) *GetResourcePoolRequesterQoSInfoRequest
	GetQosRequester() *string
	SetResourcePool(v string) *GetResourcePoolRequesterQoSInfoRequest
	GetResourcePool() *string
}

type GetResourcePoolRequesterQoSInfoRequest struct {
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

func (s GetResourcePoolRequesterQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolRequesterQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePoolRequesterQoSInfoRequest) GetQosRequester() *string {
	return s.QosRequester
}

func (s *GetResourcePoolRequesterQoSInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *GetResourcePoolRequesterQoSInfoRequest) SetQosRequester(v string) *GetResourcePoolRequesterQoSInfoRequest {
	s.QosRequester = &v
	return s
}

func (s *GetResourcePoolRequesterQoSInfoRequest) SetResourcePool(v string) *GetResourcePoolRequesterQoSInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *GetResourcePoolRequesterQoSInfoRequest) Validate() error {
	return dara.Validate(s)
}
