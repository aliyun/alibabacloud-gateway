package com.aliyun.gateway.sls;

import com.aliyun.gateway.spi.models.InterceptorContext;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import java.util.Map;
import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class SignatureVersionTest {
    @Test
    @SuppressWarnings("unchecked")
    public void selectsSignatureVersion() throws Exception {
        String json = new String(Files.readAllBytes(Paths.get("../testdata/signature-version.json")), StandardCharsets.UTF_8);
        List<Map<String, Object>> cases = (List<Map<String, Object>>) com.aliyun.teautil.Common.parseJSON(json);
        Client client = new Client();
        for (Map<String, Object> test : cases) {
            InterceptorContext context = new InterceptorContext();
            context.request = new InterceptorContext.InterceptorContextRequest();
            context.configuration = new InterceptorContext.InterceptorContextConfiguration();
            context.request.signatureVersion = (String) test.get("version");
            context.configuration.regionId = (String) test.get("region");
            context.configuration.endpoint = (String) test.get("endpoint");
            String name = (String) test.get("name");
            assertEquals(name, test.get("want"), client.getSignatureVersion(context));
            assertEquals(name, test.get("resolved"), context.configuration.regionId);
            assertEquals(name, test.get("version"), context.request.signatureVersion);
        }
    }
}
