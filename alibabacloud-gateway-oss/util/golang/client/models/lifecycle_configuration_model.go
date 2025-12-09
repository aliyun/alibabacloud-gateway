// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLifecycleConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v []*LifecycleRule) *LifecycleConfiguration
	GetRule() []*LifecycleRule
}

type LifecycleConfiguration struct {
	Rule []*LifecycleRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s LifecycleConfiguration) String() string {
	return dara.Prettify(s)
}

func (s LifecycleConfiguration) GoString() string {
	return s.String()
}

func (s *LifecycleConfiguration) GetRule() []*LifecycleRule {
	return s.Rule
}

func (s *LifecycleConfiguration) SetRule(v []*LifecycleRule) *LifecycleConfiguration {
	s.Rule = v
	return s
}

func (s *LifecycleConfiguration) Validate() error {
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
