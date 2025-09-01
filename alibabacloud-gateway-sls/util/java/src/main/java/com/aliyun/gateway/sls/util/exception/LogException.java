package com.aliyun.gateway.sls.util.exception;

import com.aliyun.tea.TeaConverter;
import com.aliyun.tea.TeaException;
import com.aliyun.tea.TeaPair;

public class LogException extends TeaException {
    public LogException(String code, String message, String requestId, int statusCode) {
        super(TeaConverter.buildMap(new TeaPair("code", code),
                new TeaPair("message", message),
                new TeaPair("data", TeaConverter.buildMap(
                        new TeaPair("httpCode", statusCode),
                        new TeaPair("requestId", requestId),
                        new TeaPair("statusCode", statusCode)))));
    }

    public LogException(String code, String message, String requestId) {
        this(code, message, requestId, -1);
    }
}
