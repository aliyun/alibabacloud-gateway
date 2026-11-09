// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"strings"

	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	spi.Client
	EndpointSuffix      *string
	SignatureTypePrefix *string
	SignPrefix          *string
	Sha256              *string
	Sm3                 *string
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
	// CLOUD4-
	client.SignatureTypePrefix = dara.String("ACS4-")
	// cloud_v4
	client.SignPrefix = dara.String("aliyun_v4")
	client.EndpointSuffix = dara.String("aliyuncs.com")
	client.Sha256 = dara.String(dara.StringValue(client.SignatureTypePrefix) + "HMAC-SHA256")
	client.Sm3 = dara.String(dara.StringValue(client.SignatureTypePrefix) + "HMAC-SM3")
	return nil
}

func (client *Client) ModifyConfiguration(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	var attributes map[string]*string
	if attributeMap != nil {
		attributes = attributeMap.Key
	}
	suffix, _err := endpointSuffixForRequest(attributes)
	if _err != nil {
		return _err
	}
	config.Endpoint, _err = client.getEndpointWithSuffix(request.ProductId, config.RegionId, config.EndpointRule, config.Network, config.Suffix, config.EndpointMap, config.Endpoint, dara.String(suffix))
	if _err != nil {
		return _err
	}

	return _err
}

func (client *Client) ModifyRequest(context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
	request := context.Request
	config := context.Configuration
	date := getTimestamp()
	request.Headers = dara.Merge(map[string]*string{
		"host":                  config.Endpoint,
		"x-acs-version":         request.Version,
		"user-agent":            request.UserAgent,
		"x-acs-date":            date,
		"x-acs-signature-nonce": getNonce(),
		"accept":                dara.String("application/json"),
	}, request.Headers)
	if !dara.BoolValue(emptyString(request.Action)) {
		request.Headers["x-acs-action"] = request.Action
	}

	var requestAttributes map[string]*string
	if attributeMap != nil {
		requestAttributes = attributeMap.Key
	}
	names := signatureNames(requestAttributes)
	signatureAlgorithm := defaultString(request.SignatureAlgorithm, dara.String(names.sha256))
	hashedRequestPayload := hexEncode(hashBytes(toBytes(dara.String("")), signatureAlgorithm))
	if !dara.BoolValue(isUnset(request.Stream)) {
		tmp, _err := dara.ReadAsBytes(request.Stream)
		if _err != nil {
			return _err
		}

		hashedRequestPayload = hexEncode(hashBytes(tmp, signatureAlgorithm))
		request.Stream = dara.ToReader(tmp)
		request.Headers["content-type"] = dara.String("application/octet-stream")
	} else {
		if !dara.BoolValue(isUnset(request.Body)) {
			if dara.BoolValue(equalString(request.ReqBodyType, dara.String("byte"))) {
				byteObj, _err := assertAsBytes(request.Body)
				if _err != nil {
					return _err
				}

				hashedRequestPayload = hexEncode(hashBytes(byteObj, signatureAlgorithm))
				request.Stream = dara.ToReader(byteObj)
			} else if dara.BoolValue(equalString(request.ReqBodyType, dara.String("json"))) {
				jsonObj := toJSONString(request.Body)
				hashedRequestPayload = hexEncode(hashBytes(toBytes(jsonObj), signatureAlgorithm))
				request.Stream = dara.ToReader(jsonObj)
				request.Headers["content-type"] = dara.String("application/json; charset=utf-8")
			} else {
				m, _err := assertAsMap(request.Body)
				if _err != nil {
					return _err
				}

				formObj := toForm(m)
				hashedRequestPayload = hexEncode(hashBytes(toBytes(formObj), signatureAlgorithm))
				request.Stream = dara.ToReader(formObj)
				request.Headers["content-type"] = dara.String("application/x-www-form-urlencoded")
			}

		}

	}

	if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sm3))) {
		request.Headers["x-acs-content-sm3"] = hashedRequestPayload
	} else {
		request.Headers["x-acs-content-sha256"] = hashedRequestPayload
	}

	if !dara.BoolValue(equalString(request.AuthType, dara.String("Anonymous"))) {
		credential := request.Credential
		if dara.BoolValue(isUnset(credential)) {
			_err = tea.NewSDKError(map[string]interface{}{
				"code":    "ParameterMissing",
				"message": "'config.credential' can not be unset",
			})
			return _err
		}

		credentialModel, _err := credential.GetCredential()
		if _err != nil {
			return _err
		}

		if !dara.BoolValue(emptyString(credentialModel.ProviderName)) {
			request.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName
		}

		authType := credentialModel.Type
		if dara.BoolValue(equalString(authType, dara.String("bearer"))) {
			bearerToken := credentialModel.BearerToken
			request.Headers["x-acs-bearer-token"] = bearerToken
			request.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
			request.Headers["Authorization"] = dara.String("Bearer " + dara.StringValue(bearerToken))
		} else if dara.BoolValue(equalString(authType, dara.String("id_token"))) {
			idToken := credentialModel.SecurityToken
			request.Headers["x-acs-zero-trust-idtoken"] = idToken
		} else {
			accessKeyId := credentialModel.AccessKeyId
			accessKeySecret := credentialModel.AccessKeySecret
			securityToken := credentialModel.SecurityToken
			if !dara.BoolValue(emptyString(securityToken)) {
				request.Headers["x-acs-accesskey-id"] = accessKeyId
				request.Headers["x-acs-security-token"] = securityToken
			}

			dateNew := subString(date, dara.Int(0), dara.Int(10))
			dateNew = replaceString(dateNew, dara.String("-"), dara.String(""), nil)
			// Derive the signing region from the selected endpoint, without rereading the environment.
			region := client.regionWithSuffix(request.ProductId, config.Endpoint, config.RegionId, client.EndpointSuffix)
			signingkey := client.signingKey(signatureAlgorithm, accessKeySecret, request.ProductId, region, dateNew, names)
			request.Headers["Authorization"], _err = client.authorization(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew, names)
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
	if dara.BoolValue(is4xx(response.StatusCode)) || dara.BoolValue(is5xx(response.StatusCode)) {
		err := map[string]interface{}{}
		if !dara.BoolValue(isUnset(response.Headers["content-type"])) && dara.BoolValue(containsString(response.Headers["content-type"], dara.String("text/xml"))) {
			_str, _err := readAsString(response.Body)
			if _err != nil {
				return _err
			}

			respMap := dara.ParseXml(dara.StringValue(_str), nil)
			err, _err = assertAsMap(respMap["Error"])
			if _err != nil {
				return _err
			}

		} else {
			_res, _err := dara.ReadAsJSON(response.Body)
			if _err != nil {
				return _err
			}

			err, _err = assertAsMap(_res)
			if _err != nil {
				return _err
			}

		}

		requestId := client.DefaultAny(err["RequestId"], err["requestId"])
		if !dara.BoolValue(isUnset(response.Headers["x-acs-request-id"])) {
			requestId = response.Headers["x-acs-request-id"]
		}

		err["statusCode"] = response.StatusCode
		_err = tea.NewSDKError(map[string]interface{}{
			"code":               dara.ToString(client.DefaultAny(err["Code"], err["code"])),
			"message":            "code: " + dara.ToString(dara.IntValue(response.StatusCode)) + ", " + dara.ToString(client.DefaultAny(err["Message"], err["message"])) + " request id: " + dara.ToString(requestId),
			"data":               err,
			"description":        dara.ToString(client.DefaultAny(err["Description"], err["description"])),
			"accessDeniedDetail": client.DefaultAny(err["AccessDeniedDetail"], err["accessDeniedDetail"]),
		})
		return _err
	}

	if dara.BoolValue(equalNumber(response.StatusCode, dara.Int(204))) {
		_, _err = readAsString(response.Body)
		if _err != nil {
			return _err
		}
	} else if dara.BoolValue(equalString(request.BodyType, dara.String("binary"))) {
		response.DeserializedBody = response.Body
	} else if dara.BoolValue(equalString(request.BodyType, dara.String("byte"))) {
		byt, _err := dara.ReadAsBytes(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = byt
	} else if dara.BoolValue(equalString(request.BodyType, dara.String("string"))) {
		str, _err := readAsString(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = str
	} else if dara.BoolValue(equalString(request.BodyType, dara.String("json"))) {
		obj, _err := dara.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		res, _err := assertAsMap(obj)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = res
	} else if dara.BoolValue(equalString(request.BodyType, dara.String("array"))) {
		arr, _err := dara.ReadAsJSON(response.Body)
		if _err != nil {
			return _err
		}

		response.DeserializedBody = arr
	} else {
		response.DeserializedBody, _err = readAsString(response.Body)
		if _err != nil {
			return _err
		}

	}

	return _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	return client.getEndpointWithSuffix(productId, regionId, endpointRule, network, suffix, endpointMap, endpoint, client.EndpointSuffix)
}

func (client *Client) getEndpointWithSuffix(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string, endpointSuffix *string) (_result *string, _err error) {
	if !dara.BoolValue(emptyString(endpoint)) {
		return applyEndpointSuffix(endpoint, endpointSuffix)
	}

	if !dara.BoolValue(isUnset(endpointMap)) && !dara.BoolValue(emptyString(endpointMap[dara.StringValue(regionId)])) {
		return applyEndpointSuffix(endpointMap[dara.StringValue(regionId)], endpointSuffix)
	}

	_body, _err := getEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	return applyEndpointSuffix(_body, endpointSuffix)
}

func (client *Client) DefaultAny(inputValue interface{}, defaultValue interface{}) (_result interface{}) {
	if dara.BoolValue(isUnset(inputValue)) {
		_result = defaultValue
		return _result
	}

	_result = inputValue
	return _result
}

func (client *Client) GetAuthorization(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, ak *string, signingkey []byte, product *string, region *string, date *string) (_result *string, _err error) {
	return client.authorization(pathname, method, query, headers, signatureAlgorithm, payload, ak, signingkey, product, region, date, requestSignature{
		signPrefix: dara.StringValue(client.SignPrefix),
		sha256:     dara.StringValue(client.Sha256),
		sm3:        dara.StringValue(client.Sm3),
	})
}

func (client *Client) authorization(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, ak *string, signingkey []byte, product *string, region *string, date *string, names requestSignature) (_result *string, _err error) {
	signature, _err := client.signature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey, names)
	if _err != nil {
		return _result, _err
	}
	signedHeaders := client.GetSignedHeaders(headers)
	signedHeadersStr := join(signedHeaders, dara.String(";"))
	_result = dara.String(dara.StringValue(signatureAlgorithm) + " Credential=" + dara.StringValue(ak) + "/" + dara.StringValue(date) + "/" + dara.StringValue(region) + "/" + dara.StringValue(product) + "/" + names.signPrefix + "_request,SignedHeaders=" + dara.StringValue(signedHeadersStr) + ",Signature=" + dara.StringValue(signature))
	return _result, _err
}

func (client *Client) GetSignature(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, signingkey []byte) (_result *string, _err error) {
	return client.signature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey, requestSignature{
		signPrefix: dara.StringValue(client.SignPrefix),
		sha256:     dara.StringValue(client.Sha256),
		sm3:        dara.StringValue(client.Sm3),
	})
}

func (client *Client) signature(pathname *string, method *string, query map[string]*string, headers map[string]*string, signatureAlgorithm *string, payload *string, signingkey []byte, names requestSignature) (_result *string, _err error) {
	canonicalURI := dara.String("/")
	if !dara.BoolValue(emptyString(pathname)) {
		canonicalURI = pathname
	}

	stringToSign := dara.String("")
	canonicalizedResource := client.BuildCanonicalizedResource(query)
	canonicalizedHeaders := client.BuildCanonicalizedHeaders(headers)
	signedHeaders := client.GetSignedHeaders(headers)
	_err = client.ValidateSignedHeaders(signedHeaders)
	if _err != nil {
		return _result, _err
	}
	signedHeadersStr := join(signedHeaders, dara.String(";"))
	stringToSign = dara.String(strings.ToUpper(dara.StringValue(method)) + "\n" + dara.StringValue(canonicalURI) + "\n" + dara.StringValue(canonicalizedResource) + "\n" + dara.StringValue(canonicalizedHeaders) + "\n" + dara.StringValue(signedHeadersStr) + "\n" + dara.StringValue(payload))
	hex := hexEncode(hashBytes(toBytes(stringToSign), signatureAlgorithm))
	stringToSign = dara.String(dara.StringValue(signatureAlgorithm) + "\n" + dara.StringValue(hex))
	signature := toBytes(dara.String(""))
	if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sha256))) {
		signature = hmacSHA256ByBytes(stringToSign, signingkey)
	} else if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sm3))) {
		signature = hmacSM3ByBytes(stringToSign, signingkey)
	}

	_body := hexEncode(signature)
	_result = _body
	return _result, _err
}

func (client *Client) ValidateSignedHeaders(signedHeaders []*string) (_err error) {
	if !dara.BoolValue(listContains(signedHeaders, dara.String("host"))) || !dara.BoolValue(listContains(signedHeaders, dara.String("x-acs-date"))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "InvalidSignedHeaders",
			"message": "signed headers must include host and x-acs-date",
		})
		return _err
	}
	return _err
}

func (client *Client) GetSigningkey(signatureAlgorithm *string, secret *string, product *string, region *string, date *string) (_result []byte) {
	return client.signingKey(signatureAlgorithm, secret, product, region, date, requestSignature{
		signPrefix: dara.StringValue(client.SignPrefix),
		sha256:     dara.StringValue(client.Sha256),
		sm3:        dara.StringValue(client.Sm3),
	})
}

func (client *Client) signingKey(signatureAlgorithm *string, secret *string, product *string, region *string, date *string, names requestSignature) (_result []byte) {
	sc1 := dara.String(names.signPrefix + dara.StringValue(secret))
	sc2 := toBytes(dara.String(""))
	if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sha256))) {
		sc2 = hmacSHA256(date, sc1)
	} else if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sm3))) {
		sc2 = hmacSM3(date, sc1)
	}

	sc3 := toBytes(dara.String(""))
	if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sha256))) {
		sc3 = hmacSHA256ByBytes(region, sc2)
	} else if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sm3))) {
		sc3 = hmacSM3ByBytes(region, sc2)
	}

	sc4 := toBytes(dara.String(""))
	if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sha256))) {
		sc4 = hmacSHA256ByBytes(product, sc3)
	} else if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sm3))) {
		sc4 = hmacSM3ByBytes(product, sc3)
	}

	hmac := toBytes(dara.String(""))
	if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sha256))) {
		hmac = hmacSHA256ByBytes(dara.String(names.signPrefix+"_request"), sc4)
	} else if dara.BoolValue(equalString(signatureAlgorithm, dara.String(names.sm3))) {
		hmac = hmacSM3ByBytes(dara.String(names.signPrefix+"_request"), sc4)
	}

	_result = hmac
	return _result
}

func (client *Client) GetRegion(product *string, endpoint *string, regionId *string) (_result *string) {
	return client.regionWithSuffix(product, endpoint, regionId, client.EndpointSuffix)
}

func (client *Client) regionWithSuffix(product *string, endpoint *string, regionId *string, endpointSuffix *string) (_result *string) {
	if !dara.BoolValue(emptyString(regionId)) {
		_result = regionId
		return _result
	}

	region := dara.String("center")
	if dara.BoolValue(emptyString(product)) || dara.BoolValue(emptyString(endpoint)) {
		_result = region
		return _result
	}

	strs := splitString(endpoint, dara.String(":"), nil)
	withoutPort := strs[0]
	preRegion := dara.String(hostWithoutDomain(dara.StringValue(withoutPort), dara.StringValue(endpointSuffix)))
	nodes := splitString(preRegion, dara.String("."), nil)
	if dara.BoolValue(equalNumber(sizeOf(nodes), dara.Int(2))) {
		region = nodes[1]
	}

	_result = region
	return _result
}

func (client *Client) BuildCanonicalizedResource(query map[string]*string) (_result *string) {
	encoded := make(map[string]*string, len(query))
	for key, value := range query {
		if key != "Signature" {
			encoded[dara.StringValue(percentEncode(dara.String(key)))] = percentEncode(value)
		}
	}
	pairs := make([]string, 0, len(encoded))
	for _, key := range ascSort(keySet(encoded)) {
		pairs = append(pairs, dara.StringValue(key)+"="+dara.StringValue(encoded[dara.StringValue(key)]))
	}
	return dara.String(strings.Join(pairs, "&"))
}

func (client *Client) BuildCanonicalizedHeaders(headers map[string]*string) (_result *string) {
	// net/http writes distinct header keys in lexical order. Preserve that order
	// when merging keys that differ only by case.
	headersArray := ascSort(keySet(headers))
	newHeaders := make(map[string]*string)
	for _, key := range headersArray {
		lowerKey := toLower(key)
		value := headers[dara.StringValue(key)]
		if !dara.BoolValue(isUnset(value)) {
			if dara.StringValue(lowerKey) == "host" && newHeaders["host"] != nil && dara.StringValue(key) != "host" {
				continue
			}
			if dara.BoolValue(isUnset(newHeaders[dara.StringValue(lowerKey)])) || dara.BoolValue(equalsString(lowerKey, dara.String("host"))) {
				newHeaders[dara.StringValue(lowerKey)] = trimString(value)
			} else {
				newHeaders[dara.StringValue(lowerKey)] = dara.String(dara.StringValue(newHeaders[dara.StringValue(lowerKey)]) + "," + dara.StringValue(trimString(value)))
			}

		}

	}
	canonicalizedHeaders := dara.String("")
	sortedHeaders := client.GetSignedHeaders(headers)
	for _, header := range sortedHeaders {
		canonicalizedHeaders = dara.String(dara.StringValue(canonicalizedHeaders) + dara.StringValue(header) + ":" + dara.StringValue(newHeaders[dara.StringValue(header)]) + "\n")
	}
	_result = canonicalizedHeaders
	return _result
}

func (client *Client) GetSignedHeaders(headers map[string]*string) (_result []*string) {
	headersArray := keySet(headers)
	newHeadersArray := []*string{}
	for _, key := range headersArray {
		lowerKey := toLower(key)
		value := headers[dara.StringValue(key)]
		if !dara.BoolValue(isUnset(value)) {
			newHeadersArray = append(newHeadersArray, lowerKey)
		}

	}
	sortedHeadersArray := ascSort(newHeadersArray)
	result := []*string{}
	for _, key := range sortedHeadersArray {
		if dara.BoolValue(hasPrefix(key, dara.String("x-acs-"))) || dara.BoolValue(equalsString(key, dara.String("host"))) || dara.BoolValue(equalsString(key, dara.String("content-type"))) {
			if !dara.BoolValue(listContains(result, key)) {
				result = append(result, key)
			}

		}

	}
	_result = result
	return _result
}
