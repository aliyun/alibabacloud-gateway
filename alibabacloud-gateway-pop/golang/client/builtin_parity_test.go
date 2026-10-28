package client

import (
	"net/url"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/alibabacloud-go/tea/dara"
)

func TestBuiltinParity(t *testing.T) {
	samples := []*string{nil, dara.String(""), dara.String("a b*~"), dara.String("中文"), dara.String("a+b"), dara.String("100%")}
	for _, sample := range samples {
		got := dara.StringValue(percentEncode(sample))
		want := legacyPercentEncode(dara.StringValue(sample))
		if got != want || got != dara.PercentEncode(dara.StringValue(sample)) {
			t.Fatalf("percentEncode(%q) = %q, legacy %q, dara %q", dara.StringValue(sample), got, want, dara.PercentEncode(dara.StringValue(sample)))
		}
	}

	for _, raw := range []interface{}{
		map[string]interface{}{"b": "1", "a": "2"},
		map[string]string{"b": "1", "a": "2"},
		map[string]*string{"b": dara.String("1"), "a": dara.String("2")},
	} {
		got := stringValues(keySet(raw))
		sort.Strings(got)
		want := append([]string{}, dara.KeySet(raw)...)
		sort.Strings(want)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("keySet(%#v) = %v, dara %v", raw, got, want)
		}
	}
	if keySet("not-a-map") != nil {
		t.Fatal("keySet should ignore values that are not string-keyed maps")
	}

	sortedInput := []*string{dara.String("b"), dara.String("a"), dara.String("")}
	gotSort := stringValues(ascSort(sortedInput))
	daraSort := stringValues(dara.SortArr(sortedInput, "asc").([]*string))
	if !reflect.DeepEqual(gotSort, []string{"", "a", "b"}) || !reflect.DeepEqual(gotSort, daraSort) {
		t.Fatalf("ascSort = %v, dara %v", gotSort, daraSort)
	}
	if !dara.BoolValue(listContains([]*string{nil}, nil)) {
		t.Fatal("listContains should match nil values")
	}

	same := dara.String("x-acs-foo")
	other := dara.String("x-acs-foo")
	if !dara.BoolValue(listContains([]*string{same}, other)) {
		t.Fatal("listContains should compare string values")
	}
	if dara.ArrContains([]*string{same}, other) {
		t.Fatal("dara.ArrContains compares pointer identity for *string and must not replace listContains")
	}
	joined := dara.StringValue(join([]*string{dara.String("a"), dara.String(""), dara.String("b")}, dara.String(",")))
	if joined != "a,,b" {
		t.Fatalf("join = %q", joined)
	}
	if dara.ArrJoin([]*string{dara.String("a"), dara.String(""), dara.String("b")}, ",") == joined {
		t.Fatal("dara.ArrJoin skips empty values and must not replace join")
	}

	form := map[string]interface{}{
		"b": "x y",
		"a": []interface{}{"m", "n"},
		"c": map[string]interface{}{"k": "v"},
	}
	if dara.StringValue(toForm(form)) != "a.1=m&a.2=n&b=x+y&c.k=v" {
		t.Fatalf("toForm = %q", dara.StringValue(toForm(form)))
	}

	regional, err := getEndpointRules(dara.String("Ecs"), dara.String("cn-hangzhou"), dara.String("regional"), dara.String(""), dara.String(""))
	if err != nil || dara.StringValue(regional) != "ecs.cn-hangzhou.aliyuncs.com" {
		t.Fatalf("regional endpoint %q %v", dara.StringValue(regional), err)
	}
	vpc, err := getEndpointRules(dara.String("Ecs"), dara.String("cn-hangzhou"), dara.String("regional"), dara.String("vpc"), dara.String("inner"))
	if err != nil || dara.StringValue(vpc) != "ecs-inner-vpc.cn-hangzhou.aliyuncs.com" {
		t.Fatalf("vpc endpoint %q %v", dara.StringValue(vpc), err)
	}
	central, err := getEndpointRules(dara.String("Ecs"), dara.String(""), dara.String("central"), dara.String("public"), nil)
	if err != nil || dara.StringValue(central) != "ecs.aliyuncs.com" {
		t.Fatalf("central endpoint %q %v", dara.StringValue(central), err)
	}
	_, err = getEndpointRules(dara.String("Ecs"), dara.String(""), dara.String("regional"), nil, nil)
	if err == nil || !strings.Contains(err.Error(), "RegionId is empty") {
		t.Fatalf("regional empty region error = %v", err)
	}

	if dara.BoolValue(isUnset(nil)) != dara.IsNil(nil) || dara.BoolValue(isUnset(dara.String(""))) != dara.IsNil(dara.String("")) {
		t.Fatal("isUnset does not match dara.IsNil")
	}
	if !strings.HasSuffix(dara.StringValue(getTimestamp()), "Z") || len(dara.StringValue(getTimestamp())) != 20 {
		t.Fatalf("timestamp %q", dara.StringValue(getTimestamp()))
	}
	if dara.StringValue(toJSONString(map[string]interface{}{"a": "<b>"})) != `{"a":"<b>"}` {
		t.Fatalf("json %q", dara.StringValue(toJSONString(map[string]interface{}{"a": "<b>"})))
	}
}

func legacyPercentEncode(raw string) string {
	uri := url.QueryEscape(raw)
	uri = strings.Replace(uri, "+", "%20", -1)
	uri = strings.Replace(uri, "*", "%2A", -1)
	uri = strings.Replace(uri, "%7E", "~", -1)
	return uri
}

func stringValues(values []*string) []string {
	out := make([]string, len(values))
	for i, value := range values {
		out[i] = dara.StringValue(value)
	}
	return out
}
