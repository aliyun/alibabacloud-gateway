// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannel interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSetContentType(v bool) *Channel
	GetAutoSetContentType() *bool
	SetDefault404Pic(v string) *Channel
	GetDefault404Pic() *string
	SetOrigPicForbidden(v bool) *Channel
	GetOrigPicForbidden() *bool
	SetSetAttachName(v bool) *Channel
	GetSetAttachName() *bool
	SetStatus(v string) *Channel
	GetStatus() *string
	SetStyleDelimiters(v string) *Channel
	GetStyleDelimiters() *string
	SetUseSrcFormat(v bool) *Channel
	GetUseSrcFormat() *bool
	SetUseStyleOnly(v bool) *Channel
	GetUseStyleOnly() *bool
}

type Channel struct {
	// example:
	//
	// true
	AutoSetContentType *bool `json:"AutoSetContentType,omitempty" xml:"AutoSetContentType,omitempty"`
	// example:
	//
	// 404.jpg
	Default404Pic *string `json:"Default404Pic,omitempty" xml:"Default404Pic,omitempty"`
	// example:
	//
	// true
	OrigPicForbidden *bool `json:"OrigPicForbidden,omitempty" xml:"OrigPicForbidden,omitempty"`
	// example:
	//
	// false
	SetAttachName *bool `json:"SetAttachName,omitempty" xml:"SetAttachName,omitempty"`
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// -,|
	StyleDelimiters *string `json:"StyleDelimiters,omitempty" xml:"StyleDelimiters,omitempty"`
	// example:
	//
	// true
	UseSrcFormat *bool `json:"UseSrcFormat,omitempty" xml:"UseSrcFormat,omitempty"`
	// example:
	//
	// false
	UseStyleOnly *bool `json:"UseStyleOnly,omitempty" xml:"UseStyleOnly,omitempty"`
}

func (s Channel) String() string {
	return dara.Prettify(s)
}

func (s Channel) GoString() string {
	return s.String()
}

func (s *Channel) GetAutoSetContentType() *bool {
	return s.AutoSetContentType
}

func (s *Channel) GetDefault404Pic() *string {
	return s.Default404Pic
}

func (s *Channel) GetOrigPicForbidden() *bool {
	return s.OrigPicForbidden
}

func (s *Channel) GetSetAttachName() *bool {
	return s.SetAttachName
}

func (s *Channel) GetStatus() *string {
	return s.Status
}

func (s *Channel) GetStyleDelimiters() *string {
	return s.StyleDelimiters
}

func (s *Channel) GetUseSrcFormat() *bool {
	return s.UseSrcFormat
}

func (s *Channel) GetUseStyleOnly() *bool {
	return s.UseStyleOnly
}

func (s *Channel) SetAutoSetContentType(v bool) *Channel {
	s.AutoSetContentType = &v
	return s
}

func (s *Channel) SetDefault404Pic(v string) *Channel {
	s.Default404Pic = &v
	return s
}

func (s *Channel) SetOrigPicForbidden(v bool) *Channel {
	s.OrigPicForbidden = &v
	return s
}

func (s *Channel) SetSetAttachName(v bool) *Channel {
	s.SetAttachName = &v
	return s
}

func (s *Channel) SetStatus(v string) *Channel {
	s.Status = &v
	return s
}

func (s *Channel) SetStyleDelimiters(v string) *Channel {
	s.StyleDelimiters = &v
	return s
}

func (s *Channel) SetUseSrcFormat(v bool) *Channel {
	s.UseSrcFormat = &v
	return s
}

func (s *Channel) SetUseStyleOnly(v bool) *Channel {
	s.UseStyleOnly = &v
	return s
}

func (s *Channel) Validate() error {
	return dara.Validate(s)
}
