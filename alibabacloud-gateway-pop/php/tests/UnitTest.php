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
        $this->assertEquals("center", $client->getRegion(null, null, null));
        $this->assertEquals("cn-hangzhou", $client->getRegion(null, null, "cn-hangzhou"));
        $this->assertEquals("center", $client->getRegion("", "", null));
        $this->assertEquals("center", $client->getRegion("test", "test", null));
        $this->assertEquals("center", $client->getRegion("test", "test.alibaba.api.com", null));
        $this->assertEquals("center", $client->getRegion("test", "test.aliyuncs.com", null));
        $this->assertEquals("center", $client->getRegion("test", "test-dualstack.aliyuncs.com", null));
        $this->assertEquals("center", $client->getRegion("test", "test-inner.aliyuncs.com", null));
        $this->assertEquals("center", $client->getRegion("test", "test-vpc.aliyuncs.com", null));
        $this->assertEquals("center", $client->getRegion("test", "test-share.aliyuncs.com", null));
        $this->assertEquals("center", $client->getRegion("test", "test-cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test.cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-inner.cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-vpc.cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-share.cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-dualstack.cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou", $client->getRegion("test", "test-proxy.cn-hangzhou.aliyuncs.com", null));
        $this->assertEquals("cn-hangzhou-acdr-ut-1", $client->getRegion("test", "test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com", null));
        $this->assertEquals("cn-edge-1", $client->getRegion("test", "test-inner.cn-edge-1.aliyuncs.com", null));
    }

    public function testGetSignedHeaders()
    {
        $client = new Client();

        // 测试空headers
        $emptyHeaders = [];
        $result = $client->getSignedHeaders($emptyHeaders);
        $this->assertEmpty($result);

        // 测试只包含需要签名的headers
        $headers = [
            "host" => "example.com",
            "content-type" => "application/json",
            "x-acs-action" => "TestAction"
        ];
        $result = $client->getSignedHeaders($headers);
        $this->assertCount(3, $result);
        $this->assertEquals("content-type", $result[0]);
        $this->assertEquals("host", $result[1]);
        $this->assertEquals("x-acs-action", $result[2]);

        // 测试包含不需要签名的headers
        $headersWithIgnore = [
            "host" => "example.com",
            "content-type" => "application/json",
            "x-acs-action" => "TestAction",
            "authorization" => "Bearer token",
            "user-agent" => "PHP-client"
        ];
        $result = $client->getSignedHeaders($headersWithIgnore);
        $this->assertCount(3, $result);
        $this->assertEquals("content-type", $result[0]);
        $this->assertEquals("host", $result[1]);
        $this->assertEquals("x-acs-action", $result[2]);

        // 测试空值header
        $headersWithEmpty = [
            "host" => "example.com",
            "content-type" => null,
            "x-acs-action" => "TestAction"
        ];
        $result = $client->getSignedHeaders($headersWithEmpty);
        $this->assertCount(2, $result);
        $this->assertEquals("host", $result[0]);
        $this->assertEquals("x-acs-action", $result[1]);

        // 测试大小写不敏感
        $headersWithCase = [
            "HOST" => "example.com",
            "Content-Type" => "application/json",
            "X-Acs-Action" => "TestAction"
        ];
        $result = $client->getSignedHeaders($headersWithCase);
        $this->assertCount(3, $result);
        $this->assertEquals("content-type", $result[0]);
        $this->assertEquals("host", $result[1]);
        $this->assertEquals("x-acs-action", $result[2]);
    }

}