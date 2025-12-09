// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectResult interface {
	dara.Model
	String() string
	GoString() string
	SetETag(v string) *CopyObjectResult
	GetETag() *string
	SetLastModified(v string) *CopyObjectResult
	GetLastModified() *string
}

type CopyObjectResult struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s CopyObjectResult) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectResult) GoString() string {
	return s.String()
}

func (s *CopyObjectResult) GetETag() *string {
	return s.ETag
}

func (s *CopyObjectResult) GetLastModified() *string {
	return s.LastModified
}

func (s *CopyObjectResult) SetETag(v string) *CopyObjectResult {
	s.ETag = &v
	return s
}

func (s *CopyObjectResult) SetLastModified(v string) *CopyObjectResult {
	s.LastModified = &v
	return s
}

func (s *CopyObjectResult) Validate() error {
	return dara.Validate(s)
}
