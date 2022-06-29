// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.fc.util;

import com.aliyun.tea.*;
import com.aliyun.credentials.*;
import org.apache.hc.core5.http.HttpRequest;
import org.apache.hc.core5.http.HttpResponse;
import org.apache.hc.core5.http.HttpHeaders;

public class Client {

    public com.aliyun.credentials.Client _credential;
    public Client(com.aliyun.credentials.Client cred) throws Exception {
        this._credential = cred;
    }


    public HttpResponse InvokeHTTPTrigger(String url, String method, byte[] body, HttpHeaders headers) throws Exception {
    }

    public HttpResponse InvokeAnonymousHTTPTrigger(String url, String method, byte[] body, HttpHeaders headers) throws Exception {
    }

    public HttpResponse SendHTTPRequestWithAuthorization(HttpRequest req) throws Exception {
    }

    public HttpResponse SendHTTPRequest(HttpRequest req) throws Exception {
    }

    public HttpRequest SignRequest(HttpRequest req) throws Exception {
    }

    public HttpRequest SignRequestWithContentMD5(HttpRequest req, String contentMD5) throws Exception {
    }

    public HttpRequest BuildHTTPRequest(String url, String method, byte[] body, HttpHeaders headers) throws Exception {
    }
}
