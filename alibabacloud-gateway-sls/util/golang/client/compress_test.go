package client

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"strconv"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
)

func testData(data []byte, t *testing.T) {
	rawSize := len(data)
	for _, typ := range []string{"lz4", "gzip", "zstd"} {
		r, err := Compress(data, &typ)
		if err != nil {
			t.Fatal(err)
		}
		r2, err := ReadAndUncompressBlock(bytes.NewReader(r), &typ, tea.String(strconv.Itoa(rawSize)))
		if err != nil {
			t.Fatal(err)
		}
		decompressed, err := ioutil.ReadAll(r2)
		if err != nil {
			t.Fatal(err)
		}

		if len(decompressed) != rawSize {
			t.Fatal("decompressed data size is not equal to original data size")
		}
		if !bytes.Equal(decompressed, data) {
			t.Fatal("decompressed data is not equal to original data")
		}
	}
}

func generateRamdomBytes(length int) []byte {
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		data[i] = byte(rand.Intn(256))
	}
	return data
}

func TestCompress(t *testing.T) {
	testData([]byte("hello world"), t)
	testData([]byte("hello world你好"), t)
	testData([]byte(""), t)
	for i := 0; i < 100; i++ {
		str := generateRamdomBytes(rand.Intn(10000))
		testData(str, t)
	}
}
