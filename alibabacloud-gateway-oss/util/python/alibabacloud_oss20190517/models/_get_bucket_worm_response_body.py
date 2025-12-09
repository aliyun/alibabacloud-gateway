# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketWormResponseBody(DaraModel):
    def __init__(
        self,
        worm_configuration: main_models.GetBucketWormResponseBodyWormConfiguration = None,
    ):
        # The container that stores the information about retention policies of the bucket.
        self.worm_configuration = worm_configuration

    def validate(self):
        if self.worm_configuration:
            self.worm_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.worm_configuration is not None:
            result['WormConfiguration'] = self.worm_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('WormConfiguration') is not None:
            temp_model = main_models.GetBucketWormResponseBodyWormConfiguration()
            self.worm_configuration = temp_model.from_map(m.get('WormConfiguration'))

        return self

class GetBucketWormResponseBodyWormConfiguration(DaraModel):
    def __init__(
        self,
        creation_date: str = None,
        expiration_date: str = None,
        retention_period_in_days: int = None,
        state: str = None,
        worm_id: str = None,
    ):
        # The time at which the retention policy was created.
        self.creation_date = creation_date
        # The time at which the retention policy will be expired.
        self.expiration_date = expiration_date
        # The number of days for which objects can be retained.
        self.retention_period_in_days = retention_period_in_days
        # The status of the retention policy. Valid values:
        # 
        # - InProgress: indicates that the retention policy is in the InProgress state. By default, a retention policy is in the InProgress state after it is created. The policy remains in this state for 24 hours.
        # - Locked: indicates that the retention policy is in the Locked state.
        self.state = state
        # The ID of the retention policy.
        # 
        # >Note If the specified retention policy ID that is used to query the retention policy configurations of the bucket does not exist, OSS returns the 404 error code.
        self.worm_id = worm_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.expiration_date is not None:
            result['ExpirationDate'] = self.expiration_date

        if self.retention_period_in_days is not None:
            result['RetentionPeriodInDays'] = self.retention_period_in_days

        if self.state is not None:
            result['State'] = self.state

        if self.worm_id is not None:
            result['WormId'] = self.worm_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('ExpirationDate') is not None:
            self.expiration_date = m.get('ExpirationDate')

        if m.get('RetentionPeriodInDays') is not None:
            self.retention_period_in_days = m.get('RetentionPeriodInDays')

        if m.get('State') is not None:
            self.state = m.get('State')

        if m.get('WormId') is not None:
            self.worm_id = m.get('WormId')

        return self

