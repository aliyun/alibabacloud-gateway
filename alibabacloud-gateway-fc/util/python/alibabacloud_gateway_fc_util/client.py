# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
import base64
import hashlib
import hmac
import logging
import urllib.parse
from typing import (
    Dict,
    List,
    Any,
)
import requests
from requests import (
    Response,
    Request,
)

from alibabacloud_credentials.client import Client as CredentialClient

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
    def __init__(self):
        pass

    @staticmethod
    def invoke_httptrigger(
            credential: CredentialClient,
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Response:
        req = Client.build_httprequest(url, method, body, headers)
        return Client.send_httprequest_with_authorization(credential, req)

    @staticmethod
    def invoke_anonymous_httptrigger(
            url: str,
            method: str,
            body: bytes,
            headers: Dict[str, str],
    ) -> Response:
        req = Client.build_httprequest(url, method, body, headers)
        return Client.send_httprequest(req)

    @staticmethod
    def send_httprequest_with_authorization(
            credential: CredentialClient,
            req: Request,
    ) -> Response:
        signedRequest = Client.sign_request(credential, req)
        return Client.send_httprequest(signedRequest)

    @staticmethod
    def send_httprequest(
            req: Request,
    ) -> Response:
        with requests.Session() as s:
            p=s.prepare_request(req)
            return s.send(p)

    @staticmethod
    def sign_request(
            credential: CredentialClient,
            req: Request,
    ) -> Request:
        # FIXME Request.data is too flexible, and we can't calculate the md5 value of it properly.
        return Client.sign_request_with_content_md5(credential, req, '')

    @staticmethod
    def sign_request_with_content_md5(
            credential: CredentialClient,
            req: Request,
            content_md5: str,
    ) -> Request:
        security_token = credential.get_security_token()
        if security_token:
            req.headers[HTTPHeaderSecurityToken] = security_token
        if content_md5:
            req.headers[HTTPHeaderContentMD5] = content_md5
        result = urllib.parse.urlparse(req.url)
        auth_string = Client.auth_string(credential,
                                         req.method, urllib.parse.unquote(result.path),
                                         req.headers, urllib.parse.unquote(result.query))
        req.headers[HTTPHeaderAuthorization] = auth_string
        return req

    @staticmethod
    def auth_string(credential: CredentialClient,
                    method: str,
                    unescaped_path: str,
                    headers: Dict[str, str],
                    unescaped_queries: str):
        """
        Sign the request. See the spec for reference.
        https://help.aliyun.com/document_detail/52877.html
        :param unescaped_queries: query of http request
        :param credential: alibabacloud_credentials.client
        :param method: method of the http request.
        :param headers: headers of the http request.
        :param unescaped_path: unescaped path without queries of the http request.
        :return: the signature string.
        """
        content_md5 = headers.get('content-md5', '')
        content_type = headers.get('content-type', '')
        date = headers.get('date', '')
        canonical_headers = Client.build_canonical_headers(headers)
        canonical_resource = unescaped_path
        access_key_id = credential.get_access_key_id()
        access_key_secret = credential.get_access_key_secret()

        if isinstance(unescaped_queries, Dict):
            canonical_resource = Client.get_sign_resource(unescaped_path, unescaped_queries)
        string_to_sign = '\n'.join(
            [method.upper(), content_md5, content_type, date, canonical_headers + canonical_resource])
        logging.debug('string to sign:{0}'.format(string_to_sign))
        h = hmac.new(access_key_secret.encode('utf-8'), string_to_sign.encode('utf-8'), hashlib.sha256)
        signature = 'FC ' + access_key_id + ':' + base64.b64encode(h.digest()).decode('utf-8')
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

    @staticmethod
    def build_httprequest(
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
