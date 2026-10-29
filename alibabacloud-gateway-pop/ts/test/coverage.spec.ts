import assert from 'assert';
import 'mocha';
import { Readable } from 'stream';
import Client from '../src/client';
import * as $SPI from '@alicloud/gateway-spi';

function readable(text: string): Readable {
  const stream = new Readable();
  stream.push(text);
  stream.push(null);
  return stream;
}

function boom(): Readable {
  return new Readable({
    read() {
      process.nextTick(() => this.emit('error', new Error('read failed')));
    }
  });
}

function credential(model: any): any {
  return {
    getCredential: async () => model
  };
}

function accessKey(): any {
  return credential({
    type: 'access_key',
    accessKeyId: 'ak',
    accessKeySecret: 'secret',
    providerName: 'static_ak'
  });
}

function request(cred: any): $SPI.InterceptorContextRequest {
  return new $SPI.InterceptorContextRequest({
    headers: {},
    query: {},
    pathname: '/',
    productId: 'ecs',
    action: 'DescribeRegions',
    version: '2014-05-26',
    protocol: 'HTTPS',
    method: 'POST',
    authType: 'AK',
    bodyType: 'json',
    reqBodyType: 'json',
    credential: cred,
    userAgent: 'test-agent'
  });
}

function configuration(endpoint?: string): $SPI.InterceptorContextConfiguration {
  return new $SPI.InterceptorContextConfiguration({
    regionId: 'cn-hangzhou',
    endpoint: endpoint
  });
}

function context(req: $SPI.InterceptorContextRequest, config: $SPI.InterceptorContextConfiguration): $SPI.InterceptorContext {
  return new $SPI.InterceptorContext({
    request: req,
    configuration: config,
    response: new $SPI.InterceptorContextResponse({})
  });
}

describe('coverage', function () {
  it('selects endpoints and applies configuration', async function () {
    const client = new Client();
    assert.strictEqual(client.getEndpoint('ecs', 'cn-hangzhou', 'regional', '', '', null, 'ecs.example.com'), 'ecs.example.com');
    assert.strictEqual(client.getEndpoint('ecs', 'cn-hangzhou', 'regional', '', '', { 'cn-hangzhou': 'ecs.mapped.example.com' }, ''), 'ecs.mapped.example.com');
    assert.strictEqual(client.getEndpoint('ecs', 'cn-hangzhou', 'regional', '', '', null, ''), 'ecs.cn-hangzhou.aliyuncs.com');
    assert.strictEqual(client.getEndpoint('ecs', '', 'central', 'public', '', null, ''), 'ecs.aliyuncs.com');
    assert.throws(() => client.getEndpoint('ecs', '', 'regional', '', '', null, ''));

    const built = new $SPI.InterceptorContextConfiguration({
      regionId: 'cn-hangzhou',
      endpointRule: 'regional',
      network: '',
      suffix: ''
    });
    const configured = context(request(accessKey()), built);
    await client.modifyConfiguration(configured, new $SPI.AttributeMap({
      key: {
        signatureTypePrefix: 'ACS4-',
        signPrefix: 'aliyun_v4',
        endpointSuffix: 'aliyuncs.com'
      }
    }));
    assert.strictEqual(configured.configuration.endpoint, 'ecs.cn-hangzhou.aliyuncs.com');

    const kept = context(request(accessKey()), configuration('already.example.com'));
    await client.modifyConfiguration(kept, new $SPI.AttributeMap({}));
    assert.strictEqual(kept.configuration.endpoint, 'already.example.com');
  });

  it('builds canonical requests and signatures', function () {
    const client = new Client();
    assert.strictEqual(client.buildCanonicalizedResource(null), '');
    const resource = client.buildCanonicalizedResource({ b: 'x y', a: 'm*n', empty: '' });
    assert.ok(resource.indexOf('a=m%2An') >= 0);
    assert.ok(resource.indexOf('b=x%20y') >= 0);
    assert.ok(resource.indexOf('empty=') >= 0);
    const canonical = client.buildCanonicalizedHeaders({
      Host: '  example.com  ',
      'X-Acs-Action': 'One',
      'x-acs-action': 'Two'
    });
    assert.ok(canonical.indexOf('host:example.com\n') >= 0);
    assert.strictEqual(client.defaultAny(null, 'fallback'), 'fallback');
    assert.strictEqual(client.defaultAny('kept', 'fallback'), 'kept');
    const signingkey = client.getSigningkey(client._sha256, 'secret', 'ecs', 'cn-hangzhou', '20240101');
    assert.ok(signingkey.length > 0);
    const unknown = client.getSigningkey('OTHER', 'secret', 'ecs', 'cn-hangzhou', '20240101');
    assert.strictEqual(unknown.length, 0);
    const headers = { host: 'example.com', 'x-acs-date': '2024-01-01T00:00:00Z' };
    const signature = client.getSignature('', 'GET', { q: 'a b' }, headers, client._sha256, 'payload', signingkey);
    assert.strictEqual(signature.length, 64);
    const authorization = client.getAuthorization('/', 'GET', null, headers, client._sha256, 'payload', 'ak', signingkey, 'ecs', 'cn-hangzhou', '20240101');
    assert.ok(authorization.indexOf('ACS4-HMAC-SHA256 Credential=ak/20240101/cn-hangzhou/ecs/aliyun_v4_request,SignedHeaders=') === 0);
    assert.ok(authorization.indexOf('host;x-acs-date') >= 0);
  });

  it('signs request bodies and credentials', async function () {
    const client = new Client();
    const json = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    json.request.body = { name: 'a b' };
    json.request.query = { InstanceId: 'i-1' };
    await client.modifyRequest(json, new $SPI.AttributeMap({}));
    assert.strictEqual(json.request.headers['content-type'], 'application/json; charset=utf-8');
    assert.strictEqual(json.request.headers['x-acs-action'], 'DescribeRegions');
    assert.ok(json.request.headers['Authorization'].indexOf('ACS4-HMAC-SHA256 Credential=ak/') === 0);
    assert.ok(json.request.headers['x-acs-content-sha256'].length > 0);
    assert.strictEqual(json.request.headers['x-acs-credentials-provider'], 'static_ak');

    const bytes = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    bytes.request.reqBodyType = 'byte';
    bytes.request.body = Buffer.from('hello');
    bytes.request.action = '';
    await client.modifyRequest(bytes, new $SPI.AttributeMap({}));
    assert.strictEqual(bytes.request.headers['x-acs-action'], undefined);

    const form = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    form.request.reqBodyType = 'form';
    form.request.body = { b: 'x y', a: ['m', 'n'] };
    await client.modifyRequest(form, new $SPI.AttributeMap({}));
    assert.strictEqual(form.request.headers['content-type'], 'application/x-www-form-urlencoded');

    const stream = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    stream.request.stream = readable('stream-body');
    await client.modifyRequest(stream, new $SPI.AttributeMap({}));
    assert.strictEqual(stream.request.headers['content-type'], 'application/octet-stream');

    await client.modifyRequest(context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com')), new $SPI.AttributeMap({}));

    const bearer = context(request(credential({ type: 'bearer', bearerToken: 'bearer-token', providerName: 'bearer' })), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    await client.modifyRequest(bearer, new $SPI.AttributeMap({}));
    assert.strictEqual(bearer.request.headers['Authorization'], 'Bearer bearer-token');

    const sts = context(request(credential({ type: 'sts', accessKeyId: 'sts-ak', accessKeySecret: 'secret', securityToken: 'token', providerName: '' })), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    await client.modifyRequest(sts, new $SPI.AttributeMap({}));
    assert.strictEqual(sts.request.headers['x-acs-security-token'], 'token');

    const idToken = context(request(credential({ type: 'id_token', securityToken: 'id-token', providerName: '' })), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    await client.modifyRequest(idToken, new $SPI.AttributeMap({}));
    assert.strictEqual(idToken.request.headers['x-acs-zero-trust-idtoken'], 'id-token');

    const anonymous = context(request(null), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    anonymous.request.authType = 'Anonymous';
    await client.modifyRequest(anonymous, new $SPI.AttributeMap({}));
    assert.strictEqual(anonymous.request.headers['Authorization'], undefined);

    await assert.rejects(client.modifyRequest(context(request(null), configuration('ecs.cn-hangzhou.aliyuncs.com')), new $SPI.AttributeMap({})));
    await assert.rejects(client.modifyRequest(context(request({ getCredential: async () => { throw new Error('credential down'); } }), configuration('ecs.cn-hangzhou.aliyuncs.com')), new $SPI.AttributeMap({})));

    const badBytes = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    badBytes.request.reqBodyType = 'byte';
    badBytes.request.body = 'not-bytes';
    await assert.rejects(client.modifyRequest(badBytes, new $SPI.AttributeMap({})));

    const badForm = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    badForm.request.reqBodyType = 'form';
    badForm.request.body = 'not-a-map';
    await assert.rejects(client.modifyRequest(badForm, new $SPI.AttributeMap({})));

    const badStream = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    badStream.request.stream = boom();
    await assert.rejects(client.modifyRequest(badStream, new $SPI.AttributeMap({})));
  });

  function response(status: number, contentType: string, body: string, bodyType: string): $SPI.InterceptorContext {
    const ctx = context(request(accessKey()), configuration('ecs.cn-hangzhou.aliyuncs.com'));
    ctx.request.bodyType = bodyType;
    ctx.response = new $SPI.InterceptorContextResponse({
      statusCode: status,
      headers: contentType ? { 'content-type': contentType } : {},
      body: readable(body)
    });
    return ctx;
  }

  it('parses responses', async function () {
    const client = new Client();
    await assert.rejects(client.modifyResponse(response(400, 'application/json', '{"Code":"Invalid","Message":"no","RequestId":"req-1","Description":"d"}', 'json'), new $SPI.AttributeMap({})), (err: any) => {
      assert.strictEqual(err.code, 'Invalid');
      assert.ok(String(err.message).indexOf('request id: req-1') >= 0);
      return true;
    });

    const denied = response(403, 'application/json', '{"code":"Denied","message":"no","requestId":"body-id"}', 'json');
    denied.response.headers['x-acs-request-id'] = 'header-id';
    await assert.rejects(client.modifyResponse(denied, new $SPI.AttributeMap({})), (err: any) => err.code === 'Denied');

    const xml = '<Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error>';
    await assert.rejects(client.modifyResponse(response(500, 'text/xml', xml, 'xml'), new $SPI.AttributeMap({})), (err: any) => err.code === 'XmlCode');

    await client.modifyResponse(response(204, '', '', 'string'), new $SPI.AttributeMap({}));
    const binary = response(200, '', 'raw', 'binary');
    await client.modifyResponse(binary, new $SPI.AttributeMap({}));
    assert.ok(binary.response.deserializedBody);
    const bytes = response(200, '', 'raw', 'byte');
    await client.modifyResponse(bytes, new $SPI.AttributeMap({}));
    assert.strictEqual(Buffer.from(bytes.response.deserializedBody).toString(), 'raw');
    const text = response(200, '', 'hello', 'string');
    await client.modifyResponse(text, new $SPI.AttributeMap({}));
    assert.strictEqual(text.response.deserializedBody, 'hello');
    const parsed = response(200, '', '{"ok":true}', 'json');
    await client.modifyResponse(parsed, new $SPI.AttributeMap({}));
    assert.strictEqual(parsed.response.deserializedBody.ok, true);
    const array = response(200, '', '[1,2]', 'array');
    await client.modifyResponse(array, new $SPI.AttributeMap({}));
    assert.strictEqual(array.response.deserializedBody.length, 2);
    const other = response(200, '', 'plain', 'other');
    await client.modifyResponse(other, new $SPI.AttributeMap({}));
    assert.strictEqual(other.response.deserializedBody, 'plain');
    await assert.rejects(client.modifyResponse(response(200, '', '{"ok":', 'json'), new $SPI.AttributeMap({})));
  });
});
