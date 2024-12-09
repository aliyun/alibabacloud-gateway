package com.aliyun.gateway.sls.util;

import java.io.ByteArrayInputStream;
import java.io.InputStream;

import org.junit.Assert;
import org.junit.Test;

public class CompressTest {

    private static final String[] compressTypes = new String[] { "zstd", "gzip", "lz4" };

    private void doCompress(byte[] data) throws Exception {
        int len = data.length;
        for (String compressType : compressTypes) {
            byte[] compressed = Client.compress(data, compressType);
            InputStream result = Client.readAndUncompressBlock(new ByteArrayInputStream(compressed), compressType,
                    String.valueOf(len));
            byte[] decompressed = com.aliyun.teautil.Common.readAsBytes(result);
            Assert.assertEquals(decompressed.length, len);
            Assert.assertArrayEquals(decompressed, data);
        }

    }

    private byte[] generateRandomBytes(int len) {
        byte[] data = new byte[len];
        for (int i = 0; i < len; i++) {
            data[i] = (byte) (Math.random() * 256);
        }
        return data;
    }

    @Test
    public void testCompress() throws Exception {
        doCompress(new byte[0]);
        doCompress("hello 你好".getBytes());
        for (int i = 0; i < 100; i++) {
            byte[] data = generateRandomBytes((int) (Math.random() * 100000));
            doCompress(data);
        }
    }

}
