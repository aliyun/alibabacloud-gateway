# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketRedundancyTypeRequest(DaraModel):
    def __init__(
        self,
        data_redundancy_type_configuration: main_models.PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration = None,
    ):
        self.data_redundancy_type_configuration = data_redundancy_type_configuration

    def validate(self):
        if self.data_redundancy_type_configuration:
            self.data_redundancy_type_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_redundancy_type_configuration is not None:
            result['DataRedundancyTypeConfiguration'] = self.data_redundancy_type_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataRedundancyTypeConfiguration') is not None:
            temp_model = main_models.PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration()
            self.data_redundancy_type_configuration = temp_model.from_map(m.get('DataRedundancyTypeConfiguration'))

        return self

class PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration(DaraModel):
    def __init__(
        self,
        type: str = None,
    ):
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self

