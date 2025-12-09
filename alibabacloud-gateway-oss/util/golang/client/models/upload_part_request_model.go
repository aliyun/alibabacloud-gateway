// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadPartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v io.Reader) *UploadPartRequest
	GetBody() io.Reader
	SetPartNumber(v int64) *UploadPartRequest
	GetPartNumber() *int64
	SetUploadId(v string) *UploadPartRequest
	GetUploadId() *string
}

type UploadPartRequest struct {
	// The request body.
	//
	// example:
	//
	// Binary content.
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	// The number that identifies a part.
	//
	// Valid values: 1 to 10000.
	//
	// The size of a part ranges from 100 KB to 5 GB.
	//
	// > In multipart upload, each part except the last part must be larger than or equal to 100 KB in size. When you call the UploadPart operation, the size of each part is not verified because not all parts have been uploaded and OSS does not know which part is the last part. The size of each part is verified only when you call CompleteMultipartUpload.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PartNumber *int64 `json:"partNumber,omitempty" xml:"partNumber,omitempty"`
	// The ID that identifies the object to which the part that you want to upload belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0004B9895DBBB6EC98E36
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s UploadPartRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadPartRequest) GoString() string {
	return s.String()
}

func (s *UploadPartRequest) GetBody() io.Reader {
	return s.Body
}

func (s *UploadPartRequest) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *UploadPartRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UploadPartRequest) SetBody(v io.Reader) *UploadPartRequest {
	s.Body = v
	return s
}

func (s *UploadPartRequest) SetPartNumber(v int64) *UploadPartRequest {
	s.PartNumber = &v
	return s
}

func (s *UploadPartRequest) SetUploadId(v string) *UploadPartRequest {
	s.UploadId = &v
	return s
}

func (s *UploadPartRequest) Validate() error {
	return dara.Validate(s)
}
