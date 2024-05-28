// This file is auto-generated, don't edit it. Thanks.
package http

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

const defaultHTTPIdleTimeout = 55 * time.Second
const defaultRequestTimeout = 60 * time.Second

var clientPool = &sync.Map{}

type HttpClient struct {
	tea.HttpClient
}

func NewHttpClient() (*HttpClient, error) {
	client := new(HttpClient)
	err := client.Init()
	return client, err
}

func (client *HttpClient) Init() (_err error) {
	return nil
}

func (client *HttpClient) Call(request *http.Request, transport *http.Transport) (_result *http.Response, _err error) {
	slsClient := getHttpClient(getClientTag(request.Host, request.URL))
	slsClient.onceFlag.Do(func() {
		if transport != nil {
			trans := transport.Clone()
			if trans.IdleConnTimeout == 0 {
				trans.IdleConnTimeout = defaultHTTPIdleTimeout
			}
			slsClient.httpClient.Transport = trans
		} else {
			slsClient.httpClient.Transport = newDefaultTransport()
		}
	})
	return slsClient.httpClient.Do(request)
}

type slsHttpClient struct {
	httpClient *http.Client
	onceFlag   sync.Once
}

func newDefaultTransport() *http.Transport {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.IdleConnTimeout = defaultHTTPIdleTimeout
	return t
}

func getHttpClient(tag string) *slsHttpClient {
	client, ok := clientPool.Load(tag)
	if client == nil && !ok {
		client = &slsHttpClient{
			httpClient: &http.Client{
				Timeout: defaultRequestTimeout,
			},
		}
		clientPool.Store(tag, client)
	}
	return client.(*slsHttpClient)
}

func getClientTag(domain string, url *url.URL) string {
	return domain + ":" + url.Scheme + ":" + url.Host + ":" + url.Path
}
