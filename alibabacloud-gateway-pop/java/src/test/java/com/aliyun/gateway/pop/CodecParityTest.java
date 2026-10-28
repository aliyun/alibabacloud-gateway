package com.aliyun.gateway.pop;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Assert;
import org.junit.Test;

public class CodecParityTest {
    @Test
    public void matchesDarabonbaHelpers() throws Exception {
        Assert.assertEquals(
                com.aliyun.darabonba.encode.Encoder.percentEncode("a b*~"),
                Codec.percentEncode("a b*~"));
        Assert.assertNull(Codec.percentEncode(null));

        byte[] raw = new byte[] {0, 15, (byte) 255, 16};
        Assert.assertEquals(com.aliyun.darabonba.encode.Encoder.hexEncode(raw), Codec.hexEncode(raw));
        Assert.assertEquals(
                com.aliyun.darabonba.encode.Encoder.hexEncode(
                        com.aliyun.darabonba.encode.Encoder.hash("acs4".getBytes("UTF-8"), "ACS4-HMAC-SHA256")),
                Codec.hexEncode(Codec.hash("acs4".getBytes("UTF-8"), "ACS4-HMAC-SHA256")));
        Assert.assertEquals(
                com.aliyun.darabonba.encode.Encoder.hexEncode(
                        com.aliyun.darabonba.encode.Encoder.hash("sm3".getBytes("UTF-8"), "ACS4-HMAC-SM3")),
                Codec.hexEncode(Codec.hash("sm3".getBytes("UTF-8"), "ACS4-HMAC-SM3")));

        Assert.assertArrayEquals(
                com.aliyun.darabonba.signature.Signer.HmacSHA256Sign("20261027", "aliyun_v4secret"),
                Codec.hmacSha256("20261027", "aliyun_v4secret"));
        byte[] key = com.aliyun.darabonba.signature.Signer.HmacSHA256Sign("20261027", "aliyun_v4secret");
        Assert.assertArrayEquals(
                com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("cn-hangzhou", key),
                Codec.hmacSha256ByBytes("cn-hangzhou", key));
        Assert.assertArrayEquals(
                com.aliyun.darabonba.signature.Signer.HmacSM3Sign("20261027", "aliyun_v4secret"),
                Codec.hmacSm3("20261027", "aliyun_v4secret"));
        byte[] sm3Key = com.aliyun.darabonba.signature.Signer.HmacSM3Sign("20261027", "aliyun_v4secret");
        Assert.assertArrayEquals(
                com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes("Ecs", sm3Key),
                Codec.hmacSm3ByBytes("Ecs", sm3Key));

        Assert.assertEquals(
                com.aliyun.darabonbastring.Client.subString("2026-10-27T00:00:00Z", 0, 10),
                Codec.subString("2026-10-27T00:00:00Z", 0, 10));
        Assert.assertEquals(
                com.aliyun.darabonbastring.Client.replace("ecs.cn-hangzhou.aliyuncs.com", ".aliyuncs.com", "", null),
                Codec.replace("ecs.cn-hangzhou.aliyuncs.com", ".aliyuncs.com", "", null));
        Assert.assertEquals(
                com.aliyun.darabonbastring.Client.split("ecs.cn-hangzhou.aliyuncs.com:443", ":", null),
                Codec.split("ecs.cn-hangzhou.aliyuncs.com:443", ":", null));

        Map<String, String> query = new HashMap<String, String>();
        query.put("B", "b b");
        query.put("A", "a*~");
        List<String> keys = Codec.keySet(query);
        Assert.assertEquals(com.aliyun.darabonba.map.Client.keySet(query), keys);
        List<String> sorted = Codec.ascSort(keys);
        Assert.assertEquals(com.aliyun.darabonba.array.Client.ascSort(keys), sorted);
        Assert.assertEquals(com.aliyun.darabonba.array.Client.join(sorted, ";"), Codec.join(sorted, ";"));
        Assert.assertFalse(Codec.listContains(Arrays.asList("x-acs-foobar"), "x-acs-foo"));
    }

    @Test
    public void missingHostOrDateIsRejected() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("host", "ecs.cn-hangzhou.aliyuncs.com");
        try {
            client.getSignature("GET", "/", new HashMap<String, String>(), headers, client._sha256, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", new byte[] {1});
            Assert.fail("expected InvalidSignedHeaders");
        } catch (Exception ex) {
            Assert.assertTrue(ex.getMessage().contains("signed headers must include host and x-acs-date"));
        }
    }

    @Test
    public void sha256AuthorizationMatchesPreviousImplementation() throws Exception {
        Client client = new Client();
        Map<String, String> headers = new HashMap<String, String>();
        headers.put("host", "ecs.cn-hangzhou.aliyuncs.com");
        headers.put("x-acs-date", "2026-10-27T01:02:03Z");
        headers.put("x-acs-action", "DescribeRegions");
        headers.put("x-acs-content-sha256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855");
        Map<String, String> query = new HashMap<String, String>();
        query.put("RegionId", "cn-hangzhou");
        byte[] signingKey = client.getSigningkey(client._sha256, "secret", "Ecs", "cn-hangzhou", "20261027");
        String authorization = client.getAuthorization(
                "/", "GET", query, headers, client._sha256, headers.get("x-acs-content-sha256"),
                "ak", signingKey, "Ecs", "cn-hangzhou", "20261027");
        Assert.assertTrue(authorization.startsWith("ACS4-HMAC-SHA256 Credential=ak/20261027/cn-hangzhou/Ecs/aliyun_v4_request,SignedHeaders="));
        Assert.assertTrue(authorization.contains("host;"));
        Assert.assertTrue(authorization.contains("x-acs-date"));
        Assert.assertTrue(authorization.contains("Signature="));
    }
}
