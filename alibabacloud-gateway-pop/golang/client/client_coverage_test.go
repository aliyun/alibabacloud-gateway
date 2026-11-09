package client

import (
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"

	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
	"github.com/aliyun/credentials-go/credentials"
)

type fixedCredential struct {
	model *credentials.CredentialModel
	err   error
}

func TestV4CanonicalQuerySignature(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	query := map[string]*string{
		"a-b": tea.String("1"), "a/b": tea.String("2"), "a": tea.String(""),
		"é": tea.String("中"), "10": tea.String("x"), "2": tea.String("y"),
		"null": nil, "punct": tea.String("!+*~"), "Signature": tea.String("ignored"),
	}
	const wantQuery = "%C3%A9=%E4%B8%AD&10=x&2=y&a=&a%2Fb=2&a-b=1&null=&punct=%21%2B%2A~"
	if got := tea.StringValue(client.BuildCanonicalizedResource(query)); got != wantQuery {
		t.Errorf("canonical query = %q, want %q", got, wantQuery)
	}
	headers := map[string]*string{
		"Host": tea.String(" example.com "), "x-acs-date": tea.String("2024-01-01T00:00:00Z"),
	}
	for _, vector := range []struct {
		algorithm *string
		payload   string
		signature string
	}{
		{client.Sha256, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "6a4bed968fa9e417380dbea97a003ca2861777b47565075631cd602af98351e2"},
		{client.Sm3, "1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b", "82b1f962d9dc4072ebcabf59cf649df29c83bd69d1d2b9646dff35fcb6c2e56b"},
	} {
		t.Run(tea.StringValue(vector.algorithm), func(t *testing.T) {
			key := client.GetSigningkey(vector.algorithm, tea.String("secret"), tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("20240101"))
			signature, err := client.GetSignature(tea.String("/"), tea.String("get"), query, headers, vector.algorithm, tea.String(vector.payload), key)
			if err != nil || tea.StringValue(signature) != vector.signature {
				t.Fatalf("signature = %q, error = %v; want %s", tea.StringValue(signature), err, vector.signature)
			}
		})
	}
}

func (f fixedCredential) GetCredential() (*credentials.CredentialModel, error) {
	return f.model, f.err
}

func (f fixedCredential) GetAccessKeyId() (*string, error) {
	if f.model == nil {
		return nil, f.err
	}
	return f.model.AccessKeyId, nil
}

func (f fixedCredential) GetAccessKeySecret() (*string, error) {
	if f.model == nil {
		return nil, f.err
	}
	return f.model.AccessKeySecret, nil
}

func (f fixedCredential) GetSecurityToken() (*string, error) {
	if f.model == nil {
		return nil, f.err
	}
	return f.model.SecurityToken, nil
}

func (f fixedCredential) GetBearerToken() *string {
	if f.model == nil {
		return nil
	}
	return f.model.BearerToken
}

func (f fixedCredential) GetType() *string {
	if f.model == nil {
		return nil
	}
	return f.model.Type
}

// A non-zero field keeps this value visible to tea-utils IsUnset.
// On Go 1.13, IsUnset treats the zero value of a struct as unset, so an
// empty reader never reaches Read.
type errReader struct {
	failed bool
}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func failingReader() errReader {
	return errReader{failed: true}
}

func akCredential() credentials.Credential {
	return fixedCredential{model: &credentials.CredentialModel{
		AccessKeyId:     tea.String("ak"),
		AccessKeySecret: tea.String("secret"),
		Type:            tea.String("access_key"),
		ProviderName:    tea.String("static_ak"),
	}}
}

func baseRequest() *spi.InterceptorContextRequest {
	return &spi.InterceptorContextRequest{
		Headers:     map[string]*string{},
		Query:       map[string]*string{},
		Pathname:    tea.String("/"),
		ProductId:   tea.String("ecs"),
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2014-05-26"),
		Protocol:    tea.String("HTTPS"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		BodyType:    tea.String("json"),
		ReqBodyType: tea.String("json"),
		Credential:  akCredential(),
		UserAgent:   tea.String("test-agent"),
	}
}

func baseConfig() *spi.InterceptorContextConfiguration {
	return &spi.InterceptorContextConfiguration{
		RegionId: tea.String("cn-hangzhou"),
		Endpoint: tea.String("ecs.cn-hangzhou.aliyuncs.com"),
	}
}

func newContext(request *spi.InterceptorContextRequest, config *spi.InterceptorContextConfiguration) *spi.InterceptorContext {
	return &spi.InterceptorContext{
		Request:       request,
		Configuration: config,
		Response:      &spi.InterceptorContextResponse{},
	}
}

func TestEndpointSelection(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	explicit, err := client.GetEndpoint(tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("regional"), tea.String(""), tea.String(""), nil, tea.String("ecs.example.com"))
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.example.com", tea.StringValue(explicit))

	mapped, err := client.GetEndpoint(tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("regional"), tea.String(""), tea.String(""), map[string]*string{
		"cn-hangzhou": tea.String("ecs.mapped.example.com"),
	}, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.mapped.example.com", tea.StringValue(mapped))

	regional, err := client.GetEndpoint(tea.String("Ecs"), tea.String("cn-hangzhou"), tea.String("regional"), tea.String(""), tea.String(""), nil, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.com", tea.StringValue(regional))

	inner, err := client.GetEndpoint(tea.String("Ecs"), tea.String("cn-hangzhou"), tea.String("regional"), tea.String("vpc"), tea.String("inner"), nil, nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs-inner-vpc.cn-hangzhou.aliyuncs.com", tea.StringValue(inner))

	central, err := client.GetEndpoint(tea.String("Ecs"), nil, tea.String("central"), tea.String("public"), nil, nil, tea.String(""))
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.aliyuncs.com", tea.StringValue(central))

	_, err = client.GetEndpoint(tea.String("Ecs"), tea.String(""), tea.String("regional"), tea.String(""), tea.String(""), nil, nil)
	utils.AssertNotNil(t, err)
}

func TestExplicitEndpointSuffix(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ input, want string }{
		{"oss-cn-hangzhou.aliyuncs.com", "oss-cn-hangzhou.aliyuncs.alibaba"},
		{"ecs.cn-hangzhou-internal.aliyuncs.com", "ecs.cn-hangzhou-internal.aliyuncs.alibaba"},
		{"ecs.aliyun.com", "ecs.aliyun.com"},
		{"custom.example.com", "custom.example.com"},
		{"ecs.aliyuncs.alibaba", "ecs.aliyuncs.alibaba"},
		{"ecs.aliyuncs.com.example.com", "ecs.aliyuncs.com.example.com"},
	} {
		for _, mapped := range []bool{false, true} {
			cfg := &spi.InterceptorContextConfiguration{RegionId: tea.String("cn-hangzhou"), Endpoint: tea.String(tc.input)}
			if mapped {
				cfg.Endpoint = nil
				cfg.EndpointMap = map[string]*string{"cn-hangzhou": tea.String(tc.input)}
			}
			attrs := &spi.AttributeMap{Key: map[string]*string{"endpointSuffix": tea.String("aliyuncs.alibaba")}}
			for i := 0; i < 2; i++ {
				if err := client.ModifyConfiguration(newContext(baseRequest(), cfg), attrs); err != nil {
					t.Fatal(err)
				}
				if tea.StringValue(cfg.Endpoint) != tc.want {
					t.Fatalf("mapped=%v input=%q got=%q want=%q", mapped, tc.input, tea.StringValue(cfg.Endpoint), tc.want)
				}
			}
		}
	}
}

func TestEndpointSuffixOverride(t *testing.T) {
	os.Setenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX", "aliyuncs.alibaba")
	defer os.Unsetenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX")

	client, err := NewClient()
	utils.AssertNil(t, err)
	config := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String("cn-hangzhou"),
		EndpointRule: tea.String("regional"),
	}
	err = client.ModifyConfiguration(newContext(baseRequest(), config), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.alibaba", tea.StringValue(config.Endpoint))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.regionWithSuffix(tea.String("ecs"), config.Endpoint, tea.String(""), tea.String("aliyuncs.alibaba"))))
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.regionWithSuffix(tea.String("ecs"), tea.String("ecs.cn-hangzhou.aliyuncs.com"), tea.String(""), tea.String("example.net"))))
	utils.AssertEqual(t, "aliyuncs.com", tea.StringValue(client.EndpointSuffix))

	overridden := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String("cn-hangzhou"),
		EndpointRule: tea.String("regional"),
	}
	err = client.ModifyConfiguration(newContext(baseRequest(), overridden), &spi.AttributeMap{
		Key: map[string]*string{
			"endpointSuffix": tea.String(" ALIYUNCS.ALIBABA "),
		},
	})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.alibaba", tea.StringValue(overridden.Endpoint))
	utils.AssertEqual(t, "ACS4-", tea.StringValue(client.SignatureTypePrefix))
	utils.AssertEqual(t, "ACS4-HMAC-SHA256", tea.StringValue(client.Sha256))
	unchanged, err := applyEndpointSuffix(tea.String("already.example.com"), tea.String("aliyuncs.alibaba"))
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "already.example.com", tea.StringValue(unchanged))
	sameHost, err := applyEndpointSuffix(tea.String("ecs.cn-hangzhou.aliyuncs.com"), tea.String(""))
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.com", tea.StringValue(sameHost))
	emptyHost, err := applyEndpointSuffix(tea.String(""), tea.String("aliyuncs.alibaba"))
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "", tea.StringValue(emptyHost))
	_, err = applyEndpointSuffix(tea.String("ecs.cn-hangzhou.aliyuncs.com"), tea.String("example.net"))
	utils.AssertNotNil(t, err)

	os.Unsetenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX")
	fallback := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String("cn-hangzhou"),
		EndpointRule: tea.String("regional"),
	}
	err = client.ModifyConfiguration(newContext(baseRequest(), fallback), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "aliyuncs.com", tea.StringValue(client.EndpointSuffix))
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.com", tea.StringValue(fallback.Endpoint))

	explicit := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String(""),
		EndpointRule: tea.String("regional"),
		Endpoint:     tea.String("ecs.cn-hangzhou.aliyuncs.com"),
	}
	os.Setenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX", "aliyuncs.alibaba")
	err = client.ModifyConfiguration(newContext(baseRequest(), explicit), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.alibaba", tea.StringValue(explicit.Endpoint))
	resolved, err := endpointSuffixForRequest(nil)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "cn-hangzhou", tea.StringValue(client.regionWithSuffix(tea.String("ecs"), explicit.Endpoint, explicit.RegionId, tea.String(resolved))))
	blank := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String("cn-hangzhou"),
		EndpointRule: tea.String("regional"),
	}
	err = client.ModifyConfiguration(newContext(baseRequest(), blank), &spi.AttributeMap{
		Key: map[string]*string{"endpointSuffix": tea.String("   ")},
	})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.alibaba", tea.StringValue(blank.Endpoint))
}

func TestEndpointSuffixRequestsDoNotCross(t *testing.T) {
	os.Unsetenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX")
	defer os.Unsetenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX")
	client, err := NewClient()
	utils.AssertNil(t, err)
	var wg sync.WaitGroup
	errCh := make(chan error, 16)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			suffix := "aliyuncs.alibaba"
			if i%2 == 1 {
				suffix = "aliyuncs.com"
			}
			built := &spi.InterceptorContextConfiguration{
				RegionId:     tea.String("cn-hangzhou"),
				EndpointRule: tea.String("regional"),
			}
			callErr := client.ModifyConfiguration(newContext(baseRequest(), built), &spi.AttributeMap{
				Key: map[string]*string{"endpointSuffix": tea.String(suffix)},
			})
			if callErr != nil {
				errCh <- callErr
				return
			}
			want := "ecs.cn-hangzhou." + suffix
			if tea.StringValue(built.Endpoint) != want {
				errCh <- errors.New(tea.StringValue(built.Endpoint) + " != " + want)
			}
		}(i)
	}
	wg.Wait()
	close(errCh)
	for callErr := range errCh {
		utils.AssertNil(t, callErr)
	}
}

func TestSignaturePrefixesDoNotCross(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)
	var wg sync.WaitGroup
	errCh := make(chan error, 16)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			prefix := "ACS4-"
			signPrefix := "aliyun_v4"
			if i%2 == 1 {
				prefix = "ACS3-"
				signPrefix = "aliyun_v3"
			}
			request := baseRequest()
			callErr := client.ModifyRequest(newContext(request, baseConfig()), &spi.AttributeMap{
				Key: map[string]*string{
					"signatureTypePrefix": tea.String(prefix),
					"signPrefix":          tea.String(signPrefix),
				},
			})
			if callErr != nil {
				errCh <- callErr
				return
			}
			authorization := tea.StringValue(request.Headers["Authorization"])
			if !strings.HasPrefix(authorization, prefix+"HMAC-SHA256 ") {
				errCh <- errors.New(authorization)
				return
			}
			if !strings.Contains(authorization, "/"+signPrefix+"_request,") {
				errCh <- errors.New(authorization)
			}
		}(i)
	}
	wg.Wait()
	close(errCh)
	for callErr := range errCh {
		utils.AssertNil(t, callErr)
	}
	utils.AssertEqual(t, "ACS4-", tea.StringValue(client.SignatureTypePrefix))
	utils.AssertEqual(t, "aliyun_v4", tea.StringValue(client.SignPrefix))
	utils.AssertEqual(t, "ACS4-HMAC-SHA256", tea.StringValue(client.Sha256))
}

func TestModifyConfiguration(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)
	request := baseRequest()
	config := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String("cn-hangzhou"),
		EndpointRule: tea.String("regional"),
		Network:      tea.String(""),
		Suffix:       tea.String(""),
	}
	err = client.ModifyConfiguration(newContext(request, config), &spi.AttributeMap{
		Key: map[string]*string{
			"signatureTypePrefix": tea.String("ACS4-"),
			"signPrefix":          tea.String("aliyun_v4"),
			"endpointSuffix":      tea.String("aliyuncs.com"),
		},
	})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "ecs.cn-hangzhou.aliyuncs.com", tea.StringValue(config.Endpoint))
	utils.AssertEqual(t, "ACS4-HMAC-SHA256", tea.StringValue(client.Sha256))

	kept := baseConfig()
	kept.Endpoint = tea.String("already.example.com")
	err = client.ModifyConfiguration(newContext(baseRequest(), kept), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "already.example.com", tea.StringValue(kept.Endpoint))

	_, err = client.GetEndpoint(tea.String("Ecs"), tea.String(""), tea.String("regional"), tea.String(""), tea.String(""), nil, nil)
	utils.AssertNotNil(t, err)
	broken := &spi.InterceptorContextConfiguration{
		RegionId:     tea.String(""),
		EndpointRule: tea.String("regional"),
	}
	err = client.ModifyConfiguration(newContext(baseRequest(), broken), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
}

func TestCanonicalRequestPieces(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	utils.AssertEqual(t, "", tea.StringValue(client.BuildCanonicalizedResource(nil)))
	resource := client.BuildCanonicalizedResource(map[string]*string{
		"b":     tea.String("x y"),
		"a":     tea.String("m*n"),
		"empty": tea.String(""),
	})
	utils.AssertEqual(t, "a=m%2An&b=x%20y&empty=", tea.StringValue(resource))

	headers := map[string]*string{
		"HOST":         tea.String("ignored-uppercase.example"),
		"Host":         tea.String("ignored-titlecase.example"),
		"host":         tea.String("  example.com  "),
		"X-Acs-Action": tea.String("One"),
		"x-acs-action": tea.String("Two"),
		"user-agent":   tea.String("skip"),
	}
	canonical := tea.StringValue(client.BuildCanonicalizedHeaders(headers))
	utils.AssertEqual(t, "host:example.com\nx-acs-action:One,Two\n", canonical)

	utils.AssertEqual(t, "fallback", client.DefaultAny(nil, "fallback"))
	utils.AssertEqual(t, "kept", client.DefaultAny("kept", "fallback"))

	signingkey := client.GetSigningkey(client.Sha256, tea.String("secret"), tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("20240101"))
	utils.AssertEqual(t, true, len(signingkey) > 0)
	sm3key := client.GetSigningkey(client.Sm3, tea.String("secret"), tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("20240101"))
	utils.AssertEqual(t, true, len(sm3key) > 0)
	unknown := client.GetSigningkey(tea.String("OTHER"), tea.String("secret"), tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("20240101"))
	utils.AssertEqual(t, 0, len(unknown))

	signature, err := client.GetSignature(tea.String(""), tea.String("GET"), map[string]*string{
		"q": tea.String("a b"),
	}, map[string]*string{
		"host":       tea.String("example.com"),
		"x-acs-date": tea.String("2024-01-01T00:00:00Z"),
	}, client.Sha256, tea.String("payload"), signingkey)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, 64, len(tea.StringValue(signature)))

	sm3Signature, err := client.GetSignature(tea.String("/"), tea.String("POST"), nil, map[string]*string{
		"host":       tea.String("example.com"),
		"x-acs-date": tea.String("2024-01-01T00:00:00Z"),
	}, client.Sm3, tea.String("payload"), sm3key)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, 64, len(tea.StringValue(sm3Signature)))

	otherSignature, err := client.GetSignature(tea.String("/"), tea.String("GET"), nil, map[string]*string{
		"host":       tea.String("example.com"),
		"x-acs-date": tea.String("2024-01-01T00:00:00Z"),
	}, tea.String("OTHER"), tea.String(""), unknown)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "", tea.StringValue(otherSignature))

	authorization, err := client.GetAuthorization(tea.String("/"), tea.String("GET"), nil, map[string]*string{
		"host":       tea.String("example.com"),
		"x-acs-date": tea.String("2024-01-01T00:00:00Z"),
	}, client.Sha256, tea.String("payload"), tea.String("ak"), signingkey, tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("20240101"))
	utils.AssertNil(t, err)
	utils.AssertContains(t, tea.StringValue(authorization), "ACS4-HMAC-SHA256 Credential=ak/20240101/cn-hangzhou/ecs/aliyun_v4_request,SignedHeaders=")
	utils.AssertContains(t, tea.StringValue(authorization), "host;x-acs-date")
}

func TestModifyRequestBodiesAndAuth(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	jsonRequest := baseRequest()
	jsonRequest.Body = map[string]interface{}{"name": "a b"}
	jsonRequest.Query = map[string]*string{"InstanceId": tea.String("i-1")}
	err = client.ModifyRequest(newContext(jsonRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "application/json; charset=utf-8", tea.StringValue(jsonRequest.Headers["content-type"]))
	utils.AssertEqual(t, "DescribeRegions", tea.StringValue(jsonRequest.Headers["x-acs-action"]))
	utils.AssertEqual(t, "static_ak", tea.StringValue(jsonRequest.Headers["x-acs-credentials-provider"]))
	utils.AssertContains(t, tea.StringValue(jsonRequest.Headers["Authorization"]), "ACS4-HMAC-SHA256 Credential=ak/")
	utils.AssertEqual(t, true, len(tea.StringValue(jsonRequest.Headers["x-acs-content-sha256"])) > 0)

	byteRequest := baseRequest()
	byteRequest.ReqBodyType = tea.String("byte")
	byteRequest.Body = []byte("hello")
	byteRequest.Action = tea.String("")
	err = client.ModifyRequest(newContext(byteRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, true, byteRequest.Headers["x-acs-action"] == nil)

	formRequest := baseRequest()
	formRequest.ReqBodyType = tea.String("form")
	formRequest.Body = map[string]interface{}{"b": "x y", "a": []interface{}{"m", "n"}}
	err = client.ModifyRequest(newContext(formRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "application/x-www-form-urlencoded", tea.StringValue(formRequest.Headers["content-type"]))

	streamRequest := baseRequest()
	streamRequest.Stream = strings.NewReader("stream-body")
	err = client.ModifyRequest(newContext(streamRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "application/octet-stream", tea.StringValue(streamRequest.Headers["content-type"]))

	emptyRequest := baseRequest()
	err = client.ModifyRequest(newContext(emptyRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)

	sm3Request := baseRequest()
	sm3Request.SignatureAlgorithm = client.Sm3
	sm3Request.Body = map[string]interface{}{"k": "v"}
	err = client.ModifyRequest(newContext(sm3Request, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, true, len(tea.StringValue(sm3Request.Headers["x-acs-content-sm3"])) > 0)
	utils.AssertContains(t, tea.StringValue(sm3Request.Headers["Authorization"]), "ACS4-HMAC-SM3 Credential=ak/")

	stsRequest := baseRequest()
	stsRequest.Credential = fixedCredential{model: &credentials.CredentialModel{
		AccessKeyId:     tea.String("sts-ak"),
		AccessKeySecret: tea.String("secret"),
		SecurityToken:   tea.String("token"),
		Type:            tea.String("sts"),
		ProviderName:    tea.String(""),
	}}
	err = client.ModifyRequest(newContext(stsRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "token", tea.StringValue(stsRequest.Headers["x-acs-security-token"]))
	utils.AssertEqual(t, "sts-ak", tea.StringValue(stsRequest.Headers["x-acs-accesskey-id"]))

	bearerRequest := baseRequest()
	bearerRequest.Credential = fixedCredential{model: &credentials.CredentialModel{
		BearerToken:  tea.String("bearer-token"),
		Type:         tea.String("bearer"),
		ProviderName: tea.String("bearer"),
	}}
	err = client.ModifyRequest(newContext(bearerRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "Bearer bearer-token", tea.StringValue(bearerRequest.Headers["Authorization"]))
	utils.AssertEqual(t, "BEARERTOKEN", tea.StringValue(bearerRequest.Headers["x-acs-signature-type"]))

	idRequest := baseRequest()
	idRequest.Credential = fixedCredential{model: &credentials.CredentialModel{
		SecurityToken: tea.String("id-token"),
		Type:          tea.String("id_token"),
	}}
	err = client.ModifyRequest(newContext(idRequest, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "id-token", tea.StringValue(idRequest.Headers["x-acs-zero-trust-idtoken"]))

	anonymous := baseRequest()
	anonymous.AuthType = tea.String("Anonymous")
	anonymous.Credential = nil
	err = client.ModifyRequest(newContext(anonymous, baseConfig()), &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, true, anonymous.Headers["Authorization"] == nil)

	missing := baseRequest()
	missing.Credential = nil
	err = client.ModifyRequest(newContext(missing, baseConfig()), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	sdkErr := err.(*tea.SDKError)
	utils.AssertEqual(t, "ParameterMissing", tea.StringValue(sdkErr.Code))
	utils.AssertEqual(t, false, tea.BoolValue(tea.Retryable(err)))

	brokenCred := baseRequest()
	brokenCred.Credential = fixedCredential{err: errors.New("credential down")}
	err = client.ModifyRequest(newContext(brokenCred, baseConfig()), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	badBytes := baseRequest()
	badBytes.ReqBodyType = tea.String("byte")
	badBytes.Body = "not-bytes"
	err = client.ModifyRequest(newContext(badBytes, baseConfig()), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	badForm := baseRequest()
	badForm.ReqBodyType = tea.String("form")
	badForm.Body = "not-a-map"
	err = client.ModifyRequest(newContext(badForm, baseConfig()), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	badStream := baseRequest()
	badStream.Stream = failingReader()
	err = client.ModifyRequest(newContext(badStream, baseConfig()), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
}

func responseContext(status int, contentType string, body io.Reader, bodyType string) *spi.InterceptorContext {
	request := baseRequest()
	request.BodyType = tea.String(bodyType)
	headers := map[string]*string{}
	if contentType != "" {
		headers["content-type"] = tea.String(contentType)
	}
	return &spi.InterceptorContext{
		Request:       request,
		Configuration: baseConfig(),
		Response: &spi.InterceptorContextResponse{
			StatusCode: tea.Int(status),
			Headers:    headers,
			Body:       body,
		},
	}
}

func TestModifyResponse(t *testing.T) {
	client, err := NewClient()
	utils.AssertNil(t, err)

	err = client.ModifyResponse(responseContext(400, "application/json", strings.NewReader(`{"Code":"Invalid","Message":"no","RequestId":"req-1","Description":"d","AccessDeniedDetail":{"Policy":"deny"}}`), "json"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	sdkErr := err.(*tea.SDKError)
	utils.AssertEqual(t, "Invalid", tea.StringValue(sdkErr.Code))
	utils.AssertContains(t, tea.StringValue(sdkErr.Message), "request id: req-1")
	utils.AssertEqual(t, false, tea.BoolValue(tea.Retryable(err)))

	withHeader := responseContext(403, "application/json", strings.NewReader(`{"code":"Denied","message":"no","requestId":"body-id"}`), "json")
	withHeader.Response.Headers["x-acs-request-id"] = tea.String("header-id")
	err = client.ModifyResponse(withHeader, &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	utils.AssertEqual(t, "Denied", tea.StringValue(err.(*tea.SDKError).Code))
	utils.AssertContains(t, tea.StringValue(err.(*tea.SDKError).Message), "request id:")
	utils.AssertEqual(t, false, tea.BoolValue(tea.Retryable(err)))

	xmlBody := `<Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error>`
	err = client.ModifyResponse(responseContext(500, "text/xml", strings.NewReader(xmlBody), "xml"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	utils.AssertEqual(t, "XmlCode", tea.StringValue(err.(*tea.SDKError).Code))
	utils.AssertEqual(t, true, tea.BoolValue(tea.Retryable(err)))

	err = client.ModifyResponse(responseContext(204, "", strings.NewReader(""), "string"), &spi.AttributeMap{})
	utils.AssertNil(t, err)

	binaryCtx := responseContext(200, "", strings.NewReader("raw"), "binary")
	err = client.ModifyResponse(binaryCtx, &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, true, binaryCtx.Response.DeserializedBody != nil)

	byteCtx := responseContext(200, "", strings.NewReader("raw"), "byte")
	err = client.ModifyResponse(byteCtx, &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "raw", string(byteCtx.Response.DeserializedBody.([]byte)))

	stringCtx := responseContext(200, "", strings.NewReader("hello"), "string")
	err = client.ModifyResponse(stringCtx, &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "hello", tea.StringValue(stringCtx.Response.DeserializedBody.(*string)))

	jsonCtx := responseContext(200, "", strings.NewReader(`{"ok":true}`), "json")
	err = client.ModifyResponse(jsonCtx, &spi.AttributeMap{})
	utils.AssertNil(t, err)
	parsed := jsonCtx.Response.DeserializedBody.(map[string]interface{})
	utils.AssertEqual(t, true, parsed["ok"])

	arrayCtx := responseContext(200, "", strings.NewReader(`[1,2]`), "array")
	err = client.ModifyResponse(arrayCtx, &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, 2, len(arrayCtx.Response.DeserializedBody.([]interface{})))

	otherCtx := responseContext(200, "", strings.NewReader("plain"), "other")
	err = client.ModifyResponse(otherCtx, &spi.AttributeMap{})
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "plain", tea.StringValue(otherCtx.Response.DeserializedBody.(*string)))

	err = client.ModifyResponse(responseContext(400, "application/json", failingReader(), "json"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(500, "text/xml", failingReader(), "xml"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(200, "", failingReader(), "string"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(200, "", failingReader(), "byte"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(200, "", failingReader(), "json"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(200, "", failingReader(), "array"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(200, "", failingReader(), "other"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	err = client.ModifyResponse(responseContext(204, "", failingReader(), "string"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	badJSON := responseContext(200, "", strings.NewReader(`{"ok":`), "json")
	err = client.ModifyResponse(badJSON, &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	notXMLError := responseContext(500, "text/xml", strings.NewReader(`<Other><Code>X</Code></Other>`), "xml")
	err = client.ModifyResponse(notXMLError, &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	err = client.ModifyResponse(responseContext(400, "application/json", strings.NewReader(`[1,2]`), "json"), &spi.AttributeMap{})
	utils.AssertNotNil(t, err)
	arrayJSON := responseContext(200, "", strings.NewReader(`[1,2]`), "json")
	err = client.ModifyResponse(arrayJSON, &spi.AttributeMap{})
	utils.AssertNotNil(t, err)

	_, err = client.GetAuthorization(tea.String("/"), tea.String("GET"), map[string]*string{}, map[string]*string{}, client.Sha256, tea.String("payload"), tea.String("ak"), []byte("key"), tea.String("ecs"), tea.String("cn-hangzhou"), tea.String("20240101"))
	utils.AssertNotNil(t, err)
	_, ok := err.(*tea.SDKError)
	utils.AssertEqual(t, true, ok)
}

func TestSlimHelperEdges(t *testing.T) {
	parts := splitString(tea.String("a-b-c"), tea.String("-"), tea.Int(2))
	utils.AssertEqual(t, "a", tea.StringValue(parts[0]))
	utils.AssertEqual(t, "b-c", tea.StringValue(parts[1]))
	utils.AssertEqual(t, "a:b-c", tea.StringValue(replaceString(tea.String("a-b-c"), tea.String("-"), tea.String(":"), tea.Int(1))))
	utils.AssertEqual(t, "bcde", tea.StringValue(subString(tea.String("abcdef"), tea.Int(1), tea.Int(-1))))
	utils.AssertEqual(t, 0, len(keySet(1)))

	pointer := tea.String("ptr")
	utils.AssertEqual(t, "ptr", tea.StringValue(toJSONString(pointer)))
	utils.AssertEqual(t, "str", tea.StringValue(toJSONString("str")))
	utils.AssertEqual(t, "hi", tea.StringValue(toJSONString([]byte("hi"))))
	utils.AssertEqual(t, "body", tea.StringValue(toJSONString(strings.NewReader("body"))))
	utils.AssertEqual(t, true, toJSONString(failingReader()) == nil)
	utils.AssertEqual(t, true, toJSONString(make(chan int)) == nil)

	form := tea.StringValue(toForm(map[string]interface{}{
		"empty":  nil,
		"nested": map[string]interface{}{"k": "v"},
		"items":  []interface{}{map[string]interface{}{"id": "1"}, "x"},
	}))
	utils.AssertContains(t, form, "nested.k=v")
	utils.AssertContains(t, form, "items.1.id=1")
	utils.AssertContains(t, form, "items.2=x")

	result := map[string]*string{}
	handleRepeatedParams(reflect.ValueOf([]interface{}(nil)), result, "items")
	utils.AssertEqual(t, 0, len(result))
}

// Changing the process environment after endpoint selection must not affect signing.
func TestSigningUsesSelectedEndpointAfterEnvironmentChange(t *testing.T) {
	previous, had := os.LookupEnv("ALIBABA_CLOUD_ENDPOINT_SUFFIX")
	defer func() {
		if had {
			os.Setenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX", previous)
		} else {
			os.Unsetenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX")
		}
	}()
	for _, selected := range []string{"aliyuncs.com", "aliyuncs.alibaba"} {
		for _, changed := range []string{"aliyuncs.com", "aliyuncs.alibaba", "invalid.example"} {
			os.Setenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX", selected)
			client, err := NewClient()
			utils.AssertNil(t, err)
			ctx := newContext(baseRequest(), baseConfig())
			ctx.Configuration.RegionId = tea.String("")
			attributes := &spi.AttributeMap{}
			utils.AssertNil(t, client.ModifyConfiguration(ctx, attributes))
			os.Setenv("ALIBABA_CLOUD_ENDPOINT_SUFFIX", changed)
			utils.AssertNil(t, client.ModifyRequest(ctx, attributes))
			utils.AssertEqual(t, "ecs.cn-hangzhou."+selected, tea.StringValue(ctx.Request.Headers["host"]))
			if !strings.Contains(tea.StringValue(ctx.Request.Headers["Authorization"]), "/cn-hangzhou/ecs/aliyun_v4_request") {
				t.Fatal("signing region does not match selected host")
			}
		}
	}
}
