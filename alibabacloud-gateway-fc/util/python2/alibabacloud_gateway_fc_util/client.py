# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_credentials.client import Client as CredentialClient


class Client(object):
    _credential = None  # type: CredentialClient

    def __init__(self, cred):
        self._credential = cred

    def invoke_httptrigger(self, url, method, body, headers):
        raise Exception('Un-implemented')

    def invoke_anonymous_httptrigger(self, url, method, body, headers):
        raise Exception('Un-implemented')

    def send_httprequest_with_authorization(self, req):
        raise Exception('Un-implemented')

    def send_httprequest(self, req):
        raise Exception('Un-implemented')

    def sign_request(self, req):
        raise Exception('Un-implemented')

    def sign_request_with_content_md5(self, req, content_md5):
        raise Exception('Un-implemented')

    def build_httprequest(self, url, method, body, headers):
        raise Exception('Un-implemented')
