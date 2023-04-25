# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.exceptions import TeaException
from Tea.core import TeaCore
from typing import Any

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient


class Client(SPIClient):
    def __init__(self):
        super().__init__()

    def modify_configuration(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        pass

    async def modify_configuration_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        pass

    def modify_request(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'user-agent': request.user_agent,
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.is_unset(request.body):
            json_obj = UtilClient.to_jsonstring(request.body)
            request.stream = json_obj
            request.headers['content-type'] = 'application/json; charset=utf-8'

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'user-agent': request.user_agent,
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.is_unset(request.body):
            json_obj = UtilClient.to_jsonstring(request.body)
            request.stream = json_obj
            request.headers['content-type'] = 'application/json; charset=utf-8'

    def modify_response(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = UtilClient.read_as_json(response.body)
            err = UtilClient.assert_as_map(_res)
            err['statusCode'] = response.status_code
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"code: {response.status_code}, {self.default_any(err.get('Message'), err.get('message'))} request id: {self.default_any(err.get('RequestId'), err.get('requestid'))}",
                'data': err,
                'description': f"{self.default_any(err.get('Description'), err.get('description'))}",
                'accessDeniedDetail': self.default_any(err.get('AccessDeniedDetail'), err.get('accessdenieddetail'))
            })
        if UtilClient.equal_number(response.status_code, 204):
            UtilClient.read_as_string(response.body)
        elif UtilClient.equal_string(request.body_type, 'binary'):
            response.deserialized_body = response.body
        elif UtilClient.equal_string(request.body_type, 'byte'):
            byt = UtilClient.read_as_bytes(response.body)
            response.deserialized_body = byt
        elif UtilClient.equal_string(request.body_type, 'string'):
            str = UtilClient.read_as_string(response.body)
            response.deserialized_body = str
        elif UtilClient.equal_string(request.body_type, 'json'):
            obj = UtilClient.read_as_json(response.body)
            res = UtilClient.assert_as_map(obj)
            response.deserialized_body = res
        elif UtilClient.equal_string(request.body_type, 'array'):
            arr = UtilClient.read_as_json(response.body)
            response.deserialized_body = arr
        else:
            response.deserialized_body = UtilClient.read_as_string(response.body)

    async def modify_response_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = await UtilClient.read_as_json_async(response.body)
            err = UtilClient.assert_as_map(_res)
            err['statusCode'] = response.status_code
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"code: {response.status_code}, {self.default_any(err.get('Message'), err.get('message'))} request id: {self.default_any(err.get('RequestId'), err.get('requestid'))}",
                'data': err,
                'description': f"{self.default_any(err.get('Description'), err.get('description'))}",
                'accessDeniedDetail': self.default_any(err.get('AccessDeniedDetail'), err.get('accessdenieddetail'))
            })
        if UtilClient.equal_number(response.status_code, 204):
            await UtilClient.read_as_string_async(response.body)
        elif UtilClient.equal_string(request.body_type, 'binary'):
            response.deserialized_body = response.body
        elif UtilClient.equal_string(request.body_type, 'byte'):
            byt = await UtilClient.read_as_bytes_async(response.body)
            response.deserialized_body = byt
        elif UtilClient.equal_string(request.body_type, 'string'):
            str = await UtilClient.read_as_string_async(response.body)
            response.deserialized_body = str
        elif UtilClient.equal_string(request.body_type, 'json'):
            obj = await UtilClient.read_as_json_async(response.body)
            res = UtilClient.assert_as_map(obj)
            response.deserialized_body = res
        elif UtilClient.equal_string(request.body_type, 'array'):
            arr = await UtilClient.read_as_json_async(response.body)
            response.deserialized_body = arr
        else:
            response.deserialized_body = await UtilClient.read_as_string_async(response.body)

    def default_any(
        self,
        input_value: Any,
        default_value: Any,
    ) -> Any:
        if UtilClient.is_unset(input_value):
            return default_value
        return input_value
