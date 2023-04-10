// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.fc.util;

import com.aliyun.credentials.utils.AcsURLEncoder;
import com.aliyun.tea.okhttp.OkHttpClientBuilder;
import okhttp3.*;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.UnsupportedEncodingException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.text.SimpleDateFormat;
import java.util.*;

import static java.lang.String.format;


public class Client {

    // HTTPHeaderContentMD5 key in request headers
    public static final String HTTPHeaderContentMD5 = "Content-MD5";
    // HTTPHeaderPrefix key prefix in request headers
    public static final String HTTPHeaderPrefix = "x-fc-";
    // HTTPHeaderDate key in request headers
    public static final String HTTPHeaderDate = "Date";
    // AuthQueryKeyExpires key in request headers
    public static final String AuthQueryKeyExpires = "x-fc-expires";
    // HTTPHeaderContentType key in request headers
    public static final String HTTPHeaderContentType = "Content-Type";
    // HTTPHeaderAuthorization key in request headers
    public static final String HTTPHeaderAuthorization = "Authorization";
    // HTTPHeaderSecurityToken key in request headers
    public static final String HTTPHeaderSecurityToken = "x-fc-security-token";

    private static final String FORMAT_RFC2616 = "EEE, dd MMM yyyy HH:mm:ss zzz";
    private static final String TIME_ZONE = "GMT";


    private final OkHttpClient defaultHttpClient;
    public com.aliyun.credentials.Client _credential;

    public Client(com.aliyun.credentials.Client cred) throws Exception {
        this.defaultHttpClient = new OkHttpClientBuilder().buildOkHttpClient();
        this._credential = cred;
    }


    public Response InvokeHTTPTrigger(String url, String method, byte[] body, Headers headers) throws Exception {
        Request req = BuildHTTPRequest(url, method, body, headers);
        return SendHTTPRequestWithAuthorization(req);
    }

    public Response InvokeAnonymousHTTPTrigger(String url, String method, byte[] body, Headers headers) throws Exception {
        Request req = BuildHTTPRequest(url, method, body, headers);
        return SendHTTPRequest(req);
    }

    public Response SendHTTPRequestWithAuthorization(Request req) throws Exception {
        Request signedRequest = SignRequest(req);
        return SendHTTPRequest(signedRequest);
    }

    public Response SendHTTPRequest(Request req) throws Exception {
        return defaultHttpClient.newCall(req).execute();
    }

    public Request SignRequest(Request req) throws Exception {
        // TODO md5 check
        return SignRequestWithContentMD5(req, "");
    }

    public Request SignRequestWithContentMD5(Request req, String contentMD5) throws Exception {
        Request.Builder reqBuilder = req.newBuilder();
        Map<String, String> headerParams = new HashMap<>();
        String d = getRFC2616Date(null);
        reqBuilder.addHeader(HTTPHeaderDate, d);
        headerParams.put(HTTPHeaderDate, d);
        if (contentMD5 != null && contentMD5.length() != 0) {
            reqBuilder.addHeader(HTTPHeaderContentMD5, contentMD5);
        }
        Map<String, List<String>> mmap = req.headers().toMultimap();
        for (String key : mmap.keySet()) {
            List<String> v = mmap.get(key);
            StringBuilder sb = new StringBuilder();
            headerParams.put(key, join(",", v));
        }
        if (req.body() != null && req.body().contentType() != null) {
            headerParams.put(HTTPHeaderContentType, req.body().contentType().toString());
        }
        String path = req.url().url().getPath();
        if (path == null || path.length() == 0) {
            path = "/";
        }
        Map<String, List<String>> params = new HashMap<>();
        for (String name : req.url().queryParameterNames()) {
            params.put(name, req.url().queryParameterValues(name));
        }
        String accessKeyId = _credential.getAccessKeyId();
        String accessKeySecret = _credential.getAccessKeySecret();
        String securityToken = _credential.getSecurityToken();
        if (securityToken != null && securityToken.length() != 0) {
            reqBuilder.addHeader(HTTPHeaderSecurityToken, securityToken);
            headerParams.put(HTTPHeaderSecurityToken, securityToken);
        }
        String authStr = getAuthString(accessKeyId, accessKeySecret, req.method(), headerParams, path, params);
        reqBuilder.addHeader(HTTPHeaderAuthorization, authStr);
        return reqBuilder.build();
    }

    public static String getAuthString(String accessKeyId, String accessKeySecret, String method, Map<String, String> headerParams, String path, Map<String, List<String>> params) {
        return "FC " + accessKeyId + ":" + getSignature(accessKeySecret, method, headerParams, path, params);
    }


    public static String getRFC2616Date(Date date) {
        Date nowDate = date;
        if (null == date) {
            nowDate = new Date();
        }
        SimpleDateFormat df = new SimpleDateFormat(FORMAT_RFC2616, Locale.ENGLISH);
        df.setTimeZone(new SimpleTimeZone(0, TIME_ZONE));
        return df.format(nowDate);
    }


    private static String getSignature(String accessKeySecret, String method, Map<String, String> headers, String path, Map<String, List<String>> params) {
        String stringToSign = composeStringToSignWithMultiValue(method, path, headers, params);
        return signString(stringToSign, accessKeySecret);
    }

    public static String signString(String source, String accessSecret) {
        try {
            Mac mac = Mac.getInstance("HmacSHA256");
            mac.init(new SecretKeySpec(
                    accessSecret.getBytes(AcsURLEncoder.URL_ENCODING), "HmacSHA256"));
            byte[] signData = mac.doFinal(source.getBytes(AcsURLEncoder.URL_ENCODING));
            return Base64Helper.encode(signData);
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException("HMAC-SHA1 not supported.");
        } catch (UnsupportedEncodingException e) {
            throw new RuntimeException("UTF-8 not supported.");
        } catch (InvalidKeyException e) {
            throw new RuntimeException("Invalid key " + accessSecret);
        }

    }

    private static String composeCanonicalizationFCHeaders(Map<String, String> headers,
                                                           Map<String, List<String>> queries) {
        StringBuilder sb = new StringBuilder();
        if (headers != null && headers.get(HTTPHeaderContentMD5) != null) {
            sb.append(headers.get(HTTPHeaderContentMD5));
        }
        sb.append("\n");
        if (headers != null && headers.get(HTTPHeaderContentType) != null) {
            sb.append(headers.get(HTTPHeaderContentType));
        }

        sb.append("\n");
        String expires = getExpiresFromURLQueries(queries);
        if (expires != null) {
            sb.append(expires);
        } else if (headers != null && headers.get(HTTPHeaderDate) != null) {
            sb.append(headers.get(HTTPHeaderDate));
        }

        return sb.toString();
    }

    private static String getExpiresFromURLQueries(Map<String, List<String>> queries) {
        if (queries == null) {
            return null;
        }
        List<String> vals = queries.get(AuthQueryKeyExpires);
        if (vals == null || vals.size() < 1) {
            return null;
        }
        return vals.get(0);
    }

    private static String composeCanonicalizedResource(String path, Map<String, List<String>> queries) {
        StringBuilder sb = new StringBuilder();
        sb.append(path);
        if (queries != null) {
            sb.append("\n");

            List<String> params = new ArrayList<String>(queries.size());

            for (Map.Entry<String, List<String>> query : queries.entrySet()) {
                String key = query.getKey();
                List<String> values = query.getValue();

                if (values == null || values.size() == 0) {
                    params.add(key);
                    continue;
                }
                for (String value : values) {
                    if (value == null) {
                        params.add(key);
                    } else {
                        params.add(format("%s=%s", key, value));
                    }
                }
            }
            Collections.sort(params);
            sb.append(join("\n", params));
        }

        return sb.toString();
    }

    static String composeStringToSignWithMultiValue(String method, String path,
                                                    Map<String, String> headers, Map<String, List<String>> queries) {
        StringBuilder sb = new StringBuilder();
        sb.append(method);
        sb.append("\n");
        HashMap<String, String> headersClone = new HashMap<>();
        if (headers != null) {
            for (String key : headers.keySet()) {
                String keyClone = key;
                if (keyClone.equalsIgnoreCase(HTTPHeaderDate)) {
                    keyClone = HTTPHeaderDate;
                }
                if (keyClone.equalsIgnoreCase(HTTPHeaderContentType)) {
                    keyClone = HTTPHeaderContentType;
                }
                if (keyClone.equalsIgnoreCase(HTTPHeaderContentMD5)) {
                    keyClone = HTTPHeaderContentMD5;
                }
                headersClone.put(keyClone, headers.get(key));
            }
        }
        sb.append(composeCanonicalizationFCHeaders(headersClone, queries));
        sb.append("\n");
        sb.append(buildCanonicalHeaders(headersClone, HTTPHeaderPrefix));

        sb.append(composeCanonicalizedResource(path, queries));

        return sb.toString();
    }

    public static String buildCanonicalHeaders(Map<String, String> headers, String headerBegin) {
        if (headers == null) return "";

        Map<String, String> sortMap = new TreeMap<String, String>();
        for (Map.Entry<String, String> e : headers.entrySet()) {
            String key = e.getKey().toLowerCase();
            String val = e.getValue();
            if (key.startsWith(headerBegin)) {
                sortMap.put(key, val);
            }
        }
        StringBuilder headerBuilder = new StringBuilder();
        for (Map.Entry<String, String> e : sortMap.entrySet()) {
            headerBuilder.append(e.getKey());
            headerBuilder.append(':').append(e.getValue());
            headerBuilder.append("\n");
        }
        return headerBuilder.toString();
    }

    private static String getSignQueries(Map<String, List<String>> params) {
        List<String> paramsList = new ArrayList<>();
        for (String key : params.keySet()) {
            List<String> value = params.get(key);
            if (value.size() == 0) {
                paramsList.add(key);
                continue;
            }
            for (String v : value) {
                paramsList.add(key + "=" + v);
            }
        }
        Collections.sort(paramsList);
        return join("\n", paramsList);
    }

    private static String join(String separator, List<String> input) {
        if (input == null || input.size() == 0) {
            return "";
        }
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < input.size(); i++) {
            sb.append(input.get(i));
            // if not the last item
            if (i != input.size() - 1) {
                sb.append(separator);
            }
        }
        return sb.toString();
    }

    public Request BuildHTTPRequest(String url, String method, byte[] body, Headers headers) throws Exception {
        Request.Builder builder = new Request.Builder();
        builder.url(url);
        if (body != null) {
            builder.method(method, RequestBody.create(MediaType.parse("application/json"), body));
        } else {
            builder.method(method, null);
        }
        builder.headers(headers);
        return builder.build();
    }
}
