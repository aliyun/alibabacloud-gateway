// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOutputSerialization interface {
	dara.Model
	String() string
	GoString() string
	SetCsv(v *CSVOutput) *OutputSerialization
	GetCsv() *CSVOutput
	SetEnablePayloadCrc(v bool) *OutputSerialization
	GetEnablePayloadCrc() *bool
	SetJson(v *JSONOutput) *OutputSerialization
	GetJson() *JSONOutput
	SetKeepAllColumns(v bool) *OutputSerialization
	GetKeepAllColumns() *bool
	SetOutputHeader(v bool) *OutputSerialization
	GetOutputHeader() *bool
	SetOutputRawData(v bool) *OutputSerialization
	GetOutputRawData() *bool
}

type OutputSerialization struct {
	Csv              *CSVOutput  `json:"CSV,omitempty" xml:"CSV,omitempty"`
	EnablePayloadCrc *bool       `json:"EnablePayloadCrc,omitempty" xml:"EnablePayloadCrc,omitempty"`
	Json             *JSONOutput `json:"JSON,omitempty" xml:"JSON,omitempty"`
	KeepAllColumns   *bool       `json:"KeepAllColumns,omitempty" xml:"KeepAllColumns,omitempty"`
	OutputHeader     *bool       `json:"OutputHeader,omitempty" xml:"OutputHeader,omitempty"`
	OutputRawData    *bool       `json:"OutputRawData,omitempty" xml:"OutputRawData,omitempty"`
}

func (s OutputSerialization) String() string {
	return dara.Prettify(s)
}

func (s OutputSerialization) GoString() string {
	return s.String()
}

func (s *OutputSerialization) GetCsv() *CSVOutput {
	return s.Csv
}

func (s *OutputSerialization) GetEnablePayloadCrc() *bool {
	return s.EnablePayloadCrc
}

func (s *OutputSerialization) GetJson() *JSONOutput {
	return s.Json
}

func (s *OutputSerialization) GetKeepAllColumns() *bool {
	return s.KeepAllColumns
}

func (s *OutputSerialization) GetOutputHeader() *bool {
	return s.OutputHeader
}

func (s *OutputSerialization) GetOutputRawData() *bool {
	return s.OutputRawData
}

func (s *OutputSerialization) SetCsv(v *CSVOutput) *OutputSerialization {
	s.Csv = v
	return s
}

func (s *OutputSerialization) SetEnablePayloadCrc(v bool) *OutputSerialization {
	s.EnablePayloadCrc = &v
	return s
}

func (s *OutputSerialization) SetJson(v *JSONOutput) *OutputSerialization {
	s.Json = v
	return s
}

func (s *OutputSerialization) SetKeepAllColumns(v bool) *OutputSerialization {
	s.KeepAllColumns = &v
	return s
}

func (s *OutputSerialization) SetOutputHeader(v bool) *OutputSerialization {
	s.OutputHeader = &v
	return s
}

func (s *OutputSerialization) SetOutputRawData(v bool) *OutputSerialization {
	s.OutputRawData = &v
	return s
}

func (s *OutputSerialization) Validate() error {
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
