# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListPartsResponseBody(DaraModel):
    def __init__(
        self,
        list_part_result: main_models.ListPartsResponseBodyListPartResult = None,
    ):
        # The container that stores the response of the ListParts request.
        self.list_part_result = list_part_result

    def validate(self):
        if self.list_part_result:
            self.list_part_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_part_result is not None:
            result['ListPartResult'] = self.list_part_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListPartResult') is not None:
            temp_model = main_models.ListPartsResponseBodyListPartResult()
            self.list_part_result = temp_model.from_map(m.get('ListPartResult'))

        return self

class ListPartsResponseBodyListPartResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        is_truncated: bool = None,
        key: str = None,
        max_parts: int = None,
        next_part_number_marker: int = None,
        part: List[main_models.Part] = None,
        part_number_marker: int = None,
        upload_id: str = None,
    ):
        # The name of the bucket.
        self.bucket = bucket
        # Indicates whether the list of parts returned in the response has been truncated. A value of true indicates that the response does not contain all required results. A value of false indicates that the response contains all required results.
        # 
        # Valid values: true and false.
        self.is_truncated = is_truncated
        # The name of the object.
        self.key = key
        # The maximum number of parts in the response.
        self.max_parts = max_parts
        # The NextPartNumberMarker value that is used for the PartNumberMarker value in a subsequent request when the response does not contain all required results.
        self.next_part_number_marker = next_part_number_marker
        # The list of all parts.
        self.part = part
        # The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
        self.part_number_marker = part_number_marker
        # The ID of the upload task.
        self.upload_id = upload_id

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
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.key is not None:
            result['Key'] = self.key

        if self.max_parts is not None:
            result['MaxParts'] = self.max_parts

        if self.next_part_number_marker is not None:
            result['NextPartNumberMarker'] = self.next_part_number_marker

        result['Part'] = []
        if self.part is not None:
            for k1 in self.part:
                result['Part'].append(k1.to_map() if k1 else None)

        if self.part_number_marker is not None:
            result['PartNumberMarker'] = self.part_number_marker

        if self.upload_id is not None:
            result['UploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('MaxParts') is not None:
            self.max_parts = m.get('MaxParts')

        if m.get('NextPartNumberMarker') is not None:
            self.next_part_number_marker = m.get('NextPartNumberMarker')

        self.part = []
        if m.get('Part') is not None:
            for k1 in m.get('Part'):
                temp_model = main_models.Part()
                self.part.append(temp_model.from_map(k1))

        if m.get('PartNumberMarker') is not None:
            self.part_number_marker = m.get('PartNumberMarker')

        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')

        return self

