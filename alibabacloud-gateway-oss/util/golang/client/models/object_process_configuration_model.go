// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectProcessConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedFeatures(v *ObjectProcessConfigurationAllowedFeatures) *ObjectProcessConfiguration
	GetAllowedFeatures() *ObjectProcessConfigurationAllowedFeatures
	SetTransformationConfigurations(v *ObjectProcessConfigurationTransformationConfigurations) *ObjectProcessConfiguration
	GetTransformationConfigurations() *ObjectProcessConfigurationTransformationConfigurations
}

type ObjectProcessConfiguration struct {
	AllowedFeatures              *ObjectProcessConfigurationAllowedFeatures              `json:"AllowedFeatures,omitempty" xml:"AllowedFeatures,omitempty" type:"Struct"`
	TransformationConfigurations *ObjectProcessConfigurationTransformationConfigurations `json:"TransformationConfigurations,omitempty" xml:"TransformationConfigurations,omitempty" type:"Struct"`
}

func (s ObjectProcessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfiguration) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfiguration) GetAllowedFeatures() *ObjectProcessConfigurationAllowedFeatures {
	return s.AllowedFeatures
}

func (s *ObjectProcessConfiguration) GetTransformationConfigurations() *ObjectProcessConfigurationTransformationConfigurations {
	return s.TransformationConfigurations
}

func (s *ObjectProcessConfiguration) SetAllowedFeatures(v *ObjectProcessConfigurationAllowedFeatures) *ObjectProcessConfiguration {
	s.AllowedFeatures = v
	return s
}

func (s *ObjectProcessConfiguration) SetTransformationConfigurations(v *ObjectProcessConfigurationTransformationConfigurations) *ObjectProcessConfiguration {
	s.TransformationConfigurations = v
	return s
}

func (s *ObjectProcessConfiguration) Validate() error {
	if s.AllowedFeatures != nil {
		if err := s.AllowedFeatures.Validate(); err != nil {
			return err
		}
	}
	if s.TransformationConfigurations != nil {
		if err := s.TransformationConfigurations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObjectProcessConfigurationAllowedFeatures struct {
	AllowedFeature []*string `json:"AllowedFeature,omitempty" xml:"AllowedFeature,omitempty" type:"Repeated"`
}

func (s ObjectProcessConfigurationAllowedFeatures) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationAllowedFeatures) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationAllowedFeatures) GetAllowedFeature() []*string {
	return s.AllowedFeature
}

func (s *ObjectProcessConfigurationAllowedFeatures) SetAllowedFeature(v []*string) *ObjectProcessConfigurationAllowedFeatures {
	s.AllowedFeature = v
	return s
}

func (s *ObjectProcessConfigurationAllowedFeatures) Validate() error {
	return dara.Validate(s)
}

type ObjectProcessConfigurationTransformationConfigurations struct {
	TransformationConfiguration []*ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration `json:"TransformationConfiguration,omitempty" xml:"TransformationConfiguration,omitempty" type:"Repeated"`
}

func (s ObjectProcessConfigurationTransformationConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurations) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurations) GetTransformationConfiguration() []*ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration {
	return s.TransformationConfiguration
}

func (s *ObjectProcessConfigurationTransformationConfigurations) SetTransformationConfiguration(v []*ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) *ObjectProcessConfigurationTransformationConfigurations {
	s.TransformationConfiguration = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurations) Validate() error {
	if s.TransformationConfiguration != nil {
		for _, item := range s.TransformationConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration struct {
	Actions               *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions               `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	ContentTransformation *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation `json:"ContentTransformation,omitempty" xml:"ContentTransformation,omitempty" type:"Struct"`
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) GetActions() *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions {
	return s.Actions
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) GetContentTransformation() *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation {
	return s.ContentTransformation
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) SetActions(v *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration {
	s.Actions = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) SetContentTransformation(v *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration {
	s.ContentTransformation = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration) Validate() error {
	if s.Actions != nil {
		if err := s.Actions.Validate(); err != nil {
			return err
		}
	}
	if s.ContentTransformation != nil {
		if err := s.ContentTransformation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions struct {
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions) GetAction() []*string {
	return s.Action
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions) SetAction(v []*string) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions {
	s.Action = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions) Validate() error {
	return dara.Validate(s)
}

type ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation struct {
	AdditionalFeatures *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures `json:"AdditionalFeatures,omitempty" xml:"AdditionalFeatures,omitempty" type:"Struct"`
	FunctionCompute    *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute    `json:"FunctionCompute,omitempty" xml:"FunctionCompute,omitempty" type:"Struct"`
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) GetAdditionalFeatures() *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures {
	return s.AdditionalFeatures
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) GetFunctionCompute() *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute {
	return s.FunctionCompute
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) SetAdditionalFeatures(v *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation {
	s.AdditionalFeatures = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) SetFunctionCompute(v *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation {
	s.FunctionCompute = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation) Validate() error {
	if s.AdditionalFeatures != nil {
		if err := s.AdditionalFeatures.Validate(); err != nil {
			return err
		}
	}
	if s.FunctionCompute != nil {
		if err := s.FunctionCompute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures struct {
	CustomForwardHeaders *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders `json:"CustomForwardHeaders,omitempty" xml:"CustomForwardHeaders,omitempty" type:"Struct"`
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures) GetCustomForwardHeaders() *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders {
	return s.CustomForwardHeaders
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures) SetCustomForwardHeaders(v *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures {
	s.CustomForwardHeaders = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures) Validate() error {
	if s.CustomForwardHeaders != nil {
		if err := s.CustomForwardHeaders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders struct {
	CustomForwardHeader []*string `json:"CustomForwardHeader,omitempty" xml:"CustomForwardHeader,omitempty" type:"Repeated"`
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders) GetCustomForwardHeader() []*string {
	return s.CustomForwardHeader
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders) SetCustomForwardHeader(v []*string) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders {
	s.CustomForwardHeader = v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders) Validate() error {
	return dara.Validate(s)
}

type ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute struct {
	// example:
	//
	// acs:fc:cn-qingdao:111933544165****:services/test-oss-fc.LATEST/functions/fc-01
	FunctionArn *string `json:"FunctionArn,omitempty" xml:"FunctionArn,omitempty"`
	// example:
	//
	// acs:ram::111933544165****:role/aliyunfcdefaultrole
	FunctionAssumeRoleArn *string `json:"FunctionAssumeRoleArn,omitempty" xml:"FunctionAssumeRoleArn,omitempty"`
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) String() string {
	return dara.Prettify(s)
}

func (s ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) GoString() string {
	return s.String()
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) GetFunctionAssumeRoleArn() *string {
	return s.FunctionAssumeRoleArn
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) SetFunctionArn(v string) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute {
	s.FunctionArn = &v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) SetFunctionAssumeRoleArn(v string) *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute {
	s.FunctionAssumeRoleArn = &v
	return s
}

func (s *ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute) Validate() error {
	return dara.Validate(s)
}
