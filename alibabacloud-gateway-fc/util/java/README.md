## README


### Download

+ Download from [Maven Central](https://search.maven.org/search?q=g:com.aliyun%20AND%20a:fc_open20210406), or depend via Maven or Gradle:

```xml
<dependency>
  <groupId>com.aliyun</groupId>
  <artifactId>fc_open20210406</artifactId>
  <version>1.1.6</version>
</dependency>
```


### Usage


+ Invoke HTTP Trigger
```java

package com.aliyun.example;

import com.aliyun.fc_open20210406.Client;
import com.aliyun.teaopenapi.models.Config;
import okhttp3.Headers;
import okhttp3.Response;

import java.nio.charset.StandardCharsets;


public class Main {
    public static void main(String[] args) throws Exception {
        String ak = System.getenv("ak");
        String sk = System.getenv("sk");
        String url = System.getenv("url");
        Config config = new Config().setAccessKeyId(ak).setAccessKeySecret(sk).setRegionId("cn-hangzhou");

        Client client = new Client(config);
        try (Response res = client.InvokeHTTPTrigger(url, "POST", "mybodystring".getBytes(StandardCharsets.UTF_8), new Headers.Builder().build())) {
            System.out.println(res.toString());
            System.out.println(res.body().string());
        }
    }
}



```



+ Invoke Anonymous HTTP Trigger

```java

package com.aliyun.example;

import com.aliyun.fc_open20210406.Client;
import com.aliyun.teaopenapi.models.Config;
import okhttp3.Headers;
import okhttp3.Response;

import java.nio.charset.StandardCharsets;

public class Main {
    public static void main(String[] args) throws Exception {
        String ak = System.getenv("ak");
        String sk = System.getenv("sk");
        String url = System.getenv("url");
        Config config = new Config().setAccessKeyId(ak).setAccessKeySecret(sk).setRegionId("cn-hangzhou");

        Client client = new Client(config);
        try (Response res = client.InvokeAnonymousHTTPTrigger(url, "POST", "mybodystring".getBytes(StandardCharsets.UTF_8), new Headers.Builder().build())) {
            System.out.println(res.toString());
            System.out.println(res.body().string());
        }
    }
}



```

+ Integration with your own OKHttpClient


```java

package com.aliyun.example;

import com.aliyun.fc_open20210406.Client;
import com.aliyun.tea.okhttp.OkHttpClientBuilder;
import com.aliyun.teaopenapi.models.Config;
import okhttp3.*;

import java.nio.charset.StandardCharsets;


public class Main {
    public static void main(String[] args) throws Exception {
        String ak = System.getenv("ak");
        String sk = System.getenv("sk");
        String url = System.getenv("url");
        Config config = new Config().setAccessKeyId(ak).setAccessKeySecret(sk).setRegionId("cn-hangzhou");

        Client client = new Client(config);
        OkHttpClient okHttpClient = new OkHttpClientBuilder().buildOkHttpClient();

        Request request = new Request.Builder()
                .url(url)
                .post(RequestBody.create(MediaType.parse("application/json"),
                        "mybodystring".getBytes(StandardCharsets.UTF_8)))
                .build();

        // sign your request
        request = client.SignRequest(request);
        try (Response res = okHttpClient.newCall(request).execute()) {
            System.out.println(res.toString());
            System.out.println(res.body().string());
        }
    }
}


```