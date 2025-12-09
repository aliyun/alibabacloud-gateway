// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCORSConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCORSRule(v []*CORSRule) *CORSConfiguration
	GetCORSRule() []*CORSRule
	SetResponseVary(v bool) *CORSConfiguration
	GetResponseVary() *bool
}

type CORSConfiguration struct {
	CORSRule     []*CORSRule `json:"CORSRule,omitempty" xml:"CORSRule,omitempty" type:"Repeated"`
	ResponseVary *bool       `json:"ResponseVary,omitempty" xml:"ResponseVary,omitempty"`
}

func (s CORSConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CORSConfiguration) GoString() string {
	return s.String()
}

func (s *CORSConfiguration) GetCORSRule() []*CORSRule {
	return s.CORSRule
}

func (s *CORSConfiguration) GetResponseVary() *bool {
	return s.ResponseVary
}

func (s *CORSConfiguration) SetCORSRule(v []*CORSRule) *CORSConfiguration {
	s.CORSRule = v
	return s
}

func (s *CORSConfiguration) SetResponseVary(v bool) *CORSConfiguration {
	s.ResponseVary = &v
	return s
}

func (s *CORSConfiguration) Validate() error {
	if s.CORSRule != nil {
		for _, item := range s.CORSRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
