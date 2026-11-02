package com.aliyun.gateway.pop;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.StringWriter;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.aliyun.credentials.models.Config;
import com.aliyun.gateway.spi.models.AttributeMap;
import com.aliyun.gateway.spi.models.InterceptorContext;
import com.aliyun.tea.TeaModel;
import com.google.gson.stream.JsonWriter;
import org.junit.Assert;
import org.junit.Test;

public class HelperCoverageTest {
    static class NamedModel extends TeaModel {
        @Override
        public Map<String, Object> toMap() {
            Map<String, Object> map = new HashMap<String, Object>();
            map.put("name", "n");
            return map;
        }
    }

    static class ClosingStream extends ByteArrayInputStream {
        ClosingStream(byte[] body) {
            super(body);
        }

        @Override
        public void close() throws IOException {
            throw new IOException("close failed");
        }
    }

    @Test
    public void supportEdges() throws Exception {
        Assert.assertEquals(0, Support.readAsBytes(null).length);
        try {
            Support.readAsBytes(new ClosingStream("a".getBytes("UTF-8")));
            Assert.fail("close should fail");
        } catch (RuntimeException ex) {
            Assert.assertTrue(ex.getMessage().contains("close failed"));
        }

        Assert.assertEquals("fallback", Support.defaultString("", "fallback"));
        Assert.assertEquals("fallback", Support.defaultString(null, "fallback"));
        Assert.assertEquals("already", Support.toJSONString("already"));
        Assert.assertFalse(Support.equalString(null, "a"));
        Assert.assertFalse(Support.equalString("a", null));
        Assert.assertFalse(Support.equalNumber(null, 1));
        Assert.assertFalse(Support.equalNumber(1, null));
        Assert.assertFalse(Support.is4xx(null));
        Assert.assertFalse(Support.is5xx(null));

        Assert.assertEquals(
            "ecs-inner.cn-hangzhou.aliyuncs.com",
            Support.getEndpointRules("ecs", "cn-hangzhou", "regional", "inner", null));

        Map<String, Object> form = new HashMap<String, Object>();
        form.put("empty", "");
        form.put("keep", "x");
        form.put("skip", null);
        form.put("bytes", "hi".getBytes("UTF-8"));
        List<Object> items = new ArrayList<Object>();
        items.add("a");
        form.put("items", items);
        Map<String, Object> nested = new HashMap<String, Object>();
        nested.put("k", "v");
        form.put("nested", nested);
        form.put("model", new NamedModel());
        String encoded = Support.toForm(form);
        Assert.assertTrue(encoded.contains("keep=x"));
        Assert.assertTrue(encoded.contains("bytes=hi"));
        Assert.assertTrue(encoded.contains("items.1=a"));
        Assert.assertTrue(encoded.contains("nested.k=v"));
        Assert.assertTrue(encoded.contains("model.name=n"));
        Assert.assertFalse(encoded.contains("empty="));

        Assert.assertTrue(Support.parseXml("").isEmpty());
        Assert.assertTrue(Support.parseXml(null).isEmpty());
        Map<String, Object> text = Support.parseXml("<Error>  hi  </Error>");
        Assert.assertEquals("hi", text.get("Error"));
        Map<String, Object> repeated = Support.parseXml("<Error><Item>a</Item><Item>b</Item><Item>c</Item></Error>");
        Map<String, Object> error = Support.assertAsMap(repeated.get("Error"));
        Assert.assertEquals(Arrays.asList("a", "b", "c"), error.get("Item"));

        Map<String, Object> parsed = Support.assertAsMap(Support.parseJSON(
            "{\"n\":1,\"d\":1.5,\"e\":1E2,\"ok\":true,\"nil\":null,\"arr\":[2,{\"k\":\"v\"}]}"));
        Assert.assertEquals(Long.valueOf(1L), parsed.get("n"));
        Assert.assertEquals(Double.valueOf(1.5d), parsed.get("d"));
        Assert.assertEquals(Double.valueOf(100d), parsed.get("e"));
        Assert.assertEquals(Boolean.TRUE, parsed.get("ok"));
        Assert.assertNull(parsed.get("nil"));
        Assert.assertTrue(parsed.get("arr") instanceof List);
    }

    @Test
    public void codecEdges() throws Exception {
        Assert.assertNull(Codec.hash(new byte[]{1}, null));
        Assert.assertEquals("ab", Codec.subString("abcdef", 0, -4));
        Assert.assertEquals("", Codec.subString("ab", 0, -5));
        Assert.assertEquals("a:b-c", Codec.replace("a-b-c", "-", ":", 1));
        Assert.assertEquals("", Codec.replace("", "a", "b", 1));

        try {
            Codec.keySet(null);
            Assert.fail("null map");
        } catch (IllegalArgumentException expected) {
            Assert.assertTrue(expected.getMessage().contains("not a valid value"));
        }
        try {
            Codec.join(null, ",");
            Assert.fail("null list");
        } catch (IllegalArgumentException expected) {
            Assert.assertTrue(expected.getMessage().contains("not a valid value"));
        }
        Assert.assertEquals("[]", Codec.join(new ArrayList<String>(), ","));
        try {
            Codec.size(null);
            Assert.fail("null size");
        } catch (IllegalArgumentException expected) {
            Assert.assertTrue(expected.getMessage().contains("not a valid value"));
        }
        try {
            Codec.ascSort(null);
            Assert.fail("null sort");
        } catch (IllegalArgumentException expected) {
            Assert.assertTrue(expected.getMessage().contains("not a valid value"));
        }
        try {
            Codec.listContains(null, "a");
            Assert.fail("null contains");
        } catch (IllegalArgumentException expected) {
            Assert.assertTrue(expected.getMessage().contains("not a valid value"));
        }
        try {
            Codec.listContains(Arrays.asList("a"), null);
            Assert.fail("null value");
        } catch (IllegalArgumentException expected) {
            Assert.assertTrue(expected.getMessage().contains("not a valid value"));
        }

        MapTypeAdapter adapter = new MapTypeAdapter();
        StringWriter writer = new StringWriter();
        JsonWriter json = new JsonWriter(writer);
        adapter.write(json, "x");
        json.flush();
        Assert.assertEquals("\"x\"", writer.toString());
    }

    @Test
    public void sm3RequestUsesTheOptionalProvider() throws Exception {
        Client client = new Client();
        Config config = new Config();
        config.type = "access_key";
        config.accessKeyId = "ak";
        config.accessKeySecret = "secret";
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
        request.bodyType = "string";
        request.reqBodyType = "string";
        request.credential = new com.aliyun.credentials.Client(config);
        request.userAgent = "test-agent";
        request.signatureAlgorithm = client._sm3;
        InterceptorContext.InterceptorContextConfiguration configuration = new InterceptorContext.InterceptorContextConfiguration();
        configuration.regionId = "cn-hangzhou";
        configuration.endpoint = "ecs.cn-hangzhou.aliyuncs.com";
        InterceptorContext context = new InterceptorContext();
        context.request = request;
        context.configuration = configuration;
        client.modifyRequest(context, new AttributeMap());
        Assert.assertTrue(request.headers.get("Authorization").contains("ACS4-HMAC-SM3 Credential=ak/"));
        Assert.assertNotNull(request.headers.get("x-acs-content-sm3"));
        byte[] key = client.getSigningkey(client._sm3, "secret", "ecs", "cn-hangzhou", "20240101");
        String signature = client.getSignature("/", "POST", request.query, request.headers, client._sm3, "payload", key);
        Assert.assertEquals(64, signature.length());
    }
}
