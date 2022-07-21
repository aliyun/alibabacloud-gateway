package com.aliyun.gateway.fc.util;

import com.aliyun.credentials.models.Config;
import okhttp3.Headers;
import okhttp3.Response;
import org.junit.Before;
import org.junit.Test;

import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;

/**
 * @author luoyu
 * @date 2022/6/30
 **/

public class ClientTest {


    private String ak;
    private String sk;
    private String url;
    private String host;

    Client client;

    @Before
    public void init() throws Exception {
        ak = System.getenv("ak");
        sk = System.getenv("sk");
        url = System.getenv("url");
        Config config = new Config();
        config.type = "access_key";
        config.accessKeyId = ak;
        config.accessKeySecret = sk;
        client = new Client(new com.aliyun.credentials.Client(config));
    }

    @Test
    public void testInvokeHttpTrigger() throws Exception {
        Response resp = client.InvokeHTTPTrigger(url + "/my-path?k1=v1&k2=v2", "POST", "mybodystring".getBytes(StandardCharsets.UTF_8), new Headers.Builder()
                .add("k1", "v1")
                .add("k1", "v2")
                .add("k2", "v2")
                .build());
        assert resp.toString().contains("code=200");
        String respBody = new String(resp.body().bytes(), Charset.defaultCharset());
        assert respBody.contains("\"path\": \"/my-path\",");
        assert respBody.contains("\"queries\": {\n" +
                "        \"k1\": \"v1\",\n" +
                "        \"k2\": \"v2\"\n" +
                "    },");
        assert respBody.contains("\"k1\": \"v1, v2\",\n" +
                "        \"k2\": \"v2\",");
        assert respBody.contains("\"body\": \"mybodystring\"");
    }

    @Test
    public void testInvokeHostOnly() throws Exception {
        Response resp = client.InvokeHTTPTrigger(url, "POST", "mybodystring".getBytes(StandardCharsets.UTF_8), new Headers.Builder()
                .add("k1", "v1")
                .add("k1", "v2")
                .add("k2", "v2")
                .build());
        assert resp.toString().contains("code=200");
        String respBody = new String(resp.body().bytes(), Charset.defaultCharset());
        assert respBody.contains("\"path\": \"/\",");
        assert respBody.contains("\"k1\": \"v1, v2\",\n" +
                "        \"k2\": \"v2\",");
        assert respBody.contains("\"body\": \"mybodystring\"");
    }

    @Test
    public void testInvokePostWithEmptyBody() throws Exception {
        Response resp = client.InvokeHTTPTrigger(url, "POST", new byte[]{}, new Headers.Builder()
                .add("k1", "v1")
                .add("k1", "v2")
                .add("k2", "v2")
                .build());
        assert resp.toString().contains("code=200");
        String respBody = new String(resp.body().bytes(), Charset.defaultCharset());
        assert respBody.contains("\"method\": \"POST\"");
        assert respBody.contains("\"path\": \"/\",");
        assert respBody.contains("\"k1\": \"v1, v2\",\n" +
                "        \"k2\": \"v2\",");
    }

    @Test
    public void testInvokeDELETEWithEmptyBody() throws Exception {
        Response resp = client.InvokeHTTPTrigger(url, "DELETE", null, new Headers.Builder()
                .add("k1", "v1")
                .add("k1", "v2")
                .add("k2", "v2")
                .build());
        assert resp.toString().contains("code=200");
        String respBody = new String(resp.body().bytes(), Charset.defaultCharset());
        assert respBody.contains("\"method\": \"DELETE\"");
        assert respBody.contains("\"path\": \"/\",");
        assert respBody.contains("\"k1\": \"v1, v2\",\n" +
                "        \"k2\": \"v2\",");
    }
}
