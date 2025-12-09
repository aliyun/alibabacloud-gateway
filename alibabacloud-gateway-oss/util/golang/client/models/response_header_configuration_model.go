// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResponseHeaderConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v []*ResponseHeaderConfigurationRule) *ResponseHeaderConfiguration
	GetRule() []*ResponseHeaderConfigurationRule
}

type ResponseHeaderConfiguration struct {
	Rule []*ResponseHeaderConfigurationRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s ResponseHeaderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ResponseHeaderConfiguration) GoString() string {
	return s.String()
}

func (s *ResponseHeaderConfiguration) GetRule() []*ResponseHeaderConfigurationRule {
	return s.Rule
}

func (s *ResponseHeaderConfiguration) SetRule(v []*ResponseHeaderConfigurationRule) *ResponseHeaderConfiguration {
	s.Rule = v
	return s
}

func (s *ResponseHeaderConfiguration) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResponseHeaderConfigurationRule struct {
	Filters     *ResponseHeaderConfigurationRuleFilters     `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	HideHeaders *ResponseHeaderConfigurationRuleHideHeaders `json:"HideHeaders,omitempty" xml:"HideHeaders,omitempty" type:"Struct"`
	// example:
	//
	// hiddenHeader
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ResponseHeaderConfigurationRule) String() string {
	return dara.Prettify(s)
}

func (s ResponseHeaderConfigurationRule) GoString() string {
	return s.String()
}

func (s *ResponseHeaderConfigurationRule) GetFilters() *ResponseHeaderConfigurationRuleFilters {
	return s.Filters
}

func (s *ResponseHeaderConfigurationRule) GetHideHeaders() *ResponseHeaderConfigurationRuleHideHeaders {
	return s.HideHeaders
}

func (s *ResponseHeaderConfigurationRule) GetName() *string {
	return s.Name
}

func (s *ResponseHeaderConfigurationRule) SetFilters(v *ResponseHeaderConfigurationRuleFilters) *ResponseHeaderConfigurationRule {
	s.Filters = v
	return s
}

func (s *ResponseHeaderConfigurationRule) SetHideHeaders(v *ResponseHeaderConfigurationRuleHideHeaders) *ResponseHeaderConfigurationRule {
	s.HideHeaders = v
	return s
}

func (s *ResponseHeaderConfigurationRule) SetName(v string) *ResponseHeaderConfigurationRule {
	s.Name = &v
	return s
}

func (s *ResponseHeaderConfigurationRule) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	if s.HideHeaders != nil {
		if err := s.HideHeaders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResponseHeaderConfigurationRuleFilters struct {
	Operation []*string `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Repeated"`
}

func (s ResponseHeaderConfigurationRuleFilters) String() string {
	return dara.Prettify(s)
}

func (s ResponseHeaderConfigurationRuleFilters) GoString() string {
	return s.String()
}

func (s *ResponseHeaderConfigurationRuleFilters) GetOperation() []*string {
	return s.Operation
}

func (s *ResponseHeaderConfigurationRuleFilters) SetOperation(v []*string) *ResponseHeaderConfigurationRuleFilters {
	s.Operation = v
	return s
}

func (s *ResponseHeaderConfigurationRuleFilters) Validate() error {
	return dara.Validate(s)
}

type ResponseHeaderConfigurationRuleHideHeaders struct {
	Header []*string `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
}

func (s ResponseHeaderConfigurationRuleHideHeaders) String() string {
	return dara.Prettify(s)
}

func (s ResponseHeaderConfigurationRuleHideHeaders) GoString() string {
	return s.String()
}

func (s *ResponseHeaderConfigurationRuleHideHeaders) GetHeader() []*string {
	return s.Header
}

func (s *ResponseHeaderConfigurationRuleHideHeaders) SetHeader(v []*string) *ResponseHeaderConfigurationRuleHideHeaders {
	s.Header = v
	return s
}

func (s *ResponseHeaderConfigurationRuleHideHeaders) Validate() error {
	return dara.Validate(s)
}
