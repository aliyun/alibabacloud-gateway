// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryAggregationsResult interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *MetaQueryAggregationsResult
	GetField() *string
	SetGroups(v *MetaQueryAggregationsResultGroups) *MetaQueryAggregationsResult
	GetGroups() *MetaQueryAggregationsResultGroups
	SetOperation(v string) *MetaQueryAggregationsResult
	GetOperation() *string
	SetValue(v float64) *MetaQueryAggregationsResult
	GetValue() *float64
}

type MetaQueryAggregationsResult struct {
	// example:
	//
	// Size
	Field  *string                            `json:"Field,omitempty" xml:"Field,omitempty"`
	Groups *MetaQueryAggregationsResultGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// example:
	//
	// sum
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 200
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MetaQueryAggregationsResult) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryAggregationsResult) GoString() string {
	return s.String()
}

func (s *MetaQueryAggregationsResult) GetField() *string {
	return s.Field
}

func (s *MetaQueryAggregationsResult) GetGroups() *MetaQueryAggregationsResultGroups {
	return s.Groups
}

func (s *MetaQueryAggregationsResult) GetOperation() *string {
	return s.Operation
}

func (s *MetaQueryAggregationsResult) GetValue() *float64 {
	return s.Value
}

func (s *MetaQueryAggregationsResult) SetField(v string) *MetaQueryAggregationsResult {
	s.Field = &v
	return s
}

func (s *MetaQueryAggregationsResult) SetGroups(v *MetaQueryAggregationsResultGroups) *MetaQueryAggregationsResult {
	s.Groups = v
	return s
}

func (s *MetaQueryAggregationsResult) SetOperation(v string) *MetaQueryAggregationsResult {
	s.Operation = &v
	return s
}

func (s *MetaQueryAggregationsResult) SetValue(v float64) *MetaQueryAggregationsResult {
	s.Value = &v
	return s
}

func (s *MetaQueryAggregationsResult) Validate() error {
	if s.Groups != nil {
		if err := s.Groups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryAggregationsResultGroups struct {
	Group []*MetaQueryAggregationsResultGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s MetaQueryAggregationsResultGroups) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryAggregationsResultGroups) GoString() string {
	return s.String()
}

func (s *MetaQueryAggregationsResultGroups) GetGroup() []*MetaQueryAggregationsResultGroupsGroup {
	return s.Group
}

func (s *MetaQueryAggregationsResultGroups) SetGroup(v []*MetaQueryAggregationsResultGroupsGroup) *MetaQueryAggregationsResultGroups {
	s.Group = v
	return s
}

func (s *MetaQueryAggregationsResultGroups) Validate() error {
	if s.Group != nil {
		for _, item := range s.Group {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetaQueryAggregationsResultGroupsGroup struct {
	// example:
	//
	// 5
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MetaQueryAggregationsResultGroupsGroup) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryAggregationsResultGroupsGroup) GoString() string {
	return s.String()
}

func (s *MetaQueryAggregationsResultGroupsGroup) GetCount() *int64 {
	return s.Count
}

func (s *MetaQueryAggregationsResultGroupsGroup) GetValue() *string {
	return s.Value
}

func (s *MetaQueryAggregationsResultGroupsGroup) SetCount(v int64) *MetaQueryAggregationsResultGroupsGroup {
	s.Count = &v
	return s
}

func (s *MetaQueryAggregationsResultGroupsGroup) SetValue(v string) *MetaQueryAggregationsResultGroupsGroup {
	s.Value = &v
	return s
}

func (s *MetaQueryAggregationsResultGroupsGroup) Validate() error {
	return dara.Validate(s)
}
