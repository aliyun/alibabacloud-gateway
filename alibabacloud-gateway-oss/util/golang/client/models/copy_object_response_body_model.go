// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCopyObjectResult(v *CopyObjectResponseBodyCopyObjectResult) *CopyObjectResponseBody
	GetCopyObjectResult() *CopyObjectResponseBodyCopyObjectResult
}

type CopyObjectResponseBody struct {
	// The results of the CopyObject operation.
	CopyObjectResult *CopyObjectResponseBodyCopyObjectResult `json:"CopyObjectResult,omitempty" xml:"CopyObjectResult,omitempty" type:"Struct"`
}

func (s CopyObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectResponseBody) GoString() string {
	return s.String()
}

func (s *CopyObjectResponseBody) GetCopyObjectResult() *CopyObjectResponseBodyCopyObjectResult {
	return s.CopyObjectResult
}

func (s *CopyObjectResponseBody) SetCopyObjectResult(v *CopyObjectResponseBodyCopyObjectResult) *CopyObjectResponseBody {
	s.CopyObjectResult = v
	return s
}

func (s *CopyObjectResponseBody) Validate() error {
	if s.CopyObjectResult != nil {
		if err := s.CopyObjectResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyObjectResponseBodyCopyObjectResult struct {
	// The ETag value of the destination object.
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5****
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The time when the destination object was last modified.
	//
	// example:
	//
	// Fri, 24 Feb 2012 07:18:48 GMT
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s CopyObjectResponseBodyCopyObjectResult) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectResponseBodyCopyObjectResult) GoString() string {
	return s.String()
}

func (s *CopyObjectResponseBodyCopyObjectResult) GetETag() *string {
	return s.ETag
}

func (s *CopyObjectResponseBodyCopyObjectResult) GetLastModified() *string {
	return s.LastModified
}

func (s *CopyObjectResponseBodyCopyObjectResult) SetETag(v string) *CopyObjectResponseBodyCopyObjectResult {
	s.ETag = &v
	return s
}

func (s *CopyObjectResponseBodyCopyObjectResult) SetLastModified(v string) *CopyObjectResponseBodyCopyObjectResult {
	s.LastModified = &v
	return s
}

func (s *CopyObjectResponseBodyCopyObjectResult) Validate() error {
	return dara.Validate(s)
}
