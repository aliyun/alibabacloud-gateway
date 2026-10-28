<?php

namespace Darabonba\GatewayPop;

use AlibabaCloud\Tea\Utils\Utils;
use OneSm\Sm3;

/**
 * Private signing helpers. SHA-256 uses ext-hash. SM3 keeps lizhichao/one-sm.
 */
class Codec
{
    public static function percentEncode($raw)
    {
        if (!isset($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        $str = urlencode($raw);
        $str = str_replace("+", "%20", $str);
        $str = str_replace("*", "%2A", $str);
        $str = str_replace("%7E", "~", $str);
        return $str;
    }

    public static function hexEncode($raw)
    {
        if (empty($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        $ret = '';
        foreach ($raw as $b) {
            $str = dechex($b);
            if (strlen($str) < 2) {
                $str = str_pad($str, 2, '0', STR_PAD_LEFT);
            }
            $ret .= $str;
        }
        return $ret;
    }

    public static function hash($raw, $algorithm)
    {
        if (empty($raw) || empty($algorithm)) {
            return null;
        }
        $str = Utils::toString($raw);
        if (false !== strpos($algorithm, 'HMAC-SHA256') || false !== strpos($algorithm, 'RSA-SHA256')) {
            return Utils::toBytes(hash('sha256', $str, true));
        }
        if (false !== strpos($algorithm, 'HMAC-SM3')) {
            return Utils::toBytes(hex2bin(self::sm3($str)));
        }
        return null;
    }

    public static function HmacSHA256Sign($stringToSign, $secret)
    {
        if (!isset($stringToSign) || !isset($secret)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        return Utils::toBytes(hash_hmac('sha256', $stringToSign, $secret, true));
    }

    public static function HmacSHA256SignByBytes($stringToSign, $secret)
    {
        return self::HmacSHA256Sign($stringToSign, Utils::toString($secret));
    }

    public static function HmacSM3Sign($stringToSign, $secret)
    {
        if (!isset($stringToSign) || !isset($secret)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        return Utils::toBytes(self::hmacSm3($stringToSign, $secret, true));
    }

    public static function HmacSM3SignByBytes($stringToSign, $secret)
    {
        return self::HmacSM3Sign($stringToSign, Utils::toString($secret));
    }

    public static function split($raw, $sep, $limit = null)
    {
        if (empty($raw)) {
            return [];
        }
        return null === $limit ? explode($sep, $raw) : explode($sep, $raw, $limit);
    }

    public static function replace($raw, $oldStr, $newStr, $index = null)
    {
        if (0 === $index) {
            return $raw;
        }
        return str_replace(self::value($oldStr), self::value($newStr), self::value($raw), $index);
    }

    public static function contains($s, $substr)
    {
        if (is_array($s)) {
            return in_array($substr, $s);
        }
        return false !== strpos(self::value($s), self::value($substr));
    }

    public static function hasPrefix($s, $prefix)
    {
        return 0 === strpos(self::value($s), self::value($prefix));
    }

    public static function toLower($s)
    {
        return strtolower(self::value($s));
    }

    public static function subString($s, $start, $end)
    {
        return substr($s, $start, $end - $start);
    }

    public static function equals($expect, $actual)
    {
        return false == strcmp(self::value($expect), self::value($actual));
    }

    public static function trim($str)
    {
        return trim($str);
    }

    public static function keySet($raw)
    {
        if (!isset($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        return array_keys($raw);
    }

    public static function join($raw, $sep)
    {
        if (!isset($raw) || empty($sep)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        return implode($sep, $raw);
    }

    public static function size($raw)
    {
        return count($raw);
    }

    public static function ascSort($raw)
    {
        if (!isset($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        sort($raw);
        return $raw;
    }

    public static function append(&$array, $item)
    {
        if (isset($array) && is_array($array)) {
            array_push($array, $item);
        }
    }

    private static function value($value, $default = '')
    {
        return null === $value ? $default : $value;
    }

    private static function sm3($str)
    {
        return (new Sm3())->sign($str);
    }

    private static function hmacSm3($data, $key, $rawOutput = false)
    {
        $pack = 'H' . strlen(self::sm3('test'));
        $blocksize = 64;
        if (strlen($key) > $blocksize) {
            $key = pack($pack, self::sm3($key));
        }
        $key = str_pad($key, $blocksize, chr(0x00));
        $ipad = $key ^ str_repeat(chr(0x36), $blocksize);
        $opad = $key ^ str_repeat(chr(0x5C), $blocksize);
        $hmac = self::sm3($opad . pack($pack, self::sm3($ipad . $data)));
        return $rawOutput ? pack($pack, $hmac) : $hmac;
    }
}
