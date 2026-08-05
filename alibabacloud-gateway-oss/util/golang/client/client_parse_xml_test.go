package client

import (
	"reflect"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
)

// Customer reproduction: Terraform jsonpath $.AccessControlPolicy.AccessControlList
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
		t.Fatalf("expected Grant=private, got %#v (%T)", acl["Grant"], acl["Grant"])
	}
	owner, ok := policy["Owner"].(map[string]interface{})
	if !ok {
		t.Fatalf("Owner should be map[string]interface{}, got %T", policy["Owner"])
	}
	if id, _ := owner["ID"].(string); id != "123" {
		t.Fatalf("expected Owner.ID=123, got %#v", owner["ID"])
	}
}

func TestParseXml_ListObjectsMultipleContents(t *testing.T) {
	body := `<?xml version="1.0" encoding="UTF-8"?>
<ListBucketResult>
  <Name>example-bucket</Name>
  <Prefix></Prefix>
  <Marker></Marker>
  <MaxKeys>100</MaxKeys>
  <IsTruncated>false</IsTruncated>
  <Contents>
    <Key>a.txt</Key>
    <Size>1</Size>
  </Contents>
  <Contents>
    <Key>b.txt</Key>
    <Size>2</Size>
  </Contents>
</ListBucketResult>`
	result, err := ParseXml(tea.String(body), tea.String("ListObjects"))
	if err != nil {
		t.Fatalf("ParseXml returned error: %v", err)
	}
	root := result.(map[string]interface{})
	bucket, ok := root["ListBucketResult"].(map[string]interface{})
	if !ok {
		t.Fatalf("ListBucketResult should be map, got %T (%#v)", root["ListBucketResult"], root["ListBucketResult"])
	}
	contents, ok := bucket["Contents"].([]interface{})
	if !ok {
		t.Fatalf("Contents should be []interface{}, got %T (%#v)", bucket["Contents"], bucket["Contents"])
	}
	if len(contents) != 2 {
		t.Fatalf("expected 2 Contents, got %d", len(contents))
	}
	first, ok := contents[0].(map[string]interface{})
	if !ok {
		t.Fatalf("Contents[0] should be map, got %T", contents[0])
	}
	if key, _ := first["Key"].(string); key != "a.txt" {
		t.Fatalf("expected Key=a.txt, got %#v", first["Key"])
	}
}

func TestParseXml_ListObjectsSingleContentsStaysSlice(t *testing.T) {
	// typeRegistry value: a single repeated XML node must still be []interface{}
	body := `<?xml version="1.0" encoding="UTF-8"?>
<ListBucketResult>
  <Name>example-bucket</Name>
  <MaxKeys>20</MaxKeys>
  <IsTruncated>true</IsTruncated>
  <Contents>
    <Key>only.txt</Key>
  </Contents>
</ListBucketResult>`
	result, err := ParseXml(tea.String(body), tea.String("ListObjects"))
	if err != nil {
		t.Fatalf("ParseXml returned error: %v", err)
	}
	root := result.(map[string]interface{})
	bucket := root["ListBucketResult"].(map[string]interface{})
	contents, ok := bucket["Contents"].([]interface{})
	if !ok {
		t.Fatalf("single Contents must stay []interface{}, got %T (%#v)", bucket["Contents"], bucket["Contents"])
	}
	if len(contents) != 1 {
		t.Fatalf("expected 1 Contents element, got %d", len(contents))
	}
	item := contents[0].(map[string]interface{})
	if key, _ := item["Key"].(string); key != "only.txt" {
		t.Fatalf("expected Key=only.txt, got %#v", item["Key"])
	}
}

func TestParseXml_ListObjectsPreservesScalarKinds(t *testing.T) {
	body := `<?xml version="1.0" encoding="UTF-8"?>
<ListBucketResult>
  <Name>example-bucket</Name>
  <MaxKeys>20</MaxKeys>
  <IsTruncated>true</IsTruncated>
</ListBucketResult>`
	result, err := ParseXml(tea.String(body), tea.String("ListObjects"))
	if err != nil {
		t.Fatalf("ParseXml returned error: %v", err)
	}
	bucket := result.(map[string]interface{})["ListBucketResult"].(map[string]interface{})

	maxKeys, ok := bucket["MaxKeys"].(int32)
	if !ok {
		t.Fatalf("MaxKeys should be int32 (not float64), got %T (%#v)", bucket["MaxKeys"], bucket["MaxKeys"])
	}
	if maxKeys != 20 {
		t.Fatalf("expected MaxKeys=20, got %v", maxKeys)
	}

	truncated, ok := bucket["IsTruncated"].(bool)
	if !ok {
		t.Fatalf("IsTruncated should be bool, got %T (%#v)", bucket["IsTruncated"], bucket["IsTruncated"])
	}
	if !truncated {
		t.Fatalf("expected IsTruncated=true")
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
			"nestedSlice":       []interface{}{map[string]interface{}{"K": "v"}},
		},
	}
	out := normalizeParsedXml(in)
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("generic map should round-trip, got %#v", out)
	}
}

func TestNormalizeParsedXml_FlattensTypedStructPreservingScalars(t *testing.T) {
	type nested struct {
		Grant   string `json:"Grant"`
		MaxKeys *int32 `json:"MaxKeys"`
		Flag    *bool  `json:"Flag"`
	}
	maxKeys := int32(7)
	flag := true
	in := map[string]interface{}{
		"AccessControlPolicy": &nested{Grant: "private", MaxKeys: &maxKeys, Flag: &flag},
	}
	out := normalizeParsedXml(in)
	policyMap, ok := out["AccessControlPolicy"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected nested map after normalize, got %T", out["AccessControlPolicy"])
	}
	if grant, _ := policyMap["Grant"].(string); grant != "private" {
		t.Fatalf("expected Grant=private, got %#v", policyMap["Grant"])
	}
	if v, ok := policyMap["MaxKeys"].(int32); !ok || v != 7 {
		t.Fatalf("expected MaxKeys int32(7), got %#v (%T)", policyMap["MaxKeys"], policyMap["MaxKeys"])
	}
	if v, ok := policyMap["Flag"].(bool); !ok || !v {
		t.Fatalf("expected Flag bool(true), got %#v (%T)", policyMap["Flag"], policyMap["Flag"])
	}
}

func TestNormalizeParsedXml_DoublePointerStruct(t *testing.T) {
	// tea-xml returns **T for typed roots
	type leaf struct {
		Grant string `json:"Grant"`
	}
	inner := &leaf{Grant: "private"}
	in := map[string]interface{}{"AccessControlPolicy": &inner}
	out := normalizeParsedXml(in)
	policyMap, ok := out["AccessControlPolicy"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected map from **struct, got %T (%#v)", out["AccessControlPolicy"], out["AccessControlPolicy"])
	}
	if grant, _ := policyMap["Grant"].(string); grant != "private" {
		t.Fatalf("expected Grant=private, got %#v", policyMap["Grant"])
	}
}

func TestNormalizeParsedXml_NilPointerAndScalarPointer(t *testing.T) {
	var nilPtr *string
	s := "x"
	in := map[string]interface{}{
		"nilField":  nilPtr,
		"strField":  &s,
		"plainInt":  int64(9),
		"nilRoot":   nil,
		"structVal": struct{ A string }{A: "b"},
	}
	out := normalizeParsedXml(in)
	if out["nilField"] != nil {
		t.Fatalf("nil *string should become nil, got %#v", out["nilField"])
	}
	if out["nilRoot"] != nil {
		t.Fatalf("nil value should stay nil")
	}
	if v, ok := out["strField"].(string); !ok || v != "x" {
		t.Fatalf("expected *string → string, got %#v (%T)", out["strField"], out["strField"])
	}
	if v, ok := out["plainInt"].(int64); !ok || v != 9 {
		t.Fatalf("expected int64 preserved, got %#v (%T)", out["plainInt"], out["plainInt"])
	}
	m, ok := out["structVal"].(map[string]interface{})
	if !ok || m["A"] != "b" {
		t.Fatalf("expected value struct → map, got %#v", out["structVal"])
	}
}

func TestNormalizeParsedXml_SliceOfStructs(t *testing.T) {
	type item struct {
		Key string `json:"Key"`
	}
	in := map[string]interface{}{
		"Contents": []*item{{Key: "a"}, {Key: "b"}},
	}
	out := normalizeParsedXml(in)
	contents, ok := out["Contents"].([]interface{})
	if !ok {
		t.Fatalf("expected []interface{}, got %T", out["Contents"])
	}
	if len(contents) != 2 {
		t.Fatalf("expected 2 items, got %d", len(contents))
	}
	first := contents[0].(map[string]interface{})
	if first["Key"] != "a" {
		t.Fatalf("expected Key=a, got %#v", first["Key"])
	}
}

type stubToMap struct {
	Grant string `json:"Grant"`
}

func (s *stubToMap) ToMap() map[string]interface{} {
	return map[string]interface{}{"Grant": s.Grant, "via": "ToMap"}
}

func TestNormalizeParsedXml_UsesToMapWhenAvailable(t *testing.T) {
	in := map[string]interface{}{
		"AccessControlPolicy": &stubToMap{Grant: "private"},
	}
	out := normalizeParsedXml(in)
	m := out["AccessControlPolicy"].(map[string]interface{})
	if m["via"] != "ToMap" || m["Grant"] != "private" {
		t.Fatalf("expected ToMap path, got %#v", m)
	}

	in2 := map[string]interface{}{
		"list": []interface{}{&stubToMap{Grant: "x"}},
	}
	out2 := normalizeParsedXml(in2)
	list := out2["list"].([]interface{})
	item := list[0].(map[string]interface{})
	if item["via"] != "ToMap" {
		t.Fatalf("expected ToMap in slice element, got %#v", item)
	}

	// value-receiver ToMap on non-pointer
	in3 := map[string]interface{}{"p": stubToMapVal{Grant: "v"}}
	out3 := normalizeParsedXml(in3)
	m3 := out3["p"].(map[string]interface{})
	if m3["via"] != "ToMapVal" {
		t.Fatalf("expected value-receiver ToMap, got %#v", m3)
	}
}

type stubToMapVal struct {
	Grant string `json:"Grant"`
}

func (s stubToMapVal) ToMap() map[string]interface{} {
	return map[string]interface{}{"Grant": s.Grant, "via": "ToMapVal"}
}

func TestReflectToGeneric_EdgeCases(t *testing.T) {
	if got := reflectToGeneric(reflect.Value{}); got != nil {
		t.Fatalf("invalid Value should be nil, got %#v", got)
	}

	var iface interface{}
	ifaceRV := reflect.ValueOf(&iface).Elem()
	if got := reflectToGeneric(ifaceRV); got != nil {
		t.Fatalf("nil interface should be nil, got %#v", got)
	}

	// non-string map keys are returned as-is
	intMap := map[int]string{1: "a"}
	got := reflectToGeneric(reflect.ValueOf(intMap))
	if _, ok := got.(map[int]string); !ok {
		t.Fatalf("expected map[int]string passthrough, got %T", got)
	}

	// typed nil pointer after interface box
	var p *int
	var boxed interface{} = p
	if got := toGenericValue(boxed); got != nil {
		t.Fatalf("boxed nil *int should be nil, got %#v", got)
	}
}

func TestStructToGenericMap_SkipsUnexportedAndDashTag(t *testing.T) {
	type sample struct {
		Export string `json:"Export"`
		hidden string // unexported
		Skip   string `json:"-"`
		Empty  string `json:""` // empty name → field name
	}
	out := structToGenericMap(reflect.ValueOf(sample{
		Export: "e",
		hidden: "h",
		Skip:   "s",
		Empty:  "x",
	}))
	if out["Export"] != "e" {
		t.Fatalf("expected Export=e, got %#v", out)
	}
	if _, ok := out["hidden"]; ok {
		t.Fatalf("unexported field must be skipped")
	}
	if _, ok := out["Skip"]; ok {
		t.Fatalf("json:\"-\" field must be skipped")
	}
	if out["Empty"] != "x" {
		t.Fatalf("empty json name should fall back to field name, got %#v", out)
	}
}
