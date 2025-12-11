# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ObjectProcessConfiguration(DaraModel):
    def __init__(
        self,
        allowed_features: main_models.ObjectProcessConfigurationAllowedFeatures = None,
        transformation_configurations: main_models.ObjectProcessConfigurationTransformationConfigurations = None,
    ):
        self.allowed_features = allowed_features
        self.transformation_configurations = transformation_configurations

    def validate(self):
        if self.allowed_features:
            self.allowed_features.validate()
        if self.transformation_configurations:
            self.transformation_configurations.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allowed_features is not None:
            result['AllowedFeatures'] = self.allowed_features.to_map()

        if self.transformation_configurations is not None:
            result['TransformationConfigurations'] = self.transformation_configurations.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowedFeatures') is not None:
            temp_model = main_models.ObjectProcessConfigurationAllowedFeatures()
            self.allowed_features = temp_model.from_map(m.get('AllowedFeatures'))

        if m.get('TransformationConfigurations') is not None:
            temp_model = main_models.ObjectProcessConfigurationTransformationConfigurations()
            self.transformation_configurations = temp_model.from_map(m.get('TransformationConfigurations'))

        return self

class ObjectProcessConfigurationTransformationConfigurations(DaraModel):
    def __init__(
        self,
        transformation_configuration: List[main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration] = None,
    ):
        self.transformation_configuration = transformation_configuration

    def validate(self):
        if self.transformation_configuration:
            for v1 in self.transformation_configuration:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['TransformationConfiguration'] = []
        if self.transformation_configuration is not None:
            for k1 in self.transformation_configuration:
                result['TransformationConfiguration'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.transformation_configuration = []
        if m.get('TransformationConfiguration') is not None:
            for k1 in m.get('TransformationConfiguration'):
                temp_model = main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration()
                self.transformation_configuration.append(temp_model.from_map(k1))

        return self

class ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration(DaraModel):
    def __init__(
        self,
        actions: main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions = None,
        content_transformation: main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation = None,
    ):
        self.actions = actions
        self.content_transformation = content_transformation

    def validate(self):
        if self.actions:
            self.actions.validate()
        if self.content_transformation:
            self.content_transformation.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.actions is not None:
            result['Actions'] = self.actions.to_map()

        if self.content_transformation is not None:
            result['ContentTransformation'] = self.content_transformation.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Actions') is not None:
            temp_model = main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions()
            self.actions = temp_model.from_map(m.get('Actions'))

        if m.get('ContentTransformation') is not None:
            temp_model = main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation()
            self.content_transformation = temp_model.from_map(m.get('ContentTransformation'))

        return self

class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation(DaraModel):
    def __init__(
        self,
        additional_features: main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures = None,
        function_compute: main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute = None,
    ):
        self.additional_features = additional_features
        self.function_compute = function_compute

    def validate(self):
        if self.additional_features:
            self.additional_features.validate()
        if self.function_compute:
            self.function_compute.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.additional_features is not None:
            result['AdditionalFeatures'] = self.additional_features.to_map()

        if self.function_compute is not None:
            result['FunctionCompute'] = self.function_compute.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AdditionalFeatures') is not None:
            temp_model = main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures()
            self.additional_features = temp_model.from_map(m.get('AdditionalFeatures'))

        if m.get('FunctionCompute') is not None:
            temp_model = main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute()
            self.function_compute = temp_model.from_map(m.get('FunctionCompute'))

        return self

class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute(DaraModel):
    def __init__(
        self,
        function_arn: str = None,
        function_assume_role_arn: str = None,
    ):
        self.function_arn = function_arn
        self.function_assume_role_arn = function_assume_role_arn

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.function_arn is not None:
            result['FunctionArn'] = self.function_arn

        if self.function_assume_role_arn is not None:
            result['FunctionAssumeRoleArn'] = self.function_assume_role_arn

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FunctionArn') is not None:
            self.function_arn = m.get('FunctionArn')

        if m.get('FunctionAssumeRoleArn') is not None:
            self.function_assume_role_arn = m.get('FunctionAssumeRoleArn')

        return self

class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures(DaraModel):
    def __init__(
        self,
        custom_forward_headers: main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders = None,
    ):
        self.custom_forward_headers = custom_forward_headers

    def validate(self):
        if self.custom_forward_headers:
            self.custom_forward_headers.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.custom_forward_headers is not None:
            result['CustomForwardHeaders'] = self.custom_forward_headers.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CustomForwardHeaders') is not None:
            temp_model = main_models.ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders()
            self.custom_forward_headers = temp_model.from_map(m.get('CustomForwardHeaders'))

        return self

class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders(DaraModel):
    def __init__(
        self,
        custom_forward_header: List[str] = None,
    ):
        self.custom_forward_header = custom_forward_header

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.custom_forward_header is not None:
            result['CustomForwardHeader'] = self.custom_forward_header

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CustomForwardHeader') is not None:
            self.custom_forward_header = m.get('CustomForwardHeader')

        return self

class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions(DaraModel):
    def __init__(
        self,
        action: List[str] = None,
    ):
        self.action = action

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.action is not None:
            result['Action'] = self.action

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')

        return self

class ObjectProcessConfigurationAllowedFeatures(DaraModel):
    def __init__(
        self,
        allowed_feature: List[str] = None,
    ):
        self.allowed_feature = allowed_feature

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allowed_feature is not None:
            result['AllowedFeature'] = self.allowed_feature

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowedFeature') is not None:
            self.allowed_feature = m.get('AllowedFeature')

        return self

