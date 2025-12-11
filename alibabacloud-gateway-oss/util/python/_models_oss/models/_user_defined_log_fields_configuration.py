# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class UserDefinedLogFieldsConfiguration(DaraModel):
    def __init__(
        self,
        header_set: main_models.UserDefinedLogFieldsConfigurationHeaderSet = None,
        param_set: main_models.UserDefinedLogFieldsConfigurationParamSet = None,
    ):
        self.header_set = header_set
        self.param_set = param_set

    def validate(self):
        if self.header_set:
            self.header_set.validate()
        if self.param_set:
            self.param_set.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.header_set is not None:
            result['HeaderSet'] = self.header_set.to_map()

        if self.param_set is not None:
            result['ParamSet'] = self.param_set.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HeaderSet') is not None:
            temp_model = main_models.UserDefinedLogFieldsConfigurationHeaderSet()
            self.header_set = temp_model.from_map(m.get('HeaderSet'))

        if m.get('ParamSet') is not None:
            temp_model = main_models.UserDefinedLogFieldsConfigurationParamSet()
            self.param_set = temp_model.from_map(m.get('ParamSet'))

        return self

class UserDefinedLogFieldsConfigurationParamSet(DaraModel):
    def __init__(
        self,
        parameter: List[str] = None,
    ):
        self.parameter = parameter

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.parameter is not None:
            result['parameter'] = self.parameter

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('parameter') is not None:
            self.parameter = m.get('parameter')

        return self

class UserDefinedLogFieldsConfigurationHeaderSet(DaraModel):
    def __init__(
        self,
        header: List[str] = None,
    ):
        self.header = header

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.header is not None:
            result['header'] = self.header

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('header') is not None:
            self.header = m.get('header')

        return self

