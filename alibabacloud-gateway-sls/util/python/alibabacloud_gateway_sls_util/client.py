# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from typing import BinaryIO


class Client:
    """
    Read data from a readable stream, and parse it by JSON format
    @param stream the readable stream
    @return the parsed result
    """
    def __init__(self):
        pass

    @staticmethod
    def read_and_uncompress_block(
        stream: BinaryIO,
        compress_type: str,
        body_raw_size: str,
    ) -> BinaryIO:
        raise Exception('Un-implemented')

    @staticmethod
    async def read_and_uncompress_block_async(
        stream: BinaryIO,
        compress_type: str,
        body_raw_size: str,
    ) -> BinaryIO:
        raise Exception('Un-implemented')
