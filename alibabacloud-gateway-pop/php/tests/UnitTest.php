<?php

namespace Darabonba\GatewayPop\Tests;

use Darabonba\GatewayPop\Client;
use PHPUnit\Framework\TestCase;

/**
 * @internal
 * @coversNothing
 */
final class UnitTest extends TestCase
{

    public function testGetRegion()
    {
        $client = new Client();
        $this->assertEquals("center", $client->getRegion(null, null));
        $this->assertEquals("center", $client->getRegion("", ""));
        $this->assertEquals("center", $client->getRegion("test", "test"));
        $this->assertEquals("center", $client->getRegion("test", "test.alibaba.api.com"));
        $this->assertEquals("center", $client->getRegion("test", "test.aliyuncs.com"));
        $this->assertEquals("center", $client->getRegion("test", "test-dualstack.aliyuncs.com"));
        $this->assertEquals("center", $client->getRegion("test", "test-inner.aliyuncs.com"));
        $this->assertEquals("center", $client->getRegion("test", "test-vpc.aliyuncs.com"));
        $this->assertEquals("center", $client->getRegion("test", "test-share.aliyuncs.com"));
        $this->assertEquals("center", $client->getRegion("test", "test-cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test.cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-inner.cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-vpc.cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-share.cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-dualstack.cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-proxy.cn-hangzhou.aliyuncs.com"));
        $this->assertEquals("cn-hangzhou-acdr-ut-1", $client->getRegion("test", "test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com"));
        $this->assertEquals("cn-edge-1", $client->getRegion("test", "test-inner.cn-edge-1.aliyuncs.com"));
    }

}
