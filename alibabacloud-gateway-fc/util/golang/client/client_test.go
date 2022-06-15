package client

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"

	credential "github.com/aliyun/credentials-go/credentials"
	"github.com/stretchr/testify/assert"
)

func initClient() (*Client, error) {
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	cred := &credential.AccessKeyCredential{
		AccessKeyId:     ak,
		AccessKeySecret: sk,
	}
	return NewClient(cred)
}

func TestVisitNoPathURL(t *testing.T) {
	c, err := initClient()
	assert.Nil(t, err)
	assert.Nil(t, err)
	url := os.Getenv("url")
	method := "GET"
	body := []byte("")
	h := http.Header{}
	resp, err := c.InvokeHTTPTrigger(&url, &method, body, &h)
	assert.Nil(t, err)
	str, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("resp: %+v\n", string(str))
	assert.Contains(t, string(str), "200 OK")
}

func TestVisitComplexURL(t *testing.T) {
	c, err := initClient()
	assert.Nil(t, err)
	assert.Nil(t, err)
	path := "/my-path"
	queryString := "k1=v1&k2=v2"
	url := os.Getenv("url") + path + "?" + queryString
	method := "POST"
	body := []byte("mybodystring")
	h := http.Header{
		"k1": {"v1", "v2"},
		"k2": {"v2"},
	}
	resp, err := c.InvokeHTTPTrigger(&url, &method, body, &h)
	assert.Nil(t, err)
	str, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("resp: %+v\n", string(str))
	assert.Contains(t, string(str), "200 OK")
	assert.Contains(t, string(str), " \"path\": \""+path+"\",")
	assert.Contains(t, string(str), "\"queries\": {\n        \"k1\": \"v1\",\n        \"k2\": \"v2\"\n    },\n")
	assert.Contains(t, string(str), "\"requestURI\": \"/my-path?k1=v1&k2=v2\",")
}
