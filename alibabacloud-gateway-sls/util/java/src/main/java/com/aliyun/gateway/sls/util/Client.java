// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import com.aliyun.gateway.sls.util.model.ILog;
import com.aliyun.gateway.sls.util.model.LogsResponseBody;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.Map;

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

    public static byte[] serializeToPbBytes(Object request) {
        if (request instanceof ILog) {
            return ((ILog) request).serializeToPbBytes();
        }
        return null;
    }

    public static LogsResponseBody deserializeFromPbBytes(byte[] uncompressedData, Integer statusCode, Map<String, String> headers) {
        return new LogsResponseBody(uncompressedData, statusCode, headers);
    }

    public static int[] decodeVarInt32(byte[] dataBytes, int pos, int maxPos) {
        int value[] = {0, 0, 0};
        int shift = 0;
        int b;
        for (int i = pos; i < maxPos; ++i) {
            b = dataBytes[i] & 0xff;
            value[1] |= (b & 127) << shift;
            shift += 7;
            if ((b & 128) == 0) {
                value[2] = i + 1;
                value[0] = 1;
                break;
            }
        }
        return value;
    }
}
