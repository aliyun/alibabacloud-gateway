// This file is auto-generated, don't edit it. Thanks.
package client

import (
	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	spi.Client
}

func NewClient(config *spi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) ModifyConfiguration(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	config.Endpoint, _err = client.GetEndpoint(request.ProductId, config.RegionId, config.EndpointRule, config.Network, config.Suffix, config.EndpointMap, config.Endpoint)
	if _err != nil {
		return _err
	}

	return _err
}

func (client *Client) ModifyRequest(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	credential := request.Credential
	request.Query = tea.Merge(map[string]*string{
		"Action":         request.Action,
		"Format":         tea.String("json"),
		"Version":        request.Version,
		"Timestamp":      openapiutil.GetTimestamp(),
		"SignatureNonce": util.GetNonce(),
	}, request.Query)
	request.Headers = map[string]*string{
		"host":          config.Endpoint,
		"x-acs-version": request.Version,
		"x-acs-action":  request.Action,
		"user-agent":    request.UserAgent,
	}
	var t map[string]interface{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		t = util.AssertAsMap(request.Body)
		tmp := util.AnyifyMapValue(openapiutil.Query(t))
		request.Stream = tea.ToReader(util.ToFormString(tmp))
		request.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
	}

	if !tea.BoolValue(util.EqualString(request.AuthType, tea.String("Anonymous"))) {
		accessKeyId, _err := credential.GetAccessKeyId()
		if _err != nil {
			return _err
		}

		accessKeySecret, _err := credential.GetAccessKeySecret()
		if _err != nil {
			return _err
		}

		securityToken, _err := credential.GetSecurityToken()
		if _err != nil {
			return _err
		}

		if !tea.BoolValue(util.Empty(securityToken)) {
			request.Query["SecurityToken"] = securityToken
		}

		request.Query["SignatureMethod"] = tea.String("HMAC-SHA1")
		request.Query["SignatureVersion"] = tea.String("1.0")
		request.Query["AccessKeyId"] = accessKeyId
		signedParam := tea.Merge(request.Query,
			openapiutil.Query(t))
		request.Query["Signature"] = openapiutil.GetRPCSignature(signedParam, request.Method, accessKeySecret)
	}

	return _err
}

func (client *Client) ModifyResponse(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	response := context.Response
	if tea.BoolValue(util.Is4xx(response.StatusCode)) || tea.BoolValue(util.Is5xx(response.StatusCode)) {
		_res, _err := util.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		err := util.AssertAsMap(_res)
		requestId := client.DefaultAny(err["RequestId"], err["requestId"])
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    tea.ToString(client.DefaultAny(err["Code"], err["code"])),
			"message": "code: " + tea.ToString(tea.IntValue(response.StatusCode)) + ", " + tea.ToString(client.DefaultAny(err["Message"], err["message"])) + " request id: " + tea.ToString(requestId),
			"data":    err,
		})
		return _err
	}

	if tea.BoolValue(util.EqualString(request.BodyType, tea.String("binary"))) {
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

		res := util.AssertAsMap(obj)
		response.DeserializedBody = res
	} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("array"))) {
		arr, _err := util.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = arr
	}

	return _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DefaultAny(inputValue interface{}, defaultValue interface{}) (_result interface{}) {
	if tea.BoolValue(util.IsUnset(inputValue)) {
		_result = defaultValue
		return _result
	}

	_result = inputValue
	return _result
}
