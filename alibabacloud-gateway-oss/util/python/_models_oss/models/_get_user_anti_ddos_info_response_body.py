# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetUserAntiDDosInfoResponseBody(DaraModel):
    def __init__(
        self,
        anti_ddoslist_configuration: main_models.GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration = None,
    ):
        # The container that stores the list of Anti-DDoS instances.
        self.anti_ddoslist_configuration = anti_ddoslist_configuration

    def validate(self):
        if self.anti_ddoslist_configuration:
            self.anti_ddoslist_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.anti_ddoslist_configuration is not None:
            result['AntiDDOSListConfiguration'] = self.anti_ddoslist_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AntiDDOSListConfiguration') is not None:
            temp_model = main_models.GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration()
            self.anti_ddoslist_configuration = temp_model.from_map(m.get('AntiDDOSListConfiguration'))

        return self

class GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration(DaraModel):
    def __init__(
        self,
        anti_ddosconfiguration: List[main_models.UserAntiDDOSInfo] = None,
    ):
        # The container that stores information about the Anti-DDoS instance.
        self.anti_ddosconfiguration = anti_ddosconfiguration

    def validate(self):
        if self.anti_ddosconfiguration:
            for v1 in self.anti_ddosconfiguration:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['AntiDDOSConfiguration'] = []
        if self.anti_ddosconfiguration is not None:
            for k1 in self.anti_ddosconfiguration:
                result['AntiDDOSConfiguration'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.anti_ddosconfiguration = []
        if m.get('AntiDDOSConfiguration') is not None:
            for k1 in m.get('AntiDDOSConfiguration'):
                temp_model = main_models.UserAntiDDOSInfo()
                self.anti_ddosconfiguration.append(temp_model.from_map(k1))

        return self

