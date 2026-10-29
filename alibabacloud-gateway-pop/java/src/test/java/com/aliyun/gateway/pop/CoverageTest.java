package com.aliyun.gateway.pop;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

import com.aliyun.credentials.models.Config;
import com.aliyun.credentials.models.CredentialModel;
import com.aliyun.gateway.spi.models.AttributeMap;
import com.aliyun.gateway.spi.models.InterceptorContext;
import com.aliyun.tea.TeaException;
import org.junit.Assert;
import org.junit.Test;

public class CoverageTest {
    static class BoomStream extends InputStream {
        @Override
        public int read() throws IOException {
            throw new IOException("read failed");
        }
    }

    static class FixedCredential extends com.aliyun.credentials.Client {
        private final CredentialModel model;
        private final RuntimeException error;

        FixedCredential(CredentialModel model) {
            this.model = model;
            this.error = null;
        }

        FixedCredential(RuntimeException error) {
            this.model = null;
            this.error = error;
        }

        @Override
        public CredentialModel getCredential() {
            if (error != null) {
                throw error;
            }
            return model;
        }
    }

    private static com.aliyun.credentials.Client accessKey() throws Exception {
        Config config = new Config();
        config.type = "access_key";
        config.accessKeyId = "ak";
        config.accessKeySecret = "secret";
        return new com.aliyun.credentials.Client(config);
    }

    private static InterceptorContext.InterceptorContextRequest request(com.aliyun.credentials.Client credential) {
        InterceptorContext.InterceptorContextRequest request = new InterceptorContext.InterceptorContextRequest();
        request.headers = new HashMap<String, String>();
        request.query = new HashMap<String, String>();
        request.pathname = "/";
        request.productId = "ecs";
        request.action = "DescribeRegions";
        request.version = "2014-05-26";
        request.protocol = "HTTPS";
        request.method = "POST";
        request.authType = "AK";
        request.bodyType = "json";
        request.reqBodyType = "json";
        request.credential = credential;
        request.userAgent = "test-agent";
        return request;
    }

    private static InterceptorContext.InterceptorContextConfiguration config() {
        InterceptorContext.InterceptorContextConfiguration config = new InterceptorContext.InterceptorContextConfiguration();
        config.regionId = "cn-hangzhou";
        config.endpoint = "ecs.cn-hangzhou.aliyuncs.com";
        return config;
    }

    private static InterceptorContext context(InterceptorContext.InterceptorContextRequest request, InterceptorContext.InterceptorContextConfiguration config) {
        InterceptorContext context = new InterceptorContext();
        context.request = request;
        context.configuration = config;
        context.response = new InterceptorContext.InterceptorContextResponse();
        return context;
    }

    private static CredentialModel model(String type) {
        return CredentialModel.builder()
            .type(type)
            .accessKeyId("ak")
            .accessKeySecret("secret")
            .providerName("static_ak")
            .build();
    }

    @Test
    public void endpointAndConfiguration() throws Exception {
        Client client = new Client();
        Assert.assertEquals("ecs.example.com", client.getEndpoint("ecs", "cn-hangzhou", "regional", "", "", null, "ecs.example.com"));
        Map<String, String> endpointMap = new HashMap<String, String>();
        endpointMap.put("cn-hangzhou", "ecs.mapped.example.com");
        Assert.assertEquals("ecs.mapped.example.com", client.getEndpoint("ecs", "cn-hangzhou", "regional", "", "", endpointMap, null));
        Assert.assertEquals("ecs.cn-hangzhou.aliyuncs.com", client.getEndpoint("ecs", "cn-hangzhou", "regional", "", null, null, null));
        Assert.assertEquals("ecs.aliyuncs.com", client.getEndpoint("ecs", null, "central", "public", null, null, ""));
        try {
            client.getEndpoint("Ecs", "", "regional", "", "", null, null);
            Assert.fail();
        } catch (Exception ignored) {
        }

        InterceptorContext.InterceptorContextRequest request = request(accessKey());
        InterceptorContext.InterceptorContextConfiguration built = new InterceptorContext.InterceptorContextConfiguration();
        built.regionId = "cn-hangzhou";
        built.endpointRule = "regional";
        built.network = "";
        built.suffix = null;
        AttributeMap attributes = new AttributeMap();
        attributes.key = new HashMap<String, String>();
        attributes.key.put("signatureTypePrefix", "ACS4-");
        attributes.key.put("signPrefix", "aliyun_v4");
        attributes.key.put("endpointSuffix", "aliyuncs.com");
        client.modifyConfiguration(context(request, built), attributes);
        Assert.assertEquals("ecs.cn-hangzhou.aliyuncs.com", built.endpoint);

        InterceptorContext.InterceptorContextConfiguration kept = config();
        kept.endpoint = "already.example.com";
        client.modifyConfiguration(context(request(accessKey()), kept), new AttributeMap());
        Assert.assertEquals("already.example.com", kept.endpoint);

        InterceptorContext.InterceptorContextConfiguration broken = new InterceptorContext.InterceptorContextConfiguration();
        broken.regionId = "";
        broken.endpointRule = "regional";
        try {
            client.modifyConfiguration(context(request(accessKey()), broken), new AttributeMap());
            Assert.fail();
        } catch (Exception ignored) {
        }
    }

    @Test
    public void canonicalAndSigning() throws Exception {
        Client client = new Client();
        Assert.assertEquals("", client.buildCanonicalizedResource(null));
        Map<String, String> query = new HashMap<String, String>();
        query.put("b", "x y");
        query.put("a", "m*n");
        query.put("empty", "");
        String resource = client.buildCanonicalizedResource(query);
        Assert.assertTrue(resource.contains("a=m%2An"));
        Assert.assertTrue(resource.contains("b=x%20y"));
        Assert.assertTrue(resource.contains("empty="));

        Map<String, String> headers = new HashMap<String, String>();
        headers.put("Host", "  example.com  ");
        headers.put("X-Acs-Action", "One");
        headers.put("x-acs-action", "Two");
        String canonical = client.buildCanonicalizedHeaders(headers);
        Assert.assertTrue(canonical.contains("host:example.com\n"));
        Assert.assertTrue(canonical.contains("x-acs-action:"));

        Assert.assertEquals("fallback", client.defaultAny(null, "fallback"));
        Assert.assertEquals("kept", client.defaultAny("kept", "fallback"));

        byte[] signingkey = client.getSigningkey(client._sha256, "secret", "ecs", "cn-hangzhou", "20240101");
        Assert.assertTrue(signingkey.length > 0);
        byte[] unknown = client.getSigningkey("OTHER", "secret", "ecs", "cn-hangzhou", "20240101");
        Assert.assertEquals(0, unknown.length);

        Map<String, String> signed = new HashMap<String, String>();
        signed.put("host", "example.com");
        signed.put("x-acs-date", "2024-01-01T00:00:00Z");
        String signature = client.getSignature("", "GET", query, signed, client._sha256, "payload", signingkey);
        Assert.assertEquals(64, signature.length());
        String other = client.getSignature("/", "GET", null, signed, "OTHER", "", unknown);
        Assert.assertEquals("", other);
        String authorization = client.getAuthorization("/", "GET", null, signed, client._sha256, "payload", "ak", signingkey, "ecs", "cn-hangzhou", "20240101");
        Assert.assertTrue(authorization.startsWith("ACS4-HMAC-SHA256 Credential=ak/20240101/cn-hangzhou/ecs/aliyun_v4_request,SignedHeaders="));
        Assert.assertTrue(authorization.contains("host;x-acs-date"));
    }

    @Test
    public void modifyRequestBodiesAndAuth() throws Exception {
        Client client = new Client();
        InterceptorContext.InterceptorContextRequest json = request(accessKey());
        Map<String, Object> body = new HashMap<String, Object>();
        body.put("name", "a b");
        json.body = body;
        json.query.put("InstanceId", "i-1");
        client.modifyRequest(context(json, config()), new AttributeMap());
        Assert.assertEquals("application/json; charset=utf-8", json.headers.get("content-type"));
        Assert.assertEquals("DescribeRegions", json.headers.get("x-acs-action"));
        Assert.assertTrue(json.headers.get("Authorization").contains("ACS4-HMAC-SHA256 Credential=ak/"));
        Assert.assertTrue(json.headers.get("x-acs-content-sha256").length() > 0);

        InterceptorContext.InterceptorContextRequest bytes = request(accessKey());
        bytes.reqBodyType = "byte";
        bytes.body = "hello".getBytes("UTF-8");
        bytes.action = "";
        client.modifyRequest(context(bytes, config()), new AttributeMap());
        Assert.assertFalse(bytes.headers.containsKey("x-acs-action"));

        InterceptorContext.InterceptorContextRequest form = request(accessKey());
        form.reqBodyType = "form";
        Map<String, Object> formBody = new HashMap<String, Object>();
        formBody.put("b", "x y");
        formBody.put("a", Arrays.asList("m", "n"));
        form.body = formBody;
        client.modifyRequest(context(form, config()), new AttributeMap());
        Assert.assertEquals("application/x-www-form-urlencoded", form.headers.get("content-type"));

        InterceptorContext.InterceptorContextRequest stream = request(accessKey());
        stream.stream = new ByteArrayInputStream("stream-body".getBytes("UTF-8"));
        client.modifyRequest(context(stream, config()), new AttributeMap());
        Assert.assertEquals("application/octet-stream", stream.headers.get("content-type"));

        client.modifyRequest(context(request(accessKey()), config()), new AttributeMap());

        CredentialModel bearer = model("bearer");
        bearer.bearerToken = "bearer-token";
        bearer.providerName = "bearer";
        InterceptorContext.InterceptorContextRequest bearerRequest = request(new FixedCredential(bearer));
        client.modifyRequest(context(bearerRequest, config()), new AttributeMap());
        Assert.assertEquals("Bearer bearer-token", bearerRequest.headers.get("Authorization"));
        Assert.assertEquals("BEARERTOKEN", bearerRequest.headers.get("x-acs-signature-type"));

        CredentialModel sts = model("sts");
        sts.accessKeyId = "sts-ak";
        sts.securityToken = "token";
        sts.providerName = "";
        InterceptorContext.InterceptorContextRequest stsRequest = request(new FixedCredential(sts));
        client.modifyRequest(context(stsRequest, config()), new AttributeMap());
        Assert.assertEquals("token", stsRequest.headers.get("x-acs-security-token"));
        Assert.assertEquals("sts-ak", stsRequest.headers.get("x-acs-accesskey-id"));

        CredentialModel idToken = model("id_token");
        idToken.securityToken = "id-token";
        idToken.providerName = null;
        InterceptorContext.InterceptorContextRequest idRequest = request(new FixedCredential(idToken));
        client.modifyRequest(context(idRequest, config()), new AttributeMap());
        Assert.assertEquals("id-token", idRequest.headers.get("x-acs-zero-trust-idtoken"));

        InterceptorContext.InterceptorContextRequest anonymous = request(null);
        anonymous.authType = "Anonymous";
        client.modifyRequest(context(anonymous, config()), new AttributeMap());
        Assert.assertFalse(anonymous.headers.containsKey("Authorization"));

        InterceptorContext.InterceptorContextRequest missing = request(null);
        try {
            client.modifyRequest(context(missing, config()), new AttributeMap());
            Assert.fail();
        } catch (TeaException e) {
            Assert.assertEquals("ParameterMissing", e.getCode());
        }

        try {
            client.modifyRequest(context(request(new FixedCredential(new RuntimeException("credential down"))), config()), new AttributeMap());
            Assert.fail();
        } catch (RuntimeException ignored) {
        }

        InterceptorContext.InterceptorContextRequest badBytes = request(accessKey());
        badBytes.reqBodyType = "byte";
        badBytes.body = "not-bytes";
        try {
            client.modifyRequest(context(badBytes, config()), new AttributeMap());
            Assert.fail();
        } catch (Exception ignored) {
        }

        InterceptorContext.InterceptorContextRequest badForm = request(accessKey());
        badForm.reqBodyType = "form";
        badForm.body = "not-a-map";
        try {
            client.modifyRequest(context(badForm, config()), new AttributeMap());
            Assert.fail();
        } catch (Exception ignored) {
        }

        InterceptorContext.InterceptorContextRequest badStream = request(accessKey());
        badStream.stream = new BoomStream();
        try {
            client.modifyRequest(context(badStream, config()), new AttributeMap());
            Assert.fail();
        } catch (Exception ignored) {
        }
    }

    private static InterceptorContext response(int status, String contentType, InputStream body, String bodyType) throws Exception {
        InterceptorContext context = context(request(accessKey()), config());
        context.request.bodyType = bodyType;
        context.response.statusCode = status;
        context.response.headers = new HashMap<String, String>();
        if (contentType != null) {
            context.response.headers.put("content-type", contentType);
        }
        context.response.body = body;
        return context;
    }

    @Test
    public void modifyResponse() throws Exception {
        Client client = new Client();
        try {
            client.modifyResponse(response(400, "application/json", new ByteArrayInputStream("{\"Code\":\"Invalid\",\"Message\":\"no\",\"RequestId\":\"req-1\",\"Description\":\"d\"}".getBytes("UTF-8")), "json"), new AttributeMap());
            Assert.fail();
        } catch (TeaException e) {
            Assert.assertEquals("Invalid", e.getCode());
            Assert.assertTrue(e.getMessage().contains("request id: req-1"));
        }

        InterceptorContext denied = response(403, "application/json", new ByteArrayInputStream("{\"code\":\"Denied\",\"message\":\"no\",\"requestId\":\"body-id\"}".getBytes("UTF-8")), "json");
        denied.response.headers.put("x-acs-request-id", "header-id");
        try {
            client.modifyResponse(denied, new AttributeMap());
            Assert.fail();
        } catch (TeaException e) {
            Assert.assertEquals("Denied", e.getCode());
            Assert.assertTrue(e.getMessage().contains("request id:"));
        }

        String xml = "<Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error>";
        try {
            client.modifyResponse(response(500, "text/xml", new ByteArrayInputStream(xml.getBytes("UTF-8")), "xml"), new AttributeMap());
            Assert.fail();
        } catch (TeaException e) {
            Assert.assertEquals("XmlCode", e.getCode());
        }

        client.modifyResponse(response(204, null, new ByteArrayInputStream(new byte[0]), "string"), new AttributeMap());

        InterceptorContext binary = response(200, null, new ByteArrayInputStream("raw".getBytes("UTF-8")), "binary");
        client.modifyResponse(binary, new AttributeMap());
        Assert.assertNotNull(binary.response.deserializedBody);

        InterceptorContext bytes = response(200, null, new ByteArrayInputStream("raw".getBytes("UTF-8")), "byte");
        client.modifyResponse(bytes, new AttributeMap());
        Assert.assertEquals("raw", new String((byte[]) bytes.response.deserializedBody, "UTF-8"));

        InterceptorContext text = response(200, null, new ByteArrayInputStream("hello".getBytes("UTF-8")), "string");
        client.modifyResponse(text, new AttributeMap());
        Assert.assertEquals("hello", text.response.deserializedBody);

        InterceptorContext json = response(200, null, new ByteArrayInputStream("{\"ok\":true}".getBytes("UTF-8")), "json");
        client.modifyResponse(json, new AttributeMap());
        Assert.assertEquals(true, ((Map) json.response.deserializedBody).get("ok"));

        InterceptorContext array = response(200, null, new ByteArrayInputStream("[1,2]".getBytes("UTF-8")), "array");
        client.modifyResponse(array, new AttributeMap());
        Assert.assertNotNull(array.response.deserializedBody);

        InterceptorContext other = response(200, null, new ByteArrayInputStream("plain".getBytes("UTF-8")), "other");
        client.modifyResponse(other, new AttributeMap());
        Assert.assertEquals("plain", other.response.deserializedBody);

        try {
            client.modifyResponse(response(400, "application/json", new BoomStream(), "json"), new AttributeMap());
            Assert.fail();
        } catch (Exception ignored) {
        }
        try {
            client.modifyResponse(response(200, null, new ByteArrayInputStream("{\"ok\":".getBytes("UTF-8")), "json"), new AttributeMap());
            Assert.fail();
        } catch (Exception ignored) {
        }
    }
}
