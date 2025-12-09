// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortMultipartUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUploadId(v string) *AbortMultipartUploadRequest
	GetUploadId() *string
}

type AbortMultipartUploadRequest struct {
	// The ID of the multipart upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0004B9895DBBB6E****
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s AbortMultipartUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *AbortMultipartUploadRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *AbortMultipartUploadRequest) SetUploadId(v string) *AbortMultipartUploadRequest {
	s.UploadId = &v
	return s
}

func (s *AbortMultipartUploadRequest) Validate() error {
	return dara.Validate(s)
}
