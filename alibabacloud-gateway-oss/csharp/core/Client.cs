// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewayOss
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {
        protected List<string> _default_signed_params;
        protected List<string> _except_signed_params;
        protected List<string> _default_additional_headers;

        public Client(): base()
        {
            this._default_signed_params = new List<string>
            {
                "response-content-type",
                "response-content-language",
                "response-cache-control",
                "logging",
                "response-content-encoding",
                "acl",
                "uploadId",
                "uploads",
                "partNumber",
                "group",
                "link",
                "delete",
                "website",
                "location",
                "objectInfo",
                "objectMeta",
                "response-expires",
                "response-content-disposition",
                "cors",
                "lifecycle",
                "restore",
                "qos",
                "referer",
                "stat",
                "bucketInfo",
                "append",
                "position",
                "security-token",
                "live",
                "comp",
                "status",
                "vod",
                "startTime",
                "endTime",
                "x-oss-process",
                "symlink",
                "callback",
                "callback-var",
                "tagging",
                "encryption",
                "versions",
                "versioning",
                "versionId",
                "policy",
                "requestPayment",
                "x-oss-traffic-limit",
                "qosInfo",
                "asyncFetch",
                "x-oss-request-payer",
                "sequential",
                "inventory",
                "inventoryId",
                "continuation-token",
                "callback",
                "callback-var",
                "worm",
                "wormId",
                "wormExtend",
                "replication",
                "replicationLocation",
                "replicationProgress",
                "transferAcceleration",
                "cname",
                "metaQuery",
                "x-oss-ac-source-ip",
                "x-oss-ac-subnet-mask",
                "x-oss-ac-vpc-id",
                "x-oss-ac-forward-allow",
                "resourceGroup",
                "style",
                "styleName",
                "x-oss-async-process",
                "rtc",
                "accessPoint",
                "accessPointPolicy",
                "httpsConfig",
                "regionsV2",
                "publicAccessBlock",
                "policyStatus",
                "redundancyTransition",
                "redundancyType",
                "redundancyProgress",
                "dataAccelerator",
                "verbose",
                "accessPointForObjectProcess",
                "accessPointConfigForObjectProcess",
                "accessPointPolicyForObjectProcess",
                "bucketArchiveDirectRead",
                "responseHeader",
                "userDefinedLogFieldsConfig",
                "reservedcapacity",
                "requesterQosInfo",
                "qosRequester",
                "resourcePool",
                "resourcePoolInfo",
                "resourcePoolBuckets",
                "processConfiguration",
                "img",
                "asyncFetch",
                "virtualBucket",
                "copy",
                "userRegion",
                "partSize",
                "chunkSize",
                "partUploadId",
                "chunkNumber",
                "userRegion",
                "regionList",
                "eventnotification",
                "cacheConfiguration",
                "dfs",
                "dfsadmin",
                "dfssecurity"
            };
            this._except_signed_params = new List<string>
            {
                "list-type",
                "regions"
            };
            this._default_additional_headers = new List<string>
            {
                "range",
                "if-modified-since"
            };
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = GetEndpoint(config.RegionId, config.Network, config.Endpoint);
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = await GetEndpointAsync(config.RegionId, config.Network, config.Endpoint);
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            Dictionary<string, string> hostMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.HostMap))
            {
                hostMap = request.HostMap;
            }
            string bucketName = hostMap.Get("bucket");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(bucketName))
            {
                bucketName = "";
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Headers.Get("x-oss-meta-*")))
            {
                object tmp = AlibabaCloud.TeaUtil.Common.ParseJSON(request.Headers.Get("x-oss-meta-*"));
                Dictionary<string, object> mapData = AlibabaCloud.TeaUtil.Common.AssertAsMap(tmp);
                Dictionary<string, string> metaData = AlibabaCloud.TeaUtil.Common.StringifyMapValue(mapData);
                List<string> metaKeySet = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(metaData);
                request.Headers["x-oss-meta-*"] = null;

                foreach (var key in metaKeySet) {
                    string newKey = "x-oss-meta-" + key;
                    request.Headers[newKey] = metaData.Get(key);
                }
            }
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string regionId = config.RegionId;
            if (AlibabaCloud.TeaUtil.Common.IsUnset(regionId) || AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = GetRegionIdFromEndpoint(config.Endpoint);
            }
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = credential.GetAccessKeyId();
            string accessKeySecret = credential.GetAccessKeySecret();
            string securityToken = credential.GetSecurityToken();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-oss-security-token"] = securityToken;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "xml"))
                {
                    Dictionary<string, object> reqBodyMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    // for python:
                    // xml_str = OSS_UtilClient.to_xml(req_body_map)
                    string xmlStr = AlibabaCloud.TeaXML.Client.ToXML(reqBodyMap);
                    request.Stream = TeaCore.BytesReadable(xmlStr);
                    request.Headers["content-type"] = "application/xml";
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(xmlStr));
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json"))
                {
                    string reqBodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Stream = TeaCore.BytesReadable(reqBodyStr);
                    request.Headers["content-type"] = "application/json; charset=utf-8";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    Dictionary<string, object> reqBodyForm = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    request.Stream = TeaCore.BytesReadable(AlibabaCloud.OpenApiUtil.Client.ToForm(reqBodyForm));
                    request.Headers["content-type"] = "application/x-www-form-urlencoded";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "binary"))
                {
                    attributeMap.Key = new Dictionary<string, string>
                    {
                        {"crc", ""},
                        {"md5", ""},
                    };
                    request.Stream = AlibabaCloud.OSSUtil.Common.Inject(request.Stream, attributeMap.Key);
                    request.Headers["content-type"] = "application/octet-stream";
                }
            }
            string host = GetHost(config.EndpointType, bucketName, config.Endpoint, context);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", host},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"user-agent", request.UserAgent},
                },
                request.Headers
            );
            string originPath = request.Pathname;
            Dictionary<string, string> originQuery = request.Query;
            if (!AlibabaCloud.TeaUtil.Common.Empty(originPath))
            {
                List<string> pathAndQueries = AlibabaCloud.DarabonbaString.StringUtil.Split(originPath, "?", 2);
                request.Pathname = pathAndQueries[0];
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(pathAndQueries), 2))
                {
                    List<string> pathQueries = AlibabaCloud.DarabonbaString.StringUtil.Split(pathAndQueries[1], "&", null);

                    foreach (var sub in pathQueries) {
                        List<string> item = AlibabaCloud.DarabonbaString.StringUtil.Split(sub, "=", null);
                        string queryKey = item[0];
                        string queryValue = "";
                        if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(item), 2))
                        {
                            queryValue = item[1];
                        }
                        if (AlibabaCloud.TeaUtil.Common.Empty(originQuery.Get(queryKey)))
                        {
                            request.Query[queryKey] = queryValue;
                        }
                    }
                }
            }
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v4");
            request.Headers["authorization"] = GetAuthorization(signatureVersion, bucketName, request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret, regionId);
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            Dictionary<string, string> hostMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.HostMap))
            {
                hostMap = request.HostMap;
            }
            string bucketName = hostMap.Get("bucket");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(bucketName))
            {
                bucketName = "";
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Headers.Get("x-oss-meta-*")))
            {
                object tmp = AlibabaCloud.TeaUtil.Common.ParseJSON(request.Headers.Get("x-oss-meta-*"));
                Dictionary<string, object> mapData = AlibabaCloud.TeaUtil.Common.AssertAsMap(tmp);
                Dictionary<string, string> metaData = AlibabaCloud.TeaUtil.Common.StringifyMapValue(mapData);
                List<string> metaKeySet = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(metaData);
                request.Headers["x-oss-meta-*"] = null;

                foreach (var key in metaKeySet) {
                    string newKey = "x-oss-meta-" + key;
                    request.Headers[newKey] = metaData.Get(key);
                }
            }
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string regionId = config.RegionId;
            if (AlibabaCloud.TeaUtil.Common.IsUnset(regionId) || AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = await GetRegionIdFromEndpointAsync(config.Endpoint);
            }
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = await credential.GetAccessKeyIdAsync();
            string accessKeySecret = await credential.GetAccessKeySecretAsync();
            string securityToken = await credential.GetSecurityTokenAsync();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-oss-security-token"] = securityToken;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "xml"))
                {
                    Dictionary<string, object> reqBodyMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    // for python:
                    // xml_str = OSS_UtilClient.to_xml(req_body_map)
                    string xmlStr = AlibabaCloud.TeaXML.Client.ToXML(reqBodyMap);
                    request.Stream = TeaCore.BytesReadable(xmlStr);
                    request.Headers["content-type"] = "application/xml";
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(xmlStr));
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json"))
                {
                    string reqBodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Stream = TeaCore.BytesReadable(reqBodyStr);
                    request.Headers["content-type"] = "application/json; charset=utf-8";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    Dictionary<string, object> reqBodyForm = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    request.Stream = TeaCore.BytesReadable(AlibabaCloud.OpenApiUtil.Client.ToForm(reqBodyForm));
                    request.Headers["content-type"] = "application/x-www-form-urlencoded";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "binary"))
                {
                    attributeMap.Key = new Dictionary<string, string>
                    {
                        {"crc", ""},
                        {"md5", ""},
                    };
                    request.Stream = AlibabaCloud.OSSUtil.Common.Inject(request.Stream, attributeMap.Key);
                    request.Headers["content-type"] = "application/octet-stream";
                }
            }
            string host = await GetHostAsync(config.EndpointType, bucketName, config.Endpoint, context);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", host},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"user-agent", request.UserAgent},
                },
                request.Headers
            );
            string originPath = request.Pathname;
            Dictionary<string, string> originQuery = request.Query;
            if (!AlibabaCloud.TeaUtil.Common.Empty(originPath))
            {
                List<string> pathAndQueries = AlibabaCloud.DarabonbaString.StringUtil.Split(originPath, "?", 2);
                request.Pathname = pathAndQueries[0];
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(pathAndQueries), 2))
                {
                    List<string> pathQueries = AlibabaCloud.DarabonbaString.StringUtil.Split(pathAndQueries[1], "&", null);

                    foreach (var sub in pathQueries) {
                        List<string> item = AlibabaCloud.DarabonbaString.StringUtil.Split(sub, "=", null);
                        string queryKey = item[0];
                        string queryValue = "";
                        if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(item), 2))
                        {
                            queryValue = item[1];
                        }
                        if (AlibabaCloud.TeaUtil.Common.Empty(originQuery.Get(queryKey)))
                        {
                            request.Query[queryKey] = queryValue;
                        }
                    }
                }
            }
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v4");
            request.Headers["authorization"] = await GetAuthorizationAsync(signatureVersion, bucketName, request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret, regionId);
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            string bodyStr = null;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                if (!AlibabaCloud.TeaUtil.Common.Empty(bodyStr))
                {
                    Dictionary<string, object> respMap = AlibabaCloud.TeaXML.Client.ParseXml(bodyStr, null);
                    Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(respMap.Get("Error"));
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", err.Get("Code")},
                        {"message", err.Get("Message")},
                        {"data", new Dictionary<string, object>
                        {
                            {"statusCode", response.StatusCode},
                            {"requestId", err.Get("RequestId")},
                            {"ecCode", err.Get("EC")},
                            {"Recommend", err.Get("RecommendDoc")},
                            {"hostId", err.Get("HostId")},
                            {"AccessDeniedDetail", err.Get("AccessDeniedDetail")},
                        }},
                    });
                }
                else
                {
                    Dictionary<string, string> headers = response.Headers;
                    string requestId = headers.Get("x-oss-request-id");
                    string ecCode = headers.Get("x-oss-ec-code");
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", response.StatusCode},
                        {"message", null},
                        {"data", new Dictionary<string, object>
                        {
                            {"statusCode", response.StatusCode},
                            {"requestId", "" + requestId},
                            {"ecCode", ecCode},
                        }},
                    });
                }
            }
            Dictionary<string, string> ctx = attributeMap.Key;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx))
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("crc")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-oss-hash-crc64ecma")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("crc"), response.Headers.Get("x-oss-hash-crc64ecma")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "CrcNotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientCrc", ctx.Get("crc")},
                            {"serverCrc", response.Headers.Get("x-oss-hash-crc64ecma")},
                        }},
                    });
                }
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("md5")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("content-md5")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("md5"), response.Headers.Get("content-md5")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "MD5NotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientMD5", ctx.Get("md5")},
                            {"serverMD5", response.Headers.Get("content-md5")},
                        }},
                    });
                }
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(response.StatusCode, 204))
                {
                    AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.BodyType, "xml"))
                {
                    bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                    response.DeserializedBody = bodyStr;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(bodyStr))
                    {
                        object result = AlibabaCloud.GatewayOss_Util.Common.ParseXml(bodyStr, request.Action);
                        // for no util language
                        // var result : any = XML.parseXml(bodyStr, null);
                        try
                        {
                            response.DeserializedBody = AlibabaCloud.TeaUtil.Common.AssertAsMap(result);
                        }
                        catch (TeaException error)
                        {
                            response.DeserializedBody = result;
                        }
                        catch (Exception _error)
                        {
                            TeaException error = new TeaException(new Dictionary<string, object>
                            {
                                { "message", _error.Message }
                            });
                            response.DeserializedBody = result;
                        }
                    }
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
                {
                    response.DeserializedBody = response.Body;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "byte"))
                {
                    byte[] byt = AlibabaCloud.TeaUtil.Common.ReadAsBytes(response.Body);
                    response.DeserializedBody = byt;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "string"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
                {
                    object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> res = AlibabaCloud.TeaUtil.Common.AssertAsMap(obj);
                    response.DeserializedBody = res;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                }
                else
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
            }
        }

        public async Task ModifyResponseAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            string bodyStr = null;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                if (!AlibabaCloud.TeaUtil.Common.Empty(bodyStr))
                {
                    Dictionary<string, object> respMap = AlibabaCloud.TeaXML.Client.ParseXml(bodyStr, null);
                    Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(respMap.Get("Error"));
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", err.Get("Code")},
                        {"message", err.Get("Message")},
                        {"data", new Dictionary<string, object>
                        {
                            {"statusCode", response.StatusCode},
                            {"requestId", err.Get("RequestId")},
                            {"ecCode", err.Get("EC")},
                            {"Recommend", err.Get("RecommendDoc")},
                            {"hostId", err.Get("HostId")},
                            {"AccessDeniedDetail", err.Get("AccessDeniedDetail")},
                        }},
                    });
                }
                else
                {
                    Dictionary<string, string> headers = response.Headers;
                    string requestId = headers.Get("x-oss-request-id");
                    string ecCode = headers.Get("x-oss-ec-code");
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", response.StatusCode},
                        {"message", null},
                        {"data", new Dictionary<string, object>
                        {
                            {"statusCode", response.StatusCode},
                            {"requestId", "" + requestId},
                            {"ecCode", ecCode},
                        }},
                    });
                }
            }
            Dictionary<string, string> ctx = attributeMap.Key;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx))
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("crc")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-oss-hash-crc64ecma")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("crc"), response.Headers.Get("x-oss-hash-crc64ecma")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "CrcNotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientCrc", ctx.Get("crc")},
                            {"serverCrc", response.Headers.Get("x-oss-hash-crc64ecma")},
                        }},
                    });
                }
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("md5")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("content-md5")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("md5"), response.Headers.Get("content-md5")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "MD5NotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientMD5", ctx.Get("md5")},
                            {"serverMD5", response.Headers.Get("content-md5")},
                        }},
                    });
                }
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(response.StatusCode, 204))
                {
                    AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.BodyType, "xml"))
                {
                    bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                    response.DeserializedBody = bodyStr;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(bodyStr))
                    {
                        object result = AlibabaCloud.GatewayOss_Util.Common.ParseXml(bodyStr, request.Action);
                        // for no util language
                        // var result : any = XML.parseXml(bodyStr, null);
                        try
                        {
                            response.DeserializedBody = AlibabaCloud.TeaUtil.Common.AssertAsMap(result);
                        }
                        catch (TeaException error)
                        {
                            response.DeserializedBody = result;
                        }
                        catch (Exception _error)
                        {
                            TeaException error = new TeaException(new Dictionary<string, object>
                            {
                                { "message", _error.Message }
                            });
                            response.DeserializedBody = result;
                        }
                    }
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
                {
                    response.DeserializedBody = response.Body;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "byte"))
                {
                    byte[] byt = AlibabaCloud.TeaUtil.Common.ReadAsBytes(response.Body);
                    response.DeserializedBody = byt;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "string"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
                {
                    object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> res = AlibabaCloud.TeaUtil.Common.AssertAsMap(obj);
                    response.DeserializedBody = res;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                }
                else
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
            }
        }

        public string GetRegionIdFromEndpoint(string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                int? idx = -1;
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(endpoint, "oss-") && AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 4, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".mgw.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".mgw.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".mgw-internal.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".mgw-internal.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, "-internal.oss-data-acc.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, "-internal.oss-data-acc.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".oss-dls.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".oss-dls.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
            }
            return "cn-hangzhou";
        }

        public async Task<string> GetRegionIdFromEndpointAsync(string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                int? idx = -1;
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(endpoint, "oss-") && AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 4, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".mgw.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".mgw.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".mgw-internal.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".mgw-internal.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, "-internal.oss-data-acc.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, "-internal.oss-data-acc.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(endpoint, ".oss-dls.aliyuncs.com"))
                {
                    idx = AlibabaCloud.DarabonbaString.StringUtil.Index(endpoint, ".oss-dls.aliyuncs.com");
                    return AlibabaCloud.DarabonbaString.StringUtil.SubString(endpoint, 0, idx);
                }
            }
            return "cn-hangzhou";
        }

        public string GetEndpoint(string regionId, string network, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = "cn-hangzhou";
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(network))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "internal"))
                {
                    return "oss-" + regionId + "-internal.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "ipv6"))
                {
                    return "" + regionId + ".oss.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "accelerate"))
                {
                    return "oss-" + network + ".aliyuncs.com";
                }
            }
            return "oss-" + regionId + ".aliyuncs.com";
        }

        public async Task<string> GetEndpointAsync(string regionId, string network, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = "cn-hangzhou";
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(network))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "internal"))
                {
                    return "oss-" + regionId + "-internal.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "ipv6"))
                {
                    return "" + regionId + ".oss.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "accelerate"))
                {
                    return "oss-" + network + ".aliyuncs.com";
                }
            }
            return "oss-" + regionId + ".aliyuncs.com";
        }

        public string GetHost(string endpointType, string bucketName, string endpoint, AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(endpoint, ".mgw.aliyuncs.com") && !AlibabaCloud.TeaUtil.Common.IsUnset(context.Request.HostMap.Get("userid")))
            {
                return "" + context.Request.HostMap.Get("userid") + "." + endpoint;
            }
            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(endpoint, ".mgw-internal.aliyuncs.com") && !AlibabaCloud.TeaUtil.Common.IsUnset(context.Request.HostMap.Get("userid")))
            {
                return "" + context.Request.HostMap.Get("userid") + "." + endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                return endpoint;
            }
            string host = "" + bucketName + "." + endpoint;
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpointType))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "ip"))
                {
                    host = "" + endpoint + "/" + bucketName;
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "cname"))
                {
                    host = endpoint;
                }
            }
            return host;
        }

        public async Task<string> GetHostAsync(string endpointType, string bucketName, string endpoint, AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(endpoint, ".mgw.aliyuncs.com") && !AlibabaCloud.TeaUtil.Common.IsUnset(context.Request.HostMap.Get("userid")))
            {
                return "" + context.Request.HostMap.Get("userid") + "." + endpoint;
            }
            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(endpoint, ".mgw-internal.aliyuncs.com") && !AlibabaCloud.TeaUtil.Common.IsUnset(context.Request.HostMap.Get("userid")))
            {
                return "" + context.Request.HostMap.Get("userid") + "." + endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                return endpoint;
            }
            string host = "" + bucketName + "." + endpoint;
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpointType))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "ip"))
                {
                    host = "" + endpoint + "/" + bucketName;
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "cname"))
                {
                    host = endpoint;
                }
            }
            return host;
        }

        public string GetAuthorization(string signatureVersion, string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret, string regionId)
        {
            string sign = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(signatureVersion))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v1"))
                {
                    sign = GetSignatureV1(bucketName, pathname, method, query, headers, secret);
                    return "OSS " + ak + ":" + sign;
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v2"))
                {
                    List<string> additionalHeaderNames = GetAdditionalHeaderNamesV2(headers);
                    sign = GetSignatureV2(bucketName, pathname, method, query, headers, secret, additionalHeaderNames);
                    string additionalHeaders = JoinSemicolon(additionalHeaderNames);
                    if (AlibabaCloud.TeaUtil.Common.Empty(additionalHeaders))
                    {
                        return "OSS2 AccessKeyId:" + ak + ",Signature:" + sign;
                    }
                    return "OSS2 AccessKeyId:" + ak + ",AdditionalHeaders:" + additionalHeaders + ",Signature:" + sign;
                }
            }
            string dateTime = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            dateTime = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateTime, "-", "", null);
            dateTime = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateTime, ":", "", null);
            headers["x-oss-date"] = dateTime;
            headers["x-oss-content-sha256"] = "UNSIGNED-PAYLOAD";
            string onlyDate = AlibabaCloud.DarabonbaString.StringUtil.SubString(dateTime, 0, 8);
            string cred = "" + ak + "/" + onlyDate + "/" + regionId + "/oss/aliyun_v4_request";
            sign = GetSignatureV4(bucketName, pathname, method, query, headers, onlyDate, regionId, secret);
            return "OSS4-HMAC-SHA256 Credential=" + cred + ", Signature=" + sign;
        }

        public async Task<string> GetAuthorizationAsync(string signatureVersion, string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret, string regionId)
        {
            string sign = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(signatureVersion))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v1"))
                {
                    sign = await GetSignatureV1Async(bucketName, pathname, method, query, headers, secret);
                    return "OSS " + ak + ":" + sign;
                }
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v2"))
                {
                    List<string> additionalHeaderNames = await GetAdditionalHeaderNamesV2Async(headers);
                    sign = await GetSignatureV2Async(bucketName, pathname, method, query, headers, secret, additionalHeaderNames);
                    string additionalHeaders = await JoinSemicolonAsync(additionalHeaderNames);
                    if (AlibabaCloud.TeaUtil.Common.Empty(additionalHeaders))
                    {
                        return "OSS2 AccessKeyId:" + ak + ",Signature:" + sign;
                    }
                    return "OSS2 AccessKeyId:" + ak + ",AdditionalHeaders:" + additionalHeaders + ",Signature:" + sign;
                }
            }
            string dateTime = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            dateTime = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateTime, "-", "", null);
            dateTime = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateTime, ":", "", null);
            headers["x-oss-date"] = dateTime;
            headers["x-oss-content-sha256"] = "UNSIGNED-PAYLOAD";
            string onlyDate = AlibabaCloud.DarabonbaString.StringUtil.SubString(dateTime, 0, 8);
            string cred = "" + ak + "/" + onlyDate + "/" + regionId + "/oss/aliyun_v4_request";
            sign = await GetSignatureV4Async(bucketName, pathname, method, query, headers, onlyDate, regionId, secret);
            return "OSS4-HMAC-SHA256 Credential=" + cred + ", Signature=" + sign;
        }

        public byte[] GetSignKey(string secret, string onlyDate, string regionId)
        {
            string temp = "aliyun_v4" + secret;
            byte[] res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(onlyDate, temp);
            res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(regionId, res);
            res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("oss", res);
            res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", res);
            return res;
        }

        public async Task<byte[]> GetSignKeyAsync(string secret, string onlyDate, string regionId)
        {
            string temp = "aliyun_v4" + secret;
            byte[] res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(onlyDate, temp);
            res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(regionId, res);
            res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("oss", res);
            res = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", res);
            return res;
        }

        public string GetSignatureV4(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string onlyDate, string regionId, string secret)
        {
            byte[] signingkey = GetSignKey(secret, onlyDate, regionId);
            string canonicalizedUri = pathname;
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
                {
                    canonicalizedUri = "/" + bucketName + canonicalizedUri;
                }
            }
            else
            {
                if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
                {
                    canonicalizedUri = "/" + bucketName + "/";
                }
                else
                {
                    canonicalizedUri = "/";
                }
            }
            // for java:
            // String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
            // canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
            canonicalizedUri = AlibabaCloud.OpenApiUtil.Client.GetEncodePath(canonicalizedUri);
            Dictionary<string, string> queryMap = new Dictionary<string, string>(){};

            foreach (var queryKey in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query)) {
                string queryValue = null;
                if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(queryKey)))
                {
                    queryValue = AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(queryKey));
                    queryValue = AlibabaCloud.DarabonbaString.StringUtil.Replace(queryValue, "+", "%20", null);
                }
                queryKey = AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(queryKey);
                queryKey = AlibabaCloud.DarabonbaString.StringUtil.Replace(queryKey, "+", "%20", null);
                // for go : queryMap[tea.StringValue(queryKey)] = queryValue
                queryMap[queryKey] = queryValue;
            }
            string canonicalizedQueryString = BuildCanonicalizedQueryStringV4(queryMap);
            string canonicalizedHeaders = BuildCanonicalizedHeadersV4(headers);
            string payload = "UNSIGNED-PAYLOAD";
            string canonicalRequest = "" + method + "\n" + canonicalizedUri + "\n" + canonicalizedQueryString + "\n" + canonicalizedHeaders + "\n\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(canonicalRequest), "ACS4-HMAC-SHA256"));
            string scope = "" + onlyDate + "/" + regionId + "/oss/aliyun_v4_request";
            string stringToSign = "OSS4-HMAC-SHA256\n" + headers.Get("x-oss-date") + "\n" + scope + "\n" + hex;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public async Task<string> GetSignatureV4Async(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string onlyDate, string regionId, string secret)
        {
            byte[] signingkey = await GetSignKeyAsync(secret, onlyDate, regionId);
            string canonicalizedUri = pathname;
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
                {
                    canonicalizedUri = "/" + bucketName + canonicalizedUri;
                }
            }
            else
            {
                if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
                {
                    canonicalizedUri = "/" + bucketName + "/";
                }
                else
                {
                    canonicalizedUri = "/";
                }
            }
            // for java:
            // String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
            // canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
            canonicalizedUri = AlibabaCloud.OpenApiUtil.Client.GetEncodePath(canonicalizedUri);
            Dictionary<string, string> queryMap = new Dictionary<string, string>(){};

            foreach (var queryKey in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query)) {
                string queryValue = null;
                if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(queryKey)))
                {
                    queryValue = AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(queryKey));
                    queryValue = AlibabaCloud.DarabonbaString.StringUtil.Replace(queryValue, "+", "%20", null);
                }
                queryKey = AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(queryKey);
                queryKey = AlibabaCloud.DarabonbaString.StringUtil.Replace(queryKey, "+", "%20", null);
                // for go : queryMap[tea.StringValue(queryKey)] = queryValue
                queryMap[queryKey] = queryValue;
            }
            string canonicalizedQueryString = await BuildCanonicalizedQueryStringV4Async(queryMap);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersV4Async(headers);
            string payload = "UNSIGNED-PAYLOAD";
            string canonicalRequest = "" + method + "\n" + canonicalizedUri + "\n" + canonicalizedQueryString + "\n" + canonicalizedHeaders + "\n\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(canonicalRequest), "ACS4-HMAC-SHA256"));
            string scope = "" + onlyDate + "/" + regionId + "/oss/aliyun_v4_request";
            string stringToSign = "OSS4-HMAC-SHA256\n" + headers.Get("x-oss-date") + "\n" + scope + "\n" + hex;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public string BuildCanonicalizedQueryStringV4(Dictionary<string, string> queryMap)
        {
            string canonicalizedQueryString = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(queryMap))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var key in sortedQueryArray) {
                    canonicalizedQueryString = "" + canonicalizedQueryString + separator + key;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(queryMap.Get(key)))
                    {
                        canonicalizedQueryString = "" + canonicalizedQueryString + "=" + queryMap.Get(key);
                    }
                    separator = "&";
                }
            }
            return canonicalizedQueryString;
        }

        public async Task<string> BuildCanonicalizedQueryStringV4Async(Dictionary<string, string> queryMap)
        {
            string canonicalizedQueryString = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(queryMap))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var key in sortedQueryArray) {
                    canonicalizedQueryString = "" + canonicalizedQueryString + separator + key;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(queryMap.Get(key)))
                    {
                        canonicalizedQueryString = "" + canonicalizedQueryString + "=" + queryMap.Get(key);
                    }
                    separator = "&";
                }
            }
            return canonicalizedQueryString;
        }

        public string BuildCanonicalizedHeadersV4(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> headersArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeadersArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(headersArray);

            foreach (var key in sortedHeadersArray) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-oss-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-md5"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + key + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(key)) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersV4Async(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> headersArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeadersArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(headersArray);

            foreach (var key in sortedHeadersArray) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-oss-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-md5"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + key + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(key)) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public string GetSignatureV1(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = "";
            string stringToSign = "";
            if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                resource = "/" + bucketName;
            }
            resource = "" + resource + pathname;
            string canonicalizedResource = BuildCanonicalizedResource(resource, query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureV1Async(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = "";
            string stringToSign = "";
            if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                resource = "/" + bucketName;
            }
            resource = "" + resource + pathname;
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(resource, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public string BuildCanonicalizedResource(string pathname, Dictionary<string, string> query)
        {
            string canonicalizedResource = pathname;
            List<string> queryKeys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryKeys);
            string separator = "?";

            foreach (var paramName in sortedParams) {
                if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(_default_signed_params, paramName) || AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(paramName, "x-oss-"))
                {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(paramName);
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceAsync(string pathname, Dictionary<string, string> query)
        {
            string canonicalizedResource = pathname;
            List<string> queryKeys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryKeys);
            string separator = "?";

            foreach (var paramName in sortedParams) {
                if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(_default_signed_params, paramName) || AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(paramName, "x-oss-"))
                {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(paramName);
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeaders(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            string contentType = headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string contentMd5 = headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            canonicalizedHeaders = "" + canonicalizedHeaders + contentMd5 + "\n" + contentType + "\n" + headers.Get("date") + "\n";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-oss-") && !AlibabaCloud.TeaUtil.Common.IsUnset(headers.Get(header)))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersAsync(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            string contentType = headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string contentMd5 = headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            canonicalizedHeaders = "" + canonicalizedHeaders + contentMd5 + "\n" + contentType + "\n" + headers.Get("date") + "\n";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-oss-") && !AlibabaCloud.TeaUtil.Common.IsUnset(headers.Get(header)))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public static string V2UriEncode(string value)
        {
            if (AlibabaCloud.TeaUtil.Common.Empty(value))
            {
                return "";
            }
            string encoded = AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(value);
            encoded = AlibabaCloud.DarabonbaString.StringUtil.Replace(encoded, "+", "%20", null);
            encoded = AlibabaCloud.DarabonbaString.StringUtil.Replace(encoded, "%7E", "~", null);
            encoded = AlibabaCloud.DarabonbaString.StringUtil.Replace(encoded, "%2D", "-", null);
            encoded = AlibabaCloud.DarabonbaString.StringUtil.Replace(encoded, "%5F", "_", null);
            encoded = AlibabaCloud.DarabonbaString.StringUtil.Replace(encoded, "%2E", ".", null);
            return encoded;
        }

        public string GetHeaderValue(Dictionary<string, string> headers, string name)
        {
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(headers.Get(name)))
            {
                return headers.Get(name);
            }

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), name))
                {
                    return headers.Get(header);
                }
            }
            return "";
        }

        public async Task<string> GetHeaderValueAsync(Dictionary<string, string> headers, string name)
        {
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(headers.Get(name)))
            {
                return headers.Get(name);
            }

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), name))
                {
                    return headers.Get(header);
                }
            }
            return "";
        }

        public List<string> GetAdditionalHeaderNamesV2(Dictionary<string, string> headers)
        {
            Dictionary<string, string> additionalHeaders = new Dictionary<string, string>(){};

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerHeader = AlibabaCloud.DarabonbaString.StringUtil.ToLower(header);
                if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(_default_additional_headers, lowerHeader))
                {
                    additionalHeaders[lowerHeader] = lowerHeader;
                }
            }
            return AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(additionalHeaders));
        }

        public async Task<List<string>> GetAdditionalHeaderNamesV2Async(Dictionary<string, string> headers)
        {
            Dictionary<string, string> additionalHeaders = new Dictionary<string, string>(){};

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerHeader = AlibabaCloud.DarabonbaString.StringUtil.ToLower(header);
                if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(_default_additional_headers, lowerHeader))
                {
                    additionalHeaders[lowerHeader] = lowerHeader;
                }
            }
            return AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(additionalHeaders));
        }

        public string JoinSemicolon(List<string> items)
        {
            string result = "";
            string separator = "";

            foreach (var item in items) {
                result = "" + result + separator + item;
                separator = ";";
            }
            return result;
        }

        public async Task<string> JoinSemicolonAsync(List<string> items)
        {
            string result = "";
            string separator = "";

            foreach (var item in items) {
                result = "" + result + separator + item;
                separator = ";";
            }
            return result;
        }

        public string BuildCanonicalizedOssHeadersV2(Dictionary<string, string> headers, List<string> additionalHeaderNames)
        {
            Dictionary<string, string> canonHeaders = new Dictionary<string, string>(){};

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerHeader = AlibabaCloud.DarabonbaString.StringUtil.ToLower(header);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerHeader, "x-oss-"))
                {
                    canonHeaders[lowerHeader] = headers.Get(header);
                }
            }

            foreach (var name in additionalHeaderNames) {
                canonHeaders[name] = GetHeaderValue(headers, name);
            }
            string canonicalizedHeaders = "";

            foreach (var header in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(canonHeaders))) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + canonHeaders.Get(header) + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedOssHeadersV2Async(Dictionary<string, string> headers, List<string> additionalHeaderNames)
        {
            Dictionary<string, string> canonHeaders = new Dictionary<string, string>(){};

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerHeader = AlibabaCloud.DarabonbaString.StringUtil.ToLower(header);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerHeader, "x-oss-"))
                {
                    canonHeaders[lowerHeader] = headers.Get(header);
                }
            }

            foreach (var name in additionalHeaderNames) {
                canonHeaders[name] = await GetHeaderValueAsync(headers, name);
            }
            string canonicalizedHeaders = "";

            foreach (var header in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(canonHeaders))) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + canonHeaders.Get(header) + "\n";
            }
            return canonicalizedHeaders;
        }

        public string BuildCanonicalizedQueryStringV2(Dictionary<string, string> query)
        {
            Dictionary<string, string> queryMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {

                foreach (var queryKey in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query)) {
                    string encodedKey = V2UriEncode(queryKey);
                    string encodedValue = "";
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(queryKey)))
                    {
                        encodedValue = V2UriEncode(query.Get(queryKey));
                    }
                    queryMap[encodedKey] = encodedValue;
                }
            }
            if (AlibabaCloud.TeaUtil.Common.IsUnset(queryMap) || AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap)), 0))
            {
                return "";
            }
            string canonicalizedQueryString = "?";
            string separator = "";

            foreach (var key in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap))) {
                canonicalizedQueryString = "" + canonicalizedQueryString + separator + key;
                if (!AlibabaCloud.TeaUtil.Common.Empty(queryMap.Get(key)))
                {
                    canonicalizedQueryString = "" + canonicalizedQueryString + "=" + queryMap.Get(key);
                }
                separator = "&";
            }
            return canonicalizedQueryString;
        }

        public async Task<string> BuildCanonicalizedQueryStringV2Async(Dictionary<string, string> query)
        {
            Dictionary<string, string> queryMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {

                foreach (var queryKey in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query)) {
                    string encodedKey = V2UriEncode(queryKey);
                    string encodedValue = "";
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(queryKey)))
                    {
                        encodedValue = V2UriEncode(query.Get(queryKey));
                    }
                    queryMap[encodedKey] = encodedValue;
                }
            }
            if (AlibabaCloud.TeaUtil.Common.IsUnset(queryMap) || AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap)), 0))
            {
                return "";
            }
            string canonicalizedQueryString = "?";
            string separator = "";

            foreach (var key in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap))) {
                canonicalizedQueryString = "" + canonicalizedQueryString + separator + key;
                if (!AlibabaCloud.TeaUtil.Common.Empty(queryMap.Get(key)))
                {
                    canonicalizedQueryString = "" + canonicalizedQueryString + "=" + queryMap.Get(key);
                }
                separator = "&";
            }
            return canonicalizedQueryString;
        }

        public string BuildCanonicalizedResourceV2(string bucketName, string pathname, Dictionary<string, string> query)
        {
            string resourcePath = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                resourcePath = "/" + bucketName + pathname;
            }
            else if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                resourcePath = pathname;
            }
            string canonicalizedResource = V2UriEncode(resourcePath);
            canonicalizedResource = "" + canonicalizedResource + BuildCanonicalizedQueryStringV2(query);
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceV2Async(string bucketName, string pathname, Dictionary<string, string> query)
        {
            string resourcePath = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                resourcePath = "/" + bucketName + pathname;
            }
            else if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                resourcePath = pathname;
            }
            string canonicalizedResource = V2UriEncode(resourcePath);
            canonicalizedResource = "" + canonicalizedResource + await BuildCanonicalizedQueryStringV2Async(query);
            return canonicalizedResource;
        }

        public string GetSignatureV2(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret, List<string> additionalHeaderNames)
        {
            string contentMd5 = AlibabaCloud.TeaUtil.Common.DefaultString(headers.Get("content-md5"), "");
            string contentType = AlibabaCloud.TeaUtil.Common.DefaultString(headers.Get("content-type"), "");
            string date = AlibabaCloud.TeaUtil.Common.DefaultString(headers.Get("date"), "");
            string canonicalizedOssHeaders = BuildCanonicalizedOssHeadersV2(headers, additionalHeaderNames);
            string additionalHeaders = JoinSemicolon(additionalHeaderNames);
            string canonicalizedResource = BuildCanonicalizedResourceV2(bucketName, pathname, query);
            string stringToSign = "" + method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedOssHeaders + additionalHeaders + "\n" + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureV2Async(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret, List<string> additionalHeaderNames)
        {
            string contentMd5 = AlibabaCloud.TeaUtil.Common.DefaultString(headers.Get("content-md5"), "");
            string contentType = AlibabaCloud.TeaUtil.Common.DefaultString(headers.Get("content-type"), "");
            string date = AlibabaCloud.TeaUtil.Common.DefaultString(headers.Get("date"), "");
            string canonicalizedOssHeaders = await BuildCanonicalizedOssHeadersV2Async(headers, additionalHeaderNames);
            string additionalHeaders = await JoinSemicolonAsync(additionalHeaderNames);
            string canonicalizedResource = await BuildCanonicalizedResourceV2Async(bucketName, pathname, query);
            string stringToSign = "" + method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedOssHeaders + additionalHeaders + "\n" + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret));
        }

    }
}
