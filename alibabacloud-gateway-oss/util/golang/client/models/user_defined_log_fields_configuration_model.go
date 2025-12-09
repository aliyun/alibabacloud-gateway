// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserDefinedLogFieldsConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetHeaderSet(v *UserDefinedLogFieldsConfigurationHeaderSet) *UserDefinedLogFieldsConfiguration
	GetHeaderSet() *UserDefinedLogFieldsConfigurationHeaderSet
	SetParamSet(v *UserDefinedLogFieldsConfigurationParamSet) *UserDefinedLogFieldsConfiguration
	GetParamSet() *UserDefinedLogFieldsConfigurationParamSet
}

type UserDefinedLogFieldsConfiguration struct {
	HeaderSet *UserDefinedLogFieldsConfigurationHeaderSet `json:"HeaderSet,omitempty" xml:"HeaderSet,omitempty" type:"Struct"`
	ParamSet  *UserDefinedLogFieldsConfigurationParamSet  `json:"ParamSet,omitempty" xml:"ParamSet,omitempty" type:"Struct"`
}

func (s UserDefinedLogFieldsConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UserDefinedLogFieldsConfiguration) GoString() string {
	return s.String()
}

func (s *UserDefinedLogFieldsConfiguration) GetHeaderSet() *UserDefinedLogFieldsConfigurationHeaderSet {
	return s.HeaderSet
}

func (s *UserDefinedLogFieldsConfiguration) GetParamSet() *UserDefinedLogFieldsConfigurationParamSet {
	return s.ParamSet
}

func (s *UserDefinedLogFieldsConfiguration) SetHeaderSet(v *UserDefinedLogFieldsConfigurationHeaderSet) *UserDefinedLogFieldsConfiguration {
	s.HeaderSet = v
	return s
}

func (s *UserDefinedLogFieldsConfiguration) SetParamSet(v *UserDefinedLogFieldsConfigurationParamSet) *UserDefinedLogFieldsConfiguration {
	s.ParamSet = v
	return s
}

func (s *UserDefinedLogFieldsConfiguration) Validate() error {
	if s.HeaderSet != nil {
		if err := s.HeaderSet.Validate(); err != nil {
			return err
		}
	}
	if s.ParamSet != nil {
		if err := s.ParamSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UserDefinedLogFieldsConfigurationHeaderSet struct {
	Header []*string `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
}

func (s UserDefinedLogFieldsConfigurationHeaderSet) String() string {
	return dara.Prettify(s)
}

func (s UserDefinedLogFieldsConfigurationHeaderSet) GoString() string {
	return s.String()
}

func (s *UserDefinedLogFieldsConfigurationHeaderSet) GetHeader() []*string {
	return s.Header
}

func (s *UserDefinedLogFieldsConfigurationHeaderSet) SetHeader(v []*string) *UserDefinedLogFieldsConfigurationHeaderSet {
	s.Header = v
	return s
}

func (s *UserDefinedLogFieldsConfigurationHeaderSet) Validate() error {
	return dara.Validate(s)
}

type UserDefinedLogFieldsConfigurationParamSet struct {
	// example:
	//
	// my-param
	Parameter []*string `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Repeated"`
}

func (s UserDefinedLogFieldsConfigurationParamSet) String() string {
	return dara.Prettify(s)
}

func (s UserDefinedLogFieldsConfigurationParamSet) GoString() string {
	return s.String()
}

func (s *UserDefinedLogFieldsConfigurationParamSet) GetParameter() []*string {
	return s.Parameter
}

func (s *UserDefinedLogFieldsConfigurationParamSet) SetParameter(v []*string) *UserDefinedLogFieldsConfigurationParamSet {
	s.Parameter = v
	return s
}

func (s *UserDefinedLogFieldsConfigurationParamSet) Validate() error {
	return dara.Validate(s)
}
