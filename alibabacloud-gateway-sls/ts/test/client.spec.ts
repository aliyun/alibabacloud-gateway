import Client from '../src/client';
import * as $SPI from '@alicloud/gateway-spi';
import assert from 'assert';
import 'mocha';

function context(regionId: string, endpoint: string, signatureVersion?: string): $SPI.InterceptorContext {
    return new $SPI.InterceptorContext({
        request: new $SPI.InterceptorContextRequest({
            signatureVersion: signatureVersion,
        }),
        configuration: new $SPI.InterceptorContextConfiguration({
            regionId: regionId,
            endpoint: endpoint,
        }),
    });
}

describe('Client', function () {
    it('parseRegion should ok', function () {
        const client = new Client();
        const cases: [string, string][] = [
            ['', ''],
            ['cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['cn-hangzhou-acdr-ut-1.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['http://cn-hangzhou-acdr-ut-1.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['cn-hangzhou-acdr-ut-1-intranet.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['cn-hangzhou-acdr-ut-1-share.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['cn-hangzhou-acdr-ut-1-vpc.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['cn-hangzhou-acdr-ut-1-internal.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'],
            ['cn-hangzhou.sls.aliyuncs.com', 'cn-hangzhou'],
            ['cn-hangzhou.log.aliyuncs.com', 'cn-hangzhou'],
            ['ftp://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', ''],
            ['cn-hangzhou-acdr-ut-1.oss.aliyuncs.com', ''],
            ['cn-hangzhou-acdr-ut-1.sls.aliyuncs.com.cn', ''],
            ['cn-hangzhou-acdr-ut-1.sls.aliyuncs.com/path', ''],
            ['https://https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', ''],
            ['CN-HANGZHOU.sls.aliyuncs.com', ''],
            ['cn_hangzhou.sls.aliyuncs.com', ''],
            ['cn-hangzhou-acdr-ut-1-intranet-share.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1-intranet'],
        ];
        for (const item of cases) {
            assert.strictEqual(client.parseRegion(item[0]), item[1], item[0]);
        }
    });

    it('setSignV4IfInAcdr should ok', function () {
        const client = new Client();

        const acdrRegion = context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com');
        client.setSignV4IfInAcdr(acdrRegion);
        assert.strictEqual(acdrRegion.request.signatureVersion, 'v4');
        assert.strictEqual(acdrRegion.configuration.regionId, 'cn-hangzhou-acdr-ut-1');

        const ordinaryRegion = context('cn-hangzhou', 'cn-hangzhou.log.aliyuncs.com');
        client.setSignV4IfInAcdr(ordinaryRegion);
        assert.strictEqual(ordinaryRegion.request.signatureVersion, undefined);

        const fromEndpoint = context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com');
        client.setSignV4IfInAcdr(fromEndpoint);
        assert.strictEqual(fromEndpoint.request.signatureVersion, 'v4');
        assert.strictEqual(fromEndpoint.configuration.regionId, 'cn-hangzhou-acdr-ut-1');

        const ordinaryEndpoint = context('', 'cn-hangzhou.sls.aliyuncs.com');
        client.setSignV4IfInAcdr(ordinaryEndpoint);
        assert.ok(!ordinaryEndpoint.request.signatureVersion);
        assert.strictEqual(ordinaryEndpoint.configuration.regionId, '');

        const emptyEndpoint = context('', '');
        client.setSignV4IfInAcdr(emptyEndpoint);
        assert.ok(!emptyEndpoint.request.signatureVersion);

        const explicitV1 = context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'v1');
        client.setSignV4IfInAcdr(explicitV1);
        assert.strictEqual(explicitV1.request.signatureVersion, 'v1');

        const explicitV4 = context('cn-hangzhou', 'cn-hangzhou.log.aliyuncs.com', 'v4');
        client.setSignV4IfInAcdr(explicitV4);
        assert.strictEqual(explicitV4.request.signatureVersion, 'v4');

        const intranet = context('', 'https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com');
        client.setSignV4IfInAcdr(intranet);
        assert.strictEqual(intranet.request.signatureVersion, 'v4');
        assert.strictEqual(intranet.configuration.regionId, 'cn-hangzhou-acdr-ut-1');

        const invalidAcdrHint = context('', 'not-a-valid-acdr-ut-host');
        client.setSignV4IfInAcdr(invalidAcdrHint);
        assert.ok(!invalidAcdrHint.request.signatureVersion);
        assert.strictEqual(invalidAcdrHint.configuration.regionId, '');

        const shanghai = context('cn-shanghai-acdr-ut-2', '');
        client.setSignV4IfInAcdr(shanghai);
        assert.strictEqual(shanghai.request.signatureVersion, 'v4');

        const explicitV1FromEndpoint = context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'v1');
        client.setSignV4IfInAcdr(explicitV1FromEndpoint);
        assert.strictEqual(explicitV1FromEndpoint.request.signatureVersion, 'v1');
        assert.strictEqual(explicitV1FromEndpoint.configuration.regionId, '');

        client.setSignV4IfInAcdr(acdrRegion);
        assert.strictEqual(acdrRegion.request.signatureVersion, 'v4');
        assert.strictEqual(acdrRegion.configuration.regionId, 'cn-hangzhou-acdr-ut-1');
    });

    it('modifyConfiguration should ok', async function () {
        const client = new Client();
        const attributeMap = new $SPI.AttributeMap({});

        const acdr = context('cn-hangzhou-acdr-ut-1', '');
        await client.modifyConfiguration(acdr, attributeMap);
        assert.strictEqual(acdr.configuration.endpoint, 'cn-hangzhou-acdr-ut-1.log.aliyuncs.com');
        assert.strictEqual(acdr.request.signatureVersion, 'v4');
        assert.strictEqual(acdr.configuration.regionId, 'cn-hangzhou-acdr-ut-1');

        const ordinary = context('', '');
        await client.modifyConfiguration(ordinary, attributeMap);
        assert.strictEqual(ordinary.configuration.endpoint, 'cn-hangzhou.log.aliyuncs.com');
        assert.ok(!ordinary.request.signatureVersion);

        const fromEndpoint = context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com');
        await client.modifyConfiguration(fromEndpoint, attributeMap);
        assert.strictEqual(fromEndpoint.configuration.endpoint, 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com');
        assert.strictEqual(fromEndpoint.request.signatureVersion, 'v4');
        assert.strictEqual(fromEndpoint.configuration.regionId, 'cn-hangzhou-acdr-ut-1');

        const explicit = context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com', 'v1');
        await client.modifyConfiguration(explicit, attributeMap);
        assert.strictEqual(explicit.request.signatureVersion, 'v1');

        await client.modifyConfiguration(acdr, attributeMap);
        assert.strictEqual(acdr.request.signatureVersion, 'v4');
    });
});
