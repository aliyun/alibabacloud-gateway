// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReservedCapacityUpdateConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *ReservedCapacityUpdateConfiguration
	GetCapacity() *int64
	SetStatus(v string) *ReservedCapacityUpdateConfiguration
	GetStatus() *string
	SetYears(v int64) *ReservedCapacityUpdateConfiguration
	GetYears() *int64
}

type ReservedCapacityUpdateConfiguration struct {
	// example:
	//
	// 11258999068426240
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Years *int64 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s ReservedCapacityUpdateConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ReservedCapacityUpdateConfiguration) GoString() string {
	return s.String()
}

func (s *ReservedCapacityUpdateConfiguration) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ReservedCapacityUpdateConfiguration) GetStatus() *string {
	return s.Status
}

func (s *ReservedCapacityUpdateConfiguration) GetYears() *int64 {
	return s.Years
}

func (s *ReservedCapacityUpdateConfiguration) SetCapacity(v int64) *ReservedCapacityUpdateConfiguration {
	s.Capacity = &v
	return s
}

func (s *ReservedCapacityUpdateConfiguration) SetStatus(v string) *ReservedCapacityUpdateConfiguration {
	s.Status = &v
	return s
}

func (s *ReservedCapacityUpdateConfiguration) SetYears(v int64) *ReservedCapacityUpdateConfiguration {
	s.Years = &v
	return s
}

func (s *ReservedCapacityUpdateConfiguration) Validate() error {
	return dara.Validate(s)
}
