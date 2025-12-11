# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class FunctionComputeConfiguration(DaraModel):
    def __init__(
        self,
        event: List[str] = None,
        filter: main_models.FunctionComputeConfigurationFilter = None,
        function: main_models.FunctionComputeConfigurationFunction = None,
        id: str = None,
    ):
        self.event = event
        self.filter = filter
        self.function = function
        self.id = id

    def validate(self):
        if self.filter:
            self.filter.validate()
        if self.function:
            self.function.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.event is not None:
            result['Event'] = self.event

        if self.filter is not None:
            result['Filter'] = self.filter.to_map()

        if self.function is not None:
            result['Function'] = self.function.to_map()

        if self.id is not None:
            result['ID'] = self.id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Event') is not None:
            self.event = m.get('Event')

        if m.get('Filter') is not None:
            temp_model = main_models.FunctionComputeConfigurationFilter()
            self.filter = temp_model.from_map(m.get('Filter'))

        if m.get('Function') is not None:
            temp_model = main_models.FunctionComputeConfigurationFunction()
            self.function = temp_model.from_map(m.get('Function'))

        if m.get('ID') is not None:
            self.id = m.get('ID')

        return self

class FunctionComputeConfigurationFunction(DaraModel):
    def __init__(
        self,
        arn: str = None,
        assume_role: str = None,
    ):
        self.arn = arn
        self.assume_role = assume_role

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.arn is not None:
            result['Arn'] = self.arn

        if self.assume_role is not None:
            result['AssumeRole'] = self.assume_role

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Arn') is not None:
            self.arn = m.get('Arn')

        if m.get('AssumeRole') is not None:
            self.assume_role = m.get('AssumeRole')

        return self

class FunctionComputeConfigurationFilter(DaraModel):
    def __init__(
        self,
        key: main_models.FunctionComputeConfigurationFilterKey = None,
    ):
        self.key = key

    def validate(self):
        if self.key:
            self.key.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.key is not None:
            result['Key'] = self.key.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            temp_model = main_models.FunctionComputeConfigurationFilterKey()
            self.key = temp_model.from_map(m.get('Key'))

        return self

class FunctionComputeConfigurationFilterKey(DaraModel):
    def __init__(
        self,
        prefix: str = None,
        suffix: str = None,
    ):
        self.prefix = prefix
        self.suffix = suffix

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.suffix is not None:
            result['Suffix'] = self.suffix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('Suffix') is not None:
            self.suffix = m.get('Suffix')

        return self

