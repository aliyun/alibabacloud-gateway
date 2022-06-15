# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
import base64
import email.utils
import hashlib
import hmac
import logging
import urllib.parse
from typing import (
    Dict,
)

import requests
from alibabacloud_credentials.client import Client as CredentialClient
from requests import (
    Response,
    Request,
)

# HTTPHeaderContentMD5 key in request headers
HTTPHeaderContentMD5 = "content-md5"
# HTTPHeaderPrefix key prefix in request headers
HTTPHeaderPrefix = "x-fc-"
# HTTPHeaderDate key in request headers
HTTPHeaderDate = "date"
# AuthQueryKeyExpires key in request headers
AuthQueryKeyExpires = "x-fc-expires"
# HTTPHeaderContentType key in request headers
HTTPHeaderContentType = "content-type"
# HTTPHeaderAuthorization key in request headers
HTTPHeaderAuthorization = "authorization"
# HTTPHeaderSecurityToken key in request headers
HTTPHeaderSecurityToken = "x-fc-security-token"


class Client:
    _credential: CredentialClient = None

    def __init__(
            self,
            cred: CredentialClient,
    ):
        self._credential = cred

    def invoke_httptrigger(
            self,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Response:
        req = self.build_httprequest(url, method, body, headers)
        return self.send_httprequest_with_authorization(req)

    async def invoke_httptrigger_async(
            self,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Response:
        req = await self.build_httprequest_async(url, method, body, headers)
        return await self.send_httprequest_with_authorization_async(req)

    def invoke_anonymous_httptrigger(
            self,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Response:
        req = self.build_httprequest(url, method, body, headers)
        return self.send_httprequest(req)

    async def invoke_anonymous_httptrigger_async(
            self,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Response:
        req = await self.build_httprequest_async(url, method, body, headers)
        return await self.send_httprequest_async(req)

    def send_httprequest_with_authorization(
            self,
            req: Request,
    ) -> Response:
        signedRequest = self.sign_request(req)
        return self.send_httprequest(signedRequest)

    async def send_httprequest_with_authorization_async(
            self,
            req: Request,
    ) -> Response:
        signedRequest = await self.sign_request_async(req)
        return await self.send_httprequest_async(signedRequest)

    def send_httprequest(
            self,
            req: Request,
    ) -> Response:
        with requests.Session() as s:
            p = s.prepare_request(req)
            return s.send(p)

    async def send_httprequest_async(
            self,
            req: Request,
    ) -> Response:
        with requests.Session() as s:
            p = s.prepare_request(req)
            return s.send(p)

    def sign_request(
            self,
            req: Request,
    ) -> Request:
        # FIXME Request.data is too flexible, and we can't calculate the md5 value of it properly.
        return self.sign_request_with_content_md5(req, '')

    async def sign_request_async(
            self,
            req: Request,
    ) -> Request:
        # FIXME Request.data is too flexible, and we can't calculate the md5 value of it properly.
        return await self.sign_request_with_content_md5_async(req, '')

    def sign_request_with_content_md5(
            self,
            req: Request,
            content_md5: str,
    ) -> Request:
        security_token = self._credential.get_security_token()
        if security_token:
            req.headers[HTTPHeaderSecurityToken] = security_token
        if content_md5:
            req.headers[HTTPHeaderContentMD5] = content_md5
        req.headers[HTTPHeaderDate] = email.utils.formatdate(usegmt=True)
        result = urllib.parse.urlparse(req.url)
        auth_string = self.auth_string(req.method, urllib.parse.unquote(result.path),
                                       req.headers, urllib.parse.unquote(result.query))
        req.headers[HTTPHeaderAuthorization] = auth_string
        return req

    async def sign_request_with_content_md5_async(
            self,
            req: Request,
            content_md5: str,
    ) -> Request:
        security_token = self._credential.get_security_token()
        if security_token:
            req.headers[HTTPHeaderSecurityToken] = security_token
        if content_md5:
            req.headers[HTTPHeaderContentMD5] = content_md5
        req.headers[HTTPHeaderDate] = email.utils.formatdate(usegmt=True)
        result = urllib.parse.urlparse(req.url)
        auth_string = self.auth_string(req.method, urllib.parse.unquote(result.path),
                                       req.headers, urllib.parse.unquote(result.query))
        req.headers[HTTPHeaderAuthorization] = auth_string
        return req

    def auth_string(self,
                    method: str,
                    unescaped_path: str,
                    headers: Dict[str, str],
                    unescaped_queries: str):
        """
        Sign the request. See the spec for reference.
        https://help.aliyun.com/document_detail/52877.html
        :param unescaped_queries: query of http request
        :param method: method of the http request.
        :param headers: headers of the http request.
        :param unescaped_path: unescaped path without queries of the http request.
        :return: the signature string.
        """
        content_md5 = headers.get('content-md5', '')
        content_type = headers.get('content-type', '')
        date = headers.get('date', '')
        canonical_headers = Client.build_canonical_headers(headers)
        if not unescaped_path:
            unescaped_path = '/'
        canonical_resource = unescaped_path + '\n'
        access_key_id = self._credential.get_access_key_id()
        access_key_secret = self._credential.get_access_key_secret()

        if unescaped_queries:
            canonical_resource += '\n'.join(unescaped_queries.split('&'))
        string_to_sign = '\n'.join(
            [method.upper(), content_md5, content_type, date, canonical_headers + canonical_resource])
        logging.debug('string to sign:{0}'.format(string_to_sign))
        h = hmac.new(access_key_secret.encode('utf-8'),
                     string_to_sign.encode('utf-8'), hashlib.sha256)
        signature = 'FC ' + access_key_id + ':' + \
                    base64.b64encode(h.digest()).decode('utf-8')
        return signature

    @staticmethod
    def build_canonical_headers(headers):
        """
        :param headers: :class:`Request` object
        :return: Canonicalized header string.
        :rtype: String
        """
        canonical_headers = []
        for k, v in headers.items():
            lower_key = k.lower()
            if lower_key.startswith('x-fc-'):
                canonical_headers.append((lower_key, v))
        canonical_headers.sort(key=lambda x: x[0])
        if canonical_headers:
            return '\n'.join(k + ':' + v for k, v in canonical_headers) + '\n'
        else:
            return ''

    @staticmethod
    def get_sign_resource(unescaped_path, unescaped_queries):
        if not isinstance(unescaped_queries, Dict):
            raise TypeError("`Dict` type required for queries")
        params = []
        for key, values in unescaped_queries.items():
            if isinstance(values, str):
                params.append('%s=%s' % (key, values))
                continue

            if len(values) > 0:
                for value in values:
                    params.append('%s=%s' % (key, value))
            else:
                params.append('%s' % key)
        params.sort()
        resource = unescaped_path + '\n' + '\n'.join(params)
        return resource

    def build_httprequest(
            self,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Request:
        return Request(
            url=url,
            method=method,
            headers=headers,
            data=body,
        )

    async def build_httprequest_async(
            self,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Request:
        return Request(
            url=url,
            method=method,
            headers=headers,
            data=body,
        )
