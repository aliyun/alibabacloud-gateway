# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetMetaQueryStatusResponseBody(DaraModel):
    def __init__(
        self,
        meta_query_status: main_models.GetMetaQueryStatusResponseBodyMetaQueryStatus = None,
    ):
        # The container that stores the metadata information.
        self.meta_query_status = meta_query_status

    def validate(self):
        if self.meta_query_status:
            self.meta_query_status.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.meta_query_status is not None:
            result['MetaQueryStatus'] = self.meta_query_status.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MetaQueryStatus') is not None:
            temp_model = main_models.GetMetaQueryStatusResponseBodyMetaQueryStatus()
            self.meta_query_status = temp_model.from_map(m.get('MetaQueryStatus'))

        return self

class GetMetaQueryStatusResponseBodyMetaQueryStatus(DaraModel):
    def __init__(
        self,
        create_time: str = None,
        phase: str = None,
        state: str = None,
        update_time: str = None,
    ):
        # The time when the metadata index library was created. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
        self.create_time = create_time
        # The scan type. Valid values:
        # - FullScanning: Full scanning is in progress.
        # - IncrementalScanning: Incremental scanning is in progress.
        self.phase = phase
        # The status of the metadata index library. Valid values:
        # - Ready: The metadata index library is being prepared after it is created.
        # In this case, the metadata index library cannot be used to query data.
        # 
        # - Stop: The metadata index library is paused.
        # - Running: The metadata index library is running.
        # - Retrying: The metadata index library failed to be created and is being created again.
        # - Failed: The metadata index library failed to be created.
        # - Deleted: The metadata index library is deleted.
        self.state = state
        # The time when the metadata index library was updated. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
        self.update_time = update_time

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.phase is not None:
            result['Phase'] = self.phase

        if self.state is not None:
            result['State'] = self.state

        if self.update_time is not None:
            result['UpdateTime'] = self.update_time

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('Phase') is not None:
            self.phase = m.get('Phase')

        if m.get('State') is not None:
            self.state = m.get('State')

        if m.get('UpdateTime') is not None:
            self.update_time = m.get('UpdateTime')

        return self

