# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class RoutingRuleRedirect(DaraModel):
    def __init__(
        self,
        enable_replace_prefix: bool = None,
        host_name: str = None,
        http_redirect_code: int = None,
        mirror_allow_get_image_info: bool = None,
        mirror_allow_head_object: bool = None,
        mirror_allow_video_snapshot: bool = None,
        mirror_async_status: int = None,
        mirror_auth: main_models.RoutingRuleRedirectMirrorAuth = None,
        mirror_check_md_5: bool = None,
        mirror_dst_region: str = None,
        mirror_dst_slave_vpc_id: str = None,
        mirror_dst_vpc_id: str = None,
        mirror_follow_redirect: bool = None,
        mirror_headers: main_models.RoutingRuleRedirectMirrorHeaders = None,
        mirror_is_express_tunnel: bool = None,
        mirror_multi_alternates: main_models.RoutingRuleRedirectMirrorMultiAlternates = None,
        mirror_pass_original_slashes: bool = None,
        mirror_pass_query_string: bool = None,
        mirror_proxy_pass: bool = None,
        mirror_return_headers: main_models.RoutingRuleRedirectMirrorReturnHeaders = None,
        mirror_role: str = None,
        mirror_sni: bool = None,
        mirror_save_oss_meta: bool = None,
        mirror_switch_all_errors: bool = None,
        mirror_taggings: main_models.RoutingRuleRedirectMirrorTaggings = None,
        mirror_tunnel_id: str = None,
        mirror_url: str = None,
        mirror_urlprobe: str = None,
        mirror_urlslave: str = None,
        mirror_user_last_modified: bool = None,
        mirror_using_role: bool = None,
        pass_query_string: bool = None,
        protocol: str = None,
        redirect_type: str = None,
        replace_key_prefix_with: str = None,
        replace_key_with: str = None,
        transparent_mirror_response_codes: str = None,
    ):
        # If this parameter is set to true, the prefix of the object names is replaced with the value specified by ReplaceKeyPrefixWith. If this parameter is not specified or empty, the prefix of object names is truncated.
        # 
        # >  When the ReplaceKeyWith parameter is not empty, the EnableReplacePrefix parameter cannot be set to true.
        # 
        # Default value: false.
        self.enable_replace_prefix = enable_replace_prefix
        # The domain name used for redirection. The domain name must comply with the domain naming rules. For example, if you access an object named test, Protocol is set to https, and Hostname is set to `example.com`, the value of the Location header is `https://example.com/test`.
        self.host_name = host_name
        # The HTTP redirect code in the response. This parameter takes effect only when RedirectType is set to External or AliCDN. Valid values: 301, 302, and 307.
        self.http_redirect_code = http_redirect_code
        # Whether to allow get image information in mirror-based back-to-origin.
        self.mirror_allow_get_image_info = mirror_allow_get_image_info
        # Whether to allow take HeadObject in mirror-based back-to-origin.
        self.mirror_allow_head_object = mirror_allow_head_object
        # Whether to allow take video snapshot in mirror-based back-to-origin.
        self.mirror_allow_video_snapshot = mirror_allow_video_snapshot
        # The HTTP status codes that trigger the asynchronous pull mode in mirror-based back-to-origin.
        self.mirror_async_status = mirror_async_status
        # The authentication information for the origin server in mirror-based back-to-origin.
        self.mirror_auth = mirror_auth
        # Specifies whether to check the MD5 hash of the body of the response returned by the origin. This parameter takes effect only when the value of RedirectType is Mirror. When **MirrorCheckMd5** is set to true and the response returned by the origin includes the Content-Md5 header, OSS checks whether the MD5 hash of the obtained data matches the header value. If the MD5 hash of the obtained data does not match the header value, the obtained data is not stored in OSS. Default value: false.
        self.mirror_check_md_5 = mirror_check_md_5
        # The VPC region for mirror-based back-to-origin express tunnel.
        self.mirror_dst_region = mirror_dst_region
        # The slave VPC ID for mirror-based back-to-origin express tunnel.
        self.mirror_dst_slave_vpc_id = mirror_dst_slave_vpc_id
        # The VPC ID for mirror-based back-to-origin express tunnel.
        self.mirror_dst_vpc_id = mirror_dst_vpc_id
        # Specifies whether to redirect the access to the address specified by Location if the origin returns an HTTP 3xx status code. This parameter takes effect only when the value of RedirectType is Mirror. For example, when a mirroring-based back-to-origin request is initiated, the origin returns 302 and Location is specified.
        # 
        # *   If you set MirrorFollowRedirect to true, OSS continues requesting the resource at the address specified by Location. The access can be redirected up to 10 times. If the access is redirected more than 10 times, the mirroring-based back-to-origin request fails.
        # *   If you set MirrorFollowRedirect to false, OSS returns 302 and passes through Location.
        # 
        # Default value: true.
        self.mirror_follow_redirect = mirror_follow_redirect
        # The headers contained in the response that is returned when you use mirroring-based back-to-origin. This parameter takes effect only when the value of RedirectType is Mirror.
        self.mirror_headers = mirror_headers
        # Mirror-based back-to-origin with express tunnel.
        self.mirror_is_express_tunnel = mirror_is_express_tunnel
        # The container to store the configuration for multiple origins in mirror-based back-to-origin.
        self.mirror_multi_alternates = mirror_multi_alternates
        # Whether to pass slashes to origin.
        self.mirror_pass_original_slashes = mirror_pass_original_slashes
        # This parameter plays the same role as PassQueryString and has a higher priority than PassQueryString. This parameter takes effect only when the value of RedirectType is Mirror. Default value: false.
        # 
        # Valid values:
        # 
        # *   true
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        # *   false
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        self.mirror_pass_query_string = mirror_pass_query_string
        # Not save data in web-based back-to-origin.
        self.mirror_proxy_pass = mirror_proxy_pass
        # Container to store the rules for setting response headers in mirror-based back-to-origin.
        self.mirror_return_headers = mirror_return_headers
        # The role name used for mirror-based back-to-origin.
        self.mirror_role = mirror_role
        # 是否透传SNI
        self.mirror_sni = mirror_sni
        # Whether to store the user defined metadata in mirror-based back-to-origin.
        self.mirror_save_oss_meta = mirror_save_oss_meta
        # Used for determining the state of primary-secondary switching. The logic for primary-secondary switching is based on the error code returned by the origin server. If MirrorSwitchAllErrors is set to true, all status codes except the following are considered failures: 200, 206, 301, 302, 303, 307, 404. If it is set to false, only status codes in the 5xx range or timeouts are considered failures.
        self.mirror_switch_all_errors = mirror_switch_all_errors
        # The rules for setting tags when saving files during mirror-based back-to-origin.
        self.mirror_taggings = mirror_taggings
        # The tunnel ID for mirror-based back-to-origin.
        self.mirror_tunnel_id = mirror_tunnel_id
        # The origin URL for mirroring-based back-to-origin. This parameter takes effect only when the value of RedirectType is Mirror. The origin URL must start with \\*\\*http://** or **https://\\*\\* and end with a forward slash (/). OSS adds an object name to the end of the URL to generate a back-to-origin URL. For example, the name of the object to access is myobject. If MirrorURL is set to `http://example.com/`, the back-to-origin URL is `http://example.com/myobject`. If MirrorURL is set to `http://example.com/dir1/`, the back-to-origin URL is `http://example.com/dir1/myobject`.
        # 
        # >  This parameter must be specified if RedirectType is set to Mirror.
        # 
        # Valid values:
        # 
        # *   true
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        # *   false
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        self.mirror_url = mirror_url
        # The probe URL for mirror-based back-to-origin.
        self.mirror_urlprobe = mirror_urlprobe
        # The slave URL for mirror-based back-to-origin.
        self.mirror_urlslave = mirror_urlslave
        # Use LastModifiedTime of the file from origin.
        self.mirror_user_last_modified = mirror_user_last_modified
        # Whether to use role for mirror-based back-to-origin.
        self.mirror_using_role = mirror_using_role
        # Specifies whether to include parameters of the original request in the redirection request when the system runs the redirection rule or mirroring-based back-to-origin rule. For example, if the **PassQueryString** parameter is set to true, the `?a=b&c=d` parameter string is included in a request sent to OSS, and the redirection mode is 302, this parameter is added to the Location header. For example, if the request is `Location:example.com?a=b&c=d` and the redirection type is mirroring-based back-to-origin, the ?a=b\\&c=d parameter string is also included in the back-to-origin request. Valid values: true and false (default).
        self.pass_query_string = pass_query_string
        # The protocol used for redirection. This parameter takes effect only when RedirectType is set to External or AliCDN. For example, if you access an object named test, Protocol is set to https, and Hostname is set to `example.com`, the value of the Location header is `https://example.com/test`. Valid values: **http** and **https**.
        self.protocol = protocol
        # The redirection type. Valid values:
        # 
        # *   **Mirror**: mirroring-based back-to-origin.
        # *   **External**: external redirection. OSS returns an HTTP 3xx status code and returns an address for you to redirect to.
        # *   **AliCDN**: redirection based on Alibaba Cloud CDN. Compared with external redirection, OSS adds an additional header to the request. After Alibaba Cloud CDN identifies the header, Alibaba Cloud CDN redirects the access to the specified address and returns the obtained data instead of the HTTP 3xx status code that redirects the access to another address.
        # 
        # >  This parameter must be specified if Redirect is specified.
        self.redirect_type = redirect_type
        # The string that is used to replace the prefix of the object name during redirection. If the prefix of an object name is empty, the string precedes the object name.
        # 
        # >  You can specify only one of the ReplaceKeyWith and ReplaceKeyPrefixWith parameters in a rule. For example, if you access an object named abc/test.txt, KeyPrefixEquals is set to abc/, ReplaceKeyPrefixWith is set to def/, the value of the Location header is `http://example.com/def/test.txt`.
        self.replace_key_prefix_with = replace_key_prefix_with
        # The string that is used to replace the requested object name when the request is redirected. This parameter can be set to the ${key} variable, which indicates the object name in the request. For example, if ReplaceKeyWith is set to `prefix/${key}.suffix` and the object to access is test, the value of the Location header is `http://example.com/prefix/test.suffix`.
        self.replace_key_with = replace_key_with
        # Specify which status codes returned by the origin server should be passed through to the client along with the body. The value should be HTTP status codes such as 4xx, 5xx, etc., separated by commas (,), for example, 400,404. This setting takes effect only when RedirectType is set to Mirror. When OSS requests content from the origin server, if the origin server returns one of the status codes specified in this parameter, OSS will pass through the status code and body returned by the origin server to the client.
        # > If the 404 status code is specified in this parameter, the configured ErrorDocument will be ineffective.
        self.transparent_mirror_response_codes = transparent_mirror_response_codes

    def validate(self):
        if self.mirror_auth:
            self.mirror_auth.validate()
        if self.mirror_headers:
            self.mirror_headers.validate()
        if self.mirror_multi_alternates:
            self.mirror_multi_alternates.validate()
        if self.mirror_return_headers:
            self.mirror_return_headers.validate()
        if self.mirror_taggings:
            self.mirror_taggings.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.enable_replace_prefix is not None:
            result['EnableReplacePrefix'] = self.enable_replace_prefix

        if self.host_name is not None:
            result['HostName'] = self.host_name

        if self.http_redirect_code is not None:
            result['HttpRedirectCode'] = self.http_redirect_code

        if self.mirror_allow_get_image_info is not None:
            result['MirrorAllowGetImageInfo'] = self.mirror_allow_get_image_info

        if self.mirror_allow_head_object is not None:
            result['MirrorAllowHeadObject'] = self.mirror_allow_head_object

        if self.mirror_allow_video_snapshot is not None:
            result['MirrorAllowVideoSnapshot'] = self.mirror_allow_video_snapshot

        if self.mirror_async_status is not None:
            result['MirrorAsyncStatus'] = self.mirror_async_status

        if self.mirror_auth is not None:
            result['MirrorAuth'] = self.mirror_auth.to_map()

        if self.mirror_check_md_5 is not None:
            result['MirrorCheckMd5'] = self.mirror_check_md_5

        if self.mirror_dst_region is not None:
            result['MirrorDstRegion'] = self.mirror_dst_region

        if self.mirror_dst_slave_vpc_id is not None:
            result['MirrorDstSlaveVpcId'] = self.mirror_dst_slave_vpc_id

        if self.mirror_dst_vpc_id is not None:
            result['MirrorDstVpcId'] = self.mirror_dst_vpc_id

        if self.mirror_follow_redirect is not None:
            result['MirrorFollowRedirect'] = self.mirror_follow_redirect

        if self.mirror_headers is not None:
            result['MirrorHeaders'] = self.mirror_headers.to_map()

        if self.mirror_is_express_tunnel is not None:
            result['MirrorIsExpressTunnel'] = self.mirror_is_express_tunnel

        if self.mirror_multi_alternates is not None:
            result['MirrorMultiAlternates'] = self.mirror_multi_alternates.to_map()

        if self.mirror_pass_original_slashes is not None:
            result['MirrorPassOriginalSlashes'] = self.mirror_pass_original_slashes

        if self.mirror_pass_query_string is not None:
            result['MirrorPassQueryString'] = self.mirror_pass_query_string

        if self.mirror_proxy_pass is not None:
            result['MirrorProxyPass'] = self.mirror_proxy_pass

        if self.mirror_return_headers is not None:
            result['MirrorReturnHeaders'] = self.mirror_return_headers.to_map()

        if self.mirror_role is not None:
            result['MirrorRole'] = self.mirror_role

        if self.mirror_sni is not None:
            result['MirrorSNI'] = self.mirror_sni

        if self.mirror_save_oss_meta is not None:
            result['MirrorSaveOssMeta'] = self.mirror_save_oss_meta

        if self.mirror_switch_all_errors is not None:
            result['MirrorSwitchAllErrors'] = self.mirror_switch_all_errors

        if self.mirror_taggings is not None:
            result['MirrorTaggings'] = self.mirror_taggings.to_map()

        if self.mirror_tunnel_id is not None:
            result['MirrorTunnelId'] = self.mirror_tunnel_id

        if self.mirror_url is not None:
            result['MirrorURL'] = self.mirror_url

        if self.mirror_urlprobe is not None:
            result['MirrorURLProbe'] = self.mirror_urlprobe

        if self.mirror_urlslave is not None:
            result['MirrorURLSlave'] = self.mirror_urlslave

        if self.mirror_user_last_modified is not None:
            result['MirrorUserLastModified'] = self.mirror_user_last_modified

        if self.mirror_using_role is not None:
            result['MirrorUsingRole'] = self.mirror_using_role

        if self.pass_query_string is not None:
            result['PassQueryString'] = self.pass_query_string

        if self.protocol is not None:
            result['Protocol'] = self.protocol

        if self.redirect_type is not None:
            result['RedirectType'] = self.redirect_type

        if self.replace_key_prefix_with is not None:
            result['ReplaceKeyPrefixWith'] = self.replace_key_prefix_with

        if self.replace_key_with is not None:
            result['ReplaceKeyWith'] = self.replace_key_with

        if self.transparent_mirror_response_codes is not None:
            result['TransparentMirrorResponseCodes'] = self.transparent_mirror_response_codes

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EnableReplacePrefix') is not None:
            self.enable_replace_prefix = m.get('EnableReplacePrefix')

        if m.get('HostName') is not None:
            self.host_name = m.get('HostName')

        if m.get('HttpRedirectCode') is not None:
            self.http_redirect_code = m.get('HttpRedirectCode')

        if m.get('MirrorAllowGetImageInfo') is not None:
            self.mirror_allow_get_image_info = m.get('MirrorAllowGetImageInfo')

        if m.get('MirrorAllowHeadObject') is not None:
            self.mirror_allow_head_object = m.get('MirrorAllowHeadObject')

        if m.get('MirrorAllowVideoSnapshot') is not None:
            self.mirror_allow_video_snapshot = m.get('MirrorAllowVideoSnapshot')

        if m.get('MirrorAsyncStatus') is not None:
            self.mirror_async_status = m.get('MirrorAsyncStatus')

        if m.get('MirrorAuth') is not None:
            temp_model = main_models.RoutingRuleRedirectMirrorAuth()
            self.mirror_auth = temp_model.from_map(m.get('MirrorAuth'))

        if m.get('MirrorCheckMd5') is not None:
            self.mirror_check_md_5 = m.get('MirrorCheckMd5')

        if m.get('MirrorDstRegion') is not None:
            self.mirror_dst_region = m.get('MirrorDstRegion')

        if m.get('MirrorDstSlaveVpcId') is not None:
            self.mirror_dst_slave_vpc_id = m.get('MirrorDstSlaveVpcId')

        if m.get('MirrorDstVpcId') is not None:
            self.mirror_dst_vpc_id = m.get('MirrorDstVpcId')

        if m.get('MirrorFollowRedirect') is not None:
            self.mirror_follow_redirect = m.get('MirrorFollowRedirect')

        if m.get('MirrorHeaders') is not None:
            temp_model = main_models.RoutingRuleRedirectMirrorHeaders()
            self.mirror_headers = temp_model.from_map(m.get('MirrorHeaders'))

        if m.get('MirrorIsExpressTunnel') is not None:
            self.mirror_is_express_tunnel = m.get('MirrorIsExpressTunnel')

        if m.get('MirrorMultiAlternates') is not None:
            temp_model = main_models.RoutingRuleRedirectMirrorMultiAlternates()
            self.mirror_multi_alternates = temp_model.from_map(m.get('MirrorMultiAlternates'))

        if m.get('MirrorPassOriginalSlashes') is not None:
            self.mirror_pass_original_slashes = m.get('MirrorPassOriginalSlashes')

        if m.get('MirrorPassQueryString') is not None:
            self.mirror_pass_query_string = m.get('MirrorPassQueryString')

        if m.get('MirrorProxyPass') is not None:
            self.mirror_proxy_pass = m.get('MirrorProxyPass')

        if m.get('MirrorReturnHeaders') is not None:
            temp_model = main_models.RoutingRuleRedirectMirrorReturnHeaders()
            self.mirror_return_headers = temp_model.from_map(m.get('MirrorReturnHeaders'))

        if m.get('MirrorRole') is not None:
            self.mirror_role = m.get('MirrorRole')

        if m.get('MirrorSNI') is not None:
            self.mirror_sni = m.get('MirrorSNI')

        if m.get('MirrorSaveOssMeta') is not None:
            self.mirror_save_oss_meta = m.get('MirrorSaveOssMeta')

        if m.get('MirrorSwitchAllErrors') is not None:
            self.mirror_switch_all_errors = m.get('MirrorSwitchAllErrors')

        if m.get('MirrorTaggings') is not None:
            temp_model = main_models.RoutingRuleRedirectMirrorTaggings()
            self.mirror_taggings = temp_model.from_map(m.get('MirrorTaggings'))

        if m.get('MirrorTunnelId') is not None:
            self.mirror_tunnel_id = m.get('MirrorTunnelId')

        if m.get('MirrorURL') is not None:
            self.mirror_url = m.get('MirrorURL')

        if m.get('MirrorURLProbe') is not None:
            self.mirror_urlprobe = m.get('MirrorURLProbe')

        if m.get('MirrorURLSlave') is not None:
            self.mirror_urlslave = m.get('MirrorURLSlave')

        if m.get('MirrorUserLastModified') is not None:
            self.mirror_user_last_modified = m.get('MirrorUserLastModified')

        if m.get('MirrorUsingRole') is not None:
            self.mirror_using_role = m.get('MirrorUsingRole')

        if m.get('PassQueryString') is not None:
            self.pass_query_string = m.get('PassQueryString')

        if m.get('Protocol') is not None:
            self.protocol = m.get('Protocol')

        if m.get('RedirectType') is not None:
            self.redirect_type = m.get('RedirectType')

        if m.get('ReplaceKeyPrefixWith') is not None:
            self.replace_key_prefix_with = m.get('ReplaceKeyPrefixWith')

        if m.get('ReplaceKeyWith') is not None:
            self.replace_key_with = m.get('ReplaceKeyWith')

        if m.get('TransparentMirrorResponseCodes') is not None:
            self.transparent_mirror_response_codes = m.get('TransparentMirrorResponseCodes')

        return self

class RoutingRuleRedirectMirrorTaggings(DaraModel):
    def __init__(
        self,
        taggings: List[main_models.RoutingRuleRedirectMirrorTaggingsTaggings] = None,
    ):
        # The rule list for setting tags.
        self.taggings = taggings

    def validate(self):
        if self.taggings:
            for v1 in self.taggings:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Taggings'] = []
        if self.taggings is not None:
            for k1 in self.taggings:
                result['Taggings'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.taggings = []
        if m.get('Taggings') is not None:
            for k1 in m.get('Taggings'):
                temp_model = main_models.RoutingRuleRedirectMirrorTaggingsTaggings()
                self.taggings.append(temp_model.from_map(k1))

        return self

class RoutingRuleRedirectMirrorTaggingsTaggings(DaraModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        # The tag key.
        self.key = key
        # The rule for setting tag value for a specific tag key.
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.key is not None:
            result['Key'] = self.key

        if self.value is not None:
            result['Value'] = self.value

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('Value') is not None:
            self.value = m.get('Value')

        return self

class RoutingRuleRedirectMirrorReturnHeaders(DaraModel):
    def __init__(
        self,
        return_header: List[main_models.RoutingRuleRedirectMirrorReturnHeadersReturnHeader] = None,
    ):
        # The rule list for setting response headers in mirror-based back-to-origin.
        self.return_header = return_header

    def validate(self):
        if self.return_header:
            for v1 in self.return_header:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['ReturnHeader'] = []
        if self.return_header is not None:
            for k1 in self.return_header:
                result['ReturnHeader'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.return_header = []
        if m.get('ReturnHeader') is not None:
            for k1 in m.get('ReturnHeader'):
                temp_model = main_models.RoutingRuleRedirectMirrorReturnHeadersReturnHeader()
                self.return_header.append(temp_model.from_map(k1))

        return self

class RoutingRuleRedirectMirrorReturnHeadersReturnHeader(DaraModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        # The response header.
        self.key = key
        # The rule for setting response header value for a specific header.
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.key is not None:
            result['Key'] = self.key

        if self.value is not None:
            result['Value'] = self.value

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('Value') is not None:
            self.value = m.get('Value')

        return self

class RoutingRuleRedirectMirrorMultiAlternates(DaraModel):
    def __init__(
        self,
        mirror_multi_alternate: List[main_models.RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate] = None,
    ):
        # The configuration list for multiple origins.
        self.mirror_multi_alternate = mirror_multi_alternate

    def validate(self):
        if self.mirror_multi_alternate:
            for v1 in self.mirror_multi_alternate:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['MirrorMultiAlternate'] = []
        if self.mirror_multi_alternate is not None:
            for k1 in self.mirror_multi_alternate:
                result['MirrorMultiAlternate'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.mirror_multi_alternate = []
        if m.get('MirrorMultiAlternate') is not None:
            for k1 in m.get('MirrorMultiAlternate'):
                temp_model = main_models.RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate()
                self.mirror_multi_alternate.append(temp_model.from_map(k1))

        return self

class RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate(DaraModel):
    def __init__(
        self,
        mirror_multi_alternate_dst_region: str = None,
        mirror_multi_alternate_number: int = None,
        mirror_multi_alternate_url: str = None,
        mirror_multi_alternate_vpc_id: str = None,
    ):
        # The region for a specific origin.
        self.mirror_multi_alternate_dst_region = mirror_multi_alternate_dst_region
        # The distinct number of a specific origin.
        self.mirror_multi_alternate_number = mirror_multi_alternate_number
        # The URL for a specific origin.
        self.mirror_multi_alternate_url = mirror_multi_alternate_url
        # The VPC ID for a specific origin.
        self.mirror_multi_alternate_vpc_id = mirror_multi_alternate_vpc_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.mirror_multi_alternate_dst_region is not None:
            result['MirrorMultiAlternateDstRegion'] = self.mirror_multi_alternate_dst_region

        if self.mirror_multi_alternate_number is not None:
            result['MirrorMultiAlternateNumber'] = self.mirror_multi_alternate_number

        if self.mirror_multi_alternate_url is not None:
            result['MirrorMultiAlternateURL'] = self.mirror_multi_alternate_url

        if self.mirror_multi_alternate_vpc_id is not None:
            result['MirrorMultiAlternateVpcId'] = self.mirror_multi_alternate_vpc_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MirrorMultiAlternateDstRegion') is not None:
            self.mirror_multi_alternate_dst_region = m.get('MirrorMultiAlternateDstRegion')

        if m.get('MirrorMultiAlternateNumber') is not None:
            self.mirror_multi_alternate_number = m.get('MirrorMultiAlternateNumber')

        if m.get('MirrorMultiAlternateURL') is not None:
            self.mirror_multi_alternate_url = m.get('MirrorMultiAlternateURL')

        if m.get('MirrorMultiAlternateVpcId') is not None:
            self.mirror_multi_alternate_vpc_id = m.get('MirrorMultiAlternateVpcId')

        return self

class RoutingRuleRedirectMirrorHeaders(DaraModel):
    def __init__(
        self,
        pass_: List[str] = None,
        pass_all: bool = None,
        remove: List[str] = None,
        set: List[main_models.RoutingRuleRedirectMirrorHeadersSet] = None,
    ):
        # The headers to pass through to the origin. This parameter takes effect only when the value of RedirectType is Mirror. Each specified header can be up to 1,024 bytes in length and can contain only letters, digits, and hyphens (-). You can specify up to 10 headers.
        self.pass_ = pass_
        # Specifies whether to pass through all request headers other than the following headers to the origin. This parameter takes effect only when the value of RedirectType is Mirror.
        # 
        # *   Headers such as content-length, authorization2, authorization, range, and date
        # *   Headers that start with oss-, x-oss-, and x-drs-
        # 
        # Default value: false.
        # 
        # Valid values:
        # 
        # *   true
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        # *   false
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        self.pass_all = pass_all
        # The headers that are not allowed to pass through to the origin. This parameter takes effect only when the value of RedirectType is Mirror. Each header can be up to 1,024 bytes in length and can contain only letters, digits, and hyphens (-). You can specify up to 10 headers. This parameter is used together with PassAll.
        self.remove = remove
        # The headers that are sent to the origin. The specified headers are configured in the data returned by the origin regardless of whether the headers are contained in the request. This parameter takes effect only when the value of RedirectType is Mirror. You can specify up to 10 headers.
        self.set = set

    def validate(self):
        if self.set:
            for v1 in self.set:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.pass_ is not None:
            result['Pass'] = self.pass_

        if self.pass_all is not None:
            result['PassAll'] = self.pass_all

        if self.remove is not None:
            result['Remove'] = self.remove

        result['Set'] = []
        if self.set is not None:
            for k1 in self.set:
                result['Set'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Pass') is not None:
            self.pass_ = m.get('Pass')

        if m.get('PassAll') is not None:
            self.pass_all = m.get('PassAll')

        if m.get('Remove') is not None:
            self.remove = m.get('Remove')

        self.set = []
        if m.get('Set') is not None:
            for k1 in m.get('Set'):
                temp_model = main_models.RoutingRuleRedirectMirrorHeadersSet()
                self.set.append(temp_model.from_map(k1))

        return self

class RoutingRuleRedirectMirrorHeadersSet(DaraModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        # The key of the header. The key can be up to 1,024 bytes in length and can contain only letters, digits, and hyphens (-). This parameter takes effect only when the value of RedirectType is Mirror.
        # 
        # >  This parameter must be specified if Set is specified.
        self.key = key
        # The value of the header. The value can be up to 1,024 bytes in length and cannot contain `\\r\\n`. This parameter takes effect only when the value of RedirectType is Mirror.
        # 
        # >  This parameter must be specified if Set is specified.
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.key is not None:
            result['Key'] = self.key

        if self.value is not None:
            result['Value'] = self.value

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('Value') is not None:
            self.value = m.get('Value')

        return self

class RoutingRuleRedirectMirrorAuth(DaraModel):
    def __init__(
        self,
        access_key_id: str = None,
        access_key_secret: str = None,
        auth_type: str = None,
        region: str = None,
    ):
        # The access key id for signature.
        self.access_key_id = access_key_id
        # The access key secret for signature.
        self.access_key_secret = access_key_secret
        # The authentication type.
        self.auth_type = auth_type
        # The sign region for signature.
        self.region = region

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_key_id is not None:
            result['AccessKeyId'] = self.access_key_id

        if self.access_key_secret is not None:
            result['AccessKeySecret'] = self.access_key_secret

        if self.auth_type is not None:
            result['AuthType'] = self.auth_type

        if self.region is not None:
            result['Region'] = self.region

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessKeyId') is not None:
            self.access_key_id = m.get('AccessKeyId')

        if m.get('AccessKeySecret') is not None:
            self.access_key_secret = m.get('AccessKeySecret')

        if m.get('AuthType') is not None:
            self.auth_type = m.get('AuthType')

        if m.get('Region') is not None:
            self.region = m.get('Region')

        return self

