// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCSVOutput interface {
	dara.Model
	String() string
	GoString() string
	SetFieldDelimiter(v string) *CSVOutput
	GetFieldDelimiter() *string
	SetRecordDelimiter(v string) *CSVOutput
	GetRecordDelimiter() *string
}

type CSVOutput struct {
	FieldDelimiter  *string `json:"FieldDelimiter,omitempty" xml:"FieldDelimiter,omitempty"`
	RecordDelimiter *string `json:"RecordDelimiter,omitempty" xml:"RecordDelimiter,omitempty"`
}

func (s CSVOutput) String() string {
	return dara.Prettify(s)
}

func (s CSVOutput) GoString() string {
	return s.String()
}

func (s *CSVOutput) GetFieldDelimiter() *string {
	return s.FieldDelimiter
}

func (s *CSVOutput) GetRecordDelimiter() *string {
	return s.RecordDelimiter
}

func (s *CSVOutput) SetFieldDelimiter(v string) *CSVOutput {
	s.FieldDelimiter = &v
	return s
}

func (s *CSVOutput) SetRecordDelimiter(v string) *CSVOutput {
	s.RecordDelimiter = &v
	return s
}

func (s *CSVOutput) Validate() error {
	return dara.Validate(s)
}
