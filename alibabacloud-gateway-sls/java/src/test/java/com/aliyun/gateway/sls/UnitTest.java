package com.aliyun.gateway.sls;

import com.aliyun.gateway.spi.models.InterceptorContext;
import org.junit.Assert;
import org.junit.Test;

public class UnitTest {

    @Test
    public void parseRegionTest() throws Exception {
        Client client = new Client();
        String[][] cases = new String[][]{
            {"", ""},
            {"cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"cn-hangzhou-acdr-ut-1.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"http://cn-hangzhou-acdr-ut-1.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"cn-hangzhou-acdr-ut-1-intranet.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"cn-hangzhou-acdr-ut-1-share.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"cn-hangzhou-acdr-ut-1-vpc.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"cn-hangzhou-acdr-ut-1-internal.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
            {"cn-hangzhou.sls.aliyuncs.com", "cn-hangzhou"},
            {"cn-hangzhou.log.aliyuncs.com", "cn-hangzhou"},
            {"ftp://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", ""},
            {"cn-hangzhou-acdr-ut-1.oss.aliyuncs.com", ""},
            {"cn-hangzhou-acdr-ut-1.sls.aliyuncs.com.cn", ""},
            {"cn-hangzhou-acdr-ut-1.sls.aliyuncs.com/path", ""},
            {"https://https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", ""},
            {"CN-HANGZHOU.sls.aliyuncs.com", ""},
            {"cn_hangzhou.sls.aliyuncs.com", ""},
            {"cn-hangzhou-acdr-ut-1-intranet-share.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1-intranet"},
        };
        for (int i = 0; i < cases.length; i++) {
            Assert.assertEquals(cases[i][0], cases[i][1], client.parseRegion(cases[i][0]));
        }
        Assert.assertEquals("", client.parseRegion(null));
    }

    @Test
    public void setSignV4IfInAcdrTest() throws Exception {
        Client client = new Client();

        InterceptorContext acdrRegion = context("cn-hangzhou-acdr-ut-1", "cn-hangzhou.log.aliyuncs.com", null);
        client.setSignV4IfInAcdr(acdrRegion);
        Assert.assertEquals("v4", acdrRegion.request.signatureVersion);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", acdrRegion.configuration.regionId);

        InterceptorContext ordinaryRegion = context("cn-hangzhou", "cn-hangzhou.log.aliyuncs.com", null);
        client.setSignV4IfInAcdr(ordinaryRegion);
        Assert.assertNull(ordinaryRegion.request.signatureVersion);

        InterceptorContext fromEndpoint = context(null, "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", null);
        client.setSignV4IfInAcdr(fromEndpoint);
        Assert.assertEquals("v4", fromEndpoint.request.signatureVersion);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", fromEndpoint.configuration.regionId);

        InterceptorContext ordinaryEndpoint = context("", "cn-hangzhou.sls.aliyuncs.com", null);
        client.setSignV4IfInAcdr(ordinaryEndpoint);
        Assert.assertNull(ordinaryEndpoint.request.signatureVersion);
        Assert.assertEquals("", ordinaryEndpoint.configuration.regionId);

        InterceptorContext emptyEndpoint = context(null, null, null);
        client.setSignV4IfInAcdr(emptyEndpoint);
        Assert.assertNull(emptyEndpoint.request.signatureVersion);

        InterceptorContext explicitV1 = context("cn-hangzhou-acdr-ut-1", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "v1");
        client.setSignV4IfInAcdr(explicitV1);
        Assert.assertEquals("v1", explicitV1.request.signatureVersion);

        InterceptorContext explicitV4 = context("cn-hangzhou", "cn-hangzhou.log.aliyuncs.com", "v4");
        client.setSignV4IfInAcdr(explicitV4);
        Assert.assertEquals("v4", explicitV4.request.signatureVersion);

        InterceptorContext intranet = context("", "https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com", null);
        client.setSignV4IfInAcdr(intranet);
        Assert.assertEquals("v4", intranet.request.signatureVersion);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", intranet.configuration.regionId);

        InterceptorContext invalidAcdrHint = context(null, "not-a-valid-acdr-ut-host", null);
        client.setSignV4IfInAcdr(invalidAcdrHint);
        Assert.assertNull(invalidAcdrHint.request.signatureVersion);
        Assert.assertNull(invalidAcdrHint.configuration.regionId);

        InterceptorContext shanghai = context("cn-shanghai-acdr-ut-2", null, null);
        client.setSignV4IfInAcdr(shanghai);
        Assert.assertEquals("v4", shanghai.request.signatureVersion);

        InterceptorContext explicitV1FromEndpoint = context("", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "v1");
        client.setSignV4IfInAcdr(explicitV1FromEndpoint);
        Assert.assertEquals("v1", explicitV1FromEndpoint.request.signatureVersion);
        Assert.assertEquals("", explicitV1FromEndpoint.configuration.regionId);

        client.setSignV4IfInAcdr(acdrRegion);
        Assert.assertEquals("v4", acdrRegion.request.signatureVersion);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", acdrRegion.configuration.regionId);
    }

    @Test
    public void modifyConfigurationTest() throws Exception {
        Client client = new Client();

        InterceptorContext acdr = context("cn-hangzhou-acdr-ut-1", null, null);
        client.modifyConfiguration(acdr, null);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1.log.aliyuncs.com", acdr.configuration.endpoint);
        Assert.assertEquals("v4", acdr.request.signatureVersion);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", acdr.configuration.regionId);

        InterceptorContext ordinary = context(null, null, null);
        client.modifyConfiguration(ordinary, null);
        Assert.assertEquals("cn-hangzhou.log.aliyuncs.com", ordinary.configuration.endpoint);
        Assert.assertNull(ordinary.request.signatureVersion);

        InterceptorContext fromEndpoint = context("", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", null);
        client.modifyConfiguration(fromEndpoint, null);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", fromEndpoint.configuration.endpoint);
        Assert.assertEquals("v4", fromEndpoint.request.signatureVersion);
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", fromEndpoint.configuration.regionId);

        InterceptorContext explicit = context("cn-hangzhou-acdr-ut-1", "cn-hangzhou.log.aliyuncs.com", "v1");
        client.modifyConfiguration(explicit, null);
        Assert.assertEquals("v1", explicit.request.signatureVersion);

        client.modifyConfiguration(acdr, null);
        Assert.assertEquals("v4", acdr.request.signatureVersion);
    }

    private InterceptorContext context(String regionId, String endpoint, String signatureVersion) {
        InterceptorContext context = new InterceptorContext();
        context.configuration = new InterceptorContext.InterceptorContextConfiguration();
        context.configuration.regionId = regionId;
        context.configuration.endpoint = endpoint;
        context.request = new InterceptorContext.InterceptorContextRequest();
        context.request.signatureVersion = signatureVersion;
        return context;
    }
}
