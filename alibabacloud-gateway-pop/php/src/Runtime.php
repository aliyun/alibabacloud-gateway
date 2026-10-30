<?php

namespace Darabonba\GatewayPop;

use AlibabaCloud\Dara\Url;
use AlibabaCloud\Dara\Util\BytesUtil;
use AlibabaCloud\Dara\Util\StreamUtil;
use AlibabaCloud\Tea\Model;
use OneSm\Sm3;

/**
 * Gateway-private helpers. Not part of the supported public API.
 */
trait Runtime
{
    private static function percentEncode($raw)
    {
        if (!isset($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }

        return Url::percentEncode($raw);
    }

    private static function hexEncode($raw)
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

    private static function hash($raw, $algorithm)
    {
        if (!is_array($raw) || empty($algorithm)) {
            return null;
        }
        $str = BytesUtil::toString($raw);
        if (false !== strpos($algorithm, 'HMAC-SHA256') || false !== strpos($algorithm, 'RSA-SHA256')) {
            return BytesUtil::from(hash('sha256', $str, true));
        }
        if (false !== strpos($algorithm, 'HMAC-SM3')) {
            return BytesUtil::from(hex2bin(self::sm3($str)));
        }

        return null;
    }

    private static function HmacSHA256Sign($stringToSign, $secret)
    {
        if (!isset($stringToSign) || !isset($secret)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }

        return BytesUtil::from(hash_hmac('sha256', $stringToSign, $secret, true));
    }

    private static function HmacSHA256SignByBytes($stringToSign, $secret)
    {
        return self::HmacSHA256Sign($stringToSign, BytesUtil::toString($secret));
    }

    private static function HmacSM3Sign($stringToSign, $secret)
    {
        if (!isset($stringToSign) || !isset($secret)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }

        return BytesUtil::from(self::hmacSm3($stringToSign, $secret, true));
    }

    private static function HmacSM3SignByBytes($stringToSign, $secret)
    {
        return self::HmacSM3Sign($stringToSign, BytesUtil::toString($secret));
    }

    private static function split($raw, $sep, $limit = null)
    {
        if (empty($raw)) {
            return array();
        }

        return null === $limit ? explode($sep, $raw) : explode($sep, $raw, $limit);
    }

    private static function replace($raw, $oldStr, $newStr, $index = null)
    {
        if (0 === $index) {
            return $raw;
        }

        return str_replace(self::value($oldStr), self::value($newStr), self::value($raw), $index);
    }

    private static function contains($s, $substr)
    {
        if (is_array($s)) {
            return in_array($substr, $s);
        }

        return false !== strpos(self::value($s), self::value($substr));
    }

    private static function hasPrefix($s, $prefix)
    {
        return 0 === strpos(self::value($s), self::value($prefix));
    }

    private static function toLower($s)
    {
        return strtolower(self::value($s));
    }

    private static function subString($s, $start, $end)
    {
        return substr($s, $start, $end - $start);
    }

    private static function equals($expect, $actual)
    {
        return false == strcmp(self::value($expect), self::value($actual));
    }

    private static function trim($str)
    {
        return trim($str);
    }

    private static function keySet($raw)
    {
        if (!isset($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }

        return array_keys($raw);
    }

    private static function join($raw, $sep)
    {
        if (!isset($raw) || empty($sep)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }

        return implode($sep, $raw);
    }

    private static function size($raw)
    {
        return count($raw);
    }

    private static function ascSort($raw)
    {
        if (!isset($raw)) {
            throw new \InvalidArgumentException('not a valid value for parameter');
        }
        sort($raw);

        return $raw;
    }

    private static function append(&$array, $item)
    {
        if (isset($array) && is_array($array)) {
            array_push($array, $item);
        }
    }

    private static function toBytes($string)
    {
        return BytesUtil::from($string);
    }

    private static function readAsBytes($stream)
    {
        return StreamUtil::readAsBytes($stream);
    }

    private static function readAsString($stream)
    {
        return StreamUtil::readAsString($stream);
    }

    private static function readAsJSON($stream)
    {
        return StreamUtil::readAsJSON($stream);
    }

    private static function getNonce()
    {
        return md5(uniqid() . uniqid(md5(microtime(true)), true));
    }

    private static function defaultString($real, $default = '')
    {
        return null === $real ? $default : $real;
    }

    private static function toJSONString($object)
    {
        if (is_string($object)) {
            return $object;
        }
        if ($object instanceof Model) {
            $object = $object->toMap();
        }

        return json_encode($object, JSON_UNESCAPED_UNICODE + JSON_UNESCAPED_SLASHES);
    }

    private static function empty_($val)
    {
        return empty($val);
    }

    private static function equalString($left, $right)
    {
        return $left === $right;
    }

    private static function equalNumber($left, $right)
    {
        return $left === $right;
    }

    private static function isUnset(&$value = null)
    {
        return !isset($value) || null === $value;
    }

    private static function assertAsBytes($value)
    {
        if (self::isBytes($value)) {
            return $value;
        }
        throw new \InvalidArgumentException('It is not a bytes value.');
    }

    private static function assertAsMap($any)
    {
        if (is_array($any)) {
            return $any;
        }
        throw new \InvalidArgumentException('It is not a map value.');
    }

    private static function is4xx($code)
    {
        return $code >= 400 && $code < 500;
    }

    private static function is5xx($code)
    {
        return $code >= 500 && $code < 600;
    }

    private static function getTimestamp()
    {
        return gmdate('Y-m-d\\TH:i:s\\Z');
    }

    private static function getEndpointRules($product, $regionId, $endpointType = '', $network = '', $suffix = '')
    {
        if (empty($product)) {
            throw new \InvalidArgumentException('Product name cannot be empty.');
        }
        $endpoint = '<product><suffix><network>.<region_id>.aliyuncs.com';
        if ('regional' === $endpointType) {
            if (empty($regionId)) {
                throw new \InvalidArgumentException('RegionId is empty, please set a valid RegionId');
            }
            $endpoint = self::renderEndpoint($endpoint, 'region_id', strtolower($regionId));
        } elseif ('central' === $endpointType) {
            $endpoint = '<product><suffix><network>.aliyuncs.com';
        } else {
            throw new \InvalidArgumentException('Invalid EndpointType');
        }
        if (!empty($network) && 'public' !== $network) {
            $endpoint = self::renderEndpoint($endpoint, 'network', '-' . $network);
        } else {
            $endpoint = self::renderEndpoint($endpoint, 'network', '');
        }
        if (!empty($suffix)) {
            $endpoint = self::renderEndpoint($endpoint, 'suffix', '-' . $suffix);
        } else {
            $endpoint = self::renderEndpoint($endpoint, 'suffix', '');
        }

        return self::renderEndpoint($endpoint, 'product', strtolower($product));
    }

    private static function toForm($filter)
    {
        $query = $filter;
        if (null === $query) {
            return '';
        }
        if ($query instanceof Model) {
            $query = $query->toMap();
        }
        $tmp = array();
        foreach ($query as $k => $v) {
            if (0 !== strpos($k, '_')) {
                $tmp[$k] = $v;
            }
        }
        $res = self::flatten($tmp);
        ksort($res);

        return http_build_query($res);
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

    private static function isBytes($value)
    {
        if (!is_array($value)) {
            return false;
        }
        $i = 0;
        foreach ($value as $k => $ord) {
            if ($k !== $i) {
                return false;
            }
            if (!is_int($ord)) {
                return false;
            }
            if ($ord < 0 || $ord > 255) {
                return false;
            }
            ++$i;
        }

        return true;
    }

    private static function renderEndpoint($str, $tag, $replace)
    {
        return str_replace('<' . $tag . '>', $replace, $str);
    }

    private static function flatten($items = array(), $delimiter = '.', $prepend = '')
    {
        $flatten = array();
        foreach ($items as $key => $value) {
            $pos = is_int($key) ? $key + 1 : $key;
            if ($value instanceof Model) {
                $value = $value->toMap();
            } elseif (is_object($value)) {
                $value = get_object_vars($value);
            }
            if (is_array($value) && !empty($value)) {
                $flatten = array_merge(
                    $flatten,
                    self::flatten($value, $delimiter, $prepend . $pos . $delimiter)
                );
            } else {
                if (is_bool($value)) {
                    $value = true === $value ? 'true' : 'false';
                }
                $flatten[$prepend . $pos] = $value;
            }
        }

        return $flatten;
    }

    private static function parseXml($xml)
    {
        if (PHP_VERSION_ID < 80000) {
            libxml_disable_entity_loader(true);
        }
        $parsed = simplexml_load_string($xml, 'SimpleXMLElement', LIBXML_NOCDATA);
        if ($parsed === false) {
            return array();
        }
        return json_decode(json_encode($parsed), true);
    }
}
