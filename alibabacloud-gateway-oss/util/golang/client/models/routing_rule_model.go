// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutingRule interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v *RoutingRuleCondition) *RoutingRule
	GetCondition() *RoutingRuleCondition
	SetLuaConfig(v *RoutingRuleLuaConfig) *RoutingRule
	GetLuaConfig() *RoutingRuleLuaConfig
	SetRedirect(v *RoutingRuleRedirect) *RoutingRule
	GetRedirect() *RoutingRuleRedirect
	SetRuleNumber(v int64) *RoutingRule
	GetRuleNumber() *int64
}

type RoutingRule struct {
	Condition *RoutingRuleCondition `json:"Condition,omitempty" xml:"Condition,omitempty"`
	LuaConfig *RoutingRuleLuaConfig `json:"LuaConfig,omitempty" xml:"LuaConfig,omitempty"`
	Redirect  *RoutingRuleRedirect  `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// example:
	//
	// 1
	RuleNumber *int64 `json:"RuleNumber,omitempty" xml:"RuleNumber,omitempty"`
}

func (s RoutingRule) String() string {
	return dara.Prettify(s)
}

func (s RoutingRule) GoString() string {
	return s.String()
}

func (s *RoutingRule) GetCondition() *RoutingRuleCondition {
	return s.Condition
}

func (s *RoutingRule) GetLuaConfig() *RoutingRuleLuaConfig {
	return s.LuaConfig
}

func (s *RoutingRule) GetRedirect() *RoutingRuleRedirect {
	return s.Redirect
}

func (s *RoutingRule) GetRuleNumber() *int64 {
	return s.RuleNumber
}

func (s *RoutingRule) SetCondition(v *RoutingRuleCondition) *RoutingRule {
	s.Condition = v
	return s
}

func (s *RoutingRule) SetLuaConfig(v *RoutingRuleLuaConfig) *RoutingRule {
	s.LuaConfig = v
	return s
}

func (s *RoutingRule) SetRedirect(v *RoutingRuleRedirect) *RoutingRule {
	s.Redirect = v
	return s
}

func (s *RoutingRule) SetRuleNumber(v int64) *RoutingRule {
	s.RuleNumber = &v
	return s
}

func (s *RoutingRule) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	if s.LuaConfig != nil {
		if err := s.LuaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Redirect != nil {
		if err := s.Redirect.Validate(); err != nil {
			return err
		}
	}
	return nil
}
