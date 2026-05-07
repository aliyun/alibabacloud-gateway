# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
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
        # Specifies whether to allow a request whose Referer field is empty. Valid values:
        # 
        # *   true (default)
        # *   false
        # 
        # This parameter is required.
        self.allow_empty_referer = allow_empty_referer
        # Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values:
        # 
        # *   true (default)
        # *   false
        self.allow_truncate_query_string = allow_truncate_query_string
        # The container that stores the Referer blacklist.
        self.referer_blacklist = referer_blacklist
        # The container that stores the Referer whitelist.
        # 
        # >  ****The PutBucketReferer operation overwrites the existing Referer whitelist with the Referer whitelist specified in RefererList. If RefererList is not specified in the request, which specifies that no Referer elements are included, the operation clears the existing Referer whitelist.
        # 
        # This parameter is required.
        self.referer_list = referer_list
        # Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values:
        # 
        # *   true
        # *   false
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
        # The addresses in the Referer whitelist.
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
        # The addresses in the Referer blacklist.
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

