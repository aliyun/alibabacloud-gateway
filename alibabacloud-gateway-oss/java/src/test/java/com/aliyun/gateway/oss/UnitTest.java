package com.aliyun.gateway.oss;

import java.util.HashMap;
import java.util.Map;
import org.junit.Assert;
import org.junit.Test;

public class UnitTest {

    @Test
    public void getEndpointTest() throws Exception {
        Client client = new Client();

        Assert.assertEquals("custom.endpoint.com",
            client.getEndpoint("cn-hangzhou", null, "custom.endpoint.com"));
        Assert.assertEquals("custom.endpoint.com",
            client.getEndpoint("cn-hangzhou", "ipv6", "custom.endpoint.com"));

        Assert.assertEquals("oss-cn-shanghai.aliyuncs.com",
            client.getEndpoint("cn-shanghai", null, null));
        Assert.assertEquals("oss-cn-hangzhou.aliyuncs.com",
            client.getEndpoint(null, null, null));
        Assert.assertEquals("oss-cn-hangzhou.aliyuncs.com",
            client.getEndpoint("", "", ""));

        Assert.assertEquals("oss-cn-hangzhou-internal.aliyuncs.com",
            client.getEndpoint("cn-hangzhou", "internal", null));
        Assert.assertEquals("oss-cn-beijing-internal.aliyuncs.com",
            client.getEndpoint("cn-beijing", "vpc-internal", null));

        Assert.assertEquals("cn-hangzhou.oss.aliyuncs.com",
            client.getEndpoint("cn-hangzhou", "ipv6", null));
        Assert.assertEquals("cn-shanghai.oss.aliyuncs.com",
            client.getEndpoint("cn-shanghai", "dualstack-ipv6", null));
        Assert.assertEquals("cn-hangzhou.oss.aliyuncs.com",
            client.getEndpoint(null, "ipv6", null));

        Assert.assertEquals("oss-accelerate.aliyuncs.com",
            client.getEndpoint("cn-hangzhou", "accelerate", null));
        Assert.assertEquals("oss-accelerate-overseas.aliyuncs.com",
            client.getEndpoint("ap-southeast-1", "accelerate-overseas", null));
    }

    @Test
    public void getSignatureV2Test() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("content-md5", "FxqG8Ca0qEJPOghSihJ8Ew==");
        headers.put("content-type", "text/plain");
        headers.put("date", "Wed, 15 Feb 2017 09:37:11 GMT");
        headers.put("x-oss-object-acl", "private");

        String signature = client.getSignatureV2(
            "oss-example",
            "/nelson",
            "PUT",
            new HashMap<String, String>(),
            headers,
            "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV",
            client.getAdditionalHeaderNamesV2(headers));

        Assert.assertEquals("5Am2ewK1tL0gXX7GV6dwybZtj7efOEtc0Mo2FR6CkM8=", signature);
    }

    @Test
    public void getAuthorizationV2Test() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("content-md5", "FxqG8Ca0qEJPOghSihJ8Ew==");
        headers.put("content-type", "text/plain");
        headers.put("date", "Wed, 15 Feb 2017 09:37:11 GMT");
        headers.put("x-oss-object-acl", "private");

        String authorization = client.getAuthorization(
            "v2",
            "oss-example",
            "/nelson",
            "PUT",
            new HashMap<String, String>(),
            headers,
            "44CF9590006BF252F707",
            "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV",
            "cn-hangzhou");

        Assert.assertEquals(
            "OSS2 AccessKeyId:44CF9590006BF252F707,Signature:5Am2ewK1tL0gXX7GV6dwybZtj7efOEtc0Mo2FR6CkM8=",
            authorization);
    }

    @Test
    public void getAuthorizationV2WithAdditionalHeadersTest() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("content-md5", "FxqG8Ca0qEJPOghSihJ8Ew==");
        headers.put("content-type", "text/plain");
        headers.put("date", "Wed, 15 Feb 2017 09:37:11 GMT");
        headers.put("x-oss-object-acl", "private");
        headers.put("Range", "bytes=0-9");

        String authorization = client.getAuthorization(
            "v2",
            "oss-example",
            "/nelson",
            "PUT",
            new HashMap<String, String>(),
            headers,
            "44CF9590006BF252F707",
            "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV",
            "cn-hangzhou");

        Assert.assertTrue(authorization.startsWith("OSS2 AccessKeyId:44CF9590006BF252F707,AdditionalHeaders:range,Signature:"));
        Assert.assertFalse(authorization.endsWith("Signature:"));
    }

    @Test
    public void v2UriEncodeTest() throws Exception {
        Assert.assertEquals("%2Foss-example%2Fnelson", Client.v2UriEncode("/oss-example/nelson"));
    }

    @Test
    public void buildCanonicalizedQueryStringV2Test() throws Exception {
        Client client = new Client();
        Map<String, String> query = new HashMap<String, String>();
        query.put("versionId", "123");
        query.put("acl", "");

        Assert.assertEquals("?acl&versionId=123", client.buildCanonicalizedQueryStringV2(query));
    }
}
