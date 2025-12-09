// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChannelResult interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSetContentType(v bool) *GetChannelResult
	GetAutoSetContentType() *bool
	SetCreateTime(v string) *GetChannelResult
	GetCreateTime() *string
	SetDefault404Pic(v string) *GetChannelResult
	GetDefault404Pic() *string
	SetLastModifyTime(v string) *GetChannelResult
	GetLastModifyTime() *string
	SetName(v string) *GetChannelResult
	GetName() *string
	SetOrigPicForbidden(v bool) *GetChannelResult
	GetOrigPicForbidden() *bool
	SetSetAttachName(v bool) *GetChannelResult
	GetSetAttachName() *bool
	SetStatus(v string) *GetChannelResult
	GetStatus() *string
	SetStyleDelimiters(v string) *GetChannelResult
	GetStyleDelimiters() *string
	SetUseSrcFormat(v bool) *GetChannelResult
	GetUseSrcFormat() *bool
	SetUseStyleOnly(v bool) *GetChannelResult
	GetUseStyleOnly() *bool
}

type GetChannelResult struct {
	// example:
	//
	// false
	AutoSetContentType *bool `json:"AutoSetContentType,omitempty" xml:"AutoSetContentType,omitempty"`
	// example:
	//
	// Fri, 15 Nov 2024 08:11:56 GMT
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 404.jpg
	Default404Pic *string `json:"Default404Pic,omitempty" xml:"Default404Pic,omitempty"`
	// example:
	//
	// Fri, 15 Nov 2024 08:11:56 GMT
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// example:
	//
	// test-bucket
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	OrigPicForbidden *bool `json:"OrigPicForbidden,omitempty" xml:"OrigPicForbidden,omitempty"`
	// example:
	//
	// true
	SetAttachName *bool `json:"SetAttachName,omitempty" xml:"SetAttachName,omitempty"`
	// example:
	//
	// enble
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// -,|
	StyleDelimiters *string `json:"StyleDelimiters,omitempty" xml:"StyleDelimiters,omitempty"`
	// example:
	//
	// false
	UseSrcFormat *bool `json:"UseSrcFormat,omitempty" xml:"UseSrcFormat,omitempty"`
	// example:
	//
	// true
	UseStyleOnly *bool `json:"UseStyleOnly,omitempty" xml:"UseStyleOnly,omitempty"`
}

func (s GetChannelResult) String() string {
	return dara.Prettify(s)
}

func (s GetChannelResult) GoString() string {
	return s.String()
}

func (s *GetChannelResult) GetAutoSetContentType() *bool {
	return s.AutoSetContentType
}

func (s *GetChannelResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetChannelResult) GetDefault404Pic() *string {
	return s.Default404Pic
}

func (s *GetChannelResult) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *GetChannelResult) GetName() *string {
	return s.Name
}

func (s *GetChannelResult) GetOrigPicForbidden() *bool {
	return s.OrigPicForbidden
}

func (s *GetChannelResult) GetSetAttachName() *bool {
	return s.SetAttachName
}

func (s *GetChannelResult) GetStatus() *string {
	return s.Status
}

func (s *GetChannelResult) GetStyleDelimiters() *string {
	return s.StyleDelimiters
}

func (s *GetChannelResult) GetUseSrcFormat() *bool {
	return s.UseSrcFormat
}

func (s *GetChannelResult) GetUseStyleOnly() *bool {
	return s.UseStyleOnly
}

func (s *GetChannelResult) SetAutoSetContentType(v bool) *GetChannelResult {
	s.AutoSetContentType = &v
	return s
}

func (s *GetChannelResult) SetCreateTime(v string) *GetChannelResult {
	s.CreateTime = &v
	return s
}

func (s *GetChannelResult) SetDefault404Pic(v string) *GetChannelResult {
	s.Default404Pic = &v
	return s
}

func (s *GetChannelResult) SetLastModifyTime(v string) *GetChannelResult {
	s.LastModifyTime = &v
	return s
}

func (s *GetChannelResult) SetName(v string) *GetChannelResult {
	s.Name = &v
	return s
}

func (s *GetChannelResult) SetOrigPicForbidden(v bool) *GetChannelResult {
	s.OrigPicForbidden = &v
	return s
}

func (s *GetChannelResult) SetSetAttachName(v bool) *GetChannelResult {
	s.SetAttachName = &v
	return s
}

func (s *GetChannelResult) SetStatus(v string) *GetChannelResult {
	s.Status = &v
	return s
}

func (s *GetChannelResult) SetStyleDelimiters(v string) *GetChannelResult {
	s.StyleDelimiters = &v
	return s
}

func (s *GetChannelResult) SetUseSrcFormat(v bool) *GetChannelResult {
	s.UseSrcFormat = &v
	return s
}

func (s *GetChannelResult) SetUseStyleOnly(v bool) *GetChannelResult {
	s.UseStyleOnly = &v
	return s
}

func (s *GetChannelResult) Validate() error {
	return dara.Validate(s)
}
