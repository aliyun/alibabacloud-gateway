package com.aliyun.gateway.pop;

import com.aliyun.tea.TeaModel;
import com.aliyun.tea.utils.StringUtils;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonElement;
import com.google.gson.reflect.TypeToken;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.UnsupportedEncodingException;
import java.lang.reflect.Type;
import java.net.URLEncoder;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.SimpleTimeZone;
import java.util.concurrent.ThreadLocalRandom;
import java.util.concurrent.atomic.AtomicLong;

final class Support {
    private static final AtomicLong SEQ_ID = new AtomicLong(0);
    private static final long PROCESS_START_TIME = System.currentTimeMillis();
    private static final String UTF8 = "UTF-8";

    private Support() {
    }

    static byte[] toBytes(String str) {
        try {
            return str.getBytes(UTF8);
        } catch (UnsupportedEncodingException e) {
            throw new RuntimeException(e.getMessage(), e);
        }
    }

    static Map<String, Object> assertAsMap(Object object) {
        if (null != object && Map.class.isAssignableFrom(object.getClass())) {
            return (Map<String, Object>) object;
        }
        throw new RuntimeException("The value is not a object");
    }

    static byte[] assertAsBytes(Object object) {
        if (object instanceof byte[]) {
            return (byte[]) object;
        }
        throw new RuntimeException("The value is not a byteArray");
    }

    static byte[] readAsBytes(InputStream stream) {
        try {
            if (null == stream) {
                return new byte[]{};
            }
            ByteArrayOutputStream os = new ByteArrayOutputStream();
            byte[] buff = new byte[1024];
            while (true) {
                int read = stream.read(buff);
                if (read == -1) {
                    return os.toByteArray();
                }
                os.write(buff, 0, read);
            }
        } catch (Exception e) {
            throw new RuntimeException(e.getMessage(), e);
        } finally {
            if (null != stream) {
                try {
                    stream.close();
                } catch (IOException e) {
                    throw new RuntimeException(e.getMessage(), e);
                }
            }
        }
    }

    static String readAsString(InputStream stream) {
        try {
            return new String(readAsBytes(stream), UTF8);
        } catch (UnsupportedEncodingException e) {
            throw new RuntimeException(e.getMessage(), e);
        }
    }

    static Object readAsJSON(InputStream stream) {
        String body = readAsString(stream);
        try {
            return parseJSON(body);
        } catch (Exception exception) {
            throw new RuntimeException("Error: convert to JSON, response is:\n" + body);
        }
    }

    static Object parseJSON(String json) {
        Gson gson = new GsonBuilder()
                .registerTypeAdapter(new TypeToken<Map<String, Object>>() {
                }.getType(), new MapTypeAdapter()).create();
        JsonElement jsonElement = gson.fromJson(json, JsonElement.class);
        Type mapType = new TypeToken<Map<String, Object>>() {
        }.getType();
        return jsonElement.isJsonArray() ? gson.fromJson(json, List.class) : gson.fromJson(json, mapType);
    }

    static String getNonce() {
        long threadId = Thread.currentThread().getId();
        long currentTime = System.currentTimeMillis();
        ThreadLocalRandom random = ThreadLocalRandom.current();
        long seq = SEQ_ID.getAndIncrement();
        long rand = random.nextLong();
        StringBuffer sb = new StringBuffer();
        sb.append(PROCESS_START_TIME).append('-')
                .append(threadId).append('-')
                .append(currentTime).append('-')
                .append(seq).append('-')
                .append(rand);
        try {
            MessageDigest digest = MessageDigest.getInstance("MD5");
            byte[] msg = sb.toString().getBytes();
            sb.setLength(0);
            for (byte b : digest.digest(msg)) {
                String hex = Integer.toHexString(b & 0xFF);
                if (hex.length() < 2) {
                    sb.append(0);
                }
                sb.append(hex);
            }
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException(e.getMessage(), e);
        }
        return sb.toString();
    }

    static String defaultString(String str, String defaultStr) {
        if (!StringUtils.isEmpty(str)) {
            return str;
        }
        return defaultStr;
    }

    static String toJSONString(Object object) {
        if (object instanceof String) {
            return (String) object;
        }
        Gson gson = new GsonBuilder().disableHtmlEscaping().create();
        return gson.toJson(object);
    }

    static boolean empty(String str) {
        return StringUtils.isEmpty(str);
    }

    static boolean equalString(String str, String val) {
        if (str == null || val == null) {
            return false;
        }
        return str.equals(val);
    }

    static boolean equalNumber(Number num, Number val) {
        if (num == null || val == null) {
            return false;
        }
        return num.doubleValue() == val.doubleValue();
    }

    static boolean isUnset(Object object) {
        return null == object;
    }

    static boolean is4xx(Number code) {
        if (null == code) {
            return false;
        }
        return code.intValue() >= 400 && code.intValue() < 500;
    }

    static boolean is5xx(Number code) {
        if (null == code) {
            return false;
        }
        return code.intValue() >= 500 && code.intValue() < 600;
    }

    static String getTimestamp() throws Exception {
        SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss'Z'");
        df.setTimeZone(new SimpleTimeZone(0, "UTC"));
        return df.format(new Date());
    }

    static String getEndpointRules(String product, String regionId, String endpointType, String netWork, String suffix) {
        String result;
        if (!isEmptyText(netWork) && !"public".equals(netWork)) {
            netWork = "-" + netWork;
        } else {
            netWork = "";
        }
        if (null == suffix) {
            suffix = "";
        } else {
            suffix = "-" + suffix;
        }
        if ("regional".equals(endpointType)) {
            if (isEmptyText(regionId)) {
                throw new RuntimeException("RegionId is empty, please set a valid RegionId");
            }
            result = String.format("%s%s%s.%s.aliyuncs.com", product, suffix, netWork, regionId);
        } else {
            result = String.format("%s%s%s.aliyuncs.com", product, suffix, netWork);
        }
        return result;
    }

    static String toForm(Map<String, ?> filter) throws Exception {
        Map<String, String> map = query(filter);
        StringBuilder result = new StringBuilder();
        boolean first = true;
        for (Map.Entry<String, ?> entry : map.entrySet()) {
            if (StringUtils.isEmpty(entry.getValue())) {
                continue;
            }
            if (first) {
                first = false;
            } else {
                result.append("&");
            }
            result.append(URLEncoder.encode(entry.getKey(), UTF8));
            result.append("=");
            result.append(URLEncoder.encode(String.valueOf(entry.getValue()), UTF8));
        }
        return result.toString();
    }

    private static Map<String, String> query(Map<String, ?> filter) throws Exception {
        Map<String, String> outMap = new HashMap<String, String>();
        if (null != filter) {
            processObject(outMap, "", filter);
        }
        return outMap;
    }

    private static void processObject(Map<String, String> map, String key, Object value) throws UnsupportedEncodingException {
        if (null == value) {
            return;
        }
        if (value instanceof List) {
            List list = (List) value;
            for (int i = 0; i < list.size(); i++) {
                processObject(map, key + "." + (i + 1), list.get(i));
            }
        } else if (value instanceof TeaModel) {
            Map<String, Object> subMap = (Map<String, Object>) ((TeaModel) value).toMap();
            for (Map.Entry<String, Object> entry : subMap.entrySet()) {
                processObject(map, key + "." + entry.getKey(), entry.getValue());
            }
        } else if (value instanceof Map) {
            Map<String, Object> subMap = (Map<String, Object>) value;
            for (Map.Entry<String, Object> entry : subMap.entrySet()) {
                processObject(map, key + "." + entry.getKey(), entry.getValue());
            }
        } else {
            if (key.startsWith(".")) {
                key = key.substring(1);
            }
            if (value instanceof byte[]) {
                map.put(key, new String((byte[]) value, UTF8));
            } else {
                map.put(key, String.valueOf(value));
            }
        }
    }

    private static boolean isEmptyText(final CharSequence cs) {
        return cs == null || cs.length() == 0;
    }
}
