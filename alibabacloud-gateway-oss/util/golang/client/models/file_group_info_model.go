// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileGroupInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *FileGroupInfo
	GetBucket() *string
	SetETag(v string) *FileGroupInfo
	GetETag() *string
	SetFileLength(v int64) *FileGroupInfo
	GetFileLength() *int64
	SetFilePart(v *FileGroupInfoFilePart) *FileGroupInfo
	GetFilePart() *FileGroupInfoFilePart
	SetKey(v string) *FileGroupInfo
	GetKey() *string
}

type FileGroupInfo struct {
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
	// 284
	FileLength *int64 `json:"FileLength,omitempty" xml:"FileLength,omitempty"`
	// FileGroup类型文件的信息
	FilePart *FileGroupInfoFilePart `json:"FilePart,omitempty" xml:"FilePart,omitempty" type:"Struct"`
	// example:
	//
	// file-goup-example
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s FileGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s FileGroupInfo) GoString() string {
	return s.String()
}

func (s *FileGroupInfo) GetBucket() *string {
	return s.Bucket
}

func (s *FileGroupInfo) GetETag() *string {
	return s.ETag
}

func (s *FileGroupInfo) GetFileLength() *int64 {
	return s.FileLength
}

func (s *FileGroupInfo) GetFilePart() *FileGroupInfoFilePart {
	return s.FilePart
}

func (s *FileGroupInfo) GetKey() *string {
	return s.Key
}

func (s *FileGroupInfo) SetBucket(v string) *FileGroupInfo {
	s.Bucket = &v
	return s
}

func (s *FileGroupInfo) SetETag(v string) *FileGroupInfo {
	s.ETag = &v
	return s
}

func (s *FileGroupInfo) SetFileLength(v int64) *FileGroupInfo {
	s.FileLength = &v
	return s
}

func (s *FileGroupInfo) SetFilePart(v *FileGroupInfoFilePart) *FileGroupInfo {
	s.FilePart = v
	return s
}

func (s *FileGroupInfo) SetKey(v string) *FileGroupInfo {
	s.Key = &v
	return s
}

func (s *FileGroupInfo) Validate() error {
	if s.FilePart != nil {
		if err := s.FilePart.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FileGroupInfoFilePart struct {
	Part []*FileGroupInfoFilePartPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s FileGroupInfoFilePart) String() string {
	return dara.Prettify(s)
}

func (s FileGroupInfoFilePart) GoString() string {
	return s.String()
}

func (s *FileGroupInfoFilePart) GetPart() []*FileGroupInfoFilePartPart {
	return s.Part
}

func (s *FileGroupInfoFilePart) SetPart(v []*FileGroupInfoFilePartPart) *FileGroupInfoFilePart {
	s.Part = v
	return s
}

func (s *FileGroupInfoFilePart) Validate() error {
	if s.Part != nil {
		for _, item := range s.Part {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FileGroupInfoFilePartPart struct {
	// example:
	//
	// "EB327B57BB20D17C293A966115FE27BD"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// test.txt
	PartName *string `json:"PartName,omitempty" xml:"PartName,omitempty"`
	// example:
	//
	// 3
	PartNumber *int64 `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
	// example:
	//
	// 38
	PartSize *int64 `json:"PartSize,omitempty" xml:"PartSize,omitempty"`
}

func (s FileGroupInfoFilePartPart) String() string {
	return dara.Prettify(s)
}

func (s FileGroupInfoFilePartPart) GoString() string {
	return s.String()
}

func (s *FileGroupInfoFilePartPart) GetETag() *string {
	return s.ETag
}

func (s *FileGroupInfoFilePartPart) GetPartName() *string {
	return s.PartName
}

func (s *FileGroupInfoFilePartPart) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *FileGroupInfoFilePartPart) GetPartSize() *int64 {
	return s.PartSize
}

func (s *FileGroupInfoFilePartPart) SetETag(v string) *FileGroupInfoFilePartPart {
	s.ETag = &v
	return s
}

func (s *FileGroupInfoFilePartPart) SetPartName(v string) *FileGroupInfoFilePartPart {
	s.PartName = &v
	return s
}

func (s *FileGroupInfoFilePartPart) SetPartNumber(v int64) *FileGroupInfoFilePartPart {
	s.PartNumber = &v
	return s
}

func (s *FileGroupInfoFilePartPart) SetPartSize(v int64) *FileGroupInfoFilePartPart {
	s.PartSize = &v
	return s
}

func (s *FileGroupInfoFilePartPart) Validate() error {
	return dara.Validate(s)
}
