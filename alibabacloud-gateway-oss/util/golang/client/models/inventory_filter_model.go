// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInventoryFilter interface {
	dara.Model
	String() string
	GoString() string
	SetLastModifyBeginTimeStamp(v int64) *InventoryFilter
	GetLastModifyBeginTimeStamp() *int64
	SetLastModifyEndTimeStamp(v int64) *InventoryFilter
	GetLastModifyEndTimeStamp() *int64
	SetLowerSizeBound(v int64) *InventoryFilter
	GetLowerSizeBound() *int64
	SetPrefix(v string) *InventoryFilter
	GetPrefix() *string
	SetStorageClass(v string) *InventoryFilter
	GetStorageClass() *string
	SetTags(v string) *InventoryFilter
	GetTags() *string
	SetTagsCondition(v string) *InventoryFilter
	GetTagsCondition() *string
	SetUpperSizeBound(v int64) *InventoryFilter
	GetUpperSizeBound() *int64
}

type InventoryFilter struct {
	// example:
	//
	// 1637883649
	LastModifyBeginTimeStamp *int64 `json:"LastModifyBeginTimeStamp,omitempty" xml:"LastModifyBeginTimeStamp,omitempty"`
	// example:
	//
	// 1638347592
	LastModifyEndTimeStamp *int64 `json:"LastModifyEndTimeStamp,omitempty" xml:"LastModifyEndTimeStamp,omitempty"`
	// example:
	//
	// 1024
	LowerSizeBound *int64 `json:"LowerSizeBound,omitempty" xml:"LowerSizeBound,omitempty"`
	// example:
	//
	// Pics/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// All
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// example:
	//
	// tag1#val1;tag2#val2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// OR_FILTER
	TagsCondition *string `json:"TagsCondition,omitempty" xml:"TagsCondition,omitempty"`
	// example:
	//
	// 1048576
	UpperSizeBound *int64 `json:"UpperSizeBound,omitempty" xml:"UpperSizeBound,omitempty"`
}

func (s InventoryFilter) String() string {
	return dara.Prettify(s)
}

func (s InventoryFilter) GoString() string {
	return s.String()
}

func (s *InventoryFilter) GetLastModifyBeginTimeStamp() *int64 {
	return s.LastModifyBeginTimeStamp
}

func (s *InventoryFilter) GetLastModifyEndTimeStamp() *int64 {
	return s.LastModifyEndTimeStamp
}

func (s *InventoryFilter) GetLowerSizeBound() *int64 {
	return s.LowerSizeBound
}

func (s *InventoryFilter) GetPrefix() *string {
	return s.Prefix
}

func (s *InventoryFilter) GetStorageClass() *string {
	return s.StorageClass
}

func (s *InventoryFilter) GetTags() *string {
	return s.Tags
}

func (s *InventoryFilter) GetTagsCondition() *string {
	return s.TagsCondition
}

func (s *InventoryFilter) GetUpperSizeBound() *int64 {
	return s.UpperSizeBound
}

func (s *InventoryFilter) SetLastModifyBeginTimeStamp(v int64) *InventoryFilter {
	s.LastModifyBeginTimeStamp = &v
	return s
}

func (s *InventoryFilter) SetLastModifyEndTimeStamp(v int64) *InventoryFilter {
	s.LastModifyEndTimeStamp = &v
	return s
}

func (s *InventoryFilter) SetLowerSizeBound(v int64) *InventoryFilter {
	s.LowerSizeBound = &v
	return s
}

func (s *InventoryFilter) SetPrefix(v string) *InventoryFilter {
	s.Prefix = &v
	return s
}

func (s *InventoryFilter) SetStorageClass(v string) *InventoryFilter {
	s.StorageClass = &v
	return s
}

func (s *InventoryFilter) SetTags(v string) *InventoryFilter {
	s.Tags = &v
	return s
}

func (s *InventoryFilter) SetTagsCondition(v string) *InventoryFilter {
	s.TagsCondition = &v
	return s
}

func (s *InventoryFilter) SetUpperSizeBound(v int64) *InventoryFilter {
	s.UpperSizeBound = &v
	return s
}

func (s *InventoryFilter) Validate() error {
	return dara.Validate(s)
}
