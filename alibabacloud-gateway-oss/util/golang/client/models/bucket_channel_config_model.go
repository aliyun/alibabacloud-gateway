// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketChannelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDebugInfo(v string) *BucketChannelConfig
	GetDebugInfo() *string
	SetRuleList(v *BucketChannelConfigRuleList) *BucketChannelConfig
	GetRuleList() *BucketChannelConfigRuleList
	SetVersion(v int32) *BucketChannelConfig
	GetVersion() *int32
}

type BucketChannelConfig struct {
	// example:
	//
	// testinfo
	DebugInfo *string                      `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	RuleList  *BucketChannelConfigRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// 2
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s BucketChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s BucketChannelConfig) GoString() string {
	return s.String()
}

func (s *BucketChannelConfig) GetDebugInfo() *string {
	return s.DebugInfo
}

func (s *BucketChannelConfig) GetRuleList() *BucketChannelConfigRuleList {
	return s.RuleList
}

func (s *BucketChannelConfig) GetVersion() *int32 {
	return s.Version
}

func (s *BucketChannelConfig) SetDebugInfo(v string) *BucketChannelConfig {
	s.DebugInfo = &v
	return s
}

func (s *BucketChannelConfig) SetRuleList(v *BucketChannelConfigRuleList) *BucketChannelConfig {
	s.RuleList = v
	return s
}

func (s *BucketChannelConfig) SetVersion(v int32) *BucketChannelConfig {
	s.Version = &v
	return s
}

func (s *BucketChannelConfig) Validate() error {
	if s.RuleList != nil {
		if err := s.RuleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BucketChannelConfigRuleList struct {
	Rule []*BucketChannelConfigRuleListRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s BucketChannelConfigRuleList) String() string {
	return dara.Prettify(s)
}

func (s BucketChannelConfigRuleList) GoString() string {
	return s.String()
}

func (s *BucketChannelConfigRuleList) GetRule() []*BucketChannelConfigRuleListRule {
	return s.Rule
}

func (s *BucketChannelConfigRuleList) SetRule(v []*BucketChannelConfigRuleListRule) *BucketChannelConfigRuleList {
	s.Rule = v
	return s
}

func (s *BucketChannelConfigRuleList) Validate() error {
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

type BucketChannelConfigRuleListRule struct {
	// example:
	//
	// a
	FrontContent *string `json:"FrontContent,omitempty" xml:"FrontContent,omitempty"`
	// example:
	//
	// rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// \\.webp$
	RuleRegex *string `json:"RuleRegex,omitempty" xml:"RuleRegex,omitempty"`
}

func (s BucketChannelConfigRuleListRule) String() string {
	return dara.Prettify(s)
}

func (s BucketChannelConfigRuleListRule) GoString() string {
	return s.String()
}

func (s *BucketChannelConfigRuleListRule) GetFrontContent() *string {
	return s.FrontContent
}

func (s *BucketChannelConfigRuleListRule) GetRuleName() *string {
	return s.RuleName
}

func (s *BucketChannelConfigRuleListRule) GetRuleRegex() *string {
	return s.RuleRegex
}

func (s *BucketChannelConfigRuleListRule) SetFrontContent(v string) *BucketChannelConfigRuleListRule {
	s.FrontContent = &v
	return s
}

func (s *BucketChannelConfigRuleListRule) SetRuleName(v string) *BucketChannelConfigRuleListRule {
	s.RuleName = &v
	return s
}

func (s *BucketChannelConfigRuleListRule) SetRuleRegex(v string) *BucketChannelConfigRuleListRule {
	s.RuleRegex = &v
	return s
}

func (s *BucketChannelConfigRuleListRule) Validate() error {
	return dara.Validate(s)
}
