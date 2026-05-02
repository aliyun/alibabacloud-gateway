from typing import Any
from aliyun_log_fastpb import serialize_log_group, deserialize_log_group_list


def serialize_log_group_to_pb(log_group: Any) -> bytes:
    if log_group is None:
        raise ValueError("log_group is None")
    if not isinstance(log_group, dict):
        raise ValueError("log_group is not a dict")
    return serialize_log_group(log_group)


def deserialize_log_group_list_from_pb(data: bytes) -> dict:
    return deserialize_log_group_list(data)
