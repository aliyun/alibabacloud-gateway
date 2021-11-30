// This file is auto-generated, don't edit it
import Credential from '@alicloud/credentials';
import { Readable } from 'stream';
import * as $tea from '@alicloud/tea-typescript';

export class InterceptorContext extends $tea.Model {
  request: InterceptorContextRequest;
  configuration: InterceptorContextConfiguration;
  response: InterceptorContextResponse;
  static names(): { [key: string]: string } {
    return {
      request: 'request',
      configuration: 'configuration',
      response: 'response',
    };
  }

  static types(): { [key: string]: any } {
    return {
      request: InterceptorContextRequest,
      configuration: InterceptorContextConfiguration,
      response: InterceptorContextResponse,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class AttributeMap extends $tea.Model {
  attributes: { [key: string]: any };
  key: { [key: string]: string };
  static names(): { [key: string]: string } {
    return {
      attributes: 'attributes',
      key: 'key',
    };
  }

  static types(): { [key: string]: any } {
    return {
      attributes: { 'type': 'map', 'keyType': 'string', 'valueType': 'any' },
      key: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class InterceptorContextRequest extends $tea.Model {
  headers?: { [key: string]: string };
  query?: { [key: string]: string };
  body?: any;
  stream?: Readable;
  hostMap?: { [key: string]: string };
  pathname: string;
  productId: string;
  action: string;
  version: string;
  protocol: string;
  method: string;
  authType: string;
  bodyType: string;
  reqBodyType: string;
  style?: string;
  credential: Credential;
  signatureVersion?: string;
  signatureAlgorithm?: string;
  userAgent: string;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      query: 'query',
      body: 'body',
      stream: 'stream',
      hostMap: 'hostMap',
      pathname: 'pathname',
      productId: 'productId',
      action: 'action',
      version: 'version',
      protocol: 'protocol',
      method: 'method',
      authType: 'authType',
      bodyType: 'bodyType',
      reqBodyType: 'reqBodyType',
      style: 'style',
      credential: 'credential',
      signatureVersion: 'signatureVersion',
      signatureAlgorithm: 'signatureAlgorithm',
      userAgent: 'userAgent',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      query: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: 'any',
      stream: 'Readable',
      hostMap: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      pathname: 'string',
      productId: 'string',
      action: 'string',
      version: 'string',
      protocol: 'string',
      method: 'string',
      authType: 'string',
      bodyType: 'string',
      reqBodyType: 'string',
      style: 'string',
      credential: Credential,
      signatureVersion: 'string',
      signatureAlgorithm: 'string',
      userAgent: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class InterceptorContextConfiguration extends $tea.Model {
  regionId: string;
  endpoint?: string;
  endpointRule?: string;
  endpointMap?: { [key: string]: string };
  endpointType?: string;
  network?: string;
  suffix?: string;
  static names(): { [key: string]: string } {
    return {
      regionId: 'regionId',
      endpoint: 'endpoint',
      endpointRule: 'endpointRule',
      endpointMap: 'endpointMap',
      endpointType: 'endpointType',
      network: 'network',
      suffix: 'suffix',
    };
  }

  static types(): { [key: string]: any } {
    return {
      regionId: 'string',
      endpoint: 'string',
      endpointRule: 'string',
      endpointMap: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      endpointType: 'string',
      network: 'string',
      suffix: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class InterceptorContextResponse extends $tea.Model {
  statusCode?: number;
  headers?: { [key: string]: string };
  body?: Readable;
  deserializedBody?: any;
  static names(): { [key: string]: string } {
    return {
      statusCode: 'statusCode',
      headers: 'headers',
      body: 'body',
      deserializedBody: 'deserializedBody',
    };
  }

  static types(): { [key: string]: any } {
    return {
      statusCode: 'number',
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: 'Readable',
      deserializedBody: 'any',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}


export default abstract class Client {

  constructor() {
  }

  abstract modifyConfiguration(context: InterceptorContext, attributeMap: AttributeMap): Promise<void> 

  abstract modifyRequest(context: InterceptorContext, attributeMap: AttributeMap): Promise<void> 

  abstract modifyResponse(context: InterceptorContext, attributeMap: AttributeMap): Promise<void> 

}
