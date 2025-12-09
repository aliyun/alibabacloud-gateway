// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateEndpoint(v string) *RegionInfo
	GetAccelerateEndpoint() *string
	SetInternalEndpoint(v string) *RegionInfo
	GetInternalEndpoint() *string
	SetInternetEndpoint(v string) *RegionInfo
	GetInternetEndpoint() *string
	SetRegion(v string) *RegionInfo
	GetRegion() *string
}

type RegionInfo struct {
	AccelerateEndpoint *string `json:"AccelerateEndpoint,omitempty" xml:"AccelerateEndpoint,omitempty"`
	InternalEndpoint   *string `json:"InternalEndpoint,omitempty" xml:"InternalEndpoint,omitempty"`
	InternetEndpoint   *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RegionInfo) String() string {
	return dara.Prettify(s)
}

func (s RegionInfo) GoString() string {
	return s.String()
}

func (s *RegionInfo) GetAccelerateEndpoint() *string {
	return s.AccelerateEndpoint
}

func (s *RegionInfo) GetInternalEndpoint() *string {
	return s.InternalEndpoint
}

func (s *RegionInfo) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *RegionInfo) GetRegion() *string {
	return s.Region
}

func (s *RegionInfo) SetAccelerateEndpoint(v string) *RegionInfo {
	s.AccelerateEndpoint = &v
	return s
}

func (s *RegionInfo) SetInternalEndpoint(v string) *RegionInfo {
	s.InternalEndpoint = &v
	return s
}

func (s *RegionInfo) SetInternetEndpoint(v string) *RegionInfo {
	s.InternetEndpoint = &v
	return s
}

func (s *RegionInfo) SetRegion(v string) *RegionInfo {
	s.Region = &v
	return s
}

func (s *RegionInfo) Validate() error {
	return dara.Validate(s)
}
