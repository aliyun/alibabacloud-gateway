// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLargeReservedCapacityResult interface {
	dara.Model
	String() string
	GoString() string
	SetID(v string) *CreateLargeReservedCapacityResult
	GetID() *string
	SetName(v string) *CreateLargeReservedCapacityResult
	GetName() *string
	SetOwner(v *Owner) *CreateLargeReservedCapacityResult
	GetOwner() *Owner
	SetRegion(v string) *CreateLargeReservedCapacityResult
	GetRegion() *string
}

type CreateLargeReservedCapacityResult struct {
	// example:
	//
	// c99797e7-510d-41e3-ac82-e851c17e1f5c
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// example:
	//
	// test-rc-01
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateLargeReservedCapacityResult) String() string {
	return dara.Prettify(s)
}

func (s CreateLargeReservedCapacityResult) GoString() string {
	return s.String()
}

func (s *CreateLargeReservedCapacityResult) GetID() *string {
	return s.ID
}

func (s *CreateLargeReservedCapacityResult) GetName() *string {
	return s.Name
}

func (s *CreateLargeReservedCapacityResult) GetOwner() *Owner {
	return s.Owner
}

func (s *CreateLargeReservedCapacityResult) GetRegion() *string {
	return s.Region
}

func (s *CreateLargeReservedCapacityResult) SetID(v string) *CreateLargeReservedCapacityResult {
	s.ID = &v
	return s
}

func (s *CreateLargeReservedCapacityResult) SetName(v string) *CreateLargeReservedCapacityResult {
	s.Name = &v
	return s
}

func (s *CreateLargeReservedCapacityResult) SetOwner(v *Owner) *CreateLargeReservedCapacityResult {
	s.Owner = v
	return s
}

func (s *CreateLargeReservedCapacityResult) SetRegion(v string) *CreateLargeReservedCapacityResult {
	s.Region = &v
	return s
}

func (s *CreateLargeReservedCapacityResult) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}
