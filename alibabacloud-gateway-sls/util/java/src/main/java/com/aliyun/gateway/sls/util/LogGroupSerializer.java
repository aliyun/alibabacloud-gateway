package com.aliyun.gateway.sls.util;

import java.net.InetAddress;
import java.net.NetworkInterface;
import java.net.SocketException;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.HashMap;

public class LogGroupSerializer {
    @SuppressWarnings("unchecked")
    public static byte[] serializeLogGroupToPB(Object logGroup) throws Exception {
        if (!(logGroup instanceof HashMap)) {
            throw new IllegalArgumentException("Invalid body type " + logGroup.getClass());
        }

        Logs.LogGroup.Builder logs = Logs.LogGroup.newBuilder();
        HashMap<String, Object> body = (HashMap<String, Object>) logGroup;

        String topic = (String) body.get("Topic");
        if (topic != null) {
            logs.setTopic(topic);
        }

        String source = (String) body.get("Source");
        if (source == null || source.isEmpty()) {
            source = getSourceIP();
        }
        if (source != null && !source.isEmpty()) {
            logs.setSource(source);
        }

        serializeLogTags(logs, body);

        serializeLogs(logs, body);
        return logs.build().toByteArray();
    }

    @SuppressWarnings("unchecked")
    private static void serializeLogs(Logs.LogGroup.Builder logs, HashMap<String, Object> body) {
        ArrayList<Object> logItems = (ArrayList<Object>) body.get("LogItems");
        if (logItems == null) {
            return;
        }
        for (Object obj : logItems) {
            Logs.Log.Builder logsBuilder = logs.addLogsBuilder();
            HashMap<String, Object> logItem = (HashMap<String, Object>) obj;
            if (logItem == null) {
                continue;
            }
            logsBuilder.setTime((Integer) logItem.get("Time"));
            ArrayList<Object> contents = (ArrayList<Object>) logItem.get("Contents");
            for (Object content : contents) {
                HashMap<String, String> realContent = (HashMap<String, String>) content;
                Logs.Log.Content.Builder contentBuilder = logsBuilder.addContentsBuilder();
                contentBuilder.setKey(realContent.get("Key"));
                contentBuilder.setValue(realContent.get("Value"));
            }
            Integer nanoTime = (Integer) logItem.get("TimeNs");
            if (nanoTime != null) {
                logsBuilder.setTimeNs(nanoTime);
            }
        }
    }

    @SuppressWarnings("unchecked")
    private static void serializeLogTags(Logs.LogGroup.Builder logs, HashMap<String, Object> body) {
        ArrayList<Object> logTags = (ArrayList<Object>) body.get("LogTags");
        if (logTags == null) {
            return;
        }
        for (Object obj : logTags) {
            HashMap<String, String> tag = (HashMap<String, String>) obj;
            Logs.LogTag.Builder tagBuilder = logs.addLogTagsBuilder();
            tagBuilder.setKey(tag.get("Key"));
            tagBuilder.setValue(tag.get("Value"));
        }
    }


    public static Object deserializeLogGroupListFromPB(byte[] data) throws Exception {
        Logs.LogGroupList logGroupListPB = Logs.LogGroupList.parseFrom(data);

        ArrayList<HashMap<String, Object>> logGroupList = new ArrayList<>();
        for (Logs.LogGroup lgPB : logGroupListPB.getLogGroupListList()) {
            HashMap<String, Object> logGroup = new HashMap<>();

            if (lgPB.hasTopic()) {
                logGroup.put("Topic", lgPB.getTopic());
            }
            if (lgPB.hasSource()) {
                logGroup.put("Source", lgPB.getSource());
            }

            if (lgPB.getLogTagsCount() > 0) {
                ArrayList<HashMap<String, String>> tags = new ArrayList<>();
                for (Logs.LogTag tag : lgPB.getLogTagsList()) {
                    HashMap<String, String> tagMap = new HashMap<>();
                    tagMap.put("Key", tag.getKey());
                    tagMap.put("Value", tag.getValue());
                    tags.add(tagMap);
                }
                logGroup.put("LogTags", tags);
            }

            ArrayList<HashMap<String, Object>> logItems = new ArrayList<>();
            for (Logs.Log log : lgPB.getLogsList()) {
                HashMap<String, Object> logItem = new HashMap<>();
                logItem.put("Time", log.getTime());

                if (log.hasTimeNs()) {
                    logItem.put("TimeNs", log.getTimeNs());
                }

                ArrayList<HashMap<String, String>> contents = new ArrayList<>();
                for (Logs.Log.Content content : log.getContentsList()) {
                    HashMap<String, String> contentMap = new HashMap<>();
                    contentMap.put("Key", content.getKey());
                    contentMap.put("Value", content.getValue());
                    contents.add(contentMap);
                }
                logItem.put("Contents", contents);
                logItems.add(logItem);
            }
            logGroup.put("LogItems", logItems);

            logGroupList.add(logGroup);
        }

        HashMap<String, Object> result = new HashMap<>();
        result.put("logGroupList", logGroupList);
        return result;
    }

    private static volatile String localMachineIP = null;
    private static String getSourceIP() {
        if (localMachineIP != null) {
            return localMachineIP;
        }
        String ip = getLocalMachineIP();
        synchronized (LogGroupSerializer.class) {
            localMachineIP = ip; // it's ok to overwrite
        }
        return localMachineIP;
    }

    private static String getLocalMachineIP() {
        try {
            Enumeration<NetworkInterface> networkInterfaces = NetworkInterface.getNetworkInterfaces();
            while (networkInterfaces.hasMoreElements()) {
                NetworkInterface ni = networkInterfaces.nextElement();
                if (!ni.isUp()) {
                    continue;
                }
                Enumeration<InetAddress> addresses = ni.getInetAddresses();
                while (addresses.hasMoreElements()) {
                    final InetAddress address = addresses.nextElement();
                    if (!address.isLinkLocalAddress() && address.getHostAddress() != null) {
                        String ipAddress = address.getHostAddress();
                        if (ipAddress.equals(CONST_LOCAL_IP)) {
                            continue;
                        }
                        if (isIPV4Addr(ipAddress)) {
                            return ipAddress;
                        }
                    }
                }
            }
        } catch (SocketException ex) {
            // swallow it
        } catch (Exception ex) {
            // swallow it
        }
        return "";
    }

    private static boolean isIPV4Addr(final String ipAddress) {
        if (ipAddress == null || ipAddress.isEmpty()) {
            return false;
        }
        try {
            final String[] tokens = ipAddress.split("\\.");
            if (tokens.length != 4) {
                return false;
            }
            for (String token : tokens) {
                int i = Integer.parseInt(token);
                if (i < 0 || i > 255) {
                    return false;
                }
            }
            return true;
        } catch (Exception ex) {
            return false;
        }
    }

    private static final String CONST_LOCAL_IP = "127.0.0.1";
}
