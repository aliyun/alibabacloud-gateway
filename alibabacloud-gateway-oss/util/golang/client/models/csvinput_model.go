// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCSVInput interface {
	dara.Model
	String() string
	GoString() string
	SetAllowQuotedRecordDelimiter(v bool) *CSVInput
	GetAllowQuotedRecordDelimiter() *bool
	SetCommentCharacter(v string) *CSVInput
	GetCommentCharacter() *string
	SetFieldDelimiter(v string) *CSVInput
	GetFieldDelimiter() *string
	SetFileHeaderInfo(v string) *CSVInput
	GetFileHeaderInfo() *string
	SetQuoteCharacter(v string) *CSVInput
	GetQuoteCharacter() *string
	SetRange(v string) *CSVInput
	GetRange() *string
	SetRecordDelimiter(v string) *CSVInput
	GetRecordDelimiter() *string
}

type CSVInput struct {
	AllowQuotedRecordDelimiter *bool   `json:"AllowQuotedRecordDelimiter,omitempty" xml:"AllowQuotedRecordDelimiter,omitempty"`
	CommentCharacter           *string `json:"CommentCharacter,omitempty" xml:"CommentCharacter,omitempty"`
	FieldDelimiter             *string `json:"FieldDelimiter,omitempty" xml:"FieldDelimiter,omitempty"`
	FileHeaderInfo             *string `json:"FileHeaderInfo,omitempty" xml:"FileHeaderInfo,omitempty"`
	QuoteCharacter             *string `json:"QuoteCharacter,omitempty" xml:"QuoteCharacter,omitempty"`
	Range                      *string `json:"Range,omitempty" xml:"Range,omitempty"`
	RecordDelimiter            *string `json:"RecordDelimiter,omitempty" xml:"RecordDelimiter,omitempty"`
}

func (s CSVInput) String() string {
	return dara.Prettify(s)
}

func (s CSVInput) GoString() string {
	return s.String()
}

func (s *CSVInput) GetAllowQuotedRecordDelimiter() *bool {
	return s.AllowQuotedRecordDelimiter
}

func (s *CSVInput) GetCommentCharacter() *string {
	return s.CommentCharacter
}

func (s *CSVInput) GetFieldDelimiter() *string {
	return s.FieldDelimiter
}

func (s *CSVInput) GetFileHeaderInfo() *string {
	return s.FileHeaderInfo
}

func (s *CSVInput) GetQuoteCharacter() *string {
	return s.QuoteCharacter
}

func (s *CSVInput) GetRange() *string {
	return s.Range
}

func (s *CSVInput) GetRecordDelimiter() *string {
	return s.RecordDelimiter
}

func (s *CSVInput) SetAllowQuotedRecordDelimiter(v bool) *CSVInput {
	s.AllowQuotedRecordDelimiter = &v
	return s
}

func (s *CSVInput) SetCommentCharacter(v string) *CSVInput {
	s.CommentCharacter = &v
	return s
}

func (s *CSVInput) SetFieldDelimiter(v string) *CSVInput {
	s.FieldDelimiter = &v
	return s
}

func (s *CSVInput) SetFileHeaderInfo(v string) *CSVInput {
	s.FileHeaderInfo = &v
	return s
}

func (s *CSVInput) SetQuoteCharacter(v string) *CSVInput {
	s.QuoteCharacter = &v
	return s
}

func (s *CSVInput) SetRange(v string) *CSVInput {
	s.Range = &v
	return s
}

func (s *CSVInput) SetRecordDelimiter(v string) *CSVInput {
	s.RecordDelimiter = &v
	return s
}

func (s *CSVInput) Validate() error {
	return dara.Validate(s)
}
