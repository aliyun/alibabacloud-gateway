// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPartUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkSize(v int64) *StartPartUploadRequest
	GetChunkSize() *int64
	SetEncodingType(v string) *StartPartUploadRequest
	GetEncodingType() *string
	SetPartNumber(v int64) *StartPartUploadRequest
	GetPartNumber() *int64
	SetPartSize(v int64) *StartPartUploadRequest
	GetPartSize() *int64
	SetUploadId(v string) *StartPartUploadRequest
	GetUploadId() *string
}

type StartPartUploadRequest struct {
	// This parameter is required.
	ChunkSize    *int64  `json:"chunkSize,omitempty" xml:"chunkSize,omitempty"`
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	// This parameter is required.
	PartNumber *int64 `json:"partNumber,omitempty" xml:"partNumber,omitempty"`
	// This parameter is required.
	PartSize *int64 `json:"partSize,omitempty" xml:"partSize,omitempty"`
	// This parameter is required.
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s StartPartUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPartUploadRequest) GoString() string {
	return s.String()
}

func (s *StartPartUploadRequest) GetChunkSize() *int64 {
	return s.ChunkSize
}

func (s *StartPartUploadRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *StartPartUploadRequest) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *StartPartUploadRequest) GetPartSize() *int64 {
	return s.PartSize
}

func (s *StartPartUploadRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *StartPartUploadRequest) SetChunkSize(v int64) *StartPartUploadRequest {
	s.ChunkSize = &v
	return s
}

func (s *StartPartUploadRequest) SetEncodingType(v string) *StartPartUploadRequest {
	s.EncodingType = &v
	return s
}

func (s *StartPartUploadRequest) SetPartNumber(v int64) *StartPartUploadRequest {
	s.PartNumber = &v
	return s
}

func (s *StartPartUploadRequest) SetPartSize(v int64) *StartPartUploadRequest {
	s.PartSize = &v
	return s
}

func (s *StartPartUploadRequest) SetUploadId(v string) *StartPartUploadRequest {
	s.UploadId = &v
	return s
}

func (s *StartPartUploadRequest) Validate() error {
	return dara.Validate(s)
}
