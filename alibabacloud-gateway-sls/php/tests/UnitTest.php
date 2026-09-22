<?php

namespace Darabonba\GatewaySls\Tests;

use Darabonba\GatewaySls\Client;
use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\InterceptorContext\configuration;
use Darabonba\GatewaySpi\Models\InterceptorContext\request;
use PHPUnit\Framework\TestCase;

/**
 * @internal
 * @coversNothing
 */
final class UnitTest extends TestCase
{
    public function testParseRegion()
    {
        $client = new Client();
        $cases = array(
            array('', ''),
            array('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('cn-hangzhou-acdr-ut-1.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('http://cn-hangzhou-acdr-ut-1.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('cn-hangzhou-acdr-ut-1-intranet.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('cn-hangzhou-acdr-ut-1-share.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('cn-hangzhou-acdr-ut-1-vpc.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('cn-hangzhou-acdr-ut-1-internal.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            array('cn-hangzhou.sls.aliyuncs.com', 'cn-hangzhou'),
            array('cn-hangzhou.log.aliyuncs.com', 'cn-hangzhou'),
            array('ftp://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', ''),
            array('cn-hangzhou-acdr-ut-1.oss.aliyuncs.com', ''),
            array('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com.cn', ''),
            array('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com/path', ''),
            array('https://https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', ''),
            array('CN-HANGZHOU.sls.aliyuncs.com', ''),
            array('cn_hangzhou.sls.aliyuncs.com', ''),
            array('cn-hangzhou-acdr-ut-1-intranet-share.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1-intranet'),
        );
        foreach ($cases as $item) {
            $this->assertEquals($item[1], $client->parseRegion($item[0]), $item[0]);
        }
        $this->assertEquals('', $client->parseRegion(null));
    }

    public function testSetSignV4IfInAcdr()
    {
        $client = new Client();

        $acdrRegion = $this->context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com', null);
        $client->setSignV4IfInAcdr($acdrRegion);
        $this->assertEquals('v4', $acdrRegion->request->signatureVersion);
        $this->assertEquals('cn-hangzhou-acdr-ut-1', $acdrRegion->configuration->regionId);

        $ordinaryRegion = $this->context('cn-hangzhou', 'cn-hangzhou.log.aliyuncs.com', null);
        $client->setSignV4IfInAcdr($ordinaryRegion);
        $this->assertNull($ordinaryRegion->request->signatureVersion);

        $fromEndpoint = $this->context(null, 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', null);
        $client->setSignV4IfInAcdr($fromEndpoint);
        $this->assertEquals('v4', $fromEndpoint->request->signatureVersion);
        $this->assertEquals('cn-hangzhou-acdr-ut-1', $fromEndpoint->configuration->regionId);

        $ordinaryEndpoint = $this->context('', 'cn-hangzhou.sls.aliyuncs.com', null);
        $client->setSignV4IfInAcdr($ordinaryEndpoint);
        $this->assertNull($ordinaryEndpoint->request->signatureVersion);
        $this->assertEquals('', $ordinaryEndpoint->configuration->regionId);

        $emptyEndpoint = $this->context(null, null, null);
        $client->setSignV4IfInAcdr($emptyEndpoint);
        $this->assertNull($emptyEndpoint->request->signatureVersion);

        $explicitV1 = $this->context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'v1');
        $client->setSignV4IfInAcdr($explicitV1);
        $this->assertEquals('v1', $explicitV1->request->signatureVersion);

        $explicitV4 = $this->context('cn-hangzhou', 'cn-hangzhou.log.aliyuncs.com', 'v4');
        $client->setSignV4IfInAcdr($explicitV4);
        $this->assertEquals('v4', $explicitV4->request->signatureVersion);

        $intranet = $this->context('', 'https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com', null);
        $client->setSignV4IfInAcdr($intranet);
        $this->assertEquals('v4', $intranet->request->signatureVersion);
        $this->assertEquals('cn-hangzhou-acdr-ut-1', $intranet->configuration->regionId);

        $invalidAcdrHint = $this->context(null, 'not-a-valid-acdr-ut-host', null);
        $client->setSignV4IfInAcdr($invalidAcdrHint);
        $this->assertNull($invalidAcdrHint->request->signatureVersion);
        $this->assertNull($invalidAcdrHint->configuration->regionId);

        $shanghai = $this->context('cn-shanghai-acdr-ut-2', null, null);
        $client->setSignV4IfInAcdr($shanghai);
        $this->assertEquals('v4', $shanghai->request->signatureVersion);

        $explicitV1FromEndpoint = $this->context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'v1');
        $client->setSignV4IfInAcdr($explicitV1FromEndpoint);
        $this->assertEquals('v1', $explicitV1FromEndpoint->request->signatureVersion);
        $this->assertEquals('', $explicitV1FromEndpoint->configuration->regionId);

        $client->setSignV4IfInAcdr($acdrRegion);
        $this->assertEquals('v4', $acdrRegion->request->signatureVersion);
        $this->assertEquals('cn-hangzhou-acdr-ut-1', $acdrRegion->configuration->regionId);
    }

    public function testModifyConfiguration()
    {
        $client = new Client();

        $acdr = $this->context('cn-hangzhou-acdr-ut-1', null, null);
        $client->modifyConfiguration($acdr, null);
        $this->assertEquals('cn-hangzhou-acdr-ut-1.log.aliyuncs.com', $acdr->configuration->endpoint);
        $this->assertEquals('v4', $acdr->request->signatureVersion);
        $this->assertEquals('cn-hangzhou-acdr-ut-1', $acdr->configuration->regionId);

        $ordinary = $this->context(null, null, null);
        $client->modifyConfiguration($ordinary, null);
        $this->assertEquals('cn-hangzhou.log.aliyuncs.com', $ordinary->configuration->endpoint);
        $this->assertNull($ordinary->request->signatureVersion);

        $fromEndpoint = $this->context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', null);
        $client->modifyConfiguration($fromEndpoint, null);
        $this->assertEquals('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', $fromEndpoint->configuration->endpoint);
        $this->assertEquals('v4', $fromEndpoint->request->signatureVersion);
        $this->assertEquals('cn-hangzhou-acdr-ut-1', $fromEndpoint->configuration->regionId);

        $explicit = $this->context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com', 'v1');
        $client->modifyConfiguration($explicit, null);
        $this->assertEquals('v1', $explicit->request->signatureVersion);

        $client->modifyConfiguration($acdr, null);
        $this->assertEquals('v4', $acdr->request->signatureVersion);
    }

    private function context($regionId, $endpoint, $signatureVersion)
    {
        $context = new InterceptorContext(array());
        $context->request = new request(array());
        $context->request->signatureVersion = $signatureVersion;
        $context->configuration = new configuration(array());
        $context->configuration->regionId = $regionId;
        $context->configuration->endpoint = $endpoint;
        return $context;
    }
}
