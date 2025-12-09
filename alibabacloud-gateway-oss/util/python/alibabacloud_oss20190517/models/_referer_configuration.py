# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class RefererConfiguration(DaraModel):
    def __init__(
        self,
        allow_empty_referer: bool = None,
        allow_truncate_query_string: bool = None,
        referer_blacklist: main_models.RefererConfigurationRefererBlacklist = None,
        referer_list: main_models.RefererConfigurationRefererList = None,
        truncate_path: bool = None,
    ):
        # This parameter is required.
        self.allow_empty_referer = allow_empty_referer
        self.allow_truncate_query_string = allow_truncate_query_string
        self.referer_blacklist = referer_blacklist
        # This parameter is required.
        self.referer_list = referer_list
        self.truncate_path = truncate_path

    def validate(self):
        if self.referer_blacklist:
            self.referer_blacklist.validate()
        if self.referer_list:
            self.referer_list.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allow_empty_referer is not None:
            result['AllowEmptyReferer'] = self.allow_empty_referer

        if self.allow_truncate_query_string is not None:
            result['AllowTruncateQueryString'] = self.allow_truncate_query_string

        if self.referer_blacklist is not None:
            result['RefererBlacklist'] = self.referer_blacklist.to_map()

        if self.referer_list is not None:
            result['RefererList'] = self.referer_list.to_map()

        if self.truncate_path is not None:
            result['TruncatePath'] = self.truncate_path

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowEmptyReferer') is not None:
            self.allow_empty_referer = m.get('AllowEmptyReferer')

        if m.get('AllowTruncateQueryString') is not None:
            self.allow_truncate_query_string = m.get('AllowTruncateQueryString')

        if m.get('RefererBlacklist') is not None:
            temp_model = main_models.RefererConfigurationRefererBlacklist()
            self.referer_blacklist = temp_model.from_map(m.get('RefererBlacklist'))

        if m.get('RefererList') is not None:
            temp_model = main_models.RefererConfigurationRefererList()
            self.referer_list = temp_model.from_map(m.get('RefererList'))

        if m.get('TruncatePath') is not None:
            self.truncate_path = m.get('TruncatePath')

        return self

class RefererConfigurationRefererList(DaraModel):
    def __init__(
        self,
        referer: List[str] = None,
    ):
        self.referer = referer

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.referer is not None:
            result['Referer'] = self.referer

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Referer') is not None:
            self.referer = m.get('Referer')

        return self

class RefererConfigurationRefererBlacklist(DaraModel):
    def __init__(
        self,
        referer: List[str] = None,
    ):
        self.referer = referer

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.referer is not None:
            result['Referer'] = self.referer

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Referer') is not None:
            self.referer = m.get('Referer')

        return self

