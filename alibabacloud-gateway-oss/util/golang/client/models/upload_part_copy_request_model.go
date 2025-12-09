// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartCopyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPartNumber(v int64) *UploadPartCopyRequest
	GetPartNumber() *int64
	SetUploadId(v string) *UploadPartCopyRequest
	GetUploadId() *string
}

type UploadPartCopyRequest struct {
	// The number of parts.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	PartNumber *int64 `json:"partNumber,omitempty" xml:"partNumber,omitempty"`
	// The ID that identifies the object to which the parts to upload belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0004B999EF5A239BB9138C6227D69F95
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s UploadPartCopyRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadPartCopyRequest) GoString() string {
	return s.String()
}

func (s *UploadPartCopyRequest) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *UploadPartCopyRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UploadPartCopyRequest) SetPartNumber(v int64) *UploadPartCopyRequest {
	s.PartNumber = &v
	return s
}

func (s *UploadPartCopyRequest) SetUploadId(v string) *UploadPartCopyRequest {
	s.UploadId = &v
	return s
}

func (s *UploadPartCopyRequest) Validate() error {
	return dara.Validate(s)
}
