// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJSONInput interface {
	dara.Model
	String() string
	GoString() string
	SetParseJsonNumberAsString(v bool) *JSONInput
	GetParseJsonNumberAsString() *bool
	SetRange(v string) *JSONInput
	GetRange() *string
	SetType(v string) *JSONInput
	GetType() *string
}

type JSONInput struct {
	ParseJsonNumberAsString *bool   `json:"ParseJsonNumberAsString,omitempty" xml:"ParseJsonNumberAsString,omitempty"`
	Range                   *string `json:"Range,omitempty" xml:"Range,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s JSONInput) String() string {
	return dara.Prettify(s)
}

func (s JSONInput) GoString() string {
	return s.String()
}

func (s *JSONInput) GetParseJsonNumberAsString() *bool {
	return s.ParseJsonNumberAsString
}

func (s *JSONInput) GetRange() *string {
	return s.Range
}

func (s *JSONInput) GetType() *string {
	return s.Type
}

func (s *JSONInput) SetParseJsonNumberAsString(v bool) *JSONInput {
	s.ParseJsonNumberAsString = &v
	return s
}

func (s *JSONInput) SetRange(v string) *JSONInput {
	s.Range = &v
	return s
}

func (s *JSONInput) SetType(v string) *JSONInput {
	s.Type = &v
	return s
}

func (s *JSONInput) Validate() error {
	return dara.Validate(s)
}
