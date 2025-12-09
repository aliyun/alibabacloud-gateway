// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectMetaStatus interface {
	dara.Model
	String() string
	GoString() string
	SetColsCount(v int64) *SelectMetaStatus
	GetColsCount() *int64
	SetErrorMessage(v string) *SelectMetaStatus
	GetErrorMessage() *string
	SetOffset(v int64) *SelectMetaStatus
	GetOffset() *int64
	SetRowsCount(v int64) *SelectMetaStatus
	GetRowsCount() *int64
	SetSplitsCount(v int64) *SelectMetaStatus
	GetSplitsCount() *int64
	SetStatus(v int64) *SelectMetaStatus
	GetStatus() *int64
	SetTotalScannedBytes(v int64) *SelectMetaStatus
	GetTotalScannedBytes() *int64
}

type SelectMetaStatus struct {
	ColsCount         *int64  `json:"ColsCount,omitempty" xml:"ColsCount,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Offset            *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RowsCount         *int64  `json:"RowsCount,omitempty" xml:"RowsCount,omitempty"`
	SplitsCount       *int64  `json:"SplitsCount,omitempty" xml:"SplitsCount,omitempty"`
	Status            *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalScannedBytes *int64  `json:"TotalScannedBytes,omitempty" xml:"TotalScannedBytes,omitempty"`
}

func (s SelectMetaStatus) String() string {
	return dara.Prettify(s)
}

func (s SelectMetaStatus) GoString() string {
	return s.String()
}

func (s *SelectMetaStatus) GetColsCount() *int64 {
	return s.ColsCount
}

func (s *SelectMetaStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SelectMetaStatus) GetOffset() *int64 {
	return s.Offset
}

func (s *SelectMetaStatus) GetRowsCount() *int64 {
	return s.RowsCount
}

func (s *SelectMetaStatus) GetSplitsCount() *int64 {
	return s.SplitsCount
}

func (s *SelectMetaStatus) GetStatus() *int64 {
	return s.Status
}

func (s *SelectMetaStatus) GetTotalScannedBytes() *int64 {
	return s.TotalScannedBytes
}

func (s *SelectMetaStatus) SetColsCount(v int64) *SelectMetaStatus {
	s.ColsCount = &v
	return s
}

func (s *SelectMetaStatus) SetErrorMessage(v string) *SelectMetaStatus {
	s.ErrorMessage = &v
	return s
}

func (s *SelectMetaStatus) SetOffset(v int64) *SelectMetaStatus {
	s.Offset = &v
	return s
}

func (s *SelectMetaStatus) SetRowsCount(v int64) *SelectMetaStatus {
	s.RowsCount = &v
	return s
}

func (s *SelectMetaStatus) SetSplitsCount(v int64) *SelectMetaStatus {
	s.SplitsCount = &v
	return s
}

func (s *SelectMetaStatus) SetStatus(v int64) *SelectMetaStatus {
	s.Status = &v
	return s
}

func (s *SelectMetaStatus) SetTotalScannedBytes(v int64) *SelectMetaStatus {
	s.TotalScannedBytes = &v
	return s
}

func (s *SelectMetaStatus) Validate() error {
	return dara.Validate(s)
}
