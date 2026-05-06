# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LiveChannelSnapshot(DaraModel):
    def __init__(
        self,
        dest_bucket: str = None,
        interval: int = None,
        notify_topic: str = None,
        role_name: str = None,
    ):
        # The bucket that stores the results of high-frequency snapshot operations. The bucket must belong to the same owner as the current bucket.
        self.dest_bucket = dest_bucket
        # The interval of high-frequency snapshot operations. If no key frame (inline frame) exists within the interval, no snapshot is captured. Unit: seconds. Valid values: [1, 100].
        self.interval = interval
        # The Message Service (MNS) topic used to notify users of the results of high-frequency snapshot operations.
        self.notify_topic = notify_topic
        # The name of the role used to perform high-frequency snapshot operations. The role must have the write permissions on DestBucket and the permissions to send messages to NotifyTopic.
        self.role_name = role_name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.dest_bucket is not None:
            result['DestBucket'] = self.dest_bucket

        if self.interval is not None:
            result['Interval'] = self.interval

        if self.notify_topic is not None:
            result['NotifyTopic'] = self.notify_topic

        if self.role_name is not None:
            result['RoleName'] = self.role_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DestBucket') is not None:
            self.dest_bucket = m.get('DestBucket')

        if m.get('Interval') is not None:
            self.interval = m.get('Interval')

        if m.get('NotifyTopic') is not None:
            self.notify_topic = m.get('NotifyTopic')

        if m.get('RoleName') is not None:
            self.role_name = m.get('RoleName')

        return self

