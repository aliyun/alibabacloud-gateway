package com.aliyun.gateway.mns;

import com.aliyun.gateway.spi.models.InterceptorContext;
import com.aliyun.tea.TeaException;
import org.junit.Assert;
import org.junit.Test;

import java.util.HashMap;
import java.util.Map;

public class UnitTest {

    private static final String RFC822_2025_06_03 = "Tue, 03 Jun 2025 11:23:48 GMT";
    private static final String RFC822_2025_12_03 = "Wed, 03 Dec 2025 11:23:48 GMT";

    @Test
    public void getEndpointTest() throws Exception {
        Client client = new Client();
        Assert.assertEquals("custom.endpoint.com", client.getEndpoint("cn-hangzhou", "custom.endpoint.com"));
        Assert.assertEquals("cn-shanghai.mns.aliyuncs.com", client.getEndpoint("cn-shanghai", null));
        Assert.assertEquals("cn-hangzhou.mns.aliyuncs.com", client.getEndpoint(null, null));
        Assert.assertEquals("cn-hangzhou.mns.aliyuncs.com", client.getEndpoint("", ""));
    }

    @Test
    public void getRegionTest() throws Exception {
        Client client = new Client();
        InterceptorContext context = new InterceptorContext();
        context.configuration = new InterceptorContext.InterceptorContextConfiguration();
        context.configuration.regionId = "eu-west-1";
        context.configuration.endpoint = "789.mns.eu-west-1.aliyuncs.com";
        Assert.assertEquals("eu-west-1", client.getRegion(context));

        context.configuration.regionId = null;
        context.configuration.endpoint = "cn-shanghai.mns.aliyuncs.com";
        Assert.assertEquals("cn-shanghai", client.getRegion(context));
    }

    @Test(expected = TeaException.class)
    public void getRegionMissingTest() throws Exception {
        Client client = new Client();
        InterceptorContext context = new InterceptorContext();
        context.configuration = new InterceptorContext.InterceptorContextConfiguration();
        context.configuration.endpoint = "invalid-host.example.com";
        client.getRegion(context);
    }

    @Test
    public void buildRequestTest() throws Exception {
        Client client = new Client();
        InterceptorContext context = new InterceptorContext();
        context.request = new InterceptorContext.InterceptorContextRequest();
        context.request.pathname = "/queues/TestQueue?num=10&empty";
        context.request.query = new HashMap<String, String>();
        client.buildRequest(context);
        Assert.assertEquals("/queues/TestQueue", context.request.pathname);
        Assert.assertEquals("10", context.request.query.get("num"));
        Assert.assertNull(context.request.query.get("empty"));
    }

    @Test
    public void testSignV2() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("date", RFC822_2025_06_03);
        headers.put("x-mns-abcDEF", "abcDEF");
        headers.put("x-kkk-def", "def");
        String authorization = client.getAuthorizationV2("/queues/TestQueue", "POST", headers, "ak1", "sk1");
        Assert.assertEquals("MNS ak1:cAThPWnbphg3hDe3ZHDADTgkgyE=", authorization);

        headers = new HashMap<String, String>();
        headers.put("date", RFC822_2025_12_03);
        headers.put("X-MNS-abcDEF", "KKKkkk");
        headers.put("x-kkk-def", "jjj");
        authorization = client.getAuthorizationV2("/queues/TestQueue123456", "GET", headers, "ak-kkkk1111111111", "sk-kkkkkk1111111111");
        Assert.assertEquals("MNS ak-kkkk1111111111:lPxnfNA+Qo/aj18MYJUnRMqa84I=", authorization);
    }

    @Test
    public void testSignV2EdgeCases() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("date", RFC822_2025_06_03);
        String authorization = client.getAuthorizationV2("/queues/TestQueue", "GET", headers, "testAk", "testSk");
        Assert.assertEquals("MNS testAk:uDJQXbk2xOhH95MDXmgbAGdtyE4=", authorization);

        headers.put("content-type", "application/json");
        authorization = client.getAuthorizationV2("/queues/Test Queue@123", "POST", headers, "testAk", "testSk");
        Assert.assertEquals("MNS testAk:GZ/I9UgZsqvbzfzjd0GmA6VHOx4=", authorization);
    }

    @Test
    public void testSignV4PublicCloud() throws Exception {
        Client client = new Client();
        String date = "20250603T105235Z";

        InterceptorContext context = buildV4Context("region1", "POST", "/queues/TestQueue/");
        context.request.query.put("param1", "abc");
        context.request.query.put("PARAM2", "d  ef");
        context.request.query.put("PARAM3", "[def]#bbb");
        context.request.query.put("PARAM4", "中文参数");
        String authorization = client.getAuthorizationV4(context, date, "ak1", "sk1");
        Assert.assertEquals(
            "MNS4-HMAC-SHA256 Credential=ak1/20250603/region1/mns/aliyun_v4_request,Signature=1fc6013f158d4c3aba09570895a2a93c8445c2c5213d82a1463baf602cf0d945",
            authorization);

        date = "20251203T112348Z";
        context = buildV4Context("region2", "PUT", "queues/TestQueue123456");
        context.request.headers.put("x-mns-abc", "kkk");
        context.request.headers.put("X-MNS-Sss", "2015-06-06xFFF");
        context.request.headers.put("x-kkk-def", "jjj");
        authorization = client.getAuthorizationV4(context, date, "akkkkk1111111111", "skkkkkkk1111111111");
        Assert.assertEquals(
            "MNS4-HMAC-SHA256 Credential=akkkkk1111111111/20251203/region2/mns/aliyun_v4_request,Signature=04f1e5f929b5aaac3af6a54970f591581ecd2aade2562f926c5f2ec5d5c8cfad",
            authorization);

        date = "20251204T063235Z";
        context = buildV4Context("eu-west-1", "GET", "/queues");
        context.request.headers.put("x-mns-version", "2015-06-30");
        context.request.headers.put("x-mns-ttt", "1");
        context.request.headers.put("X-MNS-UUUuu", "FALse");
        context.request.headers.put("bb-MNS-UUUuu", "abcdddDDD");
        context.request.headers.put("content-type", "application/json");
        context.request.headers.put("content-md5", "0d9d2ce750000593eb34f00");
        authorization = client.getAuthorizationV4(context, date, "akkkkk222222222", "skkkkkk2222222222");
        Assert.assertEquals(
            "MNS4-HMAC-SHA256 Credential=akkkkk222222222/20251204/eu-west-1/mns/aliyun_v4_request,Signature=294f776a2d5726df679b54b081ab8126a199a2545c8c8a7a875adff0b386e458",
            authorization);
    }

    @Test
    public void testSignV4EdgeCases() throws Exception {
        Client client = new Client();
        String date = "20250603T105235Z";
        InterceptorContext context = buildV4Context("testRegion", "GET", "/");
        String authorization = client.getAuthorizationV4(context, date, "testAk", "testSk");
        Assert.assertEquals(
            "MNS4-HMAC-SHA256 Credential=testAk/20250603/testRegion/mns/aliyun_v4_request,Signature=63f78964dcd1132bab91b9c38ceebdfb59509149b34874c5b2c9dfc2b12e7291",
            authorization);

        context = buildV4Context("testRegion", "POST", "/queues/Test Queue@123");
        authorization = client.getAuthorizationV4(context, date, "testAk", "testSk");
        Assert.assertEquals(
            "MNS4-HMAC-SHA256 Credential=testAk/20250603/testRegion/mns/aliyun_v4_request,Signature=87de7499365a31e767cdc0c726a893c72c860e415876f7f92914197ac3795833",
            authorization);
    }

    @Test
    public void testSignV4WithResourcePathContainsParam() throws Exception {
        Client client = new Client();
        String date = "20250603T105235Z";
        InterceptorContext context = buildV4Context("region1", "POST", "/queues/TestQueue?num=10&paramx=100");
        context.request.query.put("param1", "abc");
        context.request.query.put("PARAM2", "d  ef");
        context.request.query.put("PARAM3", "[def]#bbb");
        context.request.query.put("PARAM4", "中文参数");
        client.buildRequest(context);
        String authorization = client.getAuthorizationV4(context, date, "ak1", "sk1");
        Assert.assertEquals(
            "MNS4-HMAC-SHA256 Credential=ak1/20250603/region1/mns/aliyun_v4_request,Signature=8d555514e749185d3e6d68df416692e6cc960ba85a089c9d69a2f4d7bb607560",
            authorization);
    }

    @Test
    public void buildCanonicalizedHeadersV2Test() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("content-md5", "md5val");
        headers.put("content-type", "text/xml");
        headers.put("date", "Wed, 01 Jan 2025 00:00:00 GMT");
        headers.put("x-mns-version", "2015-06-06");
        headers.put("x-mns-zzz", "z");
        headers.put("host", "example.com");
        String canonical = client.buildCanonicalizedHeadersV2(headers);
        Assert.assertEquals("md5val\ntext/xml\nWed, 01 Jan 2025 00:00:00 GMT\nx-mns-version:2015-06-06\nx-mns-zzz:z\n", canonical);
    }

    @Test
    public void hasSignedHeaderV4Test() throws Exception {
        Client client = new Client();
        Assert.assertTrue(client.hasSignedHeaderV4("content-type"));
        Assert.assertTrue(client.hasSignedHeaderV4("x-mns-version"));
        Assert.assertFalse(client.hasSignedHeaderV4("host"));
        Assert.assertFalse(client.hasSignedHeaderV4("x-kkk-def"));
    }

    @Test
    public void base64MessageBodyTest() throws Exception {
        Client client = new Client();

        Assert.assertEquals("aGVsbG8tbW5z", client.base64Encode("hello-mns"));
        Assert.assertEquals("hello-mns", client.base64Decode("aGVsbG8tbW5z"));

        String messageBody = "中文 message body !@#";
        String encoded = client.base64Encode(messageBody);
        Assert.assertEquals(messageBody, client.base64Decode(encoded));
    }

    @Test
    public void base64EmptyInputTest() throws Exception {
        Client client = new Client();

        Assert.assertEquals("", client.base64Encode(""));
        Assert.assertEquals("", client.base64Decode(""));
        Assert.assertEquals("", client.base64Encode(null));
        Assert.assertEquals("", client.base64Decode(null));
    }

    private InterceptorContext buildV4Context(String regionId, String method, String pathname) {
        InterceptorContext context = new InterceptorContext();
        context.configuration = new InterceptorContext.InterceptorContextConfiguration();
        context.configuration.regionId = regionId;
        context.request = new InterceptorContext.InterceptorContextRequest();
        context.request.method = method;
        context.request.pathname = pathname;
        context.request.query = new HashMap<String, String>();
        context.request.headers = new HashMap<String, String>();
        return context;
    }
}
