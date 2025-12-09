// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReservedCapacityRecord interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *ReservedCapacityRecord
	GetCapacity() *int64
	SetCreateTime(v int64) *ReservedCapacityRecord
	GetCreateTime() *int64
	SetDataRedundancyType(v string) *ReservedCapacityRecord
	GetDataRedundancyType() *string
	SetDueTime(v int64) *ReservedCapacityRecord
	GetDueTime() *int64
	SetExpansionTime(v int64) *ReservedCapacityRecord
	GetExpansionTime() *int64
	SetFirstTimeEnabled(v int64) *ReservedCapacityRecord
	GetFirstTimeEnabled() *int64
	SetID(v string) *ReservedCapacityRecord
	GetID() *string
	SetLastExpansionCapacity(v int64) *ReservedCapacityRecord
	GetLastExpansionCapacity() *int64
	SetLastModifyTime(v int64) *ReservedCapacityRecord
	GetLastModifyTime() *int64
	SetName(v string) *ReservedCapacityRecord
	GetName() *string
	SetOwner(v *Owner) *ReservedCapacityRecord
	GetOwner() *Owner
	SetRegion(v string) *ReservedCapacityRecord
	GetRegion() *string
	SetStatus(v string) *ReservedCapacityRecord
	GetStatus() *string
	SetVersion(v int64) *ReservedCapacityRecord
	GetVersion() *int64
	SetYears(v int64) *ReservedCapacityRecord
	GetYears() *int64
}

type ReservedCapacityRecord struct {
	// example:
	//
	// 11258999068426240
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// 1733822455
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// example:
	//
	// 0
	DueTime *int64 `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 0
	ExpansionTime *int64 `json:"ExpansionTime,omitempty" xml:"ExpansionTime,omitempty"`
	// example:
	//
	// 0
	FirstTimeEnabled *int64 `json:"FirstTimeEnabled,omitempty" xml:"FirstTimeEnabled,omitempty"`
	// example:
	//
	// c99797e7-510d-41e3-ac82-e851c17e1f5c
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// example:
	//
	// 0
	LastExpansionCapacity *int64 `json:"LastExpansionCapacity,omitempty" xml:"LastExpansionCapacity,omitempty"`
	// example:
	//
	// 1733822455
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// example:
	//
	// test-rc-01
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 1
	Years *int64 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s ReservedCapacityRecord) String() string {
	return dara.Prettify(s)
}

func (s ReservedCapacityRecord) GoString() string {
	return s.String()
}

func (s *ReservedCapacityRecord) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ReservedCapacityRecord) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ReservedCapacityRecord) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *ReservedCapacityRecord) GetDueTime() *int64 {
	return s.DueTime
}

func (s *ReservedCapacityRecord) GetExpansionTime() *int64 {
	return s.ExpansionTime
}

func (s *ReservedCapacityRecord) GetFirstTimeEnabled() *int64 {
	return s.FirstTimeEnabled
}

func (s *ReservedCapacityRecord) GetID() *string {
	return s.ID
}

func (s *ReservedCapacityRecord) GetLastExpansionCapacity() *int64 {
	return s.LastExpansionCapacity
}

func (s *ReservedCapacityRecord) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *ReservedCapacityRecord) GetName() *string {
	return s.Name
}

func (s *ReservedCapacityRecord) GetOwner() *Owner {
	return s.Owner
}

func (s *ReservedCapacityRecord) GetRegion() *string {
	return s.Region
}

func (s *ReservedCapacityRecord) GetStatus() *string {
	return s.Status
}

func (s *ReservedCapacityRecord) GetVersion() *int64 {
	return s.Version
}

func (s *ReservedCapacityRecord) GetYears() *int64 {
	return s.Years
}

func (s *ReservedCapacityRecord) SetCapacity(v int64) *ReservedCapacityRecord {
	s.Capacity = &v
	return s
}

func (s *ReservedCapacityRecord) SetCreateTime(v int64) *ReservedCapacityRecord {
	s.CreateTime = &v
	return s
}

func (s *ReservedCapacityRecord) SetDataRedundancyType(v string) *ReservedCapacityRecord {
	s.DataRedundancyType = &v
	return s
}

func (s *ReservedCapacityRecord) SetDueTime(v int64) *ReservedCapacityRecord {
	s.DueTime = &v
	return s
}

func (s *ReservedCapacityRecord) SetExpansionTime(v int64) *ReservedCapacityRecord {
	s.ExpansionTime = &v
	return s
}

func (s *ReservedCapacityRecord) SetFirstTimeEnabled(v int64) *ReservedCapacityRecord {
	s.FirstTimeEnabled = &v
	return s
}

func (s *ReservedCapacityRecord) SetID(v string) *ReservedCapacityRecord {
	s.ID = &v
	return s
}

func (s *ReservedCapacityRecord) SetLastExpansionCapacity(v int64) *ReservedCapacityRecord {
	s.LastExpansionCapacity = &v
	return s
}

func (s *ReservedCapacityRecord) SetLastModifyTime(v int64) *ReservedCapacityRecord {
	s.LastModifyTime = &v
	return s
}

func (s *ReservedCapacityRecord) SetName(v string) *ReservedCapacityRecord {
	s.Name = &v
	return s
}

func (s *ReservedCapacityRecord) SetOwner(v *Owner) *ReservedCapacityRecord {
	s.Owner = v
	return s
}

func (s *ReservedCapacityRecord) SetRegion(v string) *ReservedCapacityRecord {
	s.Region = &v
	return s
}

func (s *ReservedCapacityRecord) SetStatus(v string) *ReservedCapacityRecord {
	s.Status = &v
	return s
}

func (s *ReservedCapacityRecord) SetVersion(v int64) *ReservedCapacityRecord {
	s.Version = &v
	return s
}

func (s *ReservedCapacityRecord) SetYears(v int64) *ReservedCapacityRecord {
	s.Years = &v
	return s
}

func (s *ReservedCapacityRecord) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}
