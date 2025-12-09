// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInputSerialization interface {
	dara.Model
	String() string
	GoString() string
	SetCsv(v *CSVInput) *InputSerialization
	GetCsv() *CSVInput
	SetCompressionType(v string) *InputSerialization
	GetCompressionType() *string
	SetJson(v *JSONInput) *InputSerialization
	GetJson() *JSONInput
}

type InputSerialization struct {
	Csv             *CSVInput  `json:"CSV,omitempty" xml:"CSV,omitempty"`
	CompressionType *string    `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	Json            *JSONInput `json:"JSON,omitempty" xml:"JSON,omitempty"`
}

func (s InputSerialization) String() string {
	return dara.Prettify(s)
}

func (s InputSerialization) GoString() string {
	return s.String()
}

func (s *InputSerialization) GetCsv() *CSVInput {
	return s.Csv
}

func (s *InputSerialization) GetCompressionType() *string {
	return s.CompressionType
}

func (s *InputSerialization) GetJson() *JSONInput {
	return s.Json
}

func (s *InputSerialization) SetCsv(v *CSVInput) *InputSerialization {
	s.Csv = v
	return s
}

func (s *InputSerialization) SetCompressionType(v string) *InputSerialization {
	s.CompressionType = &v
	return s
}

func (s *InputSerialization) SetJson(v *JSONInput) *InputSerialization {
	s.Json = v
	return s
}

func (s *InputSerialization) Validate() error {
	if s.Csv != nil {
		if err := s.Csv.Validate(); err != nil {
			return err
		}
	}
	if s.Json != nil {
		if err := s.Json.Validate(); err != nil {
			return err
		}
	}
	return nil
}
