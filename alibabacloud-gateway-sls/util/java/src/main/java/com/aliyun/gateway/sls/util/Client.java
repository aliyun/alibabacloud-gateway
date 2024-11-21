// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import java.io.ByteArrayInputStream;
import java.io.InputStream;

public class Client {

    public static InputStream readAndUncompressBlock(InputStream stream, String compressType, String bodyRawSize)
            throws Exception {
        byte[] rawData = com.aliyun.teautil.Common.readAsBytes(stream);
        int rawSize = Integer.parseInt(bodyRawSize);
        if (rawSize == 0) {
            return new ByteArrayInputStream(new byte[0]);
        }
        rawData = DecompressorFactory.getDecompressor(compressType).decompress(rawData, rawSize);
        if (rawData.length != rawSize) {
            throw new RuntimeException(String.format("unexpected uncompressed size: %d, expected: %d, compressType: %s",
                    rawData.length, rawSize, compressType));
        }
        return new ByteArrayInputStream(rawData);
    }

    /**
     * <b>description</b> :
     * <p>
     * Compress data by specified compress type, use isCompressorAvailable to check
     * if the compress type is supported.
     * </p>
     *
     * @param src          the data to be compressed
     * @param compressType the compress type
     * @return the compressed data
     * @throws Exception if the compress type is not supported or the compress
     *                   failed
     */
    public static byte[] compress(byte[] src, String compressType) throws Exception {
        return CompressorFactory.getCompressor(compressType).compress(src);
    }

    public static Boolean isCompressorAvailable(String compressType) throws Exception {
        return CompressorFactory.isCompressorAvailable(compressType);
    }

    public static Boolean isDecompressorAvailable(String compressType) throws Exception {
        return DecompressorFactory.isDecompressorAvailable(compressType);
    }

    public static Long bytesLength(byte[] src) throws Exception {
        return (long) src.length;
    }
}
