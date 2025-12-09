// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutChannelConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSetContentType(v bool) *PutChannelConfiguration
	GetAutoSetContentType() *bool
	SetDefault404Pic(v string) *PutChannelConfiguration
	GetDefault404Pic() *string
	SetOrigPicForbidden(v bool) *PutChannelConfiguration
	GetOrigPicForbidden() *bool
	SetSetAttachName(v bool) *PutChannelConfiguration
	GetSetAttachName() *bool
	SetStatus(v string) *PutChannelConfiguration
	GetStatus() *string
	SetStyleDelimiters(v string) *PutChannelConfiguration
	GetStyleDelimiters() *string
	SetUseSrcFormat(v bool) *PutChannelConfiguration
	GetUseSrcFormat() *bool
	SetUseStyleOnly(v bool) *PutChannelConfiguration
	GetUseStyleOnly() *bool
}

type PutChannelConfiguration struct {
	// example:
	//
	// False
	AutoSetContentType *bool `json:"AutoSetContentType,omitempty" xml:"AutoSetContentType,omitempty"`
	// example:
	//
	// 404.jpg
	Default404Pic *string `json:"Default404Pic,omitempty" xml:"Default404Pic,omitempty"`
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

func (s PutChannelConfiguration) String() string {
	return dara.Prettify(s)
}

func (s PutChannelConfiguration) GoString() string {
	return s.String()
}

func (s *PutChannelConfiguration) GetAutoSetContentType() *bool {
	return s.AutoSetContentType
}

func (s *PutChannelConfiguration) GetDefault404Pic() *string {
	return s.Default404Pic
}

func (s *PutChannelConfiguration) GetOrigPicForbidden() *bool {
	return s.OrigPicForbidden
}

func (s *PutChannelConfiguration) GetSetAttachName() *bool {
	return s.SetAttachName
}

func (s *PutChannelConfiguration) GetStatus() *string {
	return s.Status
}

func (s *PutChannelConfiguration) GetStyleDelimiters() *string {
	return s.StyleDelimiters
}

func (s *PutChannelConfiguration) GetUseSrcFormat() *bool {
	return s.UseSrcFormat
}

func (s *PutChannelConfiguration) GetUseStyleOnly() *bool {
	return s.UseStyleOnly
}

func (s *PutChannelConfiguration) SetAutoSetContentType(v bool) *PutChannelConfiguration {
	s.AutoSetContentType = &v
	return s
}

func (s *PutChannelConfiguration) SetDefault404Pic(v string) *PutChannelConfiguration {
	s.Default404Pic = &v
	return s
}

func (s *PutChannelConfiguration) SetOrigPicForbidden(v bool) *PutChannelConfiguration {
	s.OrigPicForbidden = &v
	return s
}

func (s *PutChannelConfiguration) SetSetAttachName(v bool) *PutChannelConfiguration {
	s.SetAttachName = &v
	return s
}

func (s *PutChannelConfiguration) SetStatus(v string) *PutChannelConfiguration {
	s.Status = &v
	return s
}

func (s *PutChannelConfiguration) SetStyleDelimiters(v string) *PutChannelConfiguration {
	s.StyleDelimiters = &v
	return s
}

func (s *PutChannelConfiguration) SetUseSrcFormat(v bool) *PutChannelConfiguration {
	s.UseSrcFormat = &v
	return s
}

func (s *PutChannelConfiguration) SetUseStyleOnly(v bool) *PutChannelConfiguration {
	s.UseStyleOnly = &v
	return s
}

func (s *PutChannelConfiguration) Validate() error {
	return dara.Validate(s)
}
