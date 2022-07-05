// This file is auto-generated, don't edit it. Thanks.
package client

import (
	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	array "github.com/alibabacloud-go/darabonba-array/client"
	encodeutil "github.com/alibabacloud-go/darabonba-encode-util/client"
	map_ "github.com/alibabacloud-go/darabonba-map/client"
	signatureutil "github.com/alibabacloud-go/darabonba-signature-util/client"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	config := context.Configuration
	config.Endpoint, _err = client.GetEndpoint(config.Network, config.Endpoint)
	if _err != nil {
		return _err
	}

	return _err
}

func (client *Client) ModifyRequest(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	request.Headers = tea.Merge(map[string]*string{
		"date":                    util.GetDateUTCString(),
		"host":                    config.Endpoint,
		"x-acs-version":           request.Version,
		"x-acs-action":            request.Action,
		"user-agent":              request.UserAgent,
		"x-acs-signature-nonce":   util.GetNonce(),
		"x-acs-signature-method":  tea.String("HMAC-SHA1"),
		"x-acs-signature-version": tea.String("1.0"),
		"accept":                  tea.String("application/json"),
	}, request.Headers)
	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		tmp, _err := util.ReadAsBytes(request.Stream)
		if _err != nil {
			return _err
		}

		request.Stream = tea.ToReader(tmp)
		request.Headers["content-type"] = tea.String("application/octet-stream")
	} else {
		if !tea.BoolValue(util.IsUnset(request.Body)) {
			if tea.BoolValue(util.EqualString(request.ReqBodyType, tea.String("json"))) {
				jsonObj := util.ToJSONString(request.Body)
				request.Stream = tea.ToReader(jsonObj)
				request.Headers["content-type"] = tea.String("application/json; charset=utf-8")
			} else {
				m := util.AssertAsMap(request.Body)
				formObj := openapiutil.ToForm(m)
				request.Stream = tea.ToReader(formObj)
				request.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			}

		}

	}

	if !tea.BoolValue(util.EqualString(request.AuthType, tea.String("Anonymous"))) {
		credential := request.Credential
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

		bearerToken := credential.GetBearerToken()
		if !tea.BoolValue(util.Empty(bearerToken)) {
			request.Headers["x-acs-bearer-token"] = bearerToken
			request.Headers["Authorization"] = tea.String("Bearer " + tea.StringValue(bearerToken))
		} else {
			if !tea.BoolValue(util.Empty(securityToken)) {
				request.Headers["x-acs-security-token"] = securityToken
			}

			request.Headers["Authorization"], _err = client.GetAuthorization(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret)
			if _err != nil {
				return _err
			}

		}

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

		err := util.AssertAsMap(_res)
		requestId := client.DefaultAny(err["RequestId"], err["requestId"])
		err["statusCode"] = response.StatusCode
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    tea.ToString(client.DefaultAny(err["Code"], err["code"])),
			"message": "code: " + tea.ToString(tea.IntValue(response.StatusCode)) + ", " + tea.ToString(client.DefaultAny(err["Message"], err["message"])) + " request id: " + tea.ToString(requestId),
			"data":    err,
		})
		return _err
	}

	if !tea.BoolValue(util.IsUnset(response.Body)) && !tea.BoolValue(util.EqualNumber(response.StatusCode, tea.Int(204))) {
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
			response.DeserializedBody, _err = util.ReadAsJSON(response.Body)
			if _err != nil {
				return _err
			}

		} else {
			response.DeserializedBody, _err = util.ReadAsString(response.Body)
			if _err != nil {
				return _err
			}

		}

	}

	return _err
}

func (client *Client) GetEndpoint(network *string, endpoint *string) (_result *string, _err error) {
	realEndpoint := tea.String("api.aliyunpds.com")
	if !tea.BoolValue(util.Empty(endpoint)) {
		realEndpoint = endpoint
	}

	if !tea.BoolValue(util.Empty(network)) && tea.BoolValue(string_.Equals(network, tea.String("vpc"))) {
		realEndpoint = string_.Replace(realEndpoint, tea.String("api.aliyunpds.com"), tea.String("api-vpc.aliyunpds.com"), nil)
		realEndpoint = string_.Replace(realEndpoint, tea.String("admin.aliyunpds.com"), tea.String("admin-vpc.aliyunpds.com"), nil)
	}

	_result = realEndpoint
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

func (client *Client) GetAuthorization(pathname *string, method *string, query map[string]*string, headers map[string]*string, ak *string, secret *string) (_result *string, _err error) {
	signature, _err := client.GetSignature(pathname, method, query, headers, secret)
	if _err != nil {
		return _result, _err
	}

	_result = tea.String("acs " + tea.StringValue(ak) + ":" + tea.StringValue(signature))
	return _result, _err
}

func (client *Client) GetSignature(pathname *string, method *string, query map[string]*string, headers map[string]*string, secret *string) (_result *string, _err error) {
	stringToSign := tea.String("")
	canonicalizedResource, _err := client.BuildCanonicalizedResource(pathname, query)
	if _err != nil {
		return _result, _err
	}

	canonicalizedHeaders, _err := client.BuildCanonicalizedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	stringToSign = tea.String(tea.StringValue(method) + "\n" + tea.StringValue(canonicalizedHeaders) + tea.StringValue(canonicalizedResource))
	signature := signatureutil.HmacSHA1Sign(stringToSign, secret)
	_body := encodeutil.Base64EncodeToString(signature)
	_result = _body
	return _result, _err
}

func (client *Client) BuildCanonicalizedResource(pathname *string, query map[string]*string) (_result *string, _err error) {
	canonicalizedResource := pathname
	if !tea.BoolValue(util.IsUnset(query)) {
		queryArray := map_.KeySet(query)
		sortedQueryArray := array.AscSort(queryArray)
		separator := tea.String("?")
		for _, key := range sortedQueryArray {
			canonicalizedResource = tea.String(tea.StringValue(canonicalizedResource) + tea.StringValue(separator) + tea.StringValue(key))
			if !tea.BoolValue(util.Empty(query[tea.StringValue(key)])) {
				canonicalizedResource = tea.String(tea.StringValue(canonicalizedResource) + "=" + tea.StringValue(query[tea.StringValue(key)]))
			}

			separator = tea.String("&")
		}
	}

	_result = canonicalizedResource
	return _result, _err
}

func (client *Client) BuildCanonicalizedHeaders(headers map[string]*string) (_result *string, _err error) {
	accept := headers["accept"]
	if tea.BoolValue(util.IsUnset(accept)) {
		accept = tea.String("")
	}

	contentMd5 := headers["content-md5"]
	if tea.BoolValue(util.IsUnset(contentMd5)) {
		contentMd5 = tea.String("")
	}

	contentType := headers["content-type"]
	if tea.BoolValue(util.IsUnset(contentType)) {
		contentType = tea.String("")
	}

	date := headers["date"]
	if tea.BoolValue(util.IsUnset(date)) {
		date = tea.String("")
	}

	canonicalizedHeaders := tea.String(tea.StringValue(accept) + "\n" + tea.StringValue(contentMd5) + "\n" + tea.StringValue(contentType) + "\n" + tea.StringValue(date) + "\n")
	sortedHeaders, _err := client.GetSignedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	for _, header := range sortedHeaders {
		value := headers[tea.StringValue(header)]
		valueTrim := string_.Trim(value)
		canonicalizedHeaders = tea.String(tea.StringValue(canonicalizedHeaders) + tea.StringValue(header) + ":" + tea.StringValue(valueTrim) + "\n")
	}
	_result = canonicalizedHeaders
	return _result, _err
}

func (client *Client) GetSignedHeaders(headers map[string]*string) (_result []*string, _err error) {
	headersArray := map_.KeySet(headers)
	sortedHeadersArray := array.AscSort(headersArray)
	tmp := tea.String("")
	separator := tea.String("")
	for _, key := range sortedHeadersArray {
		lowerKey := string_.ToLower(key)
		if tea.BoolValue(string_.HasPrefix(lowerKey, tea.String("x-acs-"))) {
			if !tea.BoolValue(string_.Contains(tmp, lowerKey)) {
				tmp = tea.String(tea.StringValue(tmp) + tea.StringValue(separator) + tea.StringValue(lowerKey))
				separator = tea.String(";")
			}

		}

	}
	_result = make([]*string, 0)
	_body := string_.Split(tmp, tea.String(";"), tea.Int(0))
	_result = _body
	return _result, _err
}
