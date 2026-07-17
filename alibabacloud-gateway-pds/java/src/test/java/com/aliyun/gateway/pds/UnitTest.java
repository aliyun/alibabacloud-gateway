package com.aliyun.gateway.pds;

import java.util.HashMap;
import java.util.regex.Pattern;

import com.aliyun.credentials.models.Config;
import com.aliyun.gateway.spi.models.AttributeMap;
import com.aliyun.gateway.spi.models.InterceptorContext;
import org.junit.Assert;
import org.junit.Test;

public class UnitTest {

    private static final Pattern ISO8601 = Pattern.compile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$");
    private static final Pattern YYYYMMDD = Pattern.compile("^\\d{8}$");

    @Test
    public void getRegionTest() throws Exception {
        Client client = new Client();
        Assert.assertEquals("center", client.getRegion(null));
        Assert.assertEquals("center", client.getRegion(""));
        Assert.assertEquals("center", client.getRegion("api.aliyunpds.com"));
        Assert.assertEquals("cn-hangzhou", client.getRegion("cn-hangzhou.admin.aliyunpds.com"));
    }

    @Test
    public void modifyRequestV4UsesIsoDateNotRfc1123() throws Exception {
        Client client = new Client();

        Config credConfig = new Config();
        credConfig.type = "access_key";
        credConfig.accessKeyId = "ak";
        credConfig.accessKeySecret = "sk";
        com.aliyun.credentials.Client credential = new com.aliyun.credentials.Client(credConfig);

        InterceptorContext.InterceptorContextRequest request = new InterceptorContext.InterceptorContextRequest();
        request.headers = new HashMap<String, String>();
        request.query = new HashMap<String, String>();
        request.pathname = "/v2/file/get";
        request.productId = "pds";
        request.action = "GetFile";
        request.version = "2022-03-01";
        request.protocol = "HTTPS";
        request.method = "POST";
        request.authType = "AK";
        request.bodyType = "json";
        request.reqBodyType = "json";
        request.userAgent = "UnitTest";
        request.signatureVersion = "v4";
        request.signatureAlgorithm = "ACS4-HMAC-SHA256";
        request.credential = credential;

        InterceptorContext.InterceptorContextConfiguration configuration =
            new InterceptorContext.InterceptorContextConfiguration();
        configuration.endpoint = "api.aliyunpds.com";

        InterceptorContext context = new InterceptorContext();
        context.request = request;
        context.configuration = configuration;
        context.response = new InterceptorContext.InterceptorContextResponse();

        client.modifyRequest(context, new AttributeMap());

        String acsDate = request.headers.get("x-acs-date");
        Assert.assertNotNull("V4 must set x-acs-date", acsDate);
        Assert.assertTrue("x-acs-date must be ISO8601, got: " + acsDate, ISO8601.matcher(acsDate).matches());

        String authorization = request.headers.get("Authorization");
        Assert.assertNotNull(authorization);
        Assert.assertTrue(authorization.startsWith("ACS4-HMAC-SHA256 Credential=ak/"));

        String credentialScopeDate = authorization.substring("ACS4-HMAC-SHA256 Credential=ak/".length(),
            "ACS4-HMAC-SHA256 Credential=ak/".length() + 8);
        Assert.assertTrue("credential scope date must be yyyyMMdd, got: " + credentialScopeDate,
            YYYYMMDD.matcher(credentialScopeDate).matches());
        Assert.assertFalse("must not slice RFC1123 Date header", credentialScopeDate.contains(","));
        Assert.assertEquals(acsDate.substring(0, 10).replace("-", ""), credentialScopeDate);
    }
}
