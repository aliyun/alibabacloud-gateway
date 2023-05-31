// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import java.io.ByteArrayInputStream;

public class Client {

    public static Object readAndUncompressBlock(java.io.InputStream stream, String compressType, String bodyRawSize) throws Exception {
        byte[] rawData = com.aliyun.teautil.Common.readAsBytes(stream);
        int rawSize = Integer.parseInt(bodyRawSize);
        rawData = DecompressorFactory.getDecompressor(compressType).decompress(rawData, rawSize);
        String data = new String(rawData, "UTF-8");
        return new ByteArrayInputStream(data.getBytes());
    }
}