<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewaySls\Util;

use GuzzleHttp\Psr7\Stream;
use GuzzleHttp\Psr7\Utils;
use \Exception;


class Client {

    /**
     * @param Stream $stream
     * @param string[] $compressType, this array parameter is for backward compatible
     * @param string[] $bodyRawSize, this array parameter is for backward compatible
     * @return Stream
     */
    public static function readAndUncompressBlock($stream, $compressType, $bodyRawSize){
        if ($compressType[0] !== 'gzip' && $compressType[0] !== 'deflate') {
            throw new Exception("Unsupported decompress type: $compressType[0]");
        }
        $rawSize = intval($bodyRawSize[0]);
        $stream->rewind();
        $uncompressedData = gzuncompress($stream);
        if ($uncompressedData === false) {
            throw new Exception("Fail to decompress using compress type: $compressType[0]");
        }
        $uncompressedSize = strlen($uncompressedData);
        if ($uncompressedSize !== $rawSize) {
            throw new Exception("Unexpected uncompressed size: $uncompressedSize , expected: $rawSize, compressType: $compressType[0]");
        }

        return Utils::streamFor($uncompressedData);
    }

    /**
     * Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.
     * @param int[] $src the data to be compressed
     * @param string $compressType the compress type
     * @return array the compressed data
     */
    public static function compress($src, $compressType){
        if ($compressType !== 'gzip' && $compressType !== 'deflate') {
            throw new Exception("Unsupported compress type: $compressType");
        }
        $compressed = gzcompress(self::toString($src));
        if ($compressed === false) {
            throw new Exception("Fail to compress using compress type: $compressType");
        }
        // performance is very low now
        return self::toBytes($compressed);
    }

    /**
     * @param string $compressType
     * @return bool
     */
    public static function isCompressorAvailable($compressType){
        return $compressType === 'gzip' || $compressType === 'deflate';
    }

    /**
     * @param string $compressType
     * @return bool
     */
    public static function isDecompressorAvailable($compressType){
        return $compressType === 'gzip' || $compressType === 'deflate';
    }

    /**
     * @param int[] $src
     * @return int
     */
    public static function bytesLength($src){
        return count($src);
    }

    /**
     * Convert a bytes to string(utf8).
     *
     * @param array $bytes
     *
     * @return string the return string
     */
    private static function toString($bytes)
    {
        if (\is_string($bytes)) {
            return $bytes;
        }
        $str = '';
        foreach ($bytes as $ch) {
            $str .= \chr($ch);
        }

        return $str;
    }

    /**
     * Convert a string(utf8) to bytes.
     *
     * @param string $string
     *
     * @return array the return bytes
     */
    private static function toBytes($string)
    {
        $bytes = [];
        for ($i = 0; $i < \strlen($string); ++$i) {
            $bytes[] = \ord($string[$i]);
        }

        return $bytes;
    }
}
