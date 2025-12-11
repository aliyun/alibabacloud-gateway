# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class MetaQueryRespAddress(DaraModel):
    def __init__(
        self,
        address_line: str = None,
        city: str = None,
        district: str = None,
        language: str = None,
        province: str = None,
        township: str = None,
    ):
        self.address_line = address_line
        self.city = city
        self.district = district
        self.language = language
        self.province = province
        self.township = township

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.address_line is not None:
            result['AddressLine'] = self.address_line

        if self.city is not None:
            result['City'] = self.city

        if self.district is not None:
            result['District'] = self.district

        if self.language is not None:
            result['Language'] = self.language

        if self.province is not None:
            result['Province'] = self.province

        if self.township is not None:
            result['Township'] = self.township

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AddressLine') is not None:
            self.address_line = m.get('AddressLine')

        if m.get('City') is not None:
            self.city = m.get('City')

        if m.get('District') is not None:
            self.district = m.get('District')

        if m.get('Language') is not None:
            self.language = m.get('Language')

        if m.get('Province') is not None:
            self.province = m.get('Province')

        if m.get('Township') is not None:
            self.township = m.get('Township')

        return self

