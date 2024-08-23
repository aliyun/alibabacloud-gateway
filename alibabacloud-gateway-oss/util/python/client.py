# -*- coding: utf-8 -*-
# import alibabacloud_tea_xml.client

from Tea.model import TeaModel
from typing import List, Dict, BinaryIO
from xml.etree import ElementTree
from collections import defaultdict
import inspect
from typing_extensions import get_origin, get_args


class Owner(TeaModel):
    def __init__(
        self,
        display_name: str = None,
        id: str = None,
    ):
        self.display_name = display_name
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.display_name is not None:
            result['DisplayName'] = self.display_name
        if self.id is not None:
            result['ID'] = self.id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DisplayName') is not None:
            self.display_name = m.get('DisplayName')
        if m.get('ID') is not None:
            self.id = m.get('ID')
        return self


class CommonPrefix(TeaModel):
    def __init__(
        self,
        prefix: str = None,
    ):
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        return self


class ObjectSummary(TeaModel):
    def __init__(
        self,
        etag: str = None,
        key: str = None,
        last_modified: str = None,
        owner: Owner = None,
        resore_info: str = None,
        size: int = None,
        storage_class: str = None,
        type: str = None,
    ):
        self.etag = etag
        self.key = key
        self.last_modified = last_modified
        self.owner = owner
        self.resore_info = resore_info
        self.size = size
        self.storage_class = storage_class
        self.type = type

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.key is not None:
            result['Key'] = self.key
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        if self.resore_info is not None:
            result['ResoreInfo'] = self.resore_info
        if self.size is not None:
            result['Size'] = self.size
        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        if m.get('ResoreInfo') is not None:
            self.resore_info = m.get('ResoreInfo')
        if m.get('Size') is not None:
            self.size = m.get('Size')
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class ListObjectsResponseBodyListBucketResult(TeaModel):
    def __init__(
        self,
        common_prefixes: List[CommonPrefix] = None,
        contents: List[ObjectSummary] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: int = None,
        name: str = None,
        next_marker: str = None,
        prefix: str = None,
    ):
        # If delimiter is specified in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.common_prefixes = common_prefixes
        # The container that stores the metadata of the returned objects.
        self.contents = contents
        # The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.
        self.encoding_type = encoding_type
        # Indicates whether the returned list in the result is truncated. Valid values:
        # - true
        # - false
        self.is_truncated = is_truncated
        # The name of the object after which the GetBucket (ListObjects) operation begins.
        self.marker = marker
        # The maximum number of returned objects in the response.
        self.max_keys = max_keys
        # The name of the bucket.
        self.name = name
        # If not all results are returned, NextMarker is included in the response to indicate the value of marker in the next request.
        self.next_marker = next_marker
        # The prefix in the names of the returned objects.
        self.prefix = prefix

    def validate(self):
        if self.common_prefixes:
            for k in self.common_prefixes:
                if k:
                    k.validate()
        if self.contents:
            for k in self.contents:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k in self.common_prefixes:
                result['CommonPrefixes'].append(k.to_map() if k else None)
        result['Contents'] = []
        if self.contents is not None:
            for k in self.contents:
                result['Contents'].append(k.to_map() if k else None)
        if self.delimiter is not None:
            result['Delimiter'] = self.delimiter
        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.marker is not None:
            result['Marker'] = self.marker
        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys
        if self.name is not None:
            result['Name'] = self.name
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k in m.get('CommonPrefixes'):
                temp_model = CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k))
        self.contents = []
        if m.get('Contents') is not None:
            for k in m.get('Contents'):
                temp_model = ObjectSummary()
                self.contents.append(temp_model.from_map(k))
        if m.get('Delimiter') is not None:
            self.delimiter = m.get('Delimiter')
        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('Marker') is not None:
            self.marker = m.get('Marker')
        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        return self


class ListObjectsResponseBody(TeaModel):
    def __init__(
        self,
        list_bucket_result: ListObjectsResponseBodyListBucketResult = None,
    ):
        # The container that stores the information about the returned objects.
        self.list_bucket_result = list_bucket_result

    def validate(self):
        if self.list_bucket_result:
            self.list_bucket_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_bucket_result is not None:
            result['ListBucketResult'] = self.list_bucket_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketResult') is not None:
            temp_model = ListObjectsResponseBodyListBucketResult()
            self.list_bucket_result = temp_model.from_map(m['ListBucketResult'])
        return self


class ListObjectsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListObjectsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListObjectsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


typeRegistry = {}
typeRegistry["ListObjects"] = ListObjectsResponseBody

basic_instance = {}
basic_instance[str] = ''
basic_instance[int] = 0
basic_instance[bool] = False
basic_instance[Dict[str, str]] = {'':''}

def build_instance_from_model(model):
    if model in basic_instance:
        return basic_instance[model]
    sig = inspect.signature(model.__init__)
    params = {}
    for para_name, param in sig.parameters.items():
        if para_name == "self":
            continue
        origin_type = get_origin(param.annotation)
        if origin_type is not None and issubclass(origin_type, list):
            params[para_name] = [build_instance_from_model(get_args(param.annotation)[0])]
        else:
            params[para_name] = build_instance_from_model(param.annotation)
    return model(**params)


instanceRegistry = {}
instanceRegistry["ListObjects"] = build_instance_from_model(ListObjectsResponseBody)


def parse_xml_impl(t, m):
    d = {t.tag: {} if t.attrib else None}
    children = list(t)
    if children:
        m = m[t.tag]
        islist = isinstance(m, list)
        mc = m[0] if islist else m
        dd = defaultdict(list)
        for dc in [parse_xml_impl(c, mc) for c in children]:
            for k, v in dc.items():
                dd[k].append(v)
        d = {t.tag: {k: v[0] if len(v) == 1 and not isinstance(mc[k], list) else v for k, v in dd.items()}}

    if t.attrib:
        d[t.tag].update(('@' + k, v) for k, v in t.attrib.items())

    if t.text:
        text = t.text.strip()
        if children or t.attrib:
            if text:
                d[t.tag]['#text'] = text
        else:
            d[t.tag] = text
    return d


def parse_xml(body, model):
    return parse_xml_impl(ElementTree.fromstring(body), model)


class Client:
    @staticmethod
    def parse_xml(bodyStr: str, apiName: str):
        d = parse_xml(bodyStr, instanceRegistry[apiName].to_map())
        return typeRegistry[apiName]().from_map(d)
