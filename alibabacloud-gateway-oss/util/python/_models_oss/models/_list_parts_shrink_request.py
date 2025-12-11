# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListPartsShrinkRequest(DaraModel):
    def __init__(
        self,
        encoding_type_shrink: str = None,
        max_parts: int = None,
        part_number_marker: int = None,
        upload_id: str = None,
    ):
        # The maximum number of parts that can be returned by OSS. 
        # 
        # Default value: 1000.
        # 
        # Maximum value: 1000.
        self.encoding_type_shrink = encoding_type_shrink
        # The maximum number of parts that can be returned by OSS.
        # 
        # Default value: 1000.
        # 
        # Maximum value: 1000.
        self.max_parts = max_parts
        # The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
        # 
        # By default, this parameter is left empty.
        self.part_number_marker = part_number_marker
        # The ID of the multipart upload task.
        # 
        # By default, this parameter is left empty.
        # 
        # This parameter is required.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.encoding_type_shrink is not None:
            result['encoding-type'] = self.encoding_type_shrink

        if self.max_parts is not None:
            result['max-parts'] = self.max_parts

        if self.part_number_marker is not None:
            result['part-number-marker'] = self.part_number_marker

        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('encoding-type') is not None:
            self.encoding_type_shrink = m.get('encoding-type')

        if m.get('max-parts') is not None:
            self.max_parts = m.get('max-parts')

        if m.get('part-number-marker') is not None:
            self.part_number_marker = m.get('part-number-marker')

        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self

