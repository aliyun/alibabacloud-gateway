<?php

namespace Darabonba\GatewayMns\Tests;

use AlibabaCloud\Tea\Utils\Utils;
use Darabonba\GatewayMns\Client;
use DateTime;
use DateTimeZone;
use PHPUnit\Framework\TestCase;
use ReflectionMethod;

/**
 * @internal
 * @coversNothing
 */
final class UnitTest extends TestCase
{
    public function testGetDateISO8601FromUTC()
    {
        $cases = array(
            'Thu, 06 Feb 2020 07:32:54 GMT' => '20200206T073254Z',
            'Tue, 03 Jun 2025 11:23:48 GMT' => '20250603T112348Z',
            'Wed, 03 Dec 2025 11:23:48 GMT' => '20251203T112348Z',
            'Thu, 01 Jan 2026 00:00:00 GMT' => '20260101T000000Z',
            'Thu, 31 Dec 2026 23:59:59 GMT' => '20261231T235959Z',
        );
        foreach ($cases as $rfc => $expected) {
            $this->assertEquals($expected, $this->invokeGetDateISO8601FromUTC($rfc));
        }
    }

    public function testGetDateISO8601FromUTCRoundTrip()
    {
        $rfc = Utils::getDateUTCString();
        $this->assertEquals(1, preg_match('/^[A-Z][a-z]{2}, \d{2} [A-Z][a-z]{2} \d{4} \d{2}:\d{2}:\d{2} GMT$/', $rfc));

        $parsed = DateTime::createFromFormat('D, d M Y H:i:s T', $rfc, new DateTimeZone('GMT'));
        $this->assertNotFalse($parsed);
        $expected = $parsed->setTimezone(new DateTimeZone('UTC'))->format('Ymd\THis\Z');

        $actual = $this->invokeGetDateISO8601FromUTC($rfc);
        $this->assertEquals($expected, $actual);
        $this->assertEquals(1, preg_match('/^\d{8}T\d{6}Z$/', $actual));
    }

    private function invokeGetDateISO8601FromUTC($utc)
    {
        $method = new ReflectionMethod(Client::class, 'getDateISO8601FromUTC');
        $method->setAccessible(true);
        return $method->invoke(new Client(), $utc);
    }
}
