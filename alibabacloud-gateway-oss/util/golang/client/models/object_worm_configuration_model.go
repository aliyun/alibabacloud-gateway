// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectWormConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetObjectWormEnabled(v string) *ObjectWormConfiguration
	GetObjectWormEnabled() *string
	SetRule(v *ObjectWormConfigurationRule) *ObjectWormConfiguration
	GetRule() *ObjectWormConfigurationRule
}

type ObjectWormConfiguration struct {
	ObjectWormEnabled *string                      `json:"ObjectWormEnabled,omitempty" xml:"ObjectWormEnabled,omitempty"`
	Rule              *ObjectWormConfigurationRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s ObjectWormConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ObjectWormConfiguration) GoString() string {
	return s.String()
}

func (s *ObjectWormConfiguration) GetObjectWormEnabled() *string {
	return s.ObjectWormEnabled
}

func (s *ObjectWormConfiguration) GetRule() *ObjectWormConfigurationRule {
	return s.Rule
}

func (s *ObjectWormConfiguration) SetObjectWormEnabled(v string) *ObjectWormConfiguration {
	s.ObjectWormEnabled = &v
	return s
}

func (s *ObjectWormConfiguration) SetRule(v *ObjectWormConfigurationRule) *ObjectWormConfiguration {
	s.Rule = v
	return s
}

func (s *ObjectWormConfiguration) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObjectWormConfigurationRule struct {
	DefaultRetention []*ObjectWormConfigurationRuleDefaultRetention `json:"DefaultRetention,omitempty" xml:"DefaultRetention,omitempty" type:"Repeated"`
}

func (s ObjectWormConfigurationRule) String() string {
	return dara.Prettify(s)
}

func (s ObjectWormConfigurationRule) GoString() string {
	return s.String()
}

func (s *ObjectWormConfigurationRule) GetDefaultRetention() []*ObjectWormConfigurationRuleDefaultRetention {
	return s.DefaultRetention
}

func (s *ObjectWormConfigurationRule) SetDefaultRetention(v []*ObjectWormConfigurationRuleDefaultRetention) *ObjectWormConfigurationRule {
	s.DefaultRetention = v
	return s
}

func (s *ObjectWormConfigurationRule) Validate() error {
	if s.DefaultRetention != nil {
		for _, item := range s.DefaultRetention {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ObjectWormConfigurationRuleDefaultRetention struct {
	Days  *int64  `json:"Days,omitempty" xml:"Days,omitempty"`
	Mode  *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Years *int64  `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s ObjectWormConfigurationRuleDefaultRetention) String() string {
	return dara.Prettify(s)
}

func (s ObjectWormConfigurationRuleDefaultRetention) GoString() string {
	return s.String()
}

func (s *ObjectWormConfigurationRuleDefaultRetention) GetDays() *int64 {
	return s.Days
}

func (s *ObjectWormConfigurationRuleDefaultRetention) GetMode() *string {
	return s.Mode
}

func (s *ObjectWormConfigurationRuleDefaultRetention) GetYears() *int64 {
	return s.Years
}

func (s *ObjectWormConfigurationRuleDefaultRetention) SetDays(v int64) *ObjectWormConfigurationRuleDefaultRetention {
	s.Days = &v
	return s
}

func (s *ObjectWormConfigurationRuleDefaultRetention) SetMode(v string) *ObjectWormConfigurationRuleDefaultRetention {
	s.Mode = &v
	return s
}

func (s *ObjectWormConfigurationRuleDefaultRetention) SetYears(v int64) *ObjectWormConfigurationRuleDefaultRetention {
	s.Years = &v
	return s
}

func (s *ObjectWormConfigurationRuleDefaultRetention) Validate() error {
	return dara.Validate(s)
}
