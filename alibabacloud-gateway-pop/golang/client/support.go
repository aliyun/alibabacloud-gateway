package client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/alibabacloud-go/tea/dara"
)

var processStartTime int64 = time.Now().UnixNano() / 1e6
var seqId int64

func isUnset(val interface{}) *bool {
	return dara.Bool(dara.IsNil(val))
}

func emptyString(val *string) *bool {
	return dara.Bool(val == nil || dara.StringValue(val) == "")
}

func equalString(val1, val2 *string) *bool {
	return dara.Bool(dara.StringValue(val1) == dara.StringValue(val2))
}

func equalNumber(val1, val2 *int) *bool {
	return dara.Bool(dara.IntValue(val1) == dara.IntValue(val2))
}

func defaultString(realStr, defaultStr *string) *string {
	if realStr == nil {
		return defaultStr
	}
	return realStr
}

func toBytes(a *string) []byte {
	return []byte(dara.StringValue(a))
}

func readAsString(body io.Reader) (*string, error) {
	value, err := dara.ReadAsString(body)
	if err != nil {
		return dara.String(""), err
	}
	return dara.String(value), nil
}

func assertAsBytes(a interface{}) ([]byte, error) {
	res, ok := a.([]byte)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%v is not a []byte", a))
	}
	return res, nil
}

func assertAsMap(a interface{}) (map[string]interface{}, error) {
	r := reflect.ValueOf(a)
	if r.Kind().String() != "map" {
		return nil, errors.New(fmt.Sprintf("%v is not a map[string]interface{}", a))
	}
	res := make(map[string]interface{})
	for _, key := range r.MapKeys() {
		res[key.String()] = r.MapIndex(key).Interface()
	}
	return res, nil
}

func toJSONString(a interface{}) *string {
	switch v := a.(type) {
	case *string:
		return v
	case string:
		return dara.String(v)
	case []byte:
		return dara.String(string(v))
	case io.Reader:
		byt, err := ioutil.ReadAll(v)
		if err != nil {
			return nil
		}
		return dara.String(string(byt))
	}
	byt := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(byt)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(a); err != nil {
		return nil
	}
	return dara.String(string(bytes.TrimSpace(byt.Bytes())))
}

func getNonce() *string {
	routineID := getGID()
	currentTime := time.Now().UnixNano() / 1e6
	seq := atomic.AddInt64(&seqId, 1)
	randNum := rand.Int63()
	msg := fmt.Sprintf("%d-%d-%d-%d-%d", processStartTime, routineID, currentTime, seq, randNum)
	h := md5.New()
	h.Write([]byte(msg))
	ret := hex.EncodeToString(h.Sum(nil))
	return &ret
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func is4xx(code *int) *bool {
	tmp := dara.IntValue(code)
	return dara.Bool(tmp >= 400 && tmp < 500)
}

func is5xx(code *int) *bool {
	tmp := dara.IntValue(code)
	return dara.Bool(tmp >= 500 && tmp < 600)
}

func getTimestamp() *string {
	gmt := time.FixedZone("GMT", 0)
	return dara.String(time.Now().In(gmt).Format("2006-01-02T15:04:05Z"))
}

func getEndpointRules(product, regionId, endpointType, network, suffix *string) (*string, error) {
	var result *string
	if dara.StringValue(endpointType) == "regional" {
		if dara.StringValue(regionId) == "" {
			return dara.String(""), fmt.Errorf("RegionId is empty, please set a valid RegionId")
		}
		result = dara.String(strings.Replace("<product><suffix><network>.<region_id>.aliyuncs.com",
			"<region_id>", dara.StringValue(regionId), 1))
	} else {
		result = dara.String("<product><suffix><network>.aliyuncs.com")
	}
	result = dara.String(strings.Replace(dara.StringValue(result),
		"<product>", strings.ToLower(dara.StringValue(product)), 1))
	if dara.StringValue(network) == "" || dara.StringValue(network) == "public" {
		result = dara.String(strings.Replace(dara.StringValue(result), "<network>", "", 1))
	} else {
		result = dara.String(strings.Replace(dara.StringValue(result),
			"<network>", "-"+dara.StringValue(network), 1))
	}
	if dara.StringValue(suffix) == "" {
		result = dara.String(strings.Replace(dara.StringValue(result), "<suffix>", "", 1))
	} else {
		result = dara.String(strings.Replace(dara.StringValue(result),
			"<suffix>", "-"+dara.StringValue(suffix), 1))
	}
	return result, nil
}

func toForm(filter map[string]interface{}) *string {
	tmp := make(map[string]interface{})
	byt, _ := json.Marshal(filter)
	decoder := json.NewDecoder(bytes.NewReader(byt))
	decoder.UseNumber()
	_ = decoder.Decode(&tmp)

	result := make(map[string]*string)
	for key, value := range tmp {
		flatRepeatedList(reflect.ValueOf(value), result, key)
	}
	return dara.String(dara.ToFormString(anyifyMapValue(result)))
}

func anyifyMapValue(a map[string]*string) map[string]interface{} {
	res := make(map[string]interface{})
	for key, value := range a {
		res[key] = dara.StringValue(value)
	}
	return res
}

func flatRepeatedList(dataValue reflect.Value, result map[string]*string, prefix string) {
	if !dataValue.IsValid() {
		return
	}
	dataType := dataValue.Type()
	if dataType.Kind().String() == "slice" {
		handleRepeatedParams(dataValue, result, prefix)
	} else if dataType.Kind().String() == "map" {
		handleFormMap(dataValue, result, prefix)
	} else {
		result[prefix] = dara.String(fmt.Sprintf("%v", dataValue.Interface()))
	}
}

func handleRepeatedParams(repeatedFieldValue reflect.Value, result map[string]*string, prefix string) {
	if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
		for m := 0; m < repeatedFieldValue.Len(); m++ {
			elementValue := repeatedFieldValue.Index(m)
			key := prefix + "." + strconv.Itoa(m+1)
			fieldValue := reflect.ValueOf(elementValue.Interface())
			if fieldValue.Kind().String() == "map" {
				handleFormMap(fieldValue, result, key)
			} else {
				result[key] = dara.String(fmt.Sprintf("%v", fieldValue.Interface()))
			}
		}
	}
}

func handleFormMap(valueField reflect.Value, result map[string]*string, prefix string) {
	if valueField.IsValid() && valueField.String() != "" {
		if valueField.Type().Kind().String() == "map" {
			byt, _ := json.Marshal(valueField.Interface())
			cache := make(map[string]interface{})
			decoder := json.NewDecoder(bytes.NewReader(byt))
			decoder.UseNumber()
			_ = decoder.Decode(&cache)
			for key, value := range cache {
				pre := key
				if prefix != "" {
					pre = prefix + "." + key
				}
				flatRepeatedList(reflect.ValueOf(value), result, pre)
			}
		}
	}
}
