# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from typing import BinaryIO
import lz4.block
import zlib

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
        if compress_type == "lz4":
            return Client._uncompress_lz4(stream, body_raw_size)
        if compress_type == "gzip":
            return Client._uncompress_gzip(stream, body_raw_size)
        
        raise Exception(f"Unsupported compress type: {compress_type}")
    
    @staticmethod
    def _uncompress_lz4(stream: BinaryIO, body_raw_size: str) -> BinaryIO:
        compressed_data = stream[:]
        try:
            decompressed_data = lz4.block.decompress(compressed_data, uncompressed_size=int(body_raw_size))
        except RuntimeError as e:
            raise RuntimeError(f"Failed to decompress LZ4 block: {e}")

        return decompressed_data

    @staticmethod
    def _uncompress_gzip(stream: BinaryIO, body_raw_size: str) -> BinaryIO:
        compressed_data = stream[:]
        try:
            decompressed_data = zlib.decompress(compressed_data)
        except RuntimeError as e:
            raise RuntimeError(f"Failed to decompress gzip block: {e}")
        if len(decompressed_data) != int(body_raw_size):
            raise RuntimeError(
                f"Failed to decompress gzip block: unexpected gzip body size {len(decompressed_data)}, expected {body_raw_size}")
        return decompressed_data

    @staticmethod
    async def read_and_uncompress_block_async(
        stream: BinaryIO,
        compress_type: str,
        body_raw_size: str,
    ) -> BinaryIO:
        raise Exception('Un-implemented')
