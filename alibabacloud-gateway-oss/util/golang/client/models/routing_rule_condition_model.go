// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutingRuleCondition interface {
	dara.Model
	String() string
	GoString() string
	SetHttpErrorCodeReturnedEquals(v int64) *RoutingRuleCondition
	GetHttpErrorCodeReturnedEquals() *int64
	SetIncludeHeader(v []*RoutingRuleConditionIncludeHeader) *RoutingRuleCondition
	GetIncludeHeader() []*RoutingRuleConditionIncludeHeader
	SetKeyPrefixEquals(v string) *RoutingRuleCondition
	GetKeyPrefixEquals() *string
	SetKeySuffixEquals(v string) *RoutingRuleCondition
	GetKeySuffixEquals() *string
}

type RoutingRuleCondition struct {
	// example:
	//
	// 404
	HttpErrorCodeReturnedEquals *int64                               `json:"HttpErrorCodeReturnedEquals,omitempty" xml:"HttpErrorCodeReturnedEquals,omitempty"`
	IncludeHeader               []*RoutingRuleConditionIncludeHeader `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty" type:"Repeated"`
	// example:
	//
	// abc/
	KeyPrefixEquals *string `json:"KeyPrefixEquals,omitempty" xml:"KeyPrefixEquals,omitempty"`
	// example:
	//
	// .txt
	KeySuffixEquals *string `json:"KeySuffixEquals,omitempty" xml:"KeySuffixEquals,omitempty"`
}

func (s RoutingRuleCondition) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleCondition) GoString() string {
	return s.String()
}

func (s *RoutingRuleCondition) GetHttpErrorCodeReturnedEquals() *int64 {
	return s.HttpErrorCodeReturnedEquals
}

func (s *RoutingRuleCondition) GetIncludeHeader() []*RoutingRuleConditionIncludeHeader {
	return s.IncludeHeader
}

func (s *RoutingRuleCondition) GetKeyPrefixEquals() *string {
	return s.KeyPrefixEquals
}

func (s *RoutingRuleCondition) GetKeySuffixEquals() *string {
	return s.KeySuffixEquals
}

func (s *RoutingRuleCondition) SetHttpErrorCodeReturnedEquals(v int64) *RoutingRuleCondition {
	s.HttpErrorCodeReturnedEquals = &v
	return s
}

func (s *RoutingRuleCondition) SetIncludeHeader(v []*RoutingRuleConditionIncludeHeader) *RoutingRuleCondition {
	s.IncludeHeader = v
	return s
}

func (s *RoutingRuleCondition) SetKeyPrefixEquals(v string) *RoutingRuleCondition {
	s.KeyPrefixEquals = &v
	return s
}

func (s *RoutingRuleCondition) SetKeySuffixEquals(v string) *RoutingRuleCondition {
	s.KeySuffixEquals = &v
	return s
}

func (s *RoutingRuleCondition) Validate() error {
	if s.IncludeHeader != nil {
		for _, item := range s.IncludeHeader {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RoutingRuleConditionIncludeHeader struct {
	// example:
	//
	// -test-suffix
	EndsWith *string `json:"EndsWith,omitempty" xml:"EndsWith,omitempty"`
	// example:
	//
	// test-value
	Equals *string `json:"Equals,omitempty" xml:"Equals,omitempty"`
	// example:
	//
	// test-header
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-prefix-
	StartsWith *string `json:"StartsWith,omitempty" xml:"StartsWith,omitempty"`
}

func (s RoutingRuleConditionIncludeHeader) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleConditionIncludeHeader) GoString() string {
	return s.String()
}

func (s *RoutingRuleConditionIncludeHeader) GetEndsWith() *string {
	return s.EndsWith
}

func (s *RoutingRuleConditionIncludeHeader) GetEquals() *string {
	return s.Equals
}

func (s *RoutingRuleConditionIncludeHeader) GetKey() *string {
	return s.Key
}

func (s *RoutingRuleConditionIncludeHeader) GetStartsWith() *string {
	return s.StartsWith
}

func (s *RoutingRuleConditionIncludeHeader) SetEndsWith(v string) *RoutingRuleConditionIncludeHeader {
	s.EndsWith = &v
	return s
}

func (s *RoutingRuleConditionIncludeHeader) SetEquals(v string) *RoutingRuleConditionIncludeHeader {
	s.Equals = &v
	return s
}

func (s *RoutingRuleConditionIncludeHeader) SetKey(v string) *RoutingRuleConditionIncludeHeader {
	s.Key = &v
	return s
}

func (s *RoutingRuleConditionIncludeHeader) SetStartsWith(v string) *RoutingRuleConditionIncludeHeader {
	s.StartsWith = &v
	return s
}

func (s *RoutingRuleConditionIncludeHeader) Validate() error {
	return dara.Validate(s)
}
