package client

import (
	"net/http"
	"regexp"
	"testing"
	"time"

	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func Test_GetDateISO8601FromUTC(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		rfc      string
		expected string
	}{
		{"Thu, 06 Feb 2020 07:32:54 GMT", "20200206T073254Z"},
		{"Tue, 03 Jun 2025 11:23:48 GMT", "20250603T112348Z"},
		{"Wed, 03 Dec 2025 11:23:48 GMT", "20251203T112348Z"},
		{"Thu, 01 Jan 2026 00:00:00 GMT", "20260101T000000Z"},
		{"Thu, 31 Dec 2026 23:59:59 GMT", "20261231T235959Z"},
	}
	for _, c := range cases {
		actual, err := client.getDateISO8601FromUTC(tea.String(c.rfc))
		if err != nil {
			t.Fatalf("rfc=%s err=%v", c.rfc, err)
		}
		if tea.StringValue(actual) != c.expected {
			t.Fatalf("rfc=%s expected=%s actual=%s", c.rfc, c.expected, tea.StringValue(actual))
		}
	}
}

func Test_GetDateISO8601FromUTCRoundTrip(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	rfc := tea.StringValue(util.GetDateUTCString())
	matched, err := regexp.MatchString(`^[A-Z][a-z]{2}, \d{2} [A-Z][a-z]{2} \d{4} \d{2}:\d{2}:\d{2} GMT$`, rfc)
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Fatalf("unexpected Date header format: %s", rfc)
	}

	parsed, err := time.Parse(http.TimeFormat, rfc)
	if err != nil {
		t.Fatalf("parse rfc=%s err=%v", rfc, err)
	}
	expected := parsed.UTC().Format("20060102T150405Z")

	actual, err := client.getDateISO8601FromUTC(tea.String(rfc))
	if err != nil {
		t.Fatal(err)
	}
	if tea.StringValue(actual) != expected {
		t.Fatalf("expected=%s actual=%s rfc=%s", expected, tea.StringValue(actual), rfc)
	}
	isoMatched, err := regexp.MatchString(`^\d{8}T\d{6}Z$`, tea.StringValue(actual))
	if err != nil {
		t.Fatal(err)
	}
	if !isoMatched {
		t.Fatalf("unexpected iso8601 format: %s", tea.StringValue(actual))
	}
}
