// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketDataAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeCacheAvailableZone(v string) *DeleteBucketDataAcceleratorRequest
	GetXOssDatalakeCacheAvailableZone() *string
}

type DeleteBucketDataAcceleratorRequest struct {
	// example:
	//
	// cn-hangzhou-a
	XOssDatalakeCacheAvailableZone *string `json:"x-oss-datalake-cache-available-zone,omitempty" xml:"x-oss-datalake-cache-available-zone,omitempty"`
}

func (s DeleteBucketDataAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketDataAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketDataAcceleratorRequest) GetXOssDatalakeCacheAvailableZone() *string {
	return s.XOssDatalakeCacheAvailableZone
}

func (s *DeleteBucketDataAcceleratorRequest) SetXOssDatalakeCacheAvailableZone(v string) *DeleteBucketDataAcceleratorRequest {
	s.XOssDatalakeCacheAvailableZone = &v
	return s
}

func (s *DeleteBucketDataAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
