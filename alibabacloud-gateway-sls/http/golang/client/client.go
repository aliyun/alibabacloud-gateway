// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	tea.HttpClient
}

func NewClient() (*Client, error) {
	client := new(Client)
	err := client.Init()
	return client, err
}

func (client *Client) Init() (_err error) {
	_err = client.HttpClient.Init()
	if _err != nil {
		return _err
	}
	return nil
}

func (client *Client) Call(request *http.Request) (response *http.Response, err error) {
	client.HttpClient.Transport = new(http.Transport)
	client.HttpClient.Transport.MaxIdleConns = 128
	client.HttpClient.Transport.MaxIdleConnsPerHost = 128
	response, err = client.HttpClient.Do(request)
	return
}
