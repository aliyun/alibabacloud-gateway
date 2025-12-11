# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListCnameResponseBody(DaraModel):
    def __init__(
        self,
        list_cname_result: main_models.ListCnameResponseBodyListCnameResult = None,
    ):
        # The container that is used to query information about all CNAME records.
        self.list_cname_result = list_cname_result

    def validate(self):
        if self.list_cname_result:
            self.list_cname_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_cname_result is not None:
            result['ListCnameResult'] = self.list_cname_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListCnameResult') is not None:
            temp_model = main_models.ListCnameResponseBodyListCnameResult()
            self.list_cname_result = temp_model.from_map(m.get('ListCnameResult'))

        return self

class ListCnameResponseBodyListCnameResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        cname: List[main_models.CnameInfo] = None,
        owner: str = None,
    ):
        # The name of the bucket to which the CNAME records you want to query are mapped.
        self.bucket = bucket
        # The container that is used to store the information about all CNAME records.
        self.cname = cname
        # The name of the bucket owner.
        self.owner = owner

    def validate(self):
        if self.cname:
            for v1 in self.cname:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        result['Cname'] = []
        if self.cname is not None:
            for k1 in self.cname:
                result['Cname'].append(k1.to_map() if k1 else None)

        if self.owner is not None:
            result['Owner'] = self.owner

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        self.cname = []
        if m.get('Cname') is not None:
            for k1 in m.get('Cname'):
                temp_model = main_models.CnameInfo()
                self.cname.append(temp_model.from_map(k1))

        if m.get('Owner') is not None:
            self.owner = m.get('Owner')

        return self

