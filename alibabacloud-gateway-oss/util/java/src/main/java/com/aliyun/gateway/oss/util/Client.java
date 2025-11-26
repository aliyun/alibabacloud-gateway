// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util;

import com.aliyun.tea.*;

public class Client {

    public static Object parseXml(String bodyStr, String apiName) throws Exception {
        Class bodyClass = null;
        try {
            try {
                bodyClass = Class.forName("com.aliyun.gateway.oss.util.models." + apiName + "ResponseBody");
            } catch (ClassNotFoundException e) {
                try {
                    bodyClass = Class.forName("com.aliyun.gateway.oss.util.hcs_mgw_models." + apiName + "ResponseBody");
                } catch (ClassNotFoundException e) {}
            }
            return com.aliyun.teaxml.Client.parseXml(bodyStr, bodyClass); 
        } catch (Exception e) {
            throw new TeaException(TeaConverter.buildMap(
                    new TeaPair("code", "ParseXmlError"),
                    new TeaPair("message", "Parse xml error: " + e.getMessage()),
                    new TeaPair("data", TeaConverter.buildMap(
                            new TeaPair("apiName", apiName),
                            new TeaPair("bodyStr", bodyStr)
                    ))
            ));
        }
    }
}
