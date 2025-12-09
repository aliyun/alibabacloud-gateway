# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
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
        self.enable_replace_prefix = enable_replace_prefix
        self.host_name = host_name
        self.http_redirect_code = http_redirect_code
        self.mirror_allow_get_image_info = mirror_allow_get_image_info
        self.mirror_allow_head_object = mirror_allow_head_object
        self.mirror_allow_video_snapshot = mirror_allow_video_snapshot
        self.mirror_async_status = mirror_async_status
        self.mirror_auth = mirror_auth
        self.mirror_check_md_5 = mirror_check_md_5
        self.mirror_dst_region = mirror_dst_region
        self.mirror_dst_slave_vpc_id = mirror_dst_slave_vpc_id
        self.mirror_dst_vpc_id = mirror_dst_vpc_id
        self.mirror_follow_redirect = mirror_follow_redirect
        self.mirror_headers = mirror_headers
        self.mirror_is_express_tunnel = mirror_is_express_tunnel
        self.mirror_multi_alternates = mirror_multi_alternates
        self.mirror_pass_original_slashes = mirror_pass_original_slashes
        self.mirror_pass_query_string = mirror_pass_query_string
        self.mirror_proxy_pass = mirror_proxy_pass
        self.mirror_return_headers = mirror_return_headers
        self.mirror_role = mirror_role
        self.mirror_sni = mirror_sni
        self.mirror_save_oss_meta = mirror_save_oss_meta
        self.mirror_switch_all_errors = mirror_switch_all_errors
        self.mirror_taggings = mirror_taggings
        self.mirror_tunnel_id = mirror_tunnel_id
        self.mirror_url = mirror_url
        self.mirror_urlprobe = mirror_urlprobe
        self.mirror_urlslave = mirror_urlslave
        self.mirror_user_last_modified = mirror_user_last_modified
        self.mirror_using_role = mirror_using_role
        self.pass_query_string = pass_query_string
        self.protocol = protocol
        self.redirect_type = redirect_type
        self.replace_key_prefix_with = replace_key_prefix_with
        self.replace_key_with = replace_key_with
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
        self.key = key
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
        self.key = key
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
        self.mirror_multi_alternate_dst_region = mirror_multi_alternate_dst_region
        self.mirror_multi_alternate_number = mirror_multi_alternate_number
        self.mirror_multi_alternate_url = mirror_multi_alternate_url
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
        self.pass_ = pass_
        self.pass_all = pass_all
        self.remove = remove
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
        self.key = key
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
        self.access_key_id = access_key_id
        self.access_key_secret = access_key_secret
        self.auth_type = auth_type
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

