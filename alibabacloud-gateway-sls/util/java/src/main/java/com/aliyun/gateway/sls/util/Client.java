// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import java.io.ByteArrayInputStream;
import java.io.InputStream;

public class Client {

    public static InputStream readAndUncompressBlock(InputStream stream, String compressType, String bodyRawSize) throws Exception {
        byte[] rawData = com.aliyun.teautil.Common.readAsBytes(stream);
        int rawSize = Integer.parseInt(bodyRawSize);
        rawData = DecompressorFactory.getDecompressor(compressType).decompress(rawData, rawSize);
        String data = new String(rawData, "UTF-8");
        return new ByteArrayInputStream(data.getBytes());
    }
}