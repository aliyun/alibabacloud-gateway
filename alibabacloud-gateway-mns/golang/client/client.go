// This file is auto-generated, don't edit it. Thanks.
package client

import (
  spi  "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  openapiutil  "github.com/alibabacloud-go/openapi-util/service"
  encodeutil  "github.com/alibabacloud-go/darabonba-encode-util/client"
  signatureutil  "github.com/alibabacloud-go/darabonba-signature-util/client"
  string_  "github.com/alibabacloud-go/darabonba-string/client"
  map_  "github.com/alibabacloud-go/darabonba-map/client"
  array  "github.com/alibabacloud-go/darabonba-array/client"
  xml  "github.com/alibabacloud-go/tea-xml/service"
  "github.com/alibabacloud-go/tea/tea"
)

type Client struct {
  spi.Client
  SignPrefix  *string
  SignSuffix  *string
  AuthPrefix  *string
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  _err = client.Client.Init(  )
  if _err != nil {
    return _err
  }
  client.SignPrefix = tea.String("aliyun_v4")
  client.SignSuffix = tea.String("aliyun_v4_request")
  client.AuthPrefix = tea.String("MNS4-HMAC-SHA256")
  return nil
}



func (client *Client) ModifyConfiguration (context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
  config := context.Configuration
  config.Endpoint, _err = client.GetEndpoint(config.RegionId, config.Endpoint)
  if _err != nil {
    return _err
  }

  return _err
}

func (client *Client) ModifyRequest (context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
  request := context.Request
  config := context.Configuration
  signatureVersion := util.DefaultString(request.SignatureVersion, tea.String("v2"))
  if !tea.BoolValue(util.IsUnset(request.Body)) {
    if tea.BoolValue(string_.Equals(request.ReqBodyType, tea.String("xml"))) {
      reqBodyMap, _err := util.AssertAsMap(request.Body)
      if _err != nil {
        return _err
      }

      xmlStr := xml.ToXML(reqBodyMap)
      request.Stream = tea.ToReader(xmlStr)
      request.Headers["content-type"] = tea.String("text/xml;charset=UTF-8")
      request.Headers["content-md5"] = encodeutil.Base64EncodeToString(signatureutil.MD5Sign(xmlStr))
    } else if tea.BoolValue(string_.Equals(request.ReqBodyType, tea.String("json"))) || tea.BoolValue(string_.Equals(request.ReqBodyType, tea.String("formData"))) {
      bodyStr := util.ToJSONString(request.Body)
      request.Stream = tea.ToReader(bodyStr)
      request.Headers["content-type"] = tea.String("application/json")
      request.Headers["content-md5"] = encodeutil.Base64EncodeToString(signatureutil.MD5Sign(bodyStr))
    } else if tea.BoolValue(string_.Equals(request.ReqBodyType, tea.String("byte"))) || tea.BoolValue(string_.Equals(request.ReqBodyType, tea.String("binary"))) {
      bodyBytes, _err := util.AssertAsBytes(request.Body)
      if _err != nil {
        return _err
      }

      request.Stream = tea.ToReader(bodyBytes)
      request.Headers["content-md5"] = encodeutil.Base64EncodeToString(signatureutil.MD5SignForBytes(bodyBytes))
    }

  }

  _err = client.BuildRequest(context)
  if _err != nil {
    return _err
  }
  request.Headers = tea.Merge(map[string]*string{
    "host": config.Endpoint,
    "x-mns-version": request.Version,
    "user-agent": request.UserAgent,
    "accept": tea.String("application/json"),
    },request.Headers)
  if !tea.BoolValue(util.EqualString(request.AuthType, tea.String("Anonymous"))) {
    credential := request.Credential
    if tea.BoolValue(util.IsUnset(credential)) {
      _err = tea.NewSDKError(map[string]interface{}{
        "code": "ParameterMissing",
        "message": "'config.credential' can not be unset",
      })
      return _err
    }

    credentialModel, _err := credential.GetCredential()
    if _err != nil {
      return _err
    }

    authType := credentialModel.Type
    if tea.BoolValue(util.EqualString(authType, tea.String("bearer"))) {
      bearerToken := credentialModel.BearerToken
      request.Headers["x-acs-bearer-token"] = bearerToken
      request.Headers["x-acs-signature-type"] = tea.String("BEARERTOKEN")
      request.Headers["Authorization"] = tea.String("Bearer " + tea.StringValue(bearerToken))
    } else {
      accessKeyId := credentialModel.AccessKeyId
      accessKeySecret := credentialModel.AccessKeySecret
      securityToken := credentialModel.SecurityToken
      if !tea.BoolValue(util.Empty(securityToken)) {
        request.Headers["security-token"] = securityToken
      }

      request.Headers["date"] = util.GetDateUTCString()
      if tea.BoolValue(string_.Equals(signatureVersion, tea.String("v4"))) {
        date, _err := client.GetDateISO8601()
        if _err != nil {
          return _err
        }

        request.Headers["authorization"], _err = client.GetAuthorizationV4(context, date, accessKeyId, accessKeySecret)
        if _err != nil {
          return _err
        }

      } else {
        request.Headers["authorization"], _err = client.GetAuthorizationV2(request.Pathname, request.Method, request.Headers, accessKeyId, accessKeySecret)
        if _err != nil {
          return _err
        }

      }

    }

  }

  return _err
}

func (client *Client) ModifyResponse (context *spi.InterceptorContext, attributeMap *spi.AttributeMap) (_err error) {
  request := context.Request
  response := context.Response
  if tea.BoolValue(util.Is4xx(response.StatusCode)) || tea.BoolValue(util.Is5xx(response.StatusCode)) {
    err := map[string]interface{}{}
    if !tea.BoolValue(util.IsUnset(response.Headers["content-type"])) && tea.BoolValue(string_.Contains(response.Headers["content-type"], tea.String("text/xml"))) {
      _str, _err := util.ReadAsString(response.Body)
      if _err != nil {
        return _err
      }

      respMap := xml.ParseXml(_str, nil)
      err, _err = util.AssertAsMap(respMap["Error"])
      if _err != nil {
        return _err
      }

    } else {
      _res, _err := util.ReadAsJSON(response.Body)
      if _err != nil {
        return _err
      }

      err, _err = util.AssertAsMap(_res)
      if _err != nil {
        return _err
      }

    }

    requestId := client.DefaultAny(err["RequestId"], err["requestId"])
    if !tea.BoolValue(util.IsUnset(response.Headers["x-mns-request-id"])) {
      requestId = response.Headers["x-mns-request-id"]
    }

    err["statusCode"] = response.StatusCode
    _err = tea.NewSDKError(map[string]interface{}{
      "code": tea.ToString(client.DefaultAny(err["Code"], err["code"])),
      "message": "code: " + tea.ToString(tea.IntValue(response.StatusCode)) + ", " + tea.ToString(client.DefaultAny(err["Message"], err["message"])) + " request id: " + tea.ToString(requestId),
      "data": err,
      "description": tea.ToString(client.DefaultAny(err["Description"], err["description"])),
      "accessDeniedDetail": client.DefaultAny(err["AccessDeniedDetail"], err["accessDeniedDetail"]),
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

func (client *Client) GetEndpoint (regionId *string, endpoint *string) (_result *string, _err error) {
  if !tea.BoolValue(util.Empty(endpoint)) {
    _result = endpoint
    return _result , _err
  }

  if tea.BoolValue(util.Empty(regionId)) {
    regionId = tea.String("cn-hangzhou")
  }

  _result = tea.String(tea.StringValue(regionId) + ".mns.aliyuncs.com")
  return _result, _err
}

func (client *Client) DefaultAny (inputValue interface{}, defaultValue interface{}) (_result interface{}) {
  if tea.BoolValue(util.IsUnset(inputValue)) {
    _result = defaultValue
    return _result
  }

  _result = inputValue
  return _result
}

func (client *Client) DefaultString (inputValue *string, defaultValue *string) (_result *string) {
  if tea.BoolValue(util.IsUnset(inputValue)) || tea.BoolValue(util.Empty(inputValue)) {
    _result = defaultValue
    return _result
  }

  _result = inputValue
  return _result
}

func Base64Encode (input *string) (_result *string) {
  if tea.BoolValue(util.IsUnset(input)) {
    _result = tea.String("")
    return _result
  }

  _body := encodeutil.Base64EncodeToString(util.ToBytes(input))
  _result = _body
  return _result
}

func Base64Decode (input *string) (_result *string) {
  if tea.BoolValue(util.IsUnset(input)) {
    _result = tea.String("")
    return _result
  }

  _body := util.ToString(encodeutil.Base64Decode(input))
  _result = _body
  return _result
}

func (client *Client) BuildRequest (context *spi.InterceptorContext) (_err error) {
  request := context.Request
  resource := request.Pathname
  if !tea.BoolValue(util.Empty(resource)) {
    paths := string_.Split(resource, tea.String("?"), tea.Int(2))
    resource = paths[0]
    if tea.BoolValue(util.EqualNumber(array.Size(paths), tea.Int(2))) {
      params := string_.Split(paths[1], tea.String("&"), nil)
      for _, sub := range params {
        item := string_.Split(sub, tea.String("="), nil)
        key := item[0]
        var value *string
        if tea.BoolValue(util.EqualNumber(array.Size(item), tea.Int(2))) {
          value = item[1]
        }

        request.Query[tea.StringValue(key)] = value
      }
    }

  }

  request.Pathname = resource
  return _err
}

func (client *Client) GetAuthorizationV2 (pathname *string, method *string, headers map[string]*string, ak *string, secret *string) (_result *string, _err error) {
  sign, _err := client.GetSignatureV2(pathname, method, headers, secret)
  if _err != nil {
    return _result, _err
  }

  _result = tea.String("MNS " + tea.StringValue(ak) + ":" + tea.StringValue(sign))
  return _result, _err
}

func (client *Client) GetSignatureV2 (pathname *string, method *string, headers map[string]*string, secret *string) (_result *string, _err error) {
  canonicalizedResource := client.DefaultString(pathname, tea.String("/"))
  canonicalizedHeaders, _err := client.BuildCanonicalizedHeadersV2(headers)
  if _err != nil {
    return _result, _err
  }

  stringToSign := tea.String(tea.StringValue(method) + "\n" + tea.StringValue(canonicalizedHeaders) + tea.StringValue(canonicalizedResource))
  _body := encodeutil.Base64EncodeToString(signatureutil.HmacSHA1Sign(stringToSign, secret))
  _result = _body
  return _result, _err
}

func (client *Client) BuildCanonicalizedHeadersV2 (headers map[string]*string) (_result *string, _err error) {
  contentMd5 := client.DefaultString(headers["content-md5"], tea.String(""))
  contentType := client.DefaultString(headers["content-type"], tea.String(""))
  date := client.DefaultString(headers["date"], tea.String(""))
  canonicalizedHeaders := tea.String(tea.StringValue(contentMd5) + "\n" + tea.StringValue(contentType) + "\n" + tea.StringValue(date) + "\n")
  mnsHeaders := make(map[string]*string)
  for _, header := range map_.KeySet(headers) {
    lowerHeader := string_.ToLower(header)
    if tea.BoolValue(string_.HasPrefix(lowerHeader, tea.String("x-mns-"))) {
      mnsHeaders[tea.StringValue(lowerHeader)] = headers[tea.StringValue(header)]
    }

  }
  for _, header := range array.AscSort(map_.KeySet(mnsHeaders)) {
    canonicalizedHeaders = tea.String(tea.StringValue(canonicalizedHeaders) + tea.StringValue(header) + ":" + tea.StringValue(mnsHeaders[tea.StringValue(header)]) + "\n")
  }
  _result = canonicalizedHeaders
  return _result , _err
}

func (client *Client) GetAuthorizationV4 (context *spi.InterceptorContext, date *string, accessKeyId *string, accessKeySecret *string) (_result *string, _err error) {
  region, _err := client.GetRegion(context)
  if _err != nil {
    return _result, _err
  }

  dateShort := string_.SubString(date, tea.Int(0), tea.Int(8))
  dateShort = string_.Replace(dateShort, tea.String("T"), tea.String(""), nil)
  scope := tea.String(tea.StringValue(dateShort) + "/" + tea.StringValue(region) + "/mns/" + tea.StringValue(client.SignSuffix))
  signingkey, _err := client.GetSigningkeyV4(accessKeySecret, region, dateShort)
  if _err != nil {
    return _result, _err
  }

  signature, _err := client.GetSignatureV4(context, date, scope, signingkey)
  if _err != nil {
    return _result, _err
  }

  _result = tea.String(tea.StringValue(client.AuthPrefix) + " Credential=" + tea.StringValue(accessKeyId) + "/" + tea.StringValue(scope) + ",Signature=" + tea.StringValue(signature))
  return _result, _err
}

func (client *Client) GetSigningkeyV4 (secret *string, region *string, date *string) (_result []byte, _err error) {
  sc1 := tea.String(tea.StringValue(client.SignPrefix) + tea.StringValue(secret))
  sc2 := signatureutil.HmacSHA256Sign(date, sc1)
  sc3 := signatureutil.HmacSHA256SignByBytes(region, sc2)
  sc4 := signatureutil.HmacSHA256SignByBytes(tea.String("mns"), sc3)
  _result = make([]byte, 0)
  _body := signatureutil.HmacSHA256SignByBytes(client.SignSuffix, sc4)
  _result = _body
  return _result, _err
}

func (client *Client) GetSignatureV4 (context *spi.InterceptorContext, date *string, scope *string, signingkey []byte) (_result *string, _err error) {
  request := context.Request
  canonicalString, _err := client.BuildCanonicalRequestV4(request.Pathname, request.Method, request.Query, request.Headers)
  if _err != nil {
    return _result, _err
  }

  stringToSign := tea.String(tea.StringValue(client.AuthPrefix) + "\n" + tea.StringValue(date) + "\n" + tea.StringValue(scope) + "\n" + tea.StringValue(canonicalString))
  signature := signatureutil.HmacSHA256SignByBytes(stringToSign, signingkey)
  _body := encodeutil.HexEncode(signature)
  _result = _body
  return _result, _err
}

func (client *Client) BuildCanonicalRequestV4 (pathname *string, method *string, query map[string]*string, headers map[string]*string) (_result *string, _err error) {
  canonicalURI := tea.String("/")
  if !tea.BoolValue(util.Empty(pathname)) {
    canonicalURI = pathname
    if !tea.BoolValue(string_.HasPrefix(canonicalURI, tea.String("/"))) {
      canonicalURI = tea.String("/" + tea.StringValue(canonicalURI))
    }

  }

  suffix := tea.String("")
  if !tea.BoolValue(string_.Equals(canonicalURI, tea.String("/"))) && tea.BoolValue(string_.HasSuffix(canonicalURI, tea.String("/"))) {
    suffix = tea.String("/")
  }

  canonicalURI = tea.String(tea.StringValue(openapiutil.GetEncodePath(canonicalURI)) + tea.StringValue(suffix))
  resources, _err := client.BuildCanonicalizedResourceV4(query)
  if _err != nil {
    return _result, _err
  }

  canonicalHeaders, _err := client.BuildCanonicalizedHeadersV4(headers)
  if _err != nil {
    return _result, _err
  }

  _result = tea.String(tea.StringValue(method) + "\n" + tea.StringValue(canonicalURI) + "\n" + tea.StringValue(resources) + "\n" + tea.StringValue(canonicalHeaders) + "\n")
  return _result, _err
}

func (client *Client) BuildCanonicalizedResourceV4 (query map[string]*string) (_result *string, _err error) {
  canonicalizedResource := tea.String("")
  if !tea.BoolValue(util.IsUnset(query)) {
    queryMap := make(map[string]*string)
    for _, key := range map_.KeySet(query) {
      encodedKey := client.PercentEncodeMns(string_.ToLower(string_.Trim(key)))
      encodedValue := tea.String("")
      if !tea.BoolValue(util.IsUnset(query[tea.StringValue(key)])) && !tea.BoolValue(util.Empty(query[tea.StringValue(key)])) {
        encodedValue = client.PercentEncodeMns(string_.Trim(query[tea.StringValue(key)]))
      }

      queryMap[tea.StringValue(encodedKey)] = encodedValue
    }
    queryArray := map_.KeySet(queryMap)
    sortedQueryArray := array.AscSort(queryArray)
    separator := tea.String("")
    for _, encodedKey := range sortedQueryArray {
      canonicalizedResource = tea.String(tea.StringValue(canonicalizedResource) + tea.StringValue(separator) + tea.StringValue(encodedKey))
      encodedValue := queryMap[tea.StringValue(encodedKey)]
      if !tea.BoolValue(util.Empty(encodedValue)) {
        canonicalizedResource = tea.String(tea.StringValue(canonicalizedResource) + "=" + tea.StringValue(encodedValue))
      }

      separator = tea.String("&")
    }
  }

  _result = canonicalizedResource
  return _result , _err
}

func (client *Client) BuildCanonicalizedHeadersV4 (headers map[string]*string) (_result *string, _err error) {
  signedHeaders := make(map[string]*string)
  for _, key := range map_.KeySet(headers) {
    lowerKey := string_.ToLower(key)
    if tea.BoolValue(client.HasSignedHeaderV4(lowerKey)) {
      signedHeaders[tea.StringValue(lowerKey)] = string_.Trim(headers[tea.StringValue(key)])
    }

  }
  canonicalizedHeaders := tea.String("")
  for _, lowerKey := range array.AscSort(map_.KeySet(signedHeaders)) {
    canonicalizedHeaders = tea.String(tea.StringValue(canonicalizedHeaders) + tea.StringValue(lowerKey) + ":" + tea.StringValue(signedHeaders[tea.StringValue(lowerKey)]) + "\n")
  }
  _result = canonicalizedHeaders
  return _result , _err
}

func (client *Client) HasSignedHeaderV4 (header *string) (_result *bool) {
  if tea.BoolValue(string_.Equals(header, tea.String("content-type"))) || tea.BoolValue(string_.Equals(header, tea.String("content-md5"))) {
    _result = tea.Bool(true)
    return _result
  }

  _body := string_.HasPrefix(header, tea.String("x-mns-"))
  _result = _body
  return _result
}

func (client *Client) PercentEncodeMns (value *string) (_result *string) {
  encoded := encodeutil.PercentEncode(value)
  _body := string_.Replace(encoded, tea.String("+"), tea.String("%20"), nil)
  _result = _body
  return _result
}

func (client *Client) GetRegion (context *spi.InterceptorContext) (_result *string, _err error) {
  config := context.Configuration
  if !tea.BoolValue(util.IsUnset(config.RegionId)) && !tea.BoolValue(util.Empty(config.RegionId)) {
    _result = config.RegionId
    return _result , _err
  }

  region := string_.Replace(config.Endpoint, tea.String(".mns.aliyuncs.com"), tea.String(""), nil)
  if tea.BoolValue(string_.Equals(region, config.Endpoint)) {
    _err = tea.NewSDKError(map[string]interface{}{
      "code": "ClientConfigError",
      "message": "The regionId configuration of mns client is missing.",
    })
    return _result, _err
  }

  _result = region
  return _result , _err
}

func (client *Client) GetDateISO8601 () (_result *string, _err error) {
  date := openapiutil.GetTimestamp()
  date = string_.Replace(date, tea.String("-"), tea.String(""), nil)
  _body := string_.Replace(date, tea.String(":"), tea.String(""), nil)
  _result = _body
  return _result, _err
}

