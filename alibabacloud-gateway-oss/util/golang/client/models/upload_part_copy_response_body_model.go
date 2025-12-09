// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartCopyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCopyPartResult(v *UploadPartCopyResponseBodyCopyPartResult) *UploadPartCopyResponseBody
	GetCopyPartResult() *UploadPartCopyResponseBodyCopyPartResult
}

type UploadPartCopyResponseBody struct {
	// The container that stores the copy result.
	CopyPartResult *UploadPartCopyResponseBodyCopyPartResult `json:"CopyPartResult,omitempty" xml:"CopyPartResult,omitempty" type:"Struct"`
}

func (s UploadPartCopyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadPartCopyResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPartCopyResponseBody) GetCopyPartResult() *UploadPartCopyResponseBodyCopyPartResult {
	return s.CopyPartResult
}

func (s *UploadPartCopyResponseBody) SetCopyPartResult(v *UploadPartCopyResponseBodyCopyPartResult) *UploadPartCopyResponseBody {
	s.CopyPartResult = v
	return s
}

func (s *UploadPartCopyResponseBody) Validate() error {
	if s.CopyPartResult != nil {
		if err := s.CopyPartResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadPartCopyResponseBodyCopyPartResult struct {
	// The ETag of the copied part.
	//
	// example:
	//
	// "5B3C1A2E053D763E1B002CC607C5****"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The last modified time of copy source.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2014-07-17T06:27:54.000Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s UploadPartCopyResponseBodyCopyPartResult) String() string {
	return dara.Prettify(s)
}

func (s UploadPartCopyResponseBodyCopyPartResult) GoString() string {
	return s.String()
}

func (s *UploadPartCopyResponseBodyCopyPartResult) GetETag() *string {
	return s.ETag
}

func (s *UploadPartCopyResponseBodyCopyPartResult) GetLastModified() *string {
	return s.LastModified
}

func (s *UploadPartCopyResponseBodyCopyPartResult) SetETag(v string) *UploadPartCopyResponseBodyCopyPartResult {
	s.ETag = &v
	return s
}

func (s *UploadPartCopyResponseBodyCopyPartResult) SetLastModified(v string) *UploadPartCopyResponseBodyCopyPartResult {
	s.LastModified = &v
	return s
}

func (s *UploadPartCopyResponseBodyCopyPartResult) Validate() error {
	return dara.Validate(s)
}
