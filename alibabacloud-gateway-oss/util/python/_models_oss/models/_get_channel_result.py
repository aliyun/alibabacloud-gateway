# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetChannelResult(DaraModel):
    def __init__(
        self,
        auto_set_content_type: bool = None,
        create_time: str = None,
        default_404pic: str = None,
        last_modify_time: str = None,
        name: str = None,
        orig_pic_forbidden: bool = None,
        set_attach_name: bool = None,
        status: str = None,
        style_delimiters: str = None,
        use_src_format: bool = None,
        use_style_only: bool = None,
    ):
        self.auto_set_content_type = auto_set_content_type
        self.create_time = create_time
        self.default_404pic = default_404pic
        self.last_modify_time = last_modify_time
        self.name = name
        self.orig_pic_forbidden = orig_pic_forbidden
        self.set_attach_name = set_attach_name
        self.status = status
        self.style_delimiters = style_delimiters
        self.use_src_format = use_src_format
        self.use_style_only = use_style_only

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.auto_set_content_type is not None:
            result['AutoSetContentType'] = self.auto_set_content_type

        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.default_404pic is not None:
            result['Default404Pic'] = self.default_404pic

        if self.last_modify_time is not None:
            result['LastModifyTime'] = self.last_modify_time

        if self.name is not None:
            result['Name'] = self.name

        if self.orig_pic_forbidden is not None:
            result['OrigPicForbidden'] = self.orig_pic_forbidden

        if self.set_attach_name is not None:
            result['SetAttachName'] = self.set_attach_name

        if self.status is not None:
            result['Status'] = self.status

        if self.style_delimiters is not None:
            result['StyleDelimiters'] = self.style_delimiters

        if self.use_src_format is not None:
            result['UseSrcFormat'] = self.use_src_format

        if self.use_style_only is not None:
            result['UseStyleOnly'] = self.use_style_only

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AutoSetContentType') is not None:
            self.auto_set_content_type = m.get('AutoSetContentType')

        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('Default404Pic') is not None:
            self.default_404pic = m.get('Default404Pic')

        if m.get('LastModifyTime') is not None:
            self.last_modify_time = m.get('LastModifyTime')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('OrigPicForbidden') is not None:
            self.orig_pic_forbidden = m.get('OrigPicForbidden')

        if m.get('SetAttachName') is not None:
            self.set_attach_name = m.get('SetAttachName')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('StyleDelimiters') is not None:
            self.style_delimiters = m.get('StyleDelimiters')

        if m.get('UseSrcFormat') is not None:
            self.use_src_format = m.get('UseSrcFormat')

        if m.get('UseStyleOnly') is not None:
            self.use_style_only = m.get('UseStyleOnly')

        return self

