import unittest
from io import BytesIO
from alibabacloud_gateway_sls_util.client import Client
import random

def generate_random_bytes(length: int) -> bytes:
    return bytes(random.randint(0, 255) for _ in range(length))

class TestClient(unittest.TestCase):

    def setUp(self):
        self.client = Client()

    def _test_data(self, data: bytes):
        for compress_type in ['gzip', 'lz4', 'zstd']:
            compressed_data = Client.compress(data, compress_type)
            stream = BytesIO(compressed_data)
            decompressed_stream = self.client.read_and_uncompress_block(stream, compress_type, str(len(data)))
            decompressed = decompressed_stream.read()
            self.assertEqual(decompressed, data)
            self.assertEqual(len(decompressed), len(data))

    def test_compress(self):
        self._test_data('hello你好？'.encode('utf-8'))
        for _ in range(100):
            data = generate_random_bytes(random.randint(0, 10000))
            self._test_data(data)

    def test_read_and_uncompress_block_zero_size(self):
        self._test_data(b'')
        
    def test_read_and_uncompress_data_corrupted(self):
        with self.assertRaises(Exception):
            self.client.read_and_uncompress_block(BytesIO(b"adsasdsaad"), "lz4", "5")
            
    def test_unsupported_decompress_type(self):
        lz4_data = Client.compress(b"adsasdsaad", "lz4")
        with self.assertRaises(Exception):
            self.client.read_and_uncompress_block(BytesIO(lz4_data), "invalid", str(len(lz4_data)))
            
    def test_unsupported_compress_type(self):
        with self.assertRaises(Exception):
            self.client.compress(b"adsadsa", "invalid")
    
    def test_supported_types(self):
        for compress_type in ['lz4', 'zstd']:
            self.assertTrue(self.client.is_compressor_available(compress_type))
        for compress_type in ['gzip', 'lz4', 'zstd']:
            self.assertTrue(self.client.is_decompressor_available(compress_type))


if __name__ == '__main__':
    unittest.main()
