// This file is auto-generated, don't edit it. Thanks.
package client

import (
	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	spi.Client
}

func NewClient() (*Client, error) {
	client := new(Client)
	err := client.Init()
	return client, err
}

func (client *Client) Init() (_err error) {
	_err = client.Client.Init()
	if _err != nil {
		return _err
	}
	return nil
}

func (client *Client) ModifyConfiguration(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	return _err
}

func (client *Client) ModifyRequest(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	request.Headers = tea.Merge(map[string]*string{
		"host":       config.Endpoint,
		"user-agent": request.UserAgent,
	}, request.Headers)
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		jsonObj := util.ToJSONString(request.Body)
		request.Stream = tea.ToReader(jsonObj)
		request.Headers["content-type"] = tea.String("application/json; charset=utf-8")
	}

	return _err
}

func (client *Client) ModifyResponse(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	response := context.Response
	if tea.BoolValue(util.Is4xx(response.StatusCode)) || tea.BoolValue(util.Is5xx(response.StatusCode)) {
		_res, _err := util.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		err, _err := util.AssertAsMap(_res)
		if _err != nil {
			return _err
		}

		err["statusCode"] = response.StatusCode
		_err = tea.NewSDKError(map[string]interface{}{
			"code":               tea.ToString(client.DefaultAny(err["Code"], err["code"])),
			"message":            "code: " + tea.ToString(tea.IntValue(response.StatusCode)) + ", " + tea.ToString(client.DefaultAny(err["Message"], err["message"])) + " request id: " + tea.ToString(client.DefaultAny(err["requestId"], err["requestid"])),
			"data":               err,
			"description":        tea.ToString(client.DefaultAny(err["Description"], err["description"])),
			"accessDeniedDetail": client.DefaultAny(err["accessDeniedDetail"], err["accessdenieddetail"]),
		})
		return _err
	}

	if tea.BoolValue(util.EqualNumber(response.StatusCode, tea.Int(204))) {
		_, _err = util.ReadAsString(response.Body)
		if _err != nil {
			return _err
		}
	} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("binary"))) {
		response.DeserializedBody = response.Body
	} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("byte"))) {
		byt, _err := util.ReadAsBytes(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = byt
	} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("string"))) {
		str, _err := util.ReadAsString(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = str
	} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("json"))) {
		obj, _err := util.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		res, _err := util.AssertAsMap(obj)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = res
	} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("array"))) {
		arr, _err := util.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = arr
	} else {
		response.DeserializedBody, _err = util.ReadAsString(response.Body)
		if _err != nil {
			return _err
		}

	}

	return _err
}

func (client *Client) DefaultAny(inputValue interface{}, defaultValue interface{}) (_result interface{}) {
	if tea.BoolValue(util.IsUnset(inputValue)) {
		_result = defaultValue
		return _result
	}

	_result = inputValue
	return _result
}
