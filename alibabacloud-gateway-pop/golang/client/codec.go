package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/tjfoc/gmsm/sm3"
)

func hexEncode(raw []byte) *string {
	return tea.String(hex.EncodeToString(raw))
}

func hashBytes(raw []byte, signatureAlgorithm *string) []byte {
	signType := tea.StringValue(signatureAlgorithm)
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
	uri := url.QueryEscape(tea.StringValue(raw))
	uri = strings.Replace(uri, "+", "%20", -1)
	uri = strings.Replace(uri, "*", "%2A", -1)
	uri = strings.Replace(uri, "%7E", "~", -1)
	return tea.String(uri)
}

func hmacSHA256(stringToSign *string, secret *string) []byte {
	return hmacSHA256ByBytes(stringToSign, []byte(tea.StringValue(secret)))
}

func hmacSHA256ByBytes(stringToSign *string, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	if _, err := h.Write([]byte(tea.StringValue(stringToSign))); err != nil {
		return nil
	}
	return h.Sum(nil)
}

func hmacSM3(stringToSign *string, secret *string) []byte {
	return hmacSM3ByBytes(stringToSign, []byte(tea.StringValue(secret)))
}

func hmacSM3ByBytes(stringToSign *string, secret []byte) []byte {
	h := hmac.New(sm3.New, secret)
	if _, err := h.Write([]byte(tea.StringValue(stringToSign))); err != nil {
		return nil
	}
	return h.Sum(nil)
}

func splitString(raw *string, sep *string, limit *int) []*string {
	if limit == nil || tea.IntValue(limit) == -1 || tea.IntValue(limit) == 0 {
		return tea.StringSlice(strings.Split(tea.StringValue(raw), tea.StringValue(sep)))
	}
	return tea.StringSlice(strings.SplitN(tea.StringValue(raw), tea.StringValue(sep), tea.IntValue(limit)))
}

func replaceString(raw *string, oldStr *string, newStr *string, count *int) *string {
	if count == nil {
		return tea.String(strings.ReplaceAll(tea.StringValue(raw), tea.StringValue(oldStr), tea.StringValue(newStr)))
	}
	return tea.String(strings.Replace(tea.StringValue(raw), tea.StringValue(oldStr), tea.StringValue(newStr), tea.IntValue(count)))
}

func containsString(s *string, substr *string) *bool {
	return tea.Bool(strings.Contains(tea.StringValue(s), tea.StringValue(substr)))
}

func hasPrefix(s *string, prefix *string) *bool {
	return tea.Bool(strings.HasPrefix(tea.StringValue(s), tea.StringValue(prefix)))
}

func toLower(s *string) *string {
	return tea.String(strings.ToLower(tea.StringValue(s)))
}

func subString(s *string, start, end *int) *string {
	endIndex := tea.IntValue(end)
	startIndex := tea.IntValue(start)
	str := tea.StringValue(s)
	if endIndex == -1 {
		return tea.String(str[startIndex : len(str)-1])
	}
	return tea.String(str[startIndex:endIndex])
}

func equalsString(expect *string, actual *string) *bool {
	return tea.Bool(tea.StringValue(expect) == tea.StringValue(actual))
}

func trimString(s *string) *string {
	return tea.String(strings.Trim(tea.StringValue(s), " "))
}

func keySet(raw interface{}) []*string {
	slice := tea.StringSliceValue(nil)
	switch v := raw.(type) {
	case map[string]interface{}:
		for k := range v {
			slice = append(slice, k)
		}
	case map[string]string:
		for k := range v {
			slice = append(slice, k)
		}
	case map[string]*string:
		for k := range v {
			slice = append(slice, k)
		}
	}
	return tea.StringSlice(slice)
}

func join(raw []*string, sep *string) *string {
	res := ""
	separator := ""
	for _, value := range raw {
		res = res + separator + tea.StringValue(value)
		separator = tea.StringValue(sep)
	}
	return tea.String(res)
}

func sizeOf(raw []*string) *int {
	return tea.Int(len(raw))
}

func ascSort(raw []*string) []*string {
	slice := tea.StringSliceValue(raw)
	sort.Strings(slice)
	return tea.StringSlice(slice)
}

func listContains(raw []*string, str *string) *bool {
	for _, v := range raw {
		if tea.StringValue(v) == tea.StringValue(str) {
			return tea.Bool(true)
		}
	}
	return tea.Bool(false)
}
