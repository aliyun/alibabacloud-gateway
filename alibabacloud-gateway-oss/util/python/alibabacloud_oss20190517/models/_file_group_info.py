# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class FileGroupInfo(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        etag: str = None,
        file_length: int = None,
        file_part: main_models.FileGroupInfoFilePart = None,
        key: str = None,
    ):
        self.bucket = bucket
        self.etag = etag
        self.file_length = file_length
        # FileGroup类型文件的信息
        self.file_part = file_part
        self.key = key

    def validate(self):
        if self.file_part:
            self.file_part.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.etag is not None:
            result['ETag'] = self.etag

        if self.file_length is not None:
            result['FileLength'] = self.file_length

        if self.file_part is not None:
            result['FilePart'] = self.file_part.to_map()

        if self.key is not None:
            result['Key'] = self.key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('FileLength') is not None:
            self.file_length = m.get('FileLength')

        if m.get('FilePart') is not None:
            temp_model = main_models.FileGroupInfoFilePart()
            self.file_part = temp_model.from_map(m.get('FilePart'))

        if m.get('Key') is not None:
            self.key = m.get('Key')

        return self

class FileGroupInfoFilePart(DaraModel):
    def __init__(
        self,
        part: List[main_models.FileGroupInfoFilePartPart] = None,
    ):
        self.part = part

    def validate(self):
        if self.part:
            for v1 in self.part:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Part'] = []
        if self.part is not None:
            for k1 in self.part:
                result['Part'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.part = []
        if m.get('Part') is not None:
            for k1 in m.get('Part'):
                temp_model = main_models.FileGroupInfoFilePartPart()
                self.part.append(temp_model.from_map(k1))

        return self

class FileGroupInfoFilePartPart(DaraModel):
    def __init__(
        self,
        etag: str = None,
        part_name: str = None,
        part_number: int = None,
        part_size: int = None,
    ):
        self.etag = etag
        self.part_name = part_name
        self.part_number = part_number
        self.part_size = part_size

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.etag is not None:
            result['ETag'] = self.etag

        if self.part_name is not None:
            result['PartName'] = self.part_name

        if self.part_number is not None:
            result['PartNumber'] = self.part_number

        if self.part_size is not None:
            result['PartSize'] = self.part_size

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('PartName') is not None:
            self.part_name = m.get('PartName')

        if m.get('PartNumber') is not None:
            self.part_number = m.get('PartNumber')

        if m.get('PartSize') is not None:
            self.part_size = m.get('PartSize')

        return self

