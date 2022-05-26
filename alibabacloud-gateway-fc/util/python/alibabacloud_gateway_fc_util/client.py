# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_credentials.client import Client as CredentialClient
from requests import Request
from requests import Response


class Client:
    def __init__(self):
        pass

    @staticmethod
    def invoke_httptrigger(
        credential: CredentialClient,
        url: str,
        method: str,
        body: bytes,
        headers: Dict[str, Any],
    ) -> Response:
        raise Exception('Un-implemented')

    @staticmethod
    def invoke_anonymous_httptrigger(
        url: str,
        method: str,
        body: bytes,
        headers: Dict[str, Any],
    ) -> Response:
        raise Exception('Un-implemented')

    @staticmethod
    def send_httprequest_with_authorization(
        credential: CredentialClient,
        req: Request,
    ) -> Response:
        raise Exception('Un-implemented')

    @staticmethod
    def send_httprequest(
        req: Request,
    ) -> Response:
        raise Exception('Un-implemented')

    @staticmethod
    def sign_request(
        credential: CredentialClient,
        req: Request,
    ) -> Request:
        raise Exception('Un-implemented')

    @staticmethod
    def sign_request_with_content_md5(
        credential: CredentialClient,
        req: Request,
        content_md5: str,
    ) -> Request:
        raise Exception('Un-implemented')

    @staticmethod
    def build_httprequest(
        url: str,
        method: str,
        body: bytes,
        headers: Dict[str, Any],
    ) -> Request:
        raise Exception('Un-implemented')
