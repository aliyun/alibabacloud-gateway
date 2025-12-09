// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCORSRule interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedHeader(v []*string) *CORSRule
	GetAllowedHeader() []*string
	SetAllowedMethod(v []*string) *CORSRule
	GetAllowedMethod() []*string
	SetAllowedOrigin(v []*string) *CORSRule
	GetAllowedOrigin() []*string
	SetExposeHeader(v []*string) *CORSRule
	GetExposeHeader() []*string
	SetMaxAgeSeconds(v int64) *CORSRule
	GetMaxAgeSeconds() *int64
}

type CORSRule struct {
	AllowedHeader []*string `json:"AllowedHeader,omitempty" xml:"AllowedHeader,omitempty" type:"Repeated"`
	AllowedMethod []*string `json:"AllowedMethod,omitempty" xml:"AllowedMethod,omitempty" type:"Repeated"`
	AllowedOrigin []*string `json:"AllowedOrigin,omitempty" xml:"AllowedOrigin,omitempty" type:"Repeated"`
	ExposeHeader  []*string `json:"ExposeHeader,omitempty" xml:"ExposeHeader,omitempty" type:"Repeated"`
	// example:
	//
	// 300
	MaxAgeSeconds *int64 `json:"MaxAgeSeconds,omitempty" xml:"MaxAgeSeconds,omitempty"`
}

func (s CORSRule) String() string {
	return dara.Prettify(s)
}

func (s CORSRule) GoString() string {
	return s.String()
}

func (s *CORSRule) GetAllowedHeader() []*string {
	return s.AllowedHeader
}

func (s *CORSRule) GetAllowedMethod() []*string {
	return s.AllowedMethod
}

func (s *CORSRule) GetAllowedOrigin() []*string {
	return s.AllowedOrigin
}

func (s *CORSRule) GetExposeHeader() []*string {
	return s.ExposeHeader
}

func (s *CORSRule) GetMaxAgeSeconds() *int64 {
	return s.MaxAgeSeconds
}

func (s *CORSRule) SetAllowedHeader(v []*string) *CORSRule {
	s.AllowedHeader = v
	return s
}

func (s *CORSRule) SetAllowedMethod(v []*string) *CORSRule {
	s.AllowedMethod = v
	return s
}

func (s *CORSRule) SetAllowedOrigin(v []*string) *CORSRule {
	s.AllowedOrigin = v
	return s
}

func (s *CORSRule) SetExposeHeader(v []*string) *CORSRule {
	s.ExposeHeader = v
	return s
}

func (s *CORSRule) SetMaxAgeSeconds(v int64) *CORSRule {
	s.MaxAgeSeconds = &v
	return s
}

func (s *CORSRule) Validate() error {
	return dara.Validate(s)
}
