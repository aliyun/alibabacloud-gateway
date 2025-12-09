// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReservedCapacityCreateConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *ReservedCapacityCreateConfiguration
	GetCapacity() *int64
	SetDataRedundancyType(v string) *ReservedCapacityCreateConfiguration
	GetDataRedundancyType() *string
	SetName(v string) *ReservedCapacityCreateConfiguration
	GetName() *string
	SetRegion(v string) *ReservedCapacityCreateConfiguration
	GetRegion() *string
	SetYears(v int64) *ReservedCapacityCreateConfiguration
	GetYears() *int64
}

type ReservedCapacityCreateConfiguration struct {
	// example:
	//
	// 11258999068426240
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// example:
	//
	// test-rc-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1
	Years *int64 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s ReservedCapacityCreateConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ReservedCapacityCreateConfiguration) GoString() string {
	return s.String()
}

func (s *ReservedCapacityCreateConfiguration) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ReservedCapacityCreateConfiguration) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *ReservedCapacityCreateConfiguration) GetName() *string {
	return s.Name
}

func (s *ReservedCapacityCreateConfiguration) GetRegion() *string {
	return s.Region
}

func (s *ReservedCapacityCreateConfiguration) GetYears() *int64 {
	return s.Years
}

func (s *ReservedCapacityCreateConfiguration) SetCapacity(v int64) *ReservedCapacityCreateConfiguration {
	s.Capacity = &v
	return s
}

func (s *ReservedCapacityCreateConfiguration) SetDataRedundancyType(v string) *ReservedCapacityCreateConfiguration {
	s.DataRedundancyType = &v
	return s
}

func (s *ReservedCapacityCreateConfiguration) SetName(v string) *ReservedCapacityCreateConfiguration {
	s.Name = &v
	return s
}

func (s *ReservedCapacityCreateConfiguration) SetRegion(v string) *ReservedCapacityCreateConfiguration {
	s.Region = &v
	return s
}

func (s *ReservedCapacityCreateConfiguration) SetYears(v int64) *ReservedCapacityCreateConfiguration {
	s.Years = &v
	return s
}

func (s *ReservedCapacityCreateConfiguration) Validate() error {
	return dara.Validate(s)
}
