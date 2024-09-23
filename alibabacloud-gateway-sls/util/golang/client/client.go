// This file is auto-generated, don't edit it. Thanks.
/**
 * Read data from a readable stream, and parse it by JSON format
 * @param stream the readable stream
 * @return the parsed result
 */
package client

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strconv"

	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/pierrec/lz4"
)

func ReadAndUncompressBlock(stream io.Reader, compressType *string, bodyRawSize *string) (_result io.Reader, _err error) {
	if *compressType == "lz4" {
		return decompressLz4(stream, bodyRawSize)
	}
	if *compressType == "gzip" {
		return decompressGzip(stream, bodyRawSize)
	}
	return nil, fmt.Errorf("unsupported compress type %s", *compressType)
}

func decompressLz4(stream io.Reader, bodyRawSize *string) (_result io.Reader, _err error) {
	rawSize, err := strconv.ParseInt(*bodyRawSize, 10, 64)
	if err != nil {
		return nil, err
	}
	out := make([]byte, rawSize)
	if rawSize != 0 {
		body, _ := util.ReadAsBytes(stream)
		len, err := lz4.UncompressBlock(body, out)
		if err != nil || int64(len) != rawSize {
			return nil, err
		}
	}
	return bytes.NewReader(out), nil
}

func decompressGzip(stream io.Reader, bodyRawSize *string) (_result io.Reader, _err error) {
	rawSize, err := strconv.ParseInt(*bodyRawSize, 10, 64)
	if err != nil {
		return nil, err
	}
	if rawSize != 0 {
		body, _ := util.ReadAsBytes(stream)
		reader, err := gzip.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		out, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		if len(out) != int(rawSize) {
			return nil, fmt.Errorf("unexpected gzip body size %d, expected %d", len(out), rawSize)
		}
		return bytes.NewReader(out), nil
	}
	return bytes.NewReader(make([]byte, 0)), nil
}
