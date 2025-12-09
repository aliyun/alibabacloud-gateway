// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefererConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAllowEmptyReferer(v bool) *RefererConfiguration
	GetAllowEmptyReferer() *bool
	SetAllowTruncateQueryString(v bool) *RefererConfiguration
	GetAllowTruncateQueryString() *bool
	SetRefererBlacklist(v *RefererConfigurationRefererBlacklist) *RefererConfiguration
	GetRefererBlacklist() *RefererConfigurationRefererBlacklist
	SetRefererList(v *RefererConfigurationRefererList) *RefererConfiguration
	GetRefererList() *RefererConfigurationRefererList
	SetTruncatePath(v bool) *RefererConfiguration
	GetTruncatePath() *bool
}

type RefererConfiguration struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	AllowEmptyReferer *bool `json:"AllowEmptyReferer,omitempty" xml:"AllowEmptyReferer,omitempty"`
	// example:
	//
	// false
	AllowTruncateQueryString *bool                                 `json:"AllowTruncateQueryString,omitempty" xml:"AllowTruncateQueryString,omitempty"`
	RefererBlacklist         *RefererConfigurationRefererBlacklist `json:"RefererBlacklist,omitempty" xml:"RefererBlacklist,omitempty" type:"Struct"`
	// This parameter is required.
	RefererList *RefererConfigurationRefererList `json:"RefererList,omitempty" xml:"RefererList,omitempty" type:"Struct"`
	// example:
	//
	// false
	TruncatePath *bool `json:"TruncatePath,omitempty" xml:"TruncatePath,omitempty"`
}

func (s RefererConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RefererConfiguration) GoString() string {
	return s.String()
}

func (s *RefererConfiguration) GetAllowEmptyReferer() *bool {
	return s.AllowEmptyReferer
}

func (s *RefererConfiguration) GetAllowTruncateQueryString() *bool {
	return s.AllowTruncateQueryString
}

func (s *RefererConfiguration) GetRefererBlacklist() *RefererConfigurationRefererBlacklist {
	return s.RefererBlacklist
}

func (s *RefererConfiguration) GetRefererList() *RefererConfigurationRefererList {
	return s.RefererList
}

func (s *RefererConfiguration) GetTruncatePath() *bool {
	return s.TruncatePath
}

func (s *RefererConfiguration) SetAllowEmptyReferer(v bool) *RefererConfiguration {
	s.AllowEmptyReferer = &v
	return s
}

func (s *RefererConfiguration) SetAllowTruncateQueryString(v bool) *RefererConfiguration {
	s.AllowTruncateQueryString = &v
	return s
}

func (s *RefererConfiguration) SetRefererBlacklist(v *RefererConfigurationRefererBlacklist) *RefererConfiguration {
	s.RefererBlacklist = v
	return s
}

func (s *RefererConfiguration) SetRefererList(v *RefererConfigurationRefererList) *RefererConfiguration {
	s.RefererList = v
	return s
}

func (s *RefererConfiguration) SetTruncatePath(v bool) *RefererConfiguration {
	s.TruncatePath = &v
	return s
}

func (s *RefererConfiguration) Validate() error {
	if s.RefererBlacklist != nil {
		if err := s.RefererBlacklist.Validate(); err != nil {
			return err
		}
	}
	if s.RefererList != nil {
		if err := s.RefererList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefererConfigurationRefererBlacklist struct {
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
}

func (s RefererConfigurationRefererBlacklist) String() string {
	return dara.Prettify(s)
}

func (s RefererConfigurationRefererBlacklist) GoString() string {
	return s.String()
}

func (s *RefererConfigurationRefererBlacklist) GetReferer() []*string {
	return s.Referer
}

func (s *RefererConfigurationRefererBlacklist) SetReferer(v []*string) *RefererConfigurationRefererBlacklist {
	s.Referer = v
	return s
}

func (s *RefererConfigurationRefererBlacklist) Validate() error {
	return dara.Validate(s)
}

type RefererConfigurationRefererList struct {
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
}

func (s RefererConfigurationRefererList) String() string {
	return dara.Prettify(s)
}

func (s RefererConfigurationRefererList) GoString() string {
	return s.String()
}

func (s *RefererConfigurationRefererList) GetReferer() []*string {
	return s.Referer
}

func (s *RefererConfigurationRefererList) SetReferer(v []*string) *RefererConfigurationRefererList {
	s.Referer = v
	return s
}

func (s *RefererConfigurationRefererList) Validate() error {
	return dara.Validate(s)
}
