// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.Map;

import com.aliyun.gateway.sls.util.model.PostLogStoreLogsRequest;
import com.aliyun.gateway.sls.util.model.PostLogStoreLogsResponse;

public class Client {

    public static InputStream readAndUncompressBlock(InputStream stream, String compressType, String bodyRawSize) throws Exception {
        byte[] rawData = com.aliyun.teautil.Common.readAsBytes(stream);
        int rawSize = Integer.parseInt(bodyRawSize);
        rawData = DecompressorFactory.getDecompressor(compressType).decompress(rawData, rawSize);
        String data = new String(rawData, "UTF-8");
        return new ByteArrayInputStream(data.getBytes());
    }
    public static byte[] SerializeToPbBytes(PostLogStoreLogsRequest request) throws Exception {
        return request.serializeToPbBytes();
    }

    public static PostLogStoreLogsResponse.PullLogsResponseBody DeserializeFromPbBytes(byte[] uncompressedData, int statusCode, Map<String, String> headers) throws Exception {
        return new PostLogStoreLogsResponse.PullLogsResponseBody(uncompressedData, statusCode, headers);
    }
}