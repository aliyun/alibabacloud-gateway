// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSetContentType(v bool) *ChannelInfo
	GetAutoSetContentType() *bool
	SetName(v string) *ChannelInfo
	GetName() *string
	SetOrigPicForbidden(v bool) *ChannelInfo
	GetOrigPicForbidden() *bool
	SetSetAttachName(v bool) *ChannelInfo
	GetSetAttachName() *bool
	SetStatus(v string) *ChannelInfo
	GetStatus() *string
	SetStyleDelimiters(v string) *ChannelInfo
	GetStyleDelimiters() *string
	SetUseSrcFormat(v bool) *ChannelInfo
	GetUseSrcFormat() *bool
	SetUseStyleOnly(v bool) *ChannelInfo
	GetUseStyleOnly() *bool
}

type ChannelInfo struct {
	// example:
	//
	// False
	AutoSetContentType *bool `json:"AutoSetContentType,omitempty" xml:"AutoSetContentType,omitempty"`
	// example:
	//
	// test-channel
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// True
	OrigPicForbidden *bool `json:"OrigPicForbidden,omitempty" xml:"OrigPicForbidden,omitempty"`
	// example:
	//
	// True
	SetAttachName *bool `json:"SetAttachName,omitempty" xml:"SetAttachName,omitempty"`
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// -
	StyleDelimiters *string `json:"StyleDelimiters,omitempty" xml:"StyleDelimiters,omitempty"`
	// example:
	//
	// True
	UseSrcFormat *bool `json:"UseSrcFormat,omitempty" xml:"UseSrcFormat,omitempty"`
	// example:
	//
	// False
	UseStyleOnly *bool `json:"UseStyleOnly,omitempty" xml:"UseStyleOnly,omitempty"`
}

func (s ChannelInfo) String() string {
	return dara.Prettify(s)
}

func (s ChannelInfo) GoString() string {
	return s.String()
}

func (s *ChannelInfo) GetAutoSetContentType() *bool {
	return s.AutoSetContentType
}

func (s *ChannelInfo) GetName() *string {
	return s.Name
}

func (s *ChannelInfo) GetOrigPicForbidden() *bool {
	return s.OrigPicForbidden
}

func (s *ChannelInfo) GetSetAttachName() *bool {
	return s.SetAttachName
}

func (s *ChannelInfo) GetStatus() *string {
	return s.Status
}

func (s *ChannelInfo) GetStyleDelimiters() *string {
	return s.StyleDelimiters
}

func (s *ChannelInfo) GetUseSrcFormat() *bool {
	return s.UseSrcFormat
}

func (s *ChannelInfo) GetUseStyleOnly() *bool {
	return s.UseStyleOnly
}

func (s *ChannelInfo) SetAutoSetContentType(v bool) *ChannelInfo {
	s.AutoSetContentType = &v
	return s
}

func (s *ChannelInfo) SetName(v string) *ChannelInfo {
	s.Name = &v
	return s
}

func (s *ChannelInfo) SetOrigPicForbidden(v bool) *ChannelInfo {
	s.OrigPicForbidden = &v
	return s
}

func (s *ChannelInfo) SetSetAttachName(v bool) *ChannelInfo {
	s.SetAttachName = &v
	return s
}

func (s *ChannelInfo) SetStatus(v string) *ChannelInfo {
	s.Status = &v
	return s
}

func (s *ChannelInfo) SetStyleDelimiters(v string) *ChannelInfo {
	s.StyleDelimiters = &v
	return s
}

func (s *ChannelInfo) SetUseSrcFormat(v bool) *ChannelInfo {
	s.UseSrcFormat = &v
	return s
}

func (s *ChannelInfo) SetUseStyleOnly(v bool) *ChannelInfo {
	s.UseStyleOnly = &v
	return s
}

func (s *ChannelInfo) Validate() error {
	return dara.Validate(s)
}
