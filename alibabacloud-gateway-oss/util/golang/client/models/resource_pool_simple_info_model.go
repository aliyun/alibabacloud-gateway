// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourcePoolSimpleInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *ResourcePoolSimpleInfo
	GetCreateTime() *string
	SetName(v string) *ResourcePoolSimpleInfo
	GetName() *string
}

type ResourcePoolSimpleInfo struct {
	// example:
	//
	// 2024-07-24T08:42:32.000Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// rp-for-api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ResourcePoolSimpleInfo) String() string {
	return dara.Prettify(s)
}

func (s ResourcePoolSimpleInfo) GoString() string {
	return s.String()
}

func (s *ResourcePoolSimpleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ResourcePoolSimpleInfo) GetName() *string {
	return s.Name
}

func (s *ResourcePoolSimpleInfo) SetCreateTime(v string) *ResourcePoolSimpleInfo {
	s.CreateTime = &v
	return s
}

func (s *ResourcePoolSimpleInfo) SetName(v string) *ResourcePoolSimpleInfo {
	s.Name = &v
	return s
}

func (s *ResourcePoolSimpleInfo) Validate() error {
	return dara.Validate(s)
}
