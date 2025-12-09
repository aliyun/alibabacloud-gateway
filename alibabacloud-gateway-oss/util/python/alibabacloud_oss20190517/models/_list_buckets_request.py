# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListBucketsRequest(DaraModel):
    def __init__(
        self,
        marker: str = None,
        max_keys: int = None,
        prefix: str = None,
        tag_key: str = None,
        tag_value: str = None,
        tagging: str = None,
    ):
        # The name of the bucket from which the buckets start to return. The buckets whose names are alphabetically after the value of marker are returned. If this parameter is not specified, all results are returned. By default, this parameter is left empty.
        self.marker = marker
        # The maximum number of buckets that can be returned. Valid values: 1 to 1000. Default value: 100
        self.max_keys = max_keys
        # The prefix that the names of returned buckets must contain. If this parameter is not specified, prefixes are not used to filter returned buckets. By default, this parameter is left empty.
        self.prefix = prefix
        # A tag key of target buckets. The listing results will only include Buckets that have been tagged with this key.
        self.tag_key = tag_key
        # A tag value for the target buckets. If this parameter is specified in the request, the tag-key must also be specified. The listing results will only include Buckets that have been tagged with this key-value pair.
        self.tag_value = tag_value
        # Tag list of target buckets. Only Buckets that match all the key-value pairs in the list will added into the listing results. The tagging parameter cannot be used with the tag-key and tag-value parameters in a request.
        self.tagging = tagging

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.marker is not None:
            result['marker'] = self.marker

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        if self.prefix is not None:
            result['prefix'] = self.prefix

        if self.tag_key is not None:
            result['tag-key'] = self.tag_key

        if self.tag_value is not None:
            result['tag-value'] = self.tag_value

        if self.tagging is not None:
            result['tagging'] = self.tagging

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('marker') is not None:
            self.marker = m.get('marker')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')

        if m.get('tag-key') is not None:
            self.tag_key = m.get('tag-key')

        if m.get('tag-value') is not None:
            self.tag_value = m.get('tag-value')

        if m.get('tagging') is not None:
            self.tagging = m.get('tagging')

        return self

