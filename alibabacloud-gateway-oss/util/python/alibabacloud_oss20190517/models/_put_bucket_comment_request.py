# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketCommentRequest(DaraModel):
    def __init__(
        self,
        comment_configuration: main_models.CommentConfiguration = None,
    ):
        self.comment_configuration = comment_configuration

    def validate(self):
        if self.comment_configuration:
            self.comment_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.comment_configuration is not None:
            result['CommentConfiguration'] = self.comment_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CommentConfiguration') is not None:
            temp_model = main_models.CommentConfiguration()
            self.comment_configuration = temp_model.from_map(m.get('CommentConfiguration'))

        return self

