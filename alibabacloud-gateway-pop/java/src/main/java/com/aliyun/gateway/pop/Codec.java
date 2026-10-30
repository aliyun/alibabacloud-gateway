package com.aliyun.gateway.pop;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.security.MessageDigest;
import java.security.Provider;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.regex.Pattern;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

/**
 * In-package replacements for darabonba-string, darabonba-array, darabonba-map,
 * darabonba-encode-util and the SHA-256 half of darabonba-signature-util.
 * Behavior matches those packages so ACS4 signatures stay byte-identical.
 * SM3 still needs Bouncy Castle at runtime; that dependency is optional.
 */
final class Codec {

    private Codec() {
    }

    static String hexEncode(byte[] raw) {
        if (raw == null) {
            return null;
        }
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < raw.length; i++) {
            String hex = Integer.toHexString(raw[i] & 0xFF);
            if (hex.length() < 2) {
                sb.append(0);
            }
            sb.append(hex);
        }
        return sb.toString();
    }

    static byte[] hash(byte[] raw, String signatureAlgorithm) throws Exception {
        if (signatureAlgorithm == null) {
            return null;
        }
        if (signatureAlgorithm.contains("HMAC-SHA256") || signatureAlgorithm.contains("RSA-SHA256")) {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            return digest.digest(raw);
        }
        if (signatureAlgorithm.contains("HMAC-SM3")) {
            MessageDigest digest = MessageDigest.getInstance("SM3", sm3Provider());
            return digest.digest(raw);
        }
        return null;
    }

    static String percentEncode(String value) throws UnsupportedEncodingException {
        if (value == null) {
            return null;
        }
        return URLEncoder.encode(value, "UTF-8")
                .replace("+", "%20")
                .replace("*", "%2A")
                .replace("%7E", "~");
    }

    static byte[] hmacSha256(String stringToSign, String secret) throws Exception {
        return hmacSha256ByBytes(stringToSign, secret.getBytes("UTF-8"));
    }

    static byte[] hmacSha256ByBytes(String stringToSign, byte[] key) throws Exception {
        Mac mac = Mac.getInstance("HmacSHA256");
        mac.init(new SecretKeySpec(key, "HmacSHA256"));
        // Match darabonba Signer.HmacSHA256SignByBytes: platform default charset.
        return mac.doFinal(stringToSign.getBytes());
    }

    static byte[] hmacSm3(String stringToSign, String secret) throws Exception {
        return hmacSm3ByBytes(stringToSign, secret.getBytes("UTF-8"));
    }

    static byte[] hmacSm3ByBytes(String stringToSign, byte[] key) throws Exception {
        Object mac = sm3Class("org.bouncycastle.crypto.macs.HMac")
                .getConstructor(sm3Class("org.bouncycastle.crypto.Digest"))
                .newInstance(sm3Class("org.bouncycastle.crypto.digests.SM3Digest").getConstructor().newInstance());
        byte[] out = new byte[(Integer) mac.getClass().getMethod("getMacSize").invoke(mac)];
        byte[] input = stringToSign.getBytes("UTF-8");
        Object keyParameter = sm3Class("org.bouncycastle.crypto.params.KeyParameter")
                .getConstructor(byte[].class)
                .newInstance(key);
        mac.getClass().getMethod("init", sm3Class("org.bouncycastle.crypto.CipherParameters"))
                .invoke(mac, keyParameter);
        mac.getClass().getMethod("update", byte[].class, int.class, int.class)
                .invoke(mac, input, 0, input.length);
        mac.getClass().getMethod("doFinal", byte[].class, int.class).invoke(mac, out, 0);
        return out;
    }

    static String subString(String value, int start, int end) {
        if (end < 0) {
            end = value.length() + end;
        }
        if (end <= 0) {
            return "";
        }
        return value.substring(start, end);
    }

    static String replace(String raw, String oldStr, String newStr, Integer count) {
        if (count == null) {
            return raw.replaceAll(Pattern.quote(oldStr), newStr == null ? "" : newStr);
        }
        String[] array = raw.split(Pattern.quote(oldStr), count + 1);
        if (array.length == 0) {
            return raw;
        }
        StringBuilder sb = new StringBuilder();
        sb.append(array[0]);
        for (int i = 1; i < array.length; i++) {
            sb.append(newStr).append(array[i]);
        }
        return sb.toString();
    }

    static boolean contains(String value, String substr) {
        return value.contains(substr);
    }

    static List<String> split(String raw, String separator, Integer limit) {
        String[] parts = limit == null
                ? raw.split(Pattern.quote(separator))
                : raw.split(Pattern.quote(separator), limit);
        return Arrays.asList(parts);
    }

    static String toLower(String value) {
        return value.toLowerCase();
    }

    static boolean equals(String expect, String actual) {
        return expect.equals(actual);
    }

    static String trim(String value) {
        return value.trim();
    }

    static boolean hasPrefix(String value, String prefix) {
        return value.startsWith(prefix);
    }

    static List<String> keySet(Map<String, ?> map) {
        if (map == null) {
            throw new IllegalArgumentException("not a valid value for parameter");
        }
        return new ArrayList<String>(map.keySet());
    }

    static String join(List<String> raw, String separator) {
        if (raw == null || separator == null) {
            throw new IllegalArgumentException("not a valid value for parameter");
        }
        if (raw.size() > 0) {
            String result = raw.get(0);
            for (int i = 1; i < raw.size(); i++) {
                result = result + separator + raw.get(i);
            }
            return result;
        }
        return raw.toString();
    }

    static int size(List<String> raw) {
        if (raw == null) {
            throw new IllegalArgumentException("not a valid value for parameter");
        }
        return raw.size();
    }

    static List<String> ascSort(List<String> raw) {
        if (raw == null) {
            throw new IllegalArgumentException("not a valid value for parameter");
        }
        String[] sorted = raw.toArray(new String[raw.size()]);
        Arrays.sort(sorted);
        return Arrays.asList(sorted);
    }

    static void append(List<String> array, String item) {
        if (array != null) {
            array.add(item);
        }
    }

    static boolean listContains(List<String> raw, String value) {
        if (raw == null || value == null) {
            throw new IllegalArgumentException("not a valid value for parameter");
        }
        return raw.contains(value);
    }

    private static Provider sm3Provider() throws Exception {
        return (Provider) sm3Class("org.bouncycastle.jce.provider.BouncyCastleProvider").getConstructor().newInstance();
    }

    private static Class<?> sm3Class(String name) throws ClassNotFoundException {
        try {
            return Class.forName(name);
        } catch (ClassNotFoundException ex) {
            throw new ClassNotFoundException(
                    "SM3 requires org.bouncycastle:bcprov-jdk15on on the classpath", ex);
        }
    }
}
