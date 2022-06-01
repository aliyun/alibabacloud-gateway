## README

### Installation

+ install with `go mod` tool.

```bash
# install alibabacloud_fc_open20210406
go get github.com/alibabacloud-go/fc-open-20210406
```

### Usage

+ Invoke HTTP Trigger

```golang
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

func main() {
	config := &openapi.Config{}
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	config.SetAccessKeyId(ak)
	config.SetAccessKeySecret(sk)
	config.SetRegionId("cn-hangzhou")
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	url := "https://xxx.fcapp.run/action?key=value"
	method := "POST"
	headers := &http.Header{}
	headers.Add("k1", "v1")
	resp, err := client.InvokeHTTPTrigger(&url, &method, []byte("abc"), headers)
	if err != nil {
		panic(err)
	}
	str, _ := httputil.DumpResponse(resp.(*http.Response), true)
	fmt.Printf("response: %+v\n", string(str))
}


```

+ Invoke Anonymous HTTP Trigger

```golang
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

func main() {
	config := &openapi.Config{}
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	config.SetAccessKeyId(ak)
	config.SetAccessKeySecret(sk)
	config.SetRegionId("cn-hangzhou")
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	url := "https://xxx.fcapp.run/action?key=value"
	method := "POST"
	headers := &http.Header{}
	headers.Add("k1", "v1")
	resp, err := client.InvokeAnonymousHTTPTrigger(&url, &method,[]byte("abc"), headers)
	if err != nil {
		panic(err)
	}
	str, _ := httputil.DumpResponse(resp.(*http.Response), true)
	fmt.Printf("response: %+v\n", string(str))
}
```

+ Integration with your own http_client

```golang

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

func main() {
	config := &openapi.Config{}
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	config.SetAccessKeyId(ak)
	config.SetAccessKeySecret(sk)
	config.SetRegionId("cn-hangzhou")
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	url := "https://xxx.fcapp.run/action?key=value"
	method := "GET"
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
    }
	request, err = client.SignRequest(request)
    if err != nil {
		panic(err)
    }
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
    }
	str, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("response: %+v\n", string(str))
}

```




