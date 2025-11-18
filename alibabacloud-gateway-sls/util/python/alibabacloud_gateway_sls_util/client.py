# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from typing import BinaryIO, List, Any
import lz4.block
import zlib
import zstd
import io
from .log_group_serializer import serialize_log_group_to_pb

# for compressor
supported_compress_types = ['lz4', 'zstd', 'gzip', 'deflate']


def lz4_compress(data: bytes) -> bytes:
    try:
        return lz4.block.compress(data)[4:]
    except Exception as e:
        raise Exception(
            f"fail to compress using compresstype lz4, error: {e}")


def zstd_compress(data: bytes) -> bytes:
    try:
        return zstd.compress(data, 1)
    except Exception as e:
        raise Exception(
            f"fail to compress using compresstype zstd, error: {e}")


def gzip_compress(data: bytes) -> bytes:
    try:
        return zlib.compress(data, 6)
    except Exception as e:
        raise Exception(
            f"fail to compress using compresstype gzip, error: {e}")


# for decompressor
supported_decompress_types = ['gzip', 'lz4', 'zstd', 'deflate']


def lz4_decompress(data: bytes, raw_size: int) -> bytes:
    try:
        return lz4.block.decompress(data, uncompressed_size=raw_size)
    except Exception as e:
        raise Exception(
            f"fail to decompress using compresstype lz4, error: {e}")


def zstd_decompress(data: bytes, raw_size: int) -> bytes:
    try:
        return zstd.decompress(data)
    except Exception as e:
        raise Exception(
            f"fail to decompress using compresstype zstd, error: {e}")


def gzip_decompress(data: bytes, raw_size: int) -> bytes:
    try:
        return zlib.decompress(data)
    except Exception as e:
        raise Exception(
            f"fail to decompress using compresstype gzip, error: {e}")


# common methods


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
        stream: BinaryIO, # actual is bytes
        compress_type: str,
        body_raw_size: str,
    ) -> BinaryIO:
        raw_size = int(body_raw_size)
        if raw_size == 0:
            return io.BytesIO(b'')
        data = stream
        
        decompressed = None
        if compress_type == 'lz4':
            decompressed = lz4_decompress(data, raw_size)
        elif compress_type == 'zstd':
            decompressed = zstd_decompress(data, raw_size)
        elif compress_type == 'gzip' or compress_type == 'deflate':
            decompressed = gzip_decompress(data, raw_size)
        else:
            raise Exception('unsupported decompress type: ' + compress_type)

        if len(decompressed) != raw_size:
            raise Exception('unexpected uncompressed size: %d, expected: %d, compressType: %s'.format(len(decompressed), raw_size, compress_type))
        
        return io.BytesIO(decompressed)

    @staticmethod
    async def read_and_uncompress_block_async(
        stream: BinaryIO,
        compress_type: str,
        body_raw_size: str,
    ) -> BinaryIO:
        raise Exception('Un-implemented')

    @staticmethod
    def compress(
        src: bytes,
        compress_type: str,
    ) -> bytes:
        """
        Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.
        @param src: the data to be compressed
        @param compress_type: the compress type
        @return: the compressed data
        @throws error if the compress type is not supported or the compress failed
        """
        if compress_type == 'zstd':
            return zstd_compress(src)
        if compress_type == 'lz4':
            return lz4_compress(src)
        if compress_type == 'gzip' or compress_type == 'deflate':
            return gzip_compress(src)
        raise Exception('unsupported compress type: ' + compress_type)

    @staticmethod
    async def compress_async(
        src: bytes,
        compress_type: str,
    ) -> bytes:
        """
        Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.
        @param src: the data to be compressed
        @param compress_type: the compress type
        @return: the compressed data
        @throws error if the compress type is not supported or the compress failed
        """
        raise Exception('Un-implemented')

    @staticmethod
    def is_compressor_available(
        compress_type: str,
    ) -> bool:
        return compress_type in supported_compress_types

    @staticmethod
    async def is_compressor_available_async(
        compress_type: str,
    ) -> bool:
        raise Exception('Un-implemented')

    @staticmethod
    def is_decompressor_available(
        compress_type: str,
    ) -> bool:
        return compress_type in supported_compress_types

    @staticmethod
    async def is_decompressor_available_async(
        compress_type: str,
    ) -> bool:
        raise Exception('Un-implemented')

    @staticmethod
    def bytes_length(
        src: bytes,
    ) -> int:
        return len(src)

    @staticmethod
    async def bytes_length_async(
        src: bytes,
    ) -> int:
        return len(src)

    @staticmethod
    def serialize_log_group_to_pb(
        log_group: Any,
    ) -> bytes:
        return serialize_log_group_to_pb(log_group)

    @staticmethod
    async def serialize_log_group_to_pb_async(
        log_group: Any,
    ) -> bytes:
        return serialize_log_group_to_pb(log_group)
