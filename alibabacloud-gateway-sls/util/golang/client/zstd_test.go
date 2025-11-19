package client

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
)

func TestZstdSimple(t *testing.T) {
	// 准备测试数据
	input := "Hello, Zstd!"
	originalBytes := []byte(input)
	compressType := "zstd"

	// 1. 压缩
	compressedBytes, err := Compress(originalBytes, &compressType)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	// 确保压缩后的数据不为空
	if len(compressedBytes) == 0 {
		t.Fatal("Compressed data is empty")
	}

	// 2. 解压
	// ReadAndUncompressBlock 需要原始大小作为字符串参数
	rawSizeStr := strconv.Itoa(len(originalBytes))

	resultReader, err := ReadAndUncompressBlock(
		bytes.NewReader(compressedBytes),
		&compressType,
		tea.String(rawSizeStr),
	)
	if err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}

	decodedBytes, err := ioutil.ReadAll(resultReader)
	if err != nil {
		t.Fatalf("Failed to read decompressed data: %v", err)
	}

	// 3. 验证
	if string(decodedBytes) != input {
		t.Errorf("Content mismatch.\nExpected: %s\nGot: %s", input, string(decodedBytes))
	}
}
