// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourcePool(v string) *GetResourcePoolInfoRequest
	GetResourcePool() *string
}

type GetResourcePoolInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rp-01
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
}

func (s GetResourcePoolInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePoolInfoRequest) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *GetResourcePoolInfoRequest) SetResourcePool(v string) *GetResourcePoolInfoRequest {
	s.ResourcePool = &v
	return s
}

func (s *GetResourcePoolInfoRequest) Validate() error {
	return dara.Validate(s)
}
