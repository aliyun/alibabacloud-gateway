# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from typing import BinaryIO
import lz4.block

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
        if compress_type != "lz4":
            raise Exception(f"Unsupported compress type: {compress_type}")

        body_size = int(body_raw_size)
        compressed_data = stream[:body_size]

        try:
            decompressed_data = lz4.block.decompress(compressed_data, uncompressed_size=body_size)
        except RuntimeError as e:
            raise RuntimeError(f"Failed to decompress LZ4 block: {e}")

        return decompressed_data

    @staticmethod
    async def read_and_uncompress_block_async(
        stream: BinaryIO,
        compress_type: str,
        body_raw_size: str,
    ) -> BinaryIO:
        raise Exception('Un-implemented')
