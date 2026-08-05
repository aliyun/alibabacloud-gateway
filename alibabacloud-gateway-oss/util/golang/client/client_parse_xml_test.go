package client

import (
	"reflect"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
)

func TestParseXml_GetBucketAclReturnsGenericMap(t *testing.T) {
	body := `<?xml version="1.0" encoding="UTF-8"?>
<AccessControlPolicy>
  <Owner>
    <ID>123</ID>
    <DisplayName>owner</DisplayName>
  </Owner>
  <AccessControlList>
    <Grant>private</Grant>
  </AccessControlList>
</AccessControlPolicy>`
	apiName := tea.String("GetBucketAcl")
	result, err := ParseXml(tea.String(body), apiName)
	if err != nil {
		t.Fatalf("ParseXml returned error: %v", err)
	}
	root, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map[string]interface{}, got %T", result)
	}
	policy, ok := root["AccessControlPolicy"].(map[string]interface{})
	if !ok {
		t.Fatalf("AccessControlPolicy should be map[string]interface{}, got %T (%#v)", root["AccessControlPolicy"], root["AccessControlPolicy"])
	}
	acl, ok := policy["AccessControlList"].(map[string]interface{})
	if !ok {
		t.Fatalf("AccessControlList should be map[string]interface{}, got %T", policy["AccessControlList"])
	}
	if grant, _ := acl["Grant"].(string); grant != "private" {
		t.Fatalf("expected Grant=private, got %#v", acl["Grant"])
	}
}

func TestParseXml_NilApiNameUsesUntypedParse(t *testing.T) {
	body := `<AccessControlPolicy><AccessControlList><Grant>public-read</Grant></AccessControlList></AccessControlPolicy>`
	result, err := ParseXml(tea.String(body), nil)
	if err != nil {
		t.Fatalf("ParseXml returned error: %v", err)
	}
	root, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map[string]interface{}, got %T", result)
	}
	policy, ok := root["AccessControlPolicy"].(map[string]interface{})
	if !ok {
		t.Fatalf("AccessControlPolicy should be nested map, got %T", root["AccessControlPolicy"])
	}
	acl, ok := policy["AccessControlList"].(map[string]interface{})
	if !ok {
		t.Fatalf("AccessControlList should be nested map, got %T", policy["AccessControlList"])
	}
	if grant, _ := acl["Grant"].(string); grant != "public-read" {
		t.Fatalf("expected Grant=public-read, got %#v", acl["Grant"])
	}
}

func TestParseXml_UnknownActionUsesUntypedParse(t *testing.T) {
	body := `<Foo><Bar>baz</Bar></Foo>`
	result, err := ParseXml(tea.String(body), tea.String("NotRegisteredAction"))
	if err != nil {
		t.Fatalf("ParseXml returned error: %v", err)
	}
	root, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map[string]interface{}, got %T", result)
	}
	foo, ok := root["Foo"].(map[string]interface{})
	if !ok {
		t.Fatalf("Foo should be nested map, got %T", root["Foo"])
	}
	if bar, _ := foo["Bar"].(string); bar != "baz" {
		t.Fatalf("expected Bar=baz, got %#v", foo["Bar"])
	}
}

func TestNormalizeParsedXml_NilAndAlreadyGeneric(t *testing.T) {
	if got := normalizeParsedXml(nil); got != nil {
		t.Fatalf("nil input should stay nil, got %#v", got)
	}
	in := map[string]interface{}{
		"AccessControlPolicy": map[string]interface{}{
			"AccessControlList": map[string]interface{}{"Grant": "private"},
		},
	}
	out := normalizeParsedXml(in)
	if !reflect.DeepEqual(out["AccessControlPolicy"], in["AccessControlPolicy"]) {
		t.Fatalf("generic map should round-trip, got %#v", out)
	}
}

func TestNormalizeParsedXml_FlattensTypedStruct(t *testing.T) {
	type nested struct {
		Grant string `json:"Grant"`
	}
	type policy struct {
		AccessControlList nested `json:"AccessControlList"`
	}
	in := map[string]interface{}{
		"AccessControlPolicy": &policy{AccessControlList: nested{Grant: "private"}},
	}
	out := normalizeParsedXml(in)
	policyMap, ok := out["AccessControlPolicy"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected nested map after normalize, got %T", out["AccessControlPolicy"])
	}
	acl, ok := policyMap["AccessControlList"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected AccessControlList map, got %T", policyMap["AccessControlList"])
	}
	if grant, _ := acl["Grant"].(string); grant != "private" {
		t.Fatalf("expected Grant=private, got %#v", acl["Grant"])
	}
}

func TestNormalizeParsedXml_MarshalErrorFallsBack(t *testing.T) {
	in := map[string]interface{}{"ch": make(chan int)}
	out := normalizeParsedXml(in)
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("marshal failure should return original map")
	}
}

func TestNormalizeParsedXml_UnmarshalErrorFallsBack(t *testing.T) {
	origMarshal := jsonMarshal
	origUnmarshal := jsonUnmarshal
	defer func() {
		jsonMarshal = origMarshal
		jsonUnmarshal = origUnmarshal
	}()
	jsonMarshal = func(v interface{}) ([]byte, error) {
		return []byte(`{"ok":true}`), nil
	}
	jsonUnmarshal = func(data []byte, v interface{}) error {
		return errUnmarshalBoom
	}
	in := map[string]interface{}{"ok": true}
	out := normalizeParsedXml(in)
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("unmarshal failure should return original map")
	}
}

type boomError string

func (e boomError) Error() string { return string(e) }

const errUnmarshalBoom = boomError("boom")
