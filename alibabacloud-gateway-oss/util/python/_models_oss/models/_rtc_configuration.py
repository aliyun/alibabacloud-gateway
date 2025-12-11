# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class RtcConfiguration(DaraModel):
    def __init__(
        self,
        id: str = None,
        rtc: main_models.RTC = None,
    ):
        self.id = id
        self.rtc = rtc

    def validate(self):
        if self.rtc:
            self.rtc.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.id is not None:
            result['ID'] = self.id

        if self.rtc is not None:
            result['RTC'] = self.rtc.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('RTC') is not None:
            temp_model = main_models.RTC()
            self.rtc = temp_model.from_map(m.get('RTC'))

        return self

