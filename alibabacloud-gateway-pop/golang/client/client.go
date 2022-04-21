// This file is auto-generated, don't edit it. Thanks.
package client

import (
	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	array "github.com/alibabacloud-go/darabonba-array/client"
	encodeutil "github.com/alibabacloud-go/darabonba-encode-util/client"
	map_ "github.com/alibabacloud-go/darabonba-map/client"
	signatureutil "github.com/alibabacloud-go/darabonba-signature-util/client"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
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
	request.Headers = tea.Merge(map[string]*string{
		"host":                  config.Endpoint,
		"x-acs-version":         request.Version,
		"x-acs-action":          request.Action,
		"user-agent":            request.UserAgent,
		"x-acs-date":            openapiutil.GetTimestamp(),
		"x-acs-signature-nonce": util.GetNonce(),
		"accept":                tea.String("application/json"),
	}, request.Headers)
	signatureAlgorithm := tea.String("ACS3-HMAC-SHA256")
	if !tea.BoolValue(util.IsUnset(request.SignatureAlgorithm)) {
		signatureAlgorithm = request.SignatureAlgorithm
	}

	hashedRequestPayload := encodeutil.HexEncode(encodeutil.Hash(util.ToBytes(tea.String("")), signatureAlgorithm))
	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		tmp, _err := util.ReadAsBytes(request.Stream)
		if _err != nil {
			return _err
		}

		hashedRequestPayload = encodeutil.HexEncode(encodeutil.Hash(tmp, signatureAlgorithm))
		request.Stream = tea.ToReader(tmp)
		request.Headers["content-type"] = tea.String("application/octet-stream")
	} else {
		if !tea.BoolValue(util.IsUnset(request.Body)) {
			if tea.BoolValue(util.EqualString(request.ReqBodyType, tea.String("json"))) {
				jsonObj := util.ToJSONString(request.Body)
				hashedRequestPayload = encodeutil.HexEncode(encodeutil.Hash(util.ToBytes(jsonObj), signatureAlgorithm))
				request.Stream = tea.ToReader(jsonObj)
				request.Headers["content-type"] = tea.String("application/json; charset=utf-8")
			} else {
				m := util.AssertAsMap(request.Body)
				formObj := openapiutil.ToForm(m)
				hashedRequestPayload = encodeutil.HexEncode(encodeutil.Hash(util.ToBytes(formObj), signatureAlgorithm))
				request.Stream = tea.ToReader(formObj)
				request.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			}

		}

	}

	request.Headers["x-acs-content-sha256"] = hashedRequestPayload
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

		if !tea.BoolValue(util.Empty(securityToken)) {
			request.Headers["x-acs-accesskey-id"] = accessKeyId
			request.Headers["x-acs-security-token"] = securityToken
		}

		request.Headers["Authorization"], _err = client.GetAuthorization(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret)
		if _err != nil {
			return _err
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
	} else {
		response.DeserializedBody, _err = util.ReadAsString(response.Body)
		if _err != nil {
			return _err
		}

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

func (client *Client) GetAuthorization(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, ak *string, secret *string) (_result *string, _err error) {
	signature, _err := client.GetSignature(pathname, method, query, headers, signatureAlgorithm, payload, secret)
	if _err != nil {
		return _result, _err
	}

	signedHeaders, _err := client.GetSignedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	signedHeadersStr := array.Join(signedHeaders, tea.String(";"))
	_result = tea.String(tea.StringValue(signatureAlgorithm) + "  Credential=" + tea.StringValue(ak) + ",SignedHeaders=" + tea.StringValue(signedHeadersStr) + ",Signature=" + tea.StringValue(signature))
	return _result, _err
}

func (client *Client) GetSignature(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, secret *string) (_result *string, _err error) {
	canonicalURI := tea.String("/")
	if !tea.BoolValue(util.Empty(pathname)) {
		canonicalURI = pathname
	}

	stringToSign := tea.String("")
	canonicalizedResource, _err := client.BuildCanonicalizedResource(query)
	if _err != nil {
		return _result, _err
	}

	canonicalizedHeaders, _err := client.BuildCanonicalizedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	signedHeaders, _err := client.GetSignedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	signedHeadersStr := array.Join(signedHeaders, tea.String(";"))
	stringToSign = tea.String(tea.StringValue(method) + "\n" + tea.StringValue(canonicalURI) + "\n" + tea.StringValue(canonicalizedResource) + "\n" + tea.StringValue(canonicalizedHeaders) + "\n" + tea.StringValue(signedHeadersStr) + "\n" + tea.StringValue(payload))
	hex := encodeutil.HexEncode(encodeutil.Hash(util.ToBytes(stringToSign), signatureAlgorithm))
	stringToSign = tea.String(tea.StringValue(signatureAlgorithm) + "\n" + tea.StringValue(hex))
	signature := util.ToBytes(tea.String(""))
	if tea.BoolValue(string_.Equals(signatureAlgorithm, tea.String("ACS3-HMAC-SHA256"))) {
		signature = signatureutil.HmacSHA256Sign(stringToSign, secret)
	} else if tea.BoolValue(string_.Equals(signatureAlgorithm, tea.String("ACS3-HMAC-SM3"))) {
		signature = signatureutil.HmacSM3Sign(stringToSign, secret)
	} else if tea.BoolValue(string_.Equals(signatureAlgorithm, tea.String("ACS3-RSA-SHA256"))) {
		signature = signatureutil.SHA256withRSASign(stringToSign, secret)
	}

	_body := encodeutil.HexEncode(signature)
	_result = _body
	return _result, _err
}

func (client *Client) BuildCanonicalizedResource(query map[string]*string) (_result *string, _err error) {
	canonicalizedResource := tea.String("")
	if !tea.BoolValue(util.IsUnset(query)) {
		queryArray := map_.KeySet(query)
		sortedQueryArray := array.AscSort(queryArray)
		separator := tea.String("")
		for _, key := range sortedQueryArray {
			canonicalizedResource = tea.String(tea.StringValue(canonicalizedResource) + tea.StringValue(separator) + tea.StringValue(encodeutil.PercentEncode(key)))
			if !tea.BoolValue(util.Empty(query[tea.StringValue(key)])) {
				canonicalizedResource = tea.String(tea.StringValue(canonicalizedResource) + "=" + tea.StringValue(encodeutil.PercentEncode(query[tea.StringValue(key)])))
			}

			separator = tea.String("&")
		}
	}

	_result = canonicalizedResource
	return _result, _err
}

func (client *Client) BuildCanonicalizedHeaders(headers map[string]*string) (_result *string, _err error) {
	canonicalizedHeaders := tea.String("")
	sortedHeaders, _err := client.GetSignedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	for _, header := range sortedHeaders {
		canonicalizedHeaders = tea.String(tea.StringValue(canonicalizedHeaders) + tea.StringValue(header) + ":" + tea.StringValue(string_.Trim(headers[tea.StringValue(header)])) + "\n")
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
		if tea.BoolValue(string_.HasPrefix(lowerKey, tea.String("x-acs-"))) || tea.BoolValue(string_.Equals(lowerKey, tea.String("host"))) || tea.BoolValue(string_.Equals(lowerKey, tea.String("content-type"))) {
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
