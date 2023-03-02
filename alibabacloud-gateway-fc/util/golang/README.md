## README

### Installation

+ install with `go mod` tool.

```bash
# install alibabacloud_fc_open20210406
go get -u github.com/alibabacloud-go/fc-open-20210406
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

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	client "github.com/alibabacloud-go/fc-open-20210406/client"

)

func main() {
	config := &openapi.Config{}
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	url := os.Getenv("url")
	
	config.SetAccessKeyId(ak)
	config.SetAccessKeySecret(sk)
	config.SetRegionId("cn-hangzhou")
	c, err := client.NewClient(config)
	if err != nil {
		panic(err)
	}
	method := "POST"
	headers := &http.Header{}
	headers.Add("k1", "v1")
	resp, err := c.InvokeHTTPTrigger(&url, &method, []byte("abc"), headers)
	if err != nil {
		panic(err)
	}
	str, _ := httputil.DumpResponse(resp, true)
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

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	client "github.com/alibabacloud-go/fc-open-20210406/client"
)

func main() {
	config := &openapi.Config{}
	ak := "dummy-ak"
	sk := "dummy-sk"
	url := os.Getenv("url")
	
	config.SetAccessKeyId(ak)
	config.SetAccessKeySecret(sk)
	config.SetRegionId("cn-hangzhou")
	c, err := client.NewClient(config)
	if err != nil {
		panic(err)
	}
	method := "POST"
	headers := &http.Header{}
	headers.Add("k1", "v1")
	resp, err := c.InvokeAnonymousHTTPTrigger(&url, &method, []byte("abc"), headers)
	if err != nil {
		panic(err)
	}
	str, _ := httputil.DumpResponse(resp, true)
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

	client "github.com/alibabacloud-go/fc-open-20210406/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

func main() {
	config := &openapi.Config{}
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	url := os.Getenv("url")

	config.SetAccessKeyId(ak)
	config.SetAccessKeySecret(sk)
	config.SetRegionId("cn-hangzhou")
	c, err := client.NewClient(config)
	if err != nil {
		panic(err)
	}
	method := "GET"
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	request, err = c.SignRequest(request)
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




