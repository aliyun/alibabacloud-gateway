import datetime
import hashlib
import random
import threading
import time
from collections import defaultdict
from typing import Any, Dict
from urllib.parse import urlencode
from xml.etree import ElementTree

from Tea.model import TeaModel
from Tea.stream import READABLE, STREAM_CLASS
from darabonba.core import DaraCore
from darabonba.utils.stream import Stream as DaraStream

_process_start_time = int(time.time() * 1000)
_seq_id = 0


def is_unset(value):
    return DaraCore.is_null(value)


def empty(val):
    return not val


def equal_string(val1, val2):
    return val1 == val2


def equal_number(val1, val2):
    return val1 == val2


def default_string(real, default):
    return real if real is not None else default


def to_bytes(val):
    if isinstance(val, bytes):
        return val
    if isinstance(val, str):
        return val.encode(encoding='utf-8')
    return str(val).encode(encoding='utf-8')


def read_as_bytes(stream):
    return DaraStream.read_as_bytes(stream)


async def read_as_bytes_async(stream):
    return await DaraStream.read_as_bytes_async(stream)


def read_as_string(stream):
    return DaraStream.read_as_string(stream)


async def read_as_string_async(stream):
    return await DaraStream.read_as_string_async(stream)


def read_as_json(stream):
    return DaraStream.read_as_json(stream)


async def read_as_json_async(stream):
    return await DaraStream.read_as_json_async(stream)


def assert_as_bytes(value):
    if not isinstance(value, bytes):
        raise ValueError(f'{value} is not a bytes')
    return value


def assert_as_map(value):
    if not isinstance(value, dict):
        raise ValueError(f'{value} is not a dict')
    return value


def to_jsonstring(val):
    return DaraCore.to_json_string(val)


def get_nonce():
    global _seq_id
    thread_id = threading.get_ident()
    current_time = int(time.time() * 1000)
    seq = _seq_id
    _seq_id += 1
    rand_num = random.getrandbits(64)
    msg = f'{_process_start_time}-{thread_id}-{current_time}-{seq}-{rand_num}'
    digest = hashlib.md5()
    digest.update(msg.encode('utf-8'))
    return digest.hexdigest()


def is_4xx(code):
    return 400 <= code < 500


def is_5xx(code):
    return 500 <= code < 600


def get_timestamp():
    return datetime.datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')


def get_endpoint_rules(product, region_id, endpoint_type, network, suffix=None):
    product = product or ''
    network = network or ''
    if endpoint_type == 'regional':
        if region_id is None or region_id == '':
            raise RuntimeError('RegionId is empty, please set a valid RegionId')
        result = '<product><network>.<region_id>.aliyuncs.com'.replace('<region_id>', region_id)
    else:
        result = '<product><network>.aliyuncs.com'
    result = result.replace('<product>', product.lower())
    if network == '' or network == 'public':
        result = result.replace('<network>', '')
    else:
        result = result.replace('<network>', '-' + network)
    return result


def to_form(filter):
    result = {}
    if filter:
        _object_handler('', filter, result)
    return _to_form_string(result)


def _object_handler(key, value, out):
    if value is None:
        return
    if isinstance(value, dict):
        for item_key, item_value in value.items():
            _object_handler('%s.%s' % (key, item_key), item_value, out)
    elif isinstance(value, TeaModel):
        for item_key, item_value in value.to_map().items():
            _object_handler('%s.%s' % (key, item_key), item_value, out)
    elif isinstance(value, (list, tuple)):
        for index, item_value in enumerate(value):
            _object_handler('%s.%s' % (key, index + 1), item_value, out)
    else:
        if key.startswith('.'):
            key = key[1:]
        if isinstance(value, bytes):
            out[key] = str(value, encoding='utf-8')
        elif not isinstance(value, STREAM_CLASS):
            out[key] = str(value)


def parse_xml(body):
    if not body:
        return {}
    return _parse_xml_element(ElementTree.fromstring(body))


def _parse_xml_element(element):
    children = list(element)
    if children:
        grouped = defaultdict(list)
        for child in children:
            parsed = _parse_xml_element(child)
            for key, value in parsed.items():
                grouped[key].append(value)
        value = {key: items[0] if len(items) == 1 else items for key, items in grouped.items()}
    elif element.attrib:
        value = {}
    else:
        text = element.text.strip() if element.text else ''
        return {element.tag: text}
    if element.attrib:
        for key, attr in element.attrib.items():
            value['@' + key] = attr
    if element.text:
        text = element.text.strip()
        if text:
            value['#text'] = text
    return {element.tag: value}


def _to_form_string(val: Dict[str, Any]):
    if not val:
        return ''
    keys = sorted(list(val))
    dic = {key: val[key] for key in keys if not isinstance(val[key], READABLE)}
    return urlencode(dic)
