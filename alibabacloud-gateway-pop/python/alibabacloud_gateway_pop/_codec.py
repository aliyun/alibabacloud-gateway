# Private replacements for darabonba string/array/map/encode/signature.
# SHA-256 uses the standard library. SM3 stays in _sm3 because Python has no SM3.
import binascii
import hashlib
import hmac
from typing import Any, Dict, List
from urllib.parse import quote

from ._sm3 import Sm3, hash_sm3


def hex_encode(raw: bytes) -> str:
    return binascii.b2a_hex(raw).decode('utf-8')


def hash_bytes(raw: bytes, signature_algorithm: str) -> bytes:
    if signature_algorithm is None:
        return None
    if 'HMAC-SHA256' in signature_algorithm or 'RSA-SHA256' in signature_algorithm:
        return hashlib.sha256(raw).digest()
    if 'HMAC-SM3' in signature_algorithm:
        return hash_sm3(raw)
    return None


def percent_encode(raw: str) -> str:
    # Same result as quote(raw, safe='~') and darabonba Url.percent_encode for strings.
    encoded = quote(raw, safe='', encoding='utf-8')
    return encoded.replace('+', '%20').replace('*', '%2A').replace('%7E', '~')


def hmac_sha256(string_to_sign: str, secret: str) -> bytes:
    return hmac_sha256_by_bytes(string_to_sign, secret.encode('utf-8'))


def hmac_sha256_by_bytes(string_to_sign: str, secret: bytes) -> bytes:
    return hmac.new(secret, string_to_sign.encode('utf-8'), hashlib.sha256).digest()


def hmac_sm3(string_to_sign: str, secret: str) -> bytes:
    return hmac_sm3_by_bytes(string_to_sign, secret.encode('utf-8'))


def hmac_sm3_by_bytes(string_to_sign: str, secret: bytes) -> bytes:
    return hmac.new(secret, string_to_sign.encode('utf-8'), Sm3).digest()


def sub_string(value: str, start: int, end: int) -> str:
    return value[start:end]


def replace(raw: str, old_str: str, new_str: str, count) -> str:
    return raw.replace(old_str, new_str, -1 if count is None else count)


def contains(value: str, substr: str) -> bool:
    return substr in value


def split(raw: str, separator: str, limit):
    return raw.split(separator, -1 if limit is None else limit - 1)


def to_lower(value: str) -> str:
    return value.lower()


def equals(expect: str, actual: str) -> bool:
    return expect == actual


def trim(value: str) -> str:
    return value.strip()


def has_prefix(value: str, prefix: str) -> bool:
    return value.startswith(prefix)


def key_set(raw: Dict[str, Any]) -> List[str]:
    return list(raw.keys())


def join(raw: List[str], separator: str) -> str:
    if raw and separator:
        return separator.join(raw)
    return None


def size(raw: List[str]) -> int:
    if raw:
        return len(raw)
    return 0


def asc_sort(raw: List[str]) -> List[str]:
    if raw:
        raw.sort()
    return raw


def append(array: Any, item: Any) -> None:
    if isinstance(array, list):
        array.append(item)


def list_contains(raw: List[str], value: str) -> bool:
    if raw:
        return raw.__contains__(value)
    return False
