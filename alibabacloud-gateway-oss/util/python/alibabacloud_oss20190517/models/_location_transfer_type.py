# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class LocationTransferType(DaraModel):
    def __init__(
        self,
        location: str = None,
        transfer_types: main_models.LocationTransferTypeTransferTypes = None,
    ):
        self.location = location
        self.transfer_types = transfer_types

    def validate(self):
        if self.transfer_types:
            self.transfer_types.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.location is not None:
            result['Location'] = self.location

        if self.transfer_types is not None:
            result['TransferTypes'] = self.transfer_types.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            self.location = m.get('Location')

        if m.get('TransferTypes') is not None:
            temp_model = main_models.LocationTransferTypeTransferTypes()
            self.transfer_types = temp_model.from_map(m.get('TransferTypes'))

        return self

class LocationTransferTypeTransferTypes(DaraModel):
    def __init__(
        self,
        type: List[str] = None,
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

