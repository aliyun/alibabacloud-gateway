// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionComputeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetEvent(v []*string) *FunctionComputeConfiguration
	GetEvent() []*string
	SetFilter(v *FunctionComputeConfigurationFilter) *FunctionComputeConfiguration
	GetFilter() *FunctionComputeConfigurationFilter
	SetFunction(v *FunctionComputeConfigurationFunction) *FunctionComputeConfiguration
	GetFunction() *FunctionComputeConfigurationFunction
	SetID(v string) *FunctionComputeConfiguration
	GetID() *string
}

type FunctionComputeConfiguration struct {
	Event    []*string                             `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
	Filter   *FunctionComputeConfigurationFilter   `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	Function *FunctionComputeConfigurationFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s FunctionComputeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s FunctionComputeConfiguration) GoString() string {
	return s.String()
}

func (s *FunctionComputeConfiguration) GetEvent() []*string {
	return s.Event
}

func (s *FunctionComputeConfiguration) GetFilter() *FunctionComputeConfigurationFilter {
	return s.Filter
}

func (s *FunctionComputeConfiguration) GetFunction() *FunctionComputeConfigurationFunction {
	return s.Function
}

func (s *FunctionComputeConfiguration) GetID() *string {
	return s.ID
}

func (s *FunctionComputeConfiguration) SetEvent(v []*string) *FunctionComputeConfiguration {
	s.Event = v
	return s
}

func (s *FunctionComputeConfiguration) SetFilter(v *FunctionComputeConfigurationFilter) *FunctionComputeConfiguration {
	s.Filter = v
	return s
}

func (s *FunctionComputeConfiguration) SetFunction(v *FunctionComputeConfigurationFunction) *FunctionComputeConfiguration {
	s.Function = v
	return s
}

func (s *FunctionComputeConfiguration) SetID(v string) *FunctionComputeConfiguration {
	s.ID = &v
	return s
}

func (s *FunctionComputeConfiguration) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FunctionComputeConfigurationFilter struct {
	Key *FunctionComputeConfigurationFilterKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
}

func (s FunctionComputeConfigurationFilter) String() string {
	return dara.Prettify(s)
}

func (s FunctionComputeConfigurationFilter) GoString() string {
	return s.String()
}

func (s *FunctionComputeConfigurationFilter) GetKey() *FunctionComputeConfigurationFilterKey {
	return s.Key
}

func (s *FunctionComputeConfigurationFilter) SetKey(v *FunctionComputeConfigurationFilterKey) *FunctionComputeConfigurationFilter {
	s.Key = v
	return s
}

func (s *FunctionComputeConfigurationFilter) Validate() error {
	if s.Key != nil {
		if err := s.Key.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FunctionComputeConfigurationFilterKey struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s FunctionComputeConfigurationFilterKey) String() string {
	return dara.Prettify(s)
}

func (s FunctionComputeConfigurationFilterKey) GoString() string {
	return s.String()
}

func (s *FunctionComputeConfigurationFilterKey) GetPrefix() *string {
	return s.Prefix
}

func (s *FunctionComputeConfigurationFilterKey) GetSuffix() *string {
	return s.Suffix
}

func (s *FunctionComputeConfigurationFilterKey) SetPrefix(v string) *FunctionComputeConfigurationFilterKey {
	s.Prefix = &v
	return s
}

func (s *FunctionComputeConfigurationFilterKey) SetSuffix(v string) *FunctionComputeConfigurationFilterKey {
	s.Suffix = &v
	return s
}

func (s *FunctionComputeConfigurationFilterKey) Validate() error {
	return dara.Validate(s)
}

type FunctionComputeConfigurationFunction struct {
	Arn        *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumeRole *string `json:"AssumeRole,omitempty" xml:"AssumeRole,omitempty"`
}

func (s FunctionComputeConfigurationFunction) String() string {
	return dara.Prettify(s)
}

func (s FunctionComputeConfigurationFunction) GoString() string {
	return s.String()
}

func (s *FunctionComputeConfigurationFunction) GetArn() *string {
	return s.Arn
}

func (s *FunctionComputeConfigurationFunction) GetAssumeRole() *string {
	return s.AssumeRole
}

func (s *FunctionComputeConfigurationFunction) SetArn(v string) *FunctionComputeConfigurationFunction {
	s.Arn = &v
	return s
}

func (s *FunctionComputeConfigurationFunction) SetAssumeRole(v string) *FunctionComputeConfigurationFunction {
	s.AssumeRole = &v
	return s
}

func (s *FunctionComputeConfigurationFunction) Validate() error {
	return dara.Validate(s)
}
