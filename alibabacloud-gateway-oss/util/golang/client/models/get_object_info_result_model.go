// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectInfoResult interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *GetObjectInfoResult
	GetBucket() *string
	SetContentType(v string) *GetObjectInfoResult
	GetContentType() *string
	SetETag(v string) *GetObjectInfoResult
	GetETag() *string
	SetEncryptFlag(v int64) *GetObjectInfoResult
	GetEncryptFlag() *int64
	SetHashCrc64ecma(v string) *GetObjectInfoResult
	GetHashCrc64ecma() *string
	SetKey(v string) *GetObjectInfoResult
	GetKey() *string
	SetLastModified(v string) *GetObjectInfoResult
	GetLastModified() *string
	SetSize(v string) *GetObjectInfoResult
	GetSize() *string
	SetStorageClass(v string) *GetObjectInfoResult
	GetStorageClass() *string
	SetType(v string) *GetObjectInfoResult
	GetType() *string
	SetUploadId(v string) *GetObjectInfoResult
	GetUploadId() *string
}

type GetObjectInfoResult struct {
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// text/plain
	ContentType *string `json:"Content-Type,omitempty" xml:"Content-Type,omitempty"`
	// example:
	//
	// "EB327B57BB20D17C293A966115FE27BD"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// 0
	EncryptFlag *int64 `json:"EncryptFlag,omitempty" xml:"EncryptFlag,omitempty"`
	// example:
	//
	// 7866203970082294162
	HashCrc64ecma *string `json:"HashCrc64ecma,omitempty" xml:"HashCrc64ecma,omitempty"`
	// example:
	//
	// test.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-11-08T10:02:11.000Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// example:
	//
	// 78
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// example:
	//
	// Multipart
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// C5DD87C5E7CD482F8F2C3D63904DA228
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s GetObjectInfoResult) String() string {
	return dara.Prettify(s)
}

func (s GetObjectInfoResult) GoString() string {
	return s.String()
}

func (s *GetObjectInfoResult) GetBucket() *string {
	return s.Bucket
}

func (s *GetObjectInfoResult) GetContentType() *string {
	return s.ContentType
}

func (s *GetObjectInfoResult) GetETag() *string {
	return s.ETag
}

func (s *GetObjectInfoResult) GetEncryptFlag() *int64 {
	return s.EncryptFlag
}

func (s *GetObjectInfoResult) GetHashCrc64ecma() *string {
	return s.HashCrc64ecma
}

func (s *GetObjectInfoResult) GetKey() *string {
	return s.Key
}

func (s *GetObjectInfoResult) GetLastModified() *string {
	return s.LastModified
}

func (s *GetObjectInfoResult) GetSize() *string {
	return s.Size
}

func (s *GetObjectInfoResult) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetObjectInfoResult) GetType() *string {
	return s.Type
}

func (s *GetObjectInfoResult) GetUploadId() *string {
	return s.UploadId
}

func (s *GetObjectInfoResult) SetBucket(v string) *GetObjectInfoResult {
	s.Bucket = &v
	return s
}

func (s *GetObjectInfoResult) SetContentType(v string) *GetObjectInfoResult {
	s.ContentType = &v
	return s
}

func (s *GetObjectInfoResult) SetETag(v string) *GetObjectInfoResult {
	s.ETag = &v
	return s
}

func (s *GetObjectInfoResult) SetEncryptFlag(v int64) *GetObjectInfoResult {
	s.EncryptFlag = &v
	return s
}

func (s *GetObjectInfoResult) SetHashCrc64ecma(v string) *GetObjectInfoResult {
	s.HashCrc64ecma = &v
	return s
}

func (s *GetObjectInfoResult) SetKey(v string) *GetObjectInfoResult {
	s.Key = &v
	return s
}

func (s *GetObjectInfoResult) SetLastModified(v string) *GetObjectInfoResult {
	s.LastModified = &v
	return s
}

func (s *GetObjectInfoResult) SetSize(v string) *GetObjectInfoResult {
	s.Size = &v
	return s
}

func (s *GetObjectInfoResult) SetStorageClass(v string) *GetObjectInfoResult {
	s.StorageClass = &v
	return s
}

func (s *GetObjectInfoResult) SetType(v string) *GetObjectInfoResult {
	s.Type = &v
	return s
}

func (s *GetObjectInfoResult) SetUploadId(v string) *GetObjectInfoResult {
	s.UploadId = &v
	return s
}

func (s *GetObjectInfoResult) Validate() error {
	return dara.Validate(s)
}
