# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListAccessPointsForObjectProcessResponseBody(DaraModel):
    def __init__(
        self,
        list_access_points_for_object_process_result: main_models.ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult = None,
    ):
        # The container that stores information about the Object FC Access Points that are returned.
        self.list_access_points_for_object_process_result = list_access_points_for_object_process_result

    def validate(self):
        if self.list_access_points_for_object_process_result:
            self.list_access_points_for_object_process_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_access_points_for_object_process_result is not None:
            result['ListAccessPointsForObjectProcessResult'] = self.list_access_points_for_object_process_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAccessPointsForObjectProcessResult') is not None:
            temp_model = main_models.ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult()
            self.list_access_points_for_object_process_result = temp_model.from_map(m.get('ListAccessPointsForObjectProcessResult'))

        return self

class ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult(DaraModel):
    def __init__(
        self,
        access_points_for_object_process: main_models.ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess = None,
        account_id: str = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
    ):
        # The container that stores information about all Object FC Access Points.
        self.access_points_for_object_process = access_points_for_object_process
        # The UID of the Alibaba Cloud account to which the Object FC Access Points belong.
        self.account_id = account_id
        # Indicates whether the returned results are truncated. Valid values:
        # 
        # true: indicates that not all results are returned for the request.
        # 
        # false: indicates that all results are returned for the request.
        self.is_truncated = is_truncated
        # Indicates that this ListAccessPointsForObjectProcess request contains subsequent results. You need to set the NextContinuationToken element to continuation-token for subsequent results.
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.access_points_for_object_process:
            self.access_points_for_object_process.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_points_for_object_process is not None:
            result['AccessPointsForObjectProcess'] = self.access_points_for_object_process.to_map()

        if self.account_id is not None:
            result['AccountId'] = self.account_id

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointsForObjectProcess') is not None:
            temp_model = main_models.ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess()
            self.access_points_for_object_process = temp_model.from_map(m.get('AccessPointsForObjectProcess'))

        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')

        return self

class ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess(DaraModel):
    def __init__(
        self,
        access_point_for_object_process: List[main_models.ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess] = None,
    ):
        # The container that stores information about a single Object FC Access Point.
        self.access_point_for_object_process = access_point_for_object_process

    def validate(self):
        if self.access_point_for_object_process:
            for v1 in self.access_point_for_object_process:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['AccessPointForObjectProcess'] = []
        if self.access_point_for_object_process is not None:
            for k1 in self.access_point_for_object_process:
                result['AccessPointForObjectProcess'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.access_point_for_object_process = []
        if m.get('AccessPointForObjectProcess') is not None:
            for k1 in m.get('AccessPointForObjectProcess'):
                temp_model = main_models.ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess()
                self.access_point_for_object_process.append(temp_model.from_map(k1))

        return self

class ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess(DaraModel):
    def __init__(
        self,
        access_point_for_object_process_alias: str = None,
        access_point_name: str = None,
        access_point_name_for_object_process: str = None,
        allow_anonymous_access_for_object_process: str = None,
        status: str = None,
    ):
        # The alias of the Object FC Access Point.
        self.access_point_for_object_process_alias = access_point_for_object_process_alias
        # The name of the access point.
        self.access_point_name = access_point_name
        # The name of the Object FC Access Point.
        self.access_point_name_for_object_process = access_point_name_for_object_process
        # Whether allow anonymous user access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The status of the Object FC Access Point. Valid values:
        # 
        # enable: The Object FC Access Point is created.
        # 
        # disable: The Object FC Access Point is disabled.
        # 
        # creating: The Object FC Access Point is being created.
        # 
        # deleting: The Object FC Access Point is deleted.
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_for_object_process_alias is not None:
            result['AccessPointForObjectProcessAlias'] = self.access_point_for_object_process_alias

        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name

        if self.access_point_name_for_object_process is not None:
            result['AccessPointNameForObjectProcess'] = self.access_point_name_for_object_process

        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointForObjectProcessAlias') is not None:
            self.access_point_for_object_process_alias = m.get('AccessPointForObjectProcessAlias')

        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')

        if m.get('AccessPointNameForObjectProcess') is not None:
            self.access_point_name_for_object_process = m.get('AccessPointNameForObjectProcess')

        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self

