// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// Description:
//
// # Read data from a readable stream, and parse it by JSON format
//
// @param stream - the readable stream
//
// @return the parsed result
func ReadAndUncompressBlock(stream io.Reader, compressType *string, bodyRawSize *string) (_result io.Reader, _err error) {
	rawSize, err := strconv.ParseInt(*bodyRawSize, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("fail to parse bodyRawSize %s, error: %w", *bodyRawSize, err)
	}
	if rawSize == 0 {
		return bytes.NewReader([]byte{}), nil
	}

	decompressor, ok := supportedDecompressors[*compressType]
	if !ok {
		return nil, fmt.Errorf("unsupported decompress type: %s", *compressType)
	}

	body, err := util.ReadAsBytes(stream)
	if err != nil {
		return nil, fmt.Errorf("fail to read body, error: %w", err)
	}

	uncompressed, err := decompressor.decompress(body, int(rawSize))
	if err != nil {
		return nil, fmt.Errorf("fail to decompress using compresstype %s, error %w", *compressType, err)
	}

	if len(uncompressed) != int(rawSize) {
		return nil, fmt.Errorf("unexpected uncompressed size: %d, expected: %d, compressType: %s", len(uncompressed), rawSize, *compressType)
	}
	return bytes.NewReader(uncompressed), nil
}

// Description:
//
// Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.
//
// @param src - the data to be compressed
//
// @param compressType - the compress type
//
// @return the compressed data
//
// @throws error if the compress type is not supported or the compress failed
func Compress(src []byte, compressType *string) (_result []byte, _err error) {
	if compressor, ok := supportedCompressors[*compressType]; ok {
		return compressor.compress(src)
	}
	return nil, fmt.Errorf("unsupported compress type: %s", *compressType)
}

func IsCompressorAvailable(compressType *string) (_result *bool) {
	if _, ok := availableCompressors[*compressType]; ok {
		return tea.Bool(true)
	}
	return tea.Bool(false)
}

func IsDecompressorAvailable(compressType *string) (_result *bool) {
	_, ok := availableDecompressors[*compressType]
	return tea.Bool(ok)
}

func BytesLength(src []byte) (_result *int64) {
	return tea.Int64(int64(len(src)))
}
