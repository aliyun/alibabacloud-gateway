package client

import (
	"bytes"
	"compress/zlib"
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4"
)

type decompressor interface {
	// @param bodyRawSize guanranteed > 0
	decompress(src []byte, bodyRawSize int) ([]byte, error)
}

var availableDecompressors = map[string]bool{"lz4": true, "zstd": true, "gzip": true, "deflate": true}

var supportedDecompressors = map[string]decompressor{
	"lz4":     &lz4Decompressor{},
	"gzip":    &gzipDecompressor{},
	"deflate": &gzipDecompressor{},
	"zstd":    newZstdDecompressor(),
}

type lz4Decompressor struct {
}

func (d *lz4Decompressor) decompress(src []byte, bodyRawSize int) ([]byte, error) {
	out := make([]byte, bodyRawSize)
	_, err := lz4.UncompressBlock(src, out)
	return out, err
}

type gzipDecompressor struct {
}

func (d *gzipDecompressor) decompress(src []byte, bodyRawSize int) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

type zstdDecompressor struct {
	reader *zstd.Decoder
}

func newZstdDecompressor() *zstdDecompressor {
	res := &zstdDecompressor{}
	res.reader, _ = zstd.NewReader(nil)
	return res
}

func (d *zstdDecompressor) decompress(src []byte, bodyRawSize int) ([]byte, error) {
	out := make([]byte, 0, bodyRawSize)
	return d.reader.DecodeAll(src, out)
}
