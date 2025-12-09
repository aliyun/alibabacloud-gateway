// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReservedCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssReservedCapacityId(v string) *DeleteReservedCapacityRequest
	GetXOssReservedCapacityId() *string
}

type DeleteReservedCapacityRequest struct {
	// This parameter is required.
	XOssReservedCapacityId *string `json:"x-oss-reserved-capacity-id,omitempty" xml:"x-oss-reserved-capacity-id,omitempty"`
}

func (s DeleteReservedCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReservedCapacityRequest) GoString() string {
	return s.String()
}

func (s *DeleteReservedCapacityRequest) GetXOssReservedCapacityId() *string {
	return s.XOssReservedCapacityId
}

func (s *DeleteReservedCapacityRequest) SetXOssReservedCapacityId(v string) *DeleteReservedCapacityRequest {
	s.XOssReservedCapacityId = &v
	return s
}

func (s *DeleteReservedCapacityRequest) Validate() error {
	return dara.Validate(s)
}
