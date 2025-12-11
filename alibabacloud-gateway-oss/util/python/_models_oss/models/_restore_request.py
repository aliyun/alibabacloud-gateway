# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class RestoreRequest(DaraModel):
    def __init__(
        self,
        days: int = None,
        job_parameters: main_models.RestoreRequestJobParameters = None,
    ):
        self.days = days
        self.job_parameters = job_parameters

    def validate(self):
        if self.job_parameters:
            self.job_parameters.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.days is not None:
            result['Days'] = self.days

        if self.job_parameters is not None:
            result['JobParameters'] = self.job_parameters.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Days') is not None:
            self.days = m.get('Days')

        if m.get('JobParameters') is not None:
            temp_model = main_models.RestoreRequestJobParameters()
            self.job_parameters = temp_model.from_map(m.get('JobParameters'))

        return self

class RestoreRequestJobParameters(DaraModel):
    def __init__(
        self,
        tier: str = None,
    ):
        self.tier = tier

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.tier is not None:
            result['Tier'] = self.tier

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tier') is not None:
            self.tier = m.get('Tier')

        return self

