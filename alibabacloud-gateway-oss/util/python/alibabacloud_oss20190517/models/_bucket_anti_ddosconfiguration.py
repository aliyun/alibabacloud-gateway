# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BucketAntiDDOSConfiguration(DaraModel):
    def __init__(
        self,
        cnames: main_models.BucketAntiDDOSConfigurationCnames = None,
    ):
        self.cnames = cnames

    def validate(self):
        if self.cnames:
            self.cnames.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cnames is not None:
            result['Cnames'] = self.cnames.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cnames') is not None:
            temp_model = main_models.BucketAntiDDOSConfigurationCnames()
            self.cnames = temp_model.from_map(m.get('Cnames'))

        return self

class BucketAntiDDOSConfigurationCnames(DaraModel):
    def __init__(
        self,
        domain: List[str] = None,
    ):
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.domain is not None:
            result['Domain'] = self.domain

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')

        return self

