// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import com.aliyun.gateway.sls.util.model.ILog;
import com.aliyun.gateway.sls.util.model.LogsResponseBody;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.Map;

public class Client {

    public static InputStream readAndUncompressBlock(InputStream stream, String compressType, String bodyRawSize) throws Exception {
        byte[] rawData = com.aliyun.teautil.Common.readAsBytes(stream);
        int rawSize = Integer.parseInt(bodyRawSize);
        rawData = DecompressorFactory.getDecompressor(compressType).decompress(rawData, rawSize);
        String data = new String(rawData, "UTF-8");
        return new ByteArrayInputStream(data.getBytes());
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