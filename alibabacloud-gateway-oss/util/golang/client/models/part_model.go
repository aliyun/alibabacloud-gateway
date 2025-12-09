// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPart interface {
	dara.Model
	String() string
	GoString() string
	SetETag(v string) *Part
	GetETag() *string
	SetLastModified(v string) *Part
	GetLastModified() *string
	SetPartNumber(v int64) *Part
	GetPartNumber() *int64
	SetSize(v int64) *Part
	GetSize() *int64
}

type Part struct {
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	PartNumber   *int64  `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s Part) String() string {
	return dara.Prettify(s)
}

func (s Part) GoString() string {
	return s.String()
}

func (s *Part) GetETag() *string {
	return s.ETag
}

func (s *Part) GetLastModified() *string {
	return s.LastModified
}

func (s *Part) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *Part) GetSize() *int64 {
	return s.Size
}

func (s *Part) SetETag(v string) *Part {
	s.ETag = &v
	return s
}

func (s *Part) SetLastModified(v string) *Part {
	s.LastModified = &v
	return s
}

func (s *Part) SetPartNumber(v int64) *Part {
	s.PartNumber = &v
	return s
}

func (s *Part) SetSize(v int64) *Part {
	s.Size = &v
	return s
}

func (s *Part) Validate() error {
	return dara.Validate(s)
}
