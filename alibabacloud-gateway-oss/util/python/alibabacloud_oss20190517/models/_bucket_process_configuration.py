# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BucketProcessConfiguration(DaraModel):
    def __init__(
        self,
        bucket_channel_config: main_models.BucketChannelConfig = None,
        complied_host: str = None,
        oss_domain_support_at_process: str = None,
        source_file_protect: str = None,
        source_file_protect_suffix: str = None,
        style_delimiters: str = None,
    ):
        self.bucket_channel_config = bucket_channel_config
        self.complied_host = complied_host
        self.oss_domain_support_at_process = oss_domain_support_at_process
        self.source_file_protect = source_file_protect
        self.source_file_protect_suffix = source_file_protect_suffix
        self.style_delimiters = style_delimiters

    def validate(self):
        if self.bucket_channel_config:
            self.bucket_channel_config.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_channel_config is not None:
            result['BucketChannelConfig'] = self.bucket_channel_config.to_map()

        if self.complied_host is not None:
            result['CompliedHost'] = self.complied_host

        if self.oss_domain_support_at_process is not None:
            result['OssDomainSupportAtProcess'] = self.oss_domain_support_at_process

        if self.source_file_protect is not None:
            result['SourceFileProtect'] = self.source_file_protect

        if self.source_file_protect_suffix is not None:
            result['SourceFileProtectSuffix'] = self.source_file_protect_suffix

        if self.style_delimiters is not None:
            result['StyleDelimiters'] = self.style_delimiters

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketChannelConfig') is not None:
            temp_model = main_models.BucketChannelConfig()
            self.bucket_channel_config = temp_model.from_map(m.get('BucketChannelConfig'))

        if m.get('CompliedHost') is not None:
            self.complied_host = m.get('CompliedHost')

        if m.get('OssDomainSupportAtProcess') is not None:
            self.oss_domain_support_at_process = m.get('OssDomainSupportAtProcess')

        if m.get('SourceFileProtect') is not None:
            self.source_file_protect = m.get('SourceFileProtect')

        if m.get('SourceFileProtectSuffix') is not None:
            self.source_file_protect_suffix = m.get('SourceFileProtectSuffix')

        if m.get('StyleDelimiters') is not None:
            self.style_delimiters = m.get('StyleDelimiters')

        return self

