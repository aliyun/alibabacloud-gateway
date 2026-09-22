package client

import (
	"testing"

	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

func Test_ParseRegion(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	cases := []struct {
		endpoint string
		want     string
	}{
		{"", ""},
		{"cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"cn-hangzhou-acdr-ut-1.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"http://cn-hangzhou-acdr-ut-1.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"cn-hangzhou-acdr-ut-1-intranet.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"cn-hangzhou-acdr-ut-1-share.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"cn-hangzhou-acdr-ut-1-vpc.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"cn-hangzhou-acdr-ut-1-internal.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com", "cn-hangzhou-acdr-ut-1"},
		{"cn-hangzhou.sls.aliyuncs.com", "cn-hangzhou"},
		{"cn-hangzhou.log.aliyuncs.com", "cn-hangzhou"},
		{"ftp://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", ""},
		{"cn-hangzhou-acdr-ut-1.oss.aliyuncs.com", ""},
		{"cn-hangzhou-acdr-ut-1.sls.aliyuncs.com.cn", ""},
		{"cn-hangzhou-acdr-ut-1.sls.aliyuncs.com/path", ""},
		{"https://https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", ""},
		{"CN-HANGZHOU.sls.aliyuncs.com", ""},
		{"cn_hangzhou.sls.aliyuncs.com", ""},
		{"cn-hangzhou-acdr-ut-1-intranet-share.sls.aliyuncs.com", "cn-hangzhou-acdr-ut-1-intranet"},
	}
	for _, c := range cases {
		utils.AssertEqual(t, c.want, client.ParseRegion(c.endpoint))
	}
}

func Test_SetSignV4IfInAcdr(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	acdrRegion := context(tea.String("cn-hangzhou-acdr-ut-1"), tea.String("cn-hangzhou.log.aliyuncs.com"), nil)
	client.SetSignV4IfInAcdr(acdrRegion)
	utils.AssertEqual(t, "v4", tea.StringValue(acdrRegion.Request.SignatureVersion))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(acdrRegion.Configuration.RegionId))

	ordinaryRegion := context(tea.String("cn-hangzhou"), tea.String("cn-hangzhou.log.aliyuncs.com"), nil)
	client.SetSignV4IfInAcdr(ordinaryRegion)
	utils.AssertEqual(t, "", tea.StringValue(ordinaryRegion.Request.SignatureVersion))

	fromEndpoint := context(nil, tea.String("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"), nil)
	client.SetSignV4IfInAcdr(fromEndpoint)
	utils.AssertEqual(t, "v4", tea.StringValue(fromEndpoint.Request.SignatureVersion))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(fromEndpoint.Configuration.RegionId))

	ordinaryEndpoint := context(tea.String(""), tea.String("cn-hangzhou.sls.aliyuncs.com"), nil)
	client.SetSignV4IfInAcdr(ordinaryEndpoint)
	utils.AssertEqual(t, "", tea.StringValue(ordinaryEndpoint.Request.SignatureVersion))
	utils.AssertEqual(t, "", tea.StringValue(ordinaryEndpoint.Configuration.RegionId))

	emptyEndpoint := context(nil, nil, nil)
	client.SetSignV4IfInAcdr(emptyEndpoint)
	utils.AssertEqual(t, "", tea.StringValue(emptyEndpoint.Request.SignatureVersion))

	explicitV1 := context(tea.String("cn-hangzhou-acdr-ut-1"), tea.String("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"), tea.String("v1"))
	client.SetSignV4IfInAcdr(explicitV1)
	utils.AssertEqual(t, "v1", tea.StringValue(explicitV1.Request.SignatureVersion))

	explicitV4 := context(tea.String("cn-hangzhou"), tea.String("cn-hangzhou.log.aliyuncs.com"), tea.String("v4"))
	client.SetSignV4IfInAcdr(explicitV4)
	utils.AssertEqual(t, "v4", tea.StringValue(explicitV4.Request.SignatureVersion))

	intranet := context(tea.String(""), tea.String("https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com"), nil)
	client.SetSignV4IfInAcdr(intranet)
	utils.AssertEqual(t, "v4", tea.StringValue(intranet.Request.SignatureVersion))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(intranet.Configuration.RegionId))

	invalidAcdrHint := context(nil, tea.String("not-a-valid-acdr-ut-host"), nil)
	client.SetSignV4IfInAcdr(invalidAcdrHint)
	utils.AssertEqual(t, "", tea.StringValue(invalidAcdrHint.Request.SignatureVersion))
	utils.AssertEqual(t, "", tea.StringValue(invalidAcdrHint.Configuration.RegionId))

	shanghai := context(tea.String("cn-shanghai-acdr-ut-2"), nil, nil)
	client.SetSignV4IfInAcdr(shanghai)
	utils.AssertEqual(t, "v4", tea.StringValue(shanghai.Request.SignatureVersion))

	explicitV1FromEndpoint := context(tea.String(""), tea.String("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"), tea.String("v1"))
	client.SetSignV4IfInAcdr(explicitV1FromEndpoint)
	utils.AssertEqual(t, "v1", tea.StringValue(explicitV1FromEndpoint.Request.SignatureVersion))
	utils.AssertEqual(t, "", tea.StringValue(explicitV1FromEndpoint.Configuration.RegionId))

	client.SetSignV4IfInAcdr(acdrRegion)
	utils.AssertEqual(t, "v4", tea.StringValue(acdrRegion.Request.SignatureVersion))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(acdrRegion.Configuration.RegionId))
}

func Test_ModifyConfiguration(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	acdr := context(tea.String("cn-hangzhou-acdr-ut-1"), nil, nil)
	err = client.ModifyConfiguration(acdr, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1.log.aliyuncs.com", tea.StringValue(acdr.Configuration.Endpoint))
	utils.AssertEqual(t, "v4", tea.StringValue(acdr.Request.SignatureVersion))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(acdr.Configuration.RegionId))

	ordinary := context(nil, nil, nil)
	err = client.ModifyConfiguration(ordinary, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "cn-hangzhou.log.aliyuncs.com", tea.StringValue(ordinary.Configuration.Endpoint))
	utils.AssertEqual(t, "", tea.StringValue(ordinary.Request.SignatureVersion))

	fromEndpoint := context(tea.String(""), tea.String("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"), nil)
	err = client.ModifyConfiguration(fromEndpoint, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", tea.StringValue(fromEndpoint.Configuration.Endpoint))
	utils.AssertEqual(t, "v4", tea.StringValue(fromEndpoint.Request.SignatureVersion))
	utils.AssertEqual(t, "cn-hangzhou-acdr-ut-1", tea.StringValue(fromEndpoint.Configuration.RegionId))

	explicit := context(tea.String("cn-hangzhou-acdr-ut-1"), tea.String("cn-hangzhou.log.aliyuncs.com"), tea.String("v1"))
	err = client.ModifyConfiguration(explicit, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "v1", tea.StringValue(explicit.Request.SignatureVersion))

	err = client.ModifyConfiguration(acdr, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "v4", tea.StringValue(acdr.Request.SignatureVersion))
}

func context(regionId, endpoint, signatureVersion *string) *spi.InterceptorContext {
	return &spi.InterceptorContext{
		Request: &spi.InterceptorContextRequest{
			SignatureVersion: signatureVersion,
		},
		Configuration: &spi.InterceptorContextConfiguration{
			RegionId: regionId,
			Endpoint: endpoint,
		},
	}
}
