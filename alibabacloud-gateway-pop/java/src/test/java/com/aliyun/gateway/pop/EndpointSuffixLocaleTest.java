package com.aliyun.gateway.pop;

import java.util.Locale;
import org.junit.Assert;
import org.junit.Test;

public class EndpointSuffixLocaleTest {
    @Test
    public void uppercaseSuffixMustNotDependOnDefaultLocale() throws Exception {
        Locale previous = Locale.getDefault();
        try {
            Locale.setDefault(new Locale("tr", "TR"));
            Assert.assertEquals("aliyuncs.alibaba", new Client().checkedEndpointSuffix("ALIYUNCS.ALIBABA"));
        } finally {
            Locale.setDefault(previous);
        }
    }
}
