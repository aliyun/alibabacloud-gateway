// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileGroupResult interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *CreateFileGroupResult
	GetBucket() *string
	SetETag(v string) *CreateFileGroupResult
	GetETag() *string
	SetKey(v string) *CreateFileGroupResult
	GetKey() *string
	SetSize(v int64) *CreateFileGroupResult
	GetSize() *int64
}

type CreateFileGroupResult struct {
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// "06A4DDABDD5F65868B8C5919E76487D6"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// test.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 384
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateFileGroupResult) String() string {
	return dara.Prettify(s)
}

func (s CreateFileGroupResult) GoString() string {
	return s.String()
}

func (s *CreateFileGroupResult) GetBucket() *string {
	return s.Bucket
}

func (s *CreateFileGroupResult) GetETag() *string {
	return s.ETag
}

func (s *CreateFileGroupResult) GetKey() *string {
	return s.Key
}

func (s *CreateFileGroupResult) GetSize() *int64 {
	return s.Size
}

func (s *CreateFileGroupResult) SetBucket(v string) *CreateFileGroupResult {
	s.Bucket = &v
	return s
}

func (s *CreateFileGroupResult) SetETag(v string) *CreateFileGroupResult {
	s.ETag = &v
	return s
}

func (s *CreateFileGroupResult) SetKey(v string) *CreateFileGroupResult {
	s.Key = &v
	return s
}

func (s *CreateFileGroupResult) SetSize(v int64) *CreateFileGroupResult {
	s.Size = &v
	return s
}

func (s *CreateFileGroupResult) Validate() error {
	return dara.Validate(s)
}
