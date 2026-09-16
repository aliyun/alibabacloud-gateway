package client

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
)

func TestSignatureVersion(t *testing.T) {
	data, err := ioutil.ReadFile("../../testdata/signature-version.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct {
		Name                      string
		Region, Endpoint, Version *string
		Want                      string
		Resolved                  *string
	}
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatal(err)
	}
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := &spi.InterceptorContext{
				Request:       &spi.InterceptorContextRequest{SignatureVersion: tc.Version},
				Configuration: &spi.InterceptorContextConfiguration{RegionId: tc.Region, Endpoint: tc.Endpoint},
			}
			got, err := c.GetSignatureVersion(ctx)
			if err != nil || tea.StringValue(got) != tc.Want {
				t.Fatalf("version = %v, %v; want %s", tea.StringValue(got), err, tc.Want)
			}
			if tea.StringValue(ctx.Configuration.RegionId) != tea.StringValue(tc.Resolved) {
				t.Fatalf("region = %q; want %q", tea.StringValue(ctx.Configuration.RegionId), tea.StringValue(tc.Resolved))
			}
			if ctx.Request.SignatureVersion != tc.Version {
				t.Fatal("explicit signature setting was modified")
			}
		})
	}
}

func TestModifyRequestAutoV4(t *testing.T) {
	for _, version := range []*string{nil, tea.String(""), tea.String("v1")} {
		for _, body := range [][]byte{nil, []byte("request body")} {
			c, err := NewClient()
			if err != nil {
				t.Fatal(err)
			}
			cred, err := credential.NewCredential(&credential.Config{Type: tea.String("access_key"), AccessKeyId: tea.String("test-id"), AccessKeySecret: tea.String("test-secret")})
			if err != nil {
				t.Fatal(err)
			}
			req := &spi.InterceptorContextRequest{SignatureVersion: version, Credential: cred, Headers: map[string]*string{}, Query: map[string]*string{}, Method: tea.String("POST"), Pathname: tea.String("/"), Action: tea.String("TestAction"), ReqBodyType: tea.String("binary"), UserAgent: tea.String("test")}
			if body != nil {
				req.Body = body
			}
			ctx := &spi.InterceptorContext{Request: req, Configuration: &spi.InterceptorContextConfiguration{Endpoint: tea.String("cn-acdr-ut-1-internal.log.aliyuncs.com")}}
			if err := c.ModifyRequest(ctx, &spi.AttributeMap{}); err != nil {
				t.Fatal(err)
			}
			auth := tea.StringValue(req.Headers["authorization"])
			if tea.StringValue(version) == "v1" {
				if !strings.HasPrefix(auth, "LOG ") {
					t.Fatalf("expected V1 authorization, got %s", auth)
				}
				if _, ok := req.Headers["x-log-content-sha256"]; ok {
					t.Fatal("V1 request received V4 hash")
				}
			} else {
				if !strings.HasPrefix(auth, "SLS4-HMAC-SHA256 ") || !strings.Contains(auth, "/cn-acdr-ut-1/sls/aliyun_v4_request") {
					t.Fatalf("unexpected V4 authorization: %s", auth)
				}
				hash := sha256.Sum256(body)
				if tea.StringValue(req.Headers["x-log-content-sha256"]) != hex.EncodeToString(hash[:]) {
					t.Fatal("incorrect V4 body hash")
				}
			}
		}
	}
}
