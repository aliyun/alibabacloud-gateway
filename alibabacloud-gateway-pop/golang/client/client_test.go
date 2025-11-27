package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

func Test_GetRegion(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(nil, nil, nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(nil, nil, tea.String("cn-hangzhou"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String(""), tea.String(""), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test.alibaba.api.com"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-dualstack.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-vpc.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-share.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-vpc.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-share.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-dualstack.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-proxy.cn-hangzhou.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com"), nil)))
	utils.AssertEqual(t, "cn-edge-1", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.cn-edge-1.aliyuncs.com"), nil)))
}

func Test_GetSignedHeaders(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	// 测试空headers
	result := client.GetSignedHeaders(make(map[string]*string))
	utils.AssertEqual(t, 1, len(result))

	// 测试只包含需要签名的headers
	headers := map[string]*string{
		"host":         tea.String("example.com"),
		"Host":         tea.String("example.com"),
		"content-type": tea.String("application/json"),
		"x-acs-action": tea.String("TestAction"),
	}
	result = client.GetSignedHeaders(headers)
	utils.AssertEqual(t, 3, len(result))
	utils.AssertEqual(t, "content-type", tea.StringValue(result[0]))
	utils.AssertEqual(t, "host", tea.StringValue(result[1]))
	utils.AssertEqual(t, "x-acs-action", tea.StringValue(result[2]))

	// 测试包含不需要签名的headers
	headersWithIgnore := map[string]*string{
		"host":          tea.String("example.com"),
		"content-type":  tea.String("application/json"),
		"x-acs-action":  tea.String("TestAction"),
		"authorization": tea.String("Bearer token"),
		"user-agent":    tea.String("Go-client"),
	}
	result = client.GetSignedHeaders(headersWithIgnore)
	utils.AssertEqual(t, 3, len(result))
	utils.AssertEqual(t, "content-type", tea.StringValue(result[0]))
	utils.AssertEqual(t, "host", tea.StringValue(result[1]))
	utils.AssertEqual(t, "x-acs-action", tea.StringValue(result[2]))

	// 测试空值header
	headersWithEmpty := map[string]*string{
		"host":         tea.String("example.com"),
		"content-type": nil,
		"x-acs-action": tea.String("TestAction"),
	}
	result = client.GetSignedHeaders(headersWithEmpty)
	utils.AssertEqual(t, 2, len(result))
	utils.AssertEqual(t, "host", tea.StringValue(result[0]))
	utils.AssertEqual(t, "x-acs-action", tea.StringValue(result[1]))
}
