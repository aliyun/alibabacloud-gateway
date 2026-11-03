package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/alibabacloud-go/tea/dara"
	"github.com/tjfoc/gmsm/sm3"
)

func hexEncode(raw []byte) *string {
	return dara.String(hex.EncodeToString(raw))
}

func hashBytes(raw []byte, signatureAlgorithm *string) []byte {
	signType := dara.StringValue(signatureAlgorithm)
	if strings.Contains(signType, "HMAC-SHA256") || strings.Contains(signType, "RSA-SHA256") {
		h := sha256.New()
		h.Write(raw)
		return h.Sum(nil)
	} else if strings.Contains(signType, "HMAC-SM3") {
		h := sm3.New()
		h.Write(raw)
		return h.Sum(nil)
	}
	return nil
}

func percentEncode(raw *string) *string {
	return dara.String(dara.PercentEncode(dara.StringValue(raw)))
}

func hmacSHA256(stringToSign *string, secret *string) []byte {
	return hmacSHA256ByBytes(stringToSign, []byte(dara.StringValue(secret)))
}

func hmacSHA256ByBytes(stringToSign *string, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	if _, err := h.Write([]byte(dara.StringValue(stringToSign))); err != nil {
		return nil
	}
	return h.Sum(nil)
}

func hmacSM3(stringToSign *string, secret *string) []byte {
	return hmacSM3ByBytes(stringToSign, []byte(dara.StringValue(secret)))
}

func hmacSM3ByBytes(stringToSign *string, secret []byte) []byte {
	h := hmac.New(sm3.New, secret)
	if _, err := h.Write([]byte(dara.StringValue(stringToSign))); err != nil {
		return nil
	}
	return h.Sum(nil)
}

func splitString(raw *string, sep *string, limit *int) []*string {
	if limit == nil || dara.IntValue(limit) == -1 || dara.IntValue(limit) == 0 {
		return dara.StringSlice(strings.Split(dara.StringValue(raw), dara.StringValue(sep)))
	}
	return dara.StringSlice(strings.SplitN(dara.StringValue(raw), dara.StringValue(sep), dara.IntValue(limit)))
}

func replaceString(raw *string, oldStr *string, newStr *string, count *int) *string {
	if count == nil {
		return dara.String(strings.ReplaceAll(dara.StringValue(raw), dara.StringValue(oldStr), dara.StringValue(newStr)))
	}
	return dara.String(strings.Replace(dara.StringValue(raw), dara.StringValue(oldStr), dara.StringValue(newStr), dara.IntValue(count)))
}

func containsString(s *string, substr *string) *bool {
	return dara.Bool(strings.Contains(dara.StringValue(s), dara.StringValue(substr)))
}

func hasPrefix(s *string, prefix *string) *bool {
	return dara.Bool(strings.HasPrefix(dara.StringValue(s), dara.StringValue(prefix)))
}

func toLower(s *string) *string {
	return dara.String(strings.ToLower(dara.StringValue(s)))
}

func subString(s *string, start, end *int) *string {
	endIndex := dara.IntValue(end)
	startIndex := dara.IntValue(start)
	str := dara.StringValue(s)
	if endIndex == -1 {
		return dara.String(str[startIndex : len(str)-1])
	}
	return dara.String(str[startIndex:endIndex])
}

func equalsString(expect *string, actual *string) *bool {
	return dara.Bool(dara.StringValue(expect) == dara.StringValue(actual))
}

func trimString(s *string) *string {
	return dara.String(strings.Trim(dara.StringValue(s), " "))
}

func keySet(raw interface{}) []*string {
	switch raw.(type) {
	case map[string]interface{}, map[string]string, map[string]*string:
		return dara.StringSlice(dara.KeySet(raw))
	default:
		return dara.StringSlice(nil)
	}
}

func join(raw []*string, sep *string) *string {
	res := ""
	separator := ""
	for _, value := range raw {
		res = res + separator + dara.StringValue(value)
		separator = dara.StringValue(sep)
	}
	return dara.String(res)
}

func sizeOf(raw []*string) *int {
	return dara.Int(len(raw))
}

func ascSort(raw []*string) []*string {
	slice := dara.StringSliceValue(raw)
	sort.Strings(slice)
	return dara.StringSlice(slice)
}

func listContains(raw []*string, str *string) *bool {
	for _, v := range raw {
		if dara.StringValue(v) == dara.StringValue(str) {
			return dara.Bool(true)
		}
	}
	return dara.Bool(false)
}
