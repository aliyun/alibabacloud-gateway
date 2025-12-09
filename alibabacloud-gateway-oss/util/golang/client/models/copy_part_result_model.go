// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyPartResult interface {
	dara.Model
	String() string
	GoString() string
	SetETag(v string) *CopyPartResult
	GetETag() *string
	SetLastModified(v string) *CopyPartResult
	GetLastModified() *string
}

type CopyPartResult struct {
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s CopyPartResult) String() string {
	return dara.Prettify(s)
}

func (s CopyPartResult) GoString() string {
	return s.String()
}

func (s *CopyPartResult) GetETag() *string {
	return s.ETag
}

func (s *CopyPartResult) GetLastModified() *string {
	return s.LastModified
}

func (s *CopyPartResult) SetETag(v string) *CopyPartResult {
	s.ETag = &v
	return s
}

func (s *CopyPartResult) SetLastModified(v string) *CopyPartResult {
	s.LastModified = &v
	return s
}

func (s *CopyPartResult) Validate() error {
	return dara.Validate(s)
}
