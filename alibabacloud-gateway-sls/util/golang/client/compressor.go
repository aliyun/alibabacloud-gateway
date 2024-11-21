package client

import (
	"bytes"
	"compress/zlib"
	"fmt"

	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4"
)

type compressor interface {
	compress(src []byte) (result []byte, err error)
}

var availableCompressors = map[string]bool{"lz4": true, "zstd": true, "gzip": true, "deflate": true}

var supportedCompressors = map[string]compressor{
	"lz4":     &lz4Compressor{},
	"gzip":    &gzipCompressor{},
	"deflate": &gzipCompressor{},
	"zstd":    newZstdCompressor(zstd.SpeedDefault),
}

type lz4Compressor struct {
}

func (c *lz4Compressor) compress(src []byte) ([]byte, error) {
	result := make([]byte, lz4.CompressBlockBound(len(src)))
	n, err := lz4.CompressBlock(src, result, nil)
	if err != nil {
		return nil, fmt.Errorf("lz4 compress error: %w", err)
	}
	return result[:n], nil
}

type gzipCompressor struct {
}

func (c *gzipCompressor) compress(src []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write(src)
	if err != nil {
		return nil, fmt.Errorf("gzip compress error: %w", err)
	}
	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("gzip compress error: %w", err)
	}
	return b.Bytes(), nil
}

type zstdCompressor struct {
	writer *zstd.Encoder
	level  zstd.EncoderLevel
}

func newZstdCompressor(level zstd.EncoderLevel) *zstdCompressor {
	res := &zstdCompressor{
		level: level,
	}
	res.writer, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(res.level))
	return res
}

func (c *zstdCompressor) compress(src []byte) ([]byte, error) {
	return c.writer.EncodeAll(src, nil), nil
}
