// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPartUploadResult interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *StartPartUploadResult
	GetBucket() *string
	SetEncodingType(v string) *StartPartUploadResult
	GetEncodingType() *string
	SetKey(v string) *StartPartUploadResult
	GetKey() *string
	SetPartUploadId(v string) *StartPartUploadResult
	GetPartUploadId() *string
	SetUploadId(v string) *StartPartUploadResult
	GetUploadId() *string
}

type StartPartUploadResult struct {
	// example:
	//
	// bucket1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// example:
	//
	// multipart.data
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 89714B59EF29136807096C6AEB0382730EDA80099D22216E4FEDE45E2B4EC622FC91ED6717B413A9B0C2A4**********
	PartUploadId *string `json:"PartUploadId,omitempty" xml:"PartUploadId,omitempty"`
	// example:
	//
	// 8C108F2EDDCE4C7E946441**********
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s StartPartUploadResult) String() string {
	return dara.Prettify(s)
}

func (s StartPartUploadResult) GoString() string {
	return s.String()
}

func (s *StartPartUploadResult) GetBucket() *string {
	return s.Bucket
}

func (s *StartPartUploadResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *StartPartUploadResult) GetKey() *string {
	return s.Key
}

func (s *StartPartUploadResult) GetPartUploadId() *string {
	return s.PartUploadId
}

func (s *StartPartUploadResult) GetUploadId() *string {
	return s.UploadId
}

func (s *StartPartUploadResult) SetBucket(v string) *StartPartUploadResult {
	s.Bucket = &v
	return s
}

func (s *StartPartUploadResult) SetEncodingType(v string) *StartPartUploadResult {
	s.EncodingType = &v
	return s
}

func (s *StartPartUploadResult) SetKey(v string) *StartPartUploadResult {
	s.Key = &v
	return s
}

func (s *StartPartUploadResult) SetPartUploadId(v string) *StartPartUploadResult {
	s.PartUploadId = &v
	return s
}

func (s *StartPartUploadResult) SetUploadId(v string) *StartPartUploadResult {
	s.UploadId = &v
	return s
}

func (s *StartPartUploadResult) Validate() error {
	return dara.Validate(s)
}
