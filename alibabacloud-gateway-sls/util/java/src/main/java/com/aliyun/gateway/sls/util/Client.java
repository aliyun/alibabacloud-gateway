// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls.util;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.net.InetAddress;
import java.util.ArrayList;
import java.util.HashMap;

public class Client {

    public static InputStream readAndUncompressBlock(InputStream stream, String compressType, String bodyRawSize) throws Exception {
        byte[] rawData = com.aliyun.teautil.Common.readAsBytes(stream);
        int rawSize = Integer.parseInt(bodyRawSize);
        rawData = DecompressorFactory.getDecompressor(compressType).decompress(rawData, rawSize);
        String data = new String(rawData, "UTF-8");
        return new ByteArrayInputStream(data.getBytes());
    }

    public static byte[] readAndCompressBlock(byte[] stream, String compressType) throws Exception {
        byte[] compressedData = CompressorFactory.getCompressor(compressType).compress(stream);
        return compressedData;
    }

    public static byte[] serializeLogGroupToPB(Object logGroup) throws Exception {
        byte[] logBytes = null;
        Logs.LogGroup.Builder logs = Logs.LogGroup.newBuilder();
        HashMap<String, Object> body;
        if (logGroup instanceof HashMap) {
            body = (HashMap<String, Object>) logGroup;
        } else {
            throw new IllegalArgumentException("Invalid body type " + logGroup.getClass());
        }

        String topic = (String) body.get("Topic");
        if (topic != null) {
            logs.setTopic((String) body.get("Topic"));
        } else {
            throw new IllegalArgumentException("Topic is null");
        }
        String source = (String) body.get("Source");
        if (source != null) {
            logs.setSource((String) body.get("Source"));
        } else {
            try {
                logs.setSource(InetAddress.getLocalHost().getHostAddress());
            } catch (Exception e) {}
        }

        ArrayList<Object> logTags = (ArrayList<Object>) body.get("LogTags");
        for (Object obj : logTags) {
            HashMap<String, String> tag = (HashMap<String, String>) obj;
            Logs.LogTag.Builder tagBuilder = logs.addLogTagsBuilder();
            tagBuilder.setKey(tag.get("Key"));
            tagBuilder.setValue(tag.get("Value"));
        }
        ArrayList<Object> logItems = (ArrayList<Object>) body.get("Logs");
        for (Object obj : logItems) {
            Logs.Log.Builder log = logs.addLogsBuilder();
            HashMap<String, Object> logItem = (HashMap<String, Object>) obj;
            log.setTime((Integer) logItem.get("Time"));
            ArrayList<Object> contents = (ArrayList<Object>) logItem.get("Contents");
            for (Object content : contents) {
                HashMap<String, String> realContent = (HashMap<String, String>) content;
                Logs.Log.Content.Builder contentBuilder = log.addContentsBuilder();
                contentBuilder.setKey(realContent.get("Key"));
                contentBuilder.setValue(realContent.get("Value"));
            }
        }
        logBytes = logs.build().toByteArray();
        return logBytes;
    }

    public static String getBytesLength(byte[] stream) throws Exception {
        return String.valueOf(stream.length);
    }
}
