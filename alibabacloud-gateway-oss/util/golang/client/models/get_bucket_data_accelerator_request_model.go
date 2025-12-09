// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDataAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v string) *GetBucketDataAcceleratorRequest
	GetVerbose() *string
	SetXOssDatalakeCacheAvailableZone(v string) *GetBucketDataAcceleratorRequest
	GetXOssDatalakeCacheAvailableZone() *string
}

type GetBucketDataAcceleratorRequest struct {
	// example:
	//
	// true
	Verbose *string `json:"verbose,omitempty" xml:"verbose,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	XOssDatalakeCacheAvailableZone *string `json:"x-oss-datalake-cache-available-zone,omitempty" xml:"x-oss-datalake-cache-available-zone,omitempty"`
}

func (s GetBucketDataAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDataAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *GetBucketDataAcceleratorRequest) GetVerbose() *string {
	return s.Verbose
}

func (s *GetBucketDataAcceleratorRequest) GetXOssDatalakeCacheAvailableZone() *string {
	return s.XOssDatalakeCacheAvailableZone
}

func (s *GetBucketDataAcceleratorRequest) SetVerbose(v string) *GetBucketDataAcceleratorRequest {
	s.Verbose = &v
	return s
}

func (s *GetBucketDataAcceleratorRequest) SetXOssDatalakeCacheAvailableZone(v string) *GetBucketDataAcceleratorRequest {
	s.XOssDatalakeCacheAvailableZone = &v
	return s
}

func (s *GetBucketDataAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
