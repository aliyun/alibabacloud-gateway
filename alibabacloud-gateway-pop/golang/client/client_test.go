package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

func Test_GetRegion(t *testing.T) {
	client := &Client{}
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(nil, nil)))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String(""), tea.String(""))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test.alibaba.api.com"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test.aliyuncs.com"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-dualstack.aliyuncs.com"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.aliyuncs.com"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-vpc.aliyuncs.com"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-share.aliyuncs.com"))))
	utils.AssertEqual(t, "center", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test.cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-vpc.cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-share.cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-dualstack.cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-proxy.cn-hangzhou.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com"))))
	utils.AssertEqual(t, "cn-edge-1", tea.StringValue(client.GetRegion(tea.String("test"), tea.String("test-inner.cn-edge-1.aliyuncs.com"))))
}
