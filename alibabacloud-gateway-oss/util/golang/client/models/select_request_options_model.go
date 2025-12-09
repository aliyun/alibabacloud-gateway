// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectRequestOptions interface {
	dara.Model
	String() string
	GoString() string
	SetMaxSkippedRecordsAllowed(v int64) *SelectRequestOptions
	GetMaxSkippedRecordsAllowed() *int64
	SetSkipPartialDataRecord(v bool) *SelectRequestOptions
	GetSkipPartialDataRecord() *bool
}

type SelectRequestOptions struct {
	MaxSkippedRecordsAllowed *int64 `json:"MaxSkippedRecordsAllowed,omitempty" xml:"MaxSkippedRecordsAllowed,omitempty"`
	SkipPartialDataRecord    *bool  `json:"SkipPartialDataRecord,omitempty" xml:"SkipPartialDataRecord,omitempty"`
}

func (s SelectRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s SelectRequestOptions) GoString() string {
	return s.String()
}

func (s *SelectRequestOptions) GetMaxSkippedRecordsAllowed() *int64 {
	return s.MaxSkippedRecordsAllowed
}

func (s *SelectRequestOptions) GetSkipPartialDataRecord() *bool {
	return s.SkipPartialDataRecord
}

func (s *SelectRequestOptions) SetMaxSkippedRecordsAllowed(v int64) *SelectRequestOptions {
	s.MaxSkippedRecordsAllowed = &v
	return s
}

func (s *SelectRequestOptions) SetSkipPartialDataRecord(v bool) *SelectRequestOptions {
	s.SkipPartialDataRecord = &v
	return s
}

func (s *SelectRequestOptions) Validate() error {
	return dara.Validate(s)
}
