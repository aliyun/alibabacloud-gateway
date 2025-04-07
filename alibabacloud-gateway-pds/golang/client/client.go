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
	date := util.GetDateUTCString()
	request.Headers = tea.Merge(map[string]*string{
		"date":                  date,
		"host":                  config.Endpoint,
		"x-acs-version":         request.Version,
		"x-acs-action":          request.Action,
		"user-agent":            request.UserAgent,
		"x-acs-signature-nonce": util.GetNonce(),
		"accept":                tea.String("application/json"),
	}, request.Headers)
	signatureAlgorithm := util.DefaultString(request.SignatureAlgorithm, tea.String("ACS4-HMAC-SHA256"))
	signatureVersion := util.DefaultString(request.SignatureVersion, tea.String("v1"))
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
				m, _err := util.AssertAsMap(request.Body)
				if _err != nil {
					return _err
				}

				formObj := openapiutil.ToForm(m)
				hashedRequestPayload = encodeutil.HexEncode(encodeutil.Hash(util.ToBytes(formObj), signatureAlgorithm))
				request.Stream = tea.ToReader(formObj)
				request.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			}

		}

	}

	if tea.BoolValue(string_.Equals(signatureVersion, tea.String("v4"))) {
		if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SM3"))) {
			request.Headers["x-acs-content-sm3"] = hashedRequestPayload
		} else {
			request.Headers["x-acs-content-sha256"] = hashedRequestPayload
		}

	} else {
		request.Headers["x-acs-signature-method"] = tea.String("HMAC-SHA1")
		request.Headers["x-acs-signature-version"] = tea.String("1.0")
	}

	if !tea.BoolValue(util.EqualString(request.AuthType, tea.String("Anonymous"))) && !tea.BoolValue(util.IsUnset(request.Credential)) {
		credential := request.Credential
		credentialModel, _err := credential.GetCredential()
		if _err != nil {
			return _err
		}

		authType := credentialModel.Type
		if tea.BoolValue(util.EqualString(authType, tea.String("bearer"))) {
			bearerToken := credentialModel.BearerToken
			request.Headers["x-acs-bearer-token"] = bearerToken
			request.Headers["Authorization"] = tea.String("Bearer " + tea.StringValue(bearerToken))
		} else {
			accessKeyId := credentialModel.AccessKeyId
			accessKeySecret := credentialModel.AccessKeySecret
			securityToken := credentialModel.SecurityToken
			if !tea.BoolValue(util.Empty(securityToken)) {
				request.Headers["x-acs-security-token"] = securityToken
			}

			headers := make(map[string]*string)
			if !tea.BoolValue(util.IsUnset(request.Headers["content-type"])) {
				headers = request.Headers
			} else if tea.BoolValue(string_.Equals(request.ReqBodyType, tea.String("formData"))) && tea.BoolValue(string_.Equals(request.Action, tea.String("DownloadFile"))) && tea.BoolValue(string_.Equals(request.Pathname, tea.String("/v2/file/download"))) {
				headersArray := map_.KeySet(request.Headers)
				for _, key := range headersArray {
					headers[tea.StringValue(key)] = request.Headers[tea.StringValue(key)]
				}
				headers["content-type"] = tea.String("application/x-www-form-urlencoded; charset=UTF-8")
			} else {
				headers = request.Headers
			}

			if tea.BoolValue(string_.Equals(signatureVersion, tea.String("v4"))) {
				dateNew := string_.SubString(date, tea.Int(0), tea.Int(10))
				region := client.GetRegion(config.Endpoint)
				signingkey, _err := client.GetSigningkey(signatureAlgorithm, accessKeySecret, region, dateNew)
				if _err != nil {
					return _err
				}

				request.Headers["Authorization"], _err = client.GetAuthorizationV4(request.Pathname, request.Method, request.Query, headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew)
				if _err != nil {
					return _err
				}

			} else {
				request.Headers["Authorization"], _err = client.GetAuthorization(request.Pathname, request.Method, request.Query, headers, accessKeyId, accessKeySecret)
				if _err != nil {
					return _err
				}

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

		err, _err := util.AssertAsMap(_res)
		if _err != nil {
			return _err
		}

		headers := response.Headers
		requestId := headers["x-ca-request-id"]
		err["statusCode"] = response.StatusCode
		err["requestId"] = requestId
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    tea.ToString(client.DefaultAny(err["Code"], err["code"])),
			"message": tea.ToString(client.DefaultAny(err["Message"], err["message"])),
			"data":    err,
		})
		return _err
	}

	if !tea.BoolValue(util.IsUnset(response.Body)) {
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
			response.DeserializedBody, _err = util.ReadAsJSON(response.Body)
			if _err != nil {
				return _err
			}

		} else if tea.BoolValue(util.EqualString(request.BodyType, tea.String("array"))) {
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
	_body := string_.Split(tmp, tea.String(";"), nil)
	_result = _body
	return _result, _err
}

func (client *Client) GetRegion(endpoint *string) (_result *string) {
	region := tea.String("center")
	if tea.BoolValue(util.Empty(endpoint)) {
		_result = region
		return _result
	}

	if tea.BoolValue(string_.Contains(endpoint, tea.String(".admin.aliyunpds.com"))) {
		region = string_.Replace(endpoint, tea.String(".admin.aliyunpds.com"), tea.String(""), nil)
	}

	_result = region
	return _result
}

func (client *Client) GetSigningkey(signatureAlgorithm *string, secret *string, region *string, date *string) (_result []byte, _err error) {
	sc1 := tea.String("aliyun_v4" + tea.StringValue(secret))
	sc2 := util.ToBytes(tea.String(""))
	if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SHA256"))) {
		sc2 = signatureutil.HmacSHA256Sign(date, sc1)
	} else if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SM3"))) {
		sc2 = signatureutil.HmacSM3Sign(date, sc1)
	}

	sc3 := util.ToBytes(tea.String(""))
	if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SHA256"))) {
		sc3 = signatureutil.HmacSHA256SignByBytes(region, sc2)
	} else if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SM3"))) {
		sc3 = signatureutil.HmacSM3SignByBytes(region, sc2)
	}

	sc4 := util.ToBytes(tea.String(""))
	if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SHA256"))) {
		sc4 = signatureutil.HmacSHA256SignByBytes(tea.String("pds"), sc3)
	} else if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SM3"))) {
		sc4 = signatureutil.HmacSM3SignByBytes(tea.String("pds"), sc3)
	}

	hmac := util.ToBytes(tea.String(""))
	if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SHA256"))) {
		hmac = signatureutil.HmacSHA256SignByBytes(tea.String("aliyun_v4_request"), sc4)
	} else if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SM3"))) {
		hmac = signatureutil.HmacSM3SignByBytes(tea.String("aliyun_v4_request"), sc4)
	}

	_result = hmac
	return _result, _err
}

func (client *Client) GetAuthorizationV4(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, ak *string, signingkey []byte, product *string, region *string, date *string) (_result *string, _err error) {
	signature, _err := client.GetSignatureV4(pathname, method, query, headers, signatureAlgorithm, payload, signingkey)
	if _err != nil {
		return _result, _err
	}

	signedHeaders, _err := client.GetSignedHeaders(headers)
	if _err != nil {
		return _result, _err
	}

	signedHeadersStr := array.Join(signedHeaders, tea.String(";"))
	_result = tea.String(tea.StringValue(signatureAlgorithm) + " Credential=" + tea.StringValue(ak) + "/" + tea.StringValue(date) + "/" + tea.StringValue(region) + "/" + tea.StringValue(product) + "/aliyun_v4_request,SignedHeaders=" + tea.StringValue(signedHeadersStr) + ",Signature=" + tea.StringValue(signature))
	return _result, _err
}

func (client *Client) GetSignatureV4(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, signingkey []byte) (_result *string, _err error) {
	stringToSign := tea.String("")
	canonicalizedResource, _err := client.BuildCanonicalizedResource(pathname, query)
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
	stringToSign = tea.String(tea.StringValue(method) + "\n" + tea.StringValue(canonicalizedResource) + "\n" + tea.StringValue(canonicalizedHeaders) + "\n" + tea.StringValue(signedHeadersStr) + "\n" + tea.StringValue(payload))
	hex := encodeutil.HexEncode(encodeutil.Hash(util.ToBytes(stringToSign), signatureAlgorithm))
	stringToSign = tea.String(tea.StringValue(signatureAlgorithm) + "\n" + tea.StringValue(hex))
	signature := util.ToBytes(tea.String(""))
	if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SHA256"))) {
		signature = signatureutil.HmacSHA256SignByBytes(stringToSign, signingkey)
	} else if tea.BoolValue(util.EqualString(signatureAlgorithm, tea.String("ACS4-HMAC-SM3"))) {
		signature = signatureutil.HmacSM3SignByBytes(stringToSign, signingkey)
	}

	_body := encodeutil.HexEncode(signature)
	_result = _body
	return _result, _err
}
