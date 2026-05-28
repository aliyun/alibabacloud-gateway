package com.aliyun.gateway.sls.util;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Assert;
import org.junit.Test;

public class LogGroupSerializerTest {

    @Test
    public void testDeserializeEmptyBytes() throws Exception {
        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(new byte[0]);
        HashMap<String, Object> map = asMap(result);
        List<?> logGroupList = (List<?>) map.get("logGroupList");
        Assert.assertNotNull(logGroupList);
        Assert.assertEquals(0, logGroupList.size());
    }

    @Test
    public void testDeserializeEmptyLogGroupList() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder().build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> map = asMap(result);
        List<?> logGroupList = (List<?>) map.get("logGroupList");
        Assert.assertNotNull(logGroupList);
        Assert.assertEquals(0, logGroupList.size());
    }

    @Test
    public void testDeserializeEmptyLogGroup() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder().build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> map = asMap(result);
        List<?> logGroupList = (List<?>) map.get("logGroupList");
        Assert.assertEquals(1, logGroupList.size());

        HashMap<String, Object> logGroup = asMap(logGroupList.get(0));
        Assert.assertFalse(logGroup.containsKey("Topic"));
        Assert.assertFalse(logGroup.containsKey("Source"));
        Assert.assertFalse(logGroup.containsKey("LogTags"));
        List<?> logItems = (List<?>) logGroup.get("LogItems");
        Assert.assertNotNull(logItems);
        Assert.assertEquals(0, logItems.size());
    }

    @Test
    public void testDeserializeWithTopicAndSource() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .setTopic("test-topic")
                        .setSource("192.168.1.1")
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> logGroup = getFirstLogGroup(result);
        Assert.assertEquals("test-topic", logGroup.get("Topic"));
        Assert.assertEquals("192.168.1.1", logGroup.get("Source"));
    }

    @Test
    public void testDeserializeWithLogTags() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .addLogTags(Logs.LogTag.newBuilder().setKey("env").setValue("prod").build())
                        .addLogTags(Logs.LogTag.newBuilder().setKey("region").setValue("cn-hangzhou").build())
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> logGroup = getFirstLogGroup(result);
        @SuppressWarnings("unchecked")
        List<HashMap<String, String>> tags = (List<HashMap<String, String>>) logGroup.get("LogTags");
        Assert.assertEquals(2, tags.size());
        Assert.assertEquals("env", tags.get(0).get("Key"));
        Assert.assertEquals("prod", tags.get(0).get("Value"));
        Assert.assertEquals("region", tags.get(1).get("Key"));
        Assert.assertEquals("cn-hangzhou", tags.get(1).get("Value"));
    }

    @Test
    public void testDeserializeWithLogs() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .addLogs(Logs.Log.newBuilder()
                                .setTime(1709736603)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("level").setValue("INFO").build())
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("message").setValue("hello world").build())
                                .build())
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> logGroup = getFirstLogGroup(result);
        @SuppressWarnings("unchecked")
        List<HashMap<String, Object>> logItems = (List<HashMap<String, Object>>) logGroup.get("LogItems");
        Assert.assertEquals(1, logItems.size());

        HashMap<String, Object> logItem = logItems.get(0);
        Assert.assertEquals(1709736603, logItem.get("Time"));
        Assert.assertFalse(logItem.containsKey("TimeNs"));

        @SuppressWarnings("unchecked")
        List<HashMap<String, String>> contents = (List<HashMap<String, String>>) logItem.get("Contents");
        Assert.assertEquals(2, contents.size());
        Assert.assertEquals("level", contents.get(0).get("Key"));
        Assert.assertEquals("INFO", contents.get(0).get("Value"));
        Assert.assertEquals("message", contents.get(1).get("Key"));
        Assert.assertEquals("hello world", contents.get(1).get("Value"));
    }

    @Test
    public void testDeserializeWithTimeNs() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .addLogs(Logs.Log.newBuilder()
                                .setTime(1709736603)
                                .setTimeNs(123456789)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("k").setValue("v").build())
                                .build())
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> logItem = getFirstLogItem(result);
        Assert.assertEquals(1709736603, logItem.get("Time"));
        Assert.assertEquals(123456789, logItem.get("TimeNs"));
    }

    @Test
    public void testDeserializeMultipleLogGroups() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .setTopic("topic-1")
                        .addLogs(Logs.Log.newBuilder().setTime(100)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("k1").setValue("v1").build())
                                .build())
                        .build())
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .setTopic("topic-2")
                        .addLogs(Logs.Log.newBuilder().setTime(200)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("k2").setValue("v2").build())
                                .build())
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> map = asMap(result);
        @SuppressWarnings("unchecked")
        List<HashMap<String, Object>> logGroupList = (List<HashMap<String, Object>>) map.get("logGroupList");
        Assert.assertEquals(2, logGroupList.size());
        Assert.assertEquals("topic-1", logGroupList.get(0).get("Topic"));
        Assert.assertEquals("topic-2", logGroupList.get(1).get("Topic"));
    }

    @Test
    public void testDeserializeMultipleLogs() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .addLogs(Logs.Log.newBuilder().setTime(100)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("a").setValue("1").build())
                                .build())
                        .addLogs(Logs.Log.newBuilder().setTime(200)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("b").setValue("2").build())
                                .build())
                        .addLogs(Logs.Log.newBuilder().setTime(300)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("c").setValue("3").build())
                                .build())
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> logGroup = getFirstLogGroup(result);
        @SuppressWarnings("unchecked")
        List<HashMap<String, Object>> logItems = (List<HashMap<String, Object>>) logGroup.get("LogItems");
        Assert.assertEquals(3, logItems.size());
        Assert.assertEquals(100, logItems.get(0).get("Time"));
        Assert.assertEquals(200, logItems.get(1).get("Time"));
        Assert.assertEquals(300, logItems.get(2).get("Time"));
    }

    @SuppressWarnings("unchecked")
    @Test
    public void testRoundTrip() throws Exception {
        HashMap<String, Object> originalLogGroup = new HashMap<>();
        originalLogGroup.put("Topic", "round-trip-topic");
        originalLogGroup.put("Source", "10.0.0.1");

        ArrayList<Object> logTags = new ArrayList<>();
        HashMap<String, String> tag = new HashMap<>();
        tag.put("Key", "app");
        tag.put("Value", "gateway");
        logTags.add(tag);
        originalLogGroup.put("LogTags", logTags);

        ArrayList<Object> logItems = new ArrayList<>();
        HashMap<String, Object> logItem = new HashMap<>();
        logItem.put("Time", 1709736603);
        ArrayList<Object> contents = new ArrayList<>();
        HashMap<String, String> content = new HashMap<>();
        content.put("Key", "msg");
        content.put("Value", "round trip test");
        contents.add(content);
        logItem.put("Contents", contents);
        logItems.add(logItem);
        originalLogGroup.put("LogItems", logItems);

        byte[] serialized = LogGroupSerializer.serializeLogGroupToPB(originalLogGroup);

        Logs.LogGroupList pbList = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.parseFrom(serialized))
                .build();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(pbList.toByteArray());
        HashMap<String, Object> map = asMap(result);
        List<HashMap<String, Object>> logGroupList = (List<HashMap<String, Object>>) map.get("logGroupList");
        Assert.assertEquals(1, logGroupList.size());

        HashMap<String, Object> deserialized = logGroupList.get(0);
        Assert.assertEquals("round-trip-topic", deserialized.get("Topic"));

        List<HashMap<String, String>> deserializedTags = (List<HashMap<String, String>>) deserialized.get("LogTags");
        Assert.assertEquals(1, deserializedTags.size());
        Assert.assertEquals("app", deserializedTags.get(0).get("Key"));
        Assert.assertEquals("gateway", deserializedTags.get(0).get("Value"));

        List<HashMap<String, Object>> deserializedItems = (List<HashMap<String, Object>>) deserialized.get("LogItems");
        Assert.assertEquals(1, deserializedItems.size());
        Assert.assertEquals(1709736603, deserializedItems.get(0).get("Time"));
        List<HashMap<String, String>> deserializedContents =
                (List<HashMap<String, String>>) deserializedItems.get(0).get("Contents");
        Assert.assertEquals("msg", deserializedContents.get(0).get("Key"));
        Assert.assertEquals("round trip test", deserializedContents.get(0).get("Value"));
    }

    @SuppressWarnings("unchecked")
    @Test
    public void testDeserializeChineseContent() throws Exception {
        Logs.LogGroupList pb = Logs.LogGroupList.newBuilder()
                .addLogGroupList(Logs.LogGroup.newBuilder()
                        .setTopic("中文主题")
                        .addLogs(Logs.Log.newBuilder()
                                .setTime(1709736603)
                                .addContents(Logs.Log.Content.newBuilder()
                                        .setKey("名称").setValue("张伟").build())
                                .build())
                        .build())
                .build();
        byte[] data = pb.toByteArray();

        Object result = LogGroupSerializer.deserializeLogGroupListFromPB(data);
        HashMap<String, Object> logGroup = getFirstLogGroup(result);
        Assert.assertEquals("中文主题", logGroup.get("Topic"));

        List<HashMap<String, Object>> logItems = (List<HashMap<String, Object>>) logGroup.get("LogItems");
        List<HashMap<String, String>> contents = (List<HashMap<String, String>>) logItems.get(0).get("Contents");
        Assert.assertEquals("名称", contents.get(0).get("Key"));
        Assert.assertEquals("张伟", contents.get(0).get("Value"));
    }

    @SuppressWarnings("unchecked")
    private HashMap<String, Object> asMap(Object obj) {
        return (HashMap<String, Object>) obj;
    }

    @SuppressWarnings("unchecked")
    private HashMap<String, Object> getFirstLogGroup(Object result) {
        HashMap<String, Object> map = asMap(result);
        List<HashMap<String, Object>> list = (List<HashMap<String, Object>>) map.get("logGroupList");
        return list.get(0);
    }

    @SuppressWarnings("unchecked")
    private HashMap<String, Object> getFirstLogItem(Object result) {
        HashMap<String, Object> logGroup = getFirstLogGroup(result);
        List<HashMap<String, Object>> logItems = (List<HashMap<String, Object>>) logGroup.get("LogItems");
        return logItems.get(0);
    }
}
