// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReservedCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssReservedCapacityId(v string) *GetReservedCapacityRequest
	GetXOssReservedCapacityId() *string
}

type GetReservedCapacityRequest struct {
	// This parameter is required.
	XOssReservedCapacityId *string `json:"x-oss-reserved-capacity-id,omitempty" xml:"x-oss-reserved-capacity-id,omitempty"`
}

func (s GetReservedCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReservedCapacityRequest) GoString() string {
	return s.String()
}

func (s *GetReservedCapacityRequest) GetXOssReservedCapacityId() *string {
	return s.XOssReservedCapacityId
}

func (s *GetReservedCapacityRequest) SetXOssReservedCapacityId(v string) *GetReservedCapacityRequest {
	s.XOssReservedCapacityId = &v
	return s
}

func (s *GetReservedCapacityRequest) Validate() error {
	return dara.Validate(s)
}
