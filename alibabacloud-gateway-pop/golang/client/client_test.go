package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/utils"
)

func Test_GetRegion(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(nil, nil, nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(nil, nil, dara.String("cn-hangzhou"))))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String(""), dara.String(""), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test.alibaba.api.com"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-dualstack.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-inner.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-vpc.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-share.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-inner.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-vpc.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-share.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-dualstack.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-proxy.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-edge-1", dara.StringValue(client.GetRegion(dara.String("test"), dara.String("test-inner.cn-edge-1.aliyuncs.com"), nil)))
}

func Test_GetSignedHeaders(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	// 测试空headers
	result := client.GetSignedHeaders(make(map[string]*string))
	utils.AssertEqual(t, 0, len(result))

	// 测试只包含需要签名的headers
	headers := map[string]*string{
		"host":         dara.String("example.com"),
		"Host":         dara.String("example.com"),
		"content-type": dara.String("application/json"),
		"x-acs-action": dara.String("TestAction"),
	}
	result = client.GetSignedHeaders(headers)
	utils.AssertEqual(t, 3, len(result))
	utils.AssertEqual(t, "content-type", dara.StringValue(result[0]))
	utils.AssertEqual(t, "host", dara.StringValue(result[1]))
	utils.AssertEqual(t, "x-acs-action", dara.StringValue(result[2]))

	// 测试包含不需要签名的headers
	headersWithIgnore := map[string]*string{
		"host":          dara.String("example.com"),
		"content-type":  dara.String("application/json"),
		"x-acs-action":  dara.String("TestAction"),
		"authorization": dara.String("Bearer token"),
		"user-agent":    dara.String("Go-client"),
	}
	result = client.GetSignedHeaders(headersWithIgnore)
	utils.AssertEqual(t, 3, len(result))
	utils.AssertEqual(t, "content-type", dara.StringValue(result[0]))
	utils.AssertEqual(t, "host", dara.StringValue(result[1]))
	utils.AssertEqual(t, "x-acs-action", dara.StringValue(result[2]))

	// 测试空值header
	headersWithEmpty := map[string]*string{
		"host":         dara.String("example.com"),
		"content-type": nil,
		"x-acs-action": dara.String("TestAction"),
	}
	result = client.GetSignedHeaders(headersWithEmpty)
	utils.AssertEqual(t, 2, len(result))
	utils.AssertEqual(t, "host", dara.StringValue(result[0]))
	utils.AssertEqual(t, "x-acs-action", dara.StringValue(result[1]))

	// Prefix pairs must not be mis-deduped via substring contains
	prefixHeaders := map[string]*string{
		"host":         dara.String("example.com"),
		"x-acs-foobar": dara.String("1"),
		"x-acs-foo":    dara.String("2"),
	}
	utils.AssertEqual(t, false, dara.BoolValue(listContains([]*string{dara.String("x-acs-foobar")}, dara.String("x-acs-foo"))))
	result = client.GetSignedHeaders(prefixHeaders)
	utils.AssertEqual(t, 3, len(result))
	utils.AssertEqual(t, "host", dara.StringValue(result[0]))
	utils.AssertEqual(t, "x-acs-foo", dara.StringValue(result[1]))
	utils.AssertEqual(t, "x-acs-foobar", dara.StringValue(result[2]))

	canonical := dara.StringValue(client.BuildCanonicalizedHeaders(prefixHeaders))
	utils.AssertContains(t, canonical, "x-acs-foo:2\n")
	utils.AssertContains(t, canonical, "x-acs-foobar:1\n")
}

func Test_ValidateSignedHeaders(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	err = client.ValidateSignedHeaders([]*string{dara.String("content-type"), dara.String("host"), dara.String("x-acs-date")})
	utils.AssertNil(t, err)

	err = client.ValidateSignedHeaders([]*string{dara.String("host")})
	utils.AssertNotNil(t, err)

	err = client.ValidateSignedHeaders([]*string{dara.String("x-acs-date")})
	utils.AssertNotNil(t, err)

	err = client.ValidateSignedHeaders([]*string{})
	utils.AssertNotNil(t, err)

	signingkey := client.GetSigningkey(client.Sha256, dara.String("secret"), dara.String("ecs"), dara.String("cn-hangzhou"), dara.String("20240101"))
	_, err = client.GetSignature(dara.String("/"), dara.String("GET"), map[string]*string{}, map[string]*string{
		"host": dara.String("example.com"),
	}, client.Sha256, dara.String(""), signingkey)
	utils.AssertNotNil(t, err)

	signature, err := client.GetSignature(dara.String("/"), dara.String("GET"), map[string]*string{}, map[string]*string{
		"host":       dara.String("example.com"),
		"x-acs-date": dara.String("2024-01-01T00:00:00Z"),
	}, client.Sha256, dara.String(""), signingkey)
	utils.AssertNil(t, err)
	utils.AssertNotNil(t, signature)
}
