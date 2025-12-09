// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutingRuleLuaConfig interface {
	dara.Model
	String() string
	GoString() string
	SetScript(v string) *RoutingRuleLuaConfig
	GetScript() *string
}

type RoutingRuleLuaConfig struct {
	// example:
	//
	// test.lua
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s RoutingRuleLuaConfig) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleLuaConfig) GoString() string {
	return s.String()
}

func (s *RoutingRuleLuaConfig) GetScript() *string {
	return s.Script
}

func (s *RoutingRuleLuaConfig) SetScript(v string) *RoutingRuleLuaConfig {
	s.Script = &v
	return s
}

func (s *RoutingRuleLuaConfig) Validate() error {
	return dara.Validate(s)
}
