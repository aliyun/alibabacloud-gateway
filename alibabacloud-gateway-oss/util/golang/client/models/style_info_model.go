// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStyleInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *StyleInfo
	GetCategory() *string
	SetContent(v string) *StyleInfo
	GetContent() *string
	SetCreateTime(v string) *StyleInfo
	GetCreateTime() *string
	SetLastModifyTime(v string) *StyleInfo
	GetLastModifyTime() *string
	SetName(v string) *StyleInfo
	GetName() *string
}

type StyleInfo struct {
	// example:
	//
	// image
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// image/resize,p_50
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// Wed, 20 May 2020 12:07:15 GMT
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Wed, 21 May 2020 12:07:15 GMT
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// example:
	//
	// first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StyleInfo) String() string {
	return dara.Prettify(s)
}

func (s StyleInfo) GoString() string {
	return s.String()
}

func (s *StyleInfo) GetCategory() *string {
	return s.Category
}

func (s *StyleInfo) GetContent() *string {
	return s.Content
}

func (s *StyleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *StyleInfo) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *StyleInfo) GetName() *string {
	return s.Name
}

func (s *StyleInfo) SetCategory(v string) *StyleInfo {
	s.Category = &v
	return s
}

func (s *StyleInfo) SetContent(v string) *StyleInfo {
	s.Content = &v
	return s
}

func (s *StyleInfo) SetCreateTime(v string) *StyleInfo {
	s.CreateTime = &v
	return s
}

func (s *StyleInfo) SetLastModifyTime(v string) *StyleInfo {
	s.LastModifyTime = &v
	return s
}

func (s *StyleInfo) SetName(v string) *StyleInfo {
	s.Name = &v
	return s
}

func (s *StyleInfo) Validate() error {
	return dara.Validate(s)
}
