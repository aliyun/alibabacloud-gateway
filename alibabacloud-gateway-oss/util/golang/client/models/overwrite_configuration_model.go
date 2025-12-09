// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOverwriteConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v []*OverwriteConfigurationRule) *OverwriteConfiguration
	GetRule() []*OverwriteConfigurationRule
}

type OverwriteConfiguration struct {
	Rule []*OverwriteConfigurationRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s OverwriteConfiguration) String() string {
	return dara.Prettify(s)
}

func (s OverwriteConfiguration) GoString() string {
	return s.String()
}

func (s *OverwriteConfiguration) GetRule() []*OverwriteConfigurationRule {
	return s.Rule
}

func (s *OverwriteConfiguration) SetRule(v []*OverwriteConfigurationRule) *OverwriteConfiguration {
	s.Rule = v
	return s
}

func (s *OverwriteConfiguration) Validate() error {
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

type OverwriteConfigurationRule struct {
	// example:
	//
	// forbid
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// forbid-write-rule1
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// example:
	//
	// a/
	Prefix     *string                               `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Principals *OverwriteConfigurationRulePrincipals `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Struct"`
	// example:
	//
	// .txt
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s OverwriteConfigurationRule) String() string {
	return dara.Prettify(s)
}

func (s OverwriteConfigurationRule) GoString() string {
	return s.String()
}

func (s *OverwriteConfigurationRule) GetAction() *string {
	return s.Action
}

func (s *OverwriteConfigurationRule) GetID() *string {
	return s.ID
}

func (s *OverwriteConfigurationRule) GetPrefix() *string {
	return s.Prefix
}

func (s *OverwriteConfigurationRule) GetPrincipals() *OverwriteConfigurationRulePrincipals {
	return s.Principals
}

func (s *OverwriteConfigurationRule) GetSuffix() *string {
	return s.Suffix
}

func (s *OverwriteConfigurationRule) SetAction(v string) *OverwriteConfigurationRule {
	s.Action = &v
	return s
}

func (s *OverwriteConfigurationRule) SetID(v string) *OverwriteConfigurationRule {
	s.ID = &v
	return s
}

func (s *OverwriteConfigurationRule) SetPrefix(v string) *OverwriteConfigurationRule {
	s.Prefix = &v
	return s
}

func (s *OverwriteConfigurationRule) SetPrincipals(v *OverwriteConfigurationRulePrincipals) *OverwriteConfigurationRule {
	s.Principals = v
	return s
}

func (s *OverwriteConfigurationRule) SetSuffix(v string) *OverwriteConfigurationRule {
	s.Suffix = &v
	return s
}

func (s *OverwriteConfigurationRule) Validate() error {
	if s.Principals != nil {
		if err := s.Principals.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OverwriteConfigurationRulePrincipals struct {
	Principal []*string `json:"Principal,omitempty" xml:"Principal,omitempty" type:"Repeated"`
}

func (s OverwriteConfigurationRulePrincipals) String() string {
	return dara.Prettify(s)
}

func (s OverwriteConfigurationRulePrincipals) GoString() string {
	return s.String()
}

func (s *OverwriteConfigurationRulePrincipals) GetPrincipal() []*string {
	return s.Principal
}

func (s *OverwriteConfigurationRulePrincipals) SetPrincipal(v []*string) *OverwriteConfigurationRulePrincipals {
	s.Principal = v
	return s
}

func (s *OverwriteConfigurationRulePrincipals) Validate() error {
	return dara.Validate(s)
}
