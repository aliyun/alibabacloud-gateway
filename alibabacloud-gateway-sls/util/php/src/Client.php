<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewaySls\Util;

use GuzzleHttp\Psr7\Stream;
use GuzzleHttp\Psr7\Utils;
use \Exception;

/**
 * Read data from a readable stream, and parse it by JSON format
 * @param stream the readable stream
 * @return the parsed result
 */
class Client {

    /**
     * @param Stream $stream
     * @param string $compressType
     * @param string $bodyRawSize
     * @return Stream
     */
    public static function readAndUncompressBlock($stream, $compressType, $bodyRawSize){        
        if ($compressType[0] !== 'gzip') {
            throw new Exception("Unknown compression type: '$compressType[0]'");
        }
        $rawSize = intval($bodyRawSize[0]);
        $stream->rewind();
        $uncompressedData = gzuncompress($stream);
        if ($uncompressedData === false || strlen($uncompressedData) !== $rawSize) {
            throw new Exception('GZIP decompression failed.');
        }

        return Utils::streamFor($uncompressedData);
    }
}
