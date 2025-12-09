// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadPartChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v io.Reader) *UploadPartChunkRequest
	GetBody() io.Reader
	SetChunkNumber(v int64) *UploadPartChunkRequest
	GetChunkNumber() *int64
	SetPartUploadId(v string) *UploadPartChunkRequest
	GetPartUploadId() *string
	SetUploadId(v string) *UploadPartChunkRequest
	GetUploadId() *string
}

type UploadPartChunkRequest struct {
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required.
	ChunkNumber *int64 `json:"chunkNumber,omitempty" xml:"chunkNumber,omitempty"`
	// This parameter is required.
	PartUploadId *string `json:"partUploadId,omitempty" xml:"partUploadId,omitempty"`
	// This parameter is required.
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s UploadPartChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadPartChunkRequest) GoString() string {
	return s.String()
}

func (s *UploadPartChunkRequest) GetBody() io.Reader {
	return s.Body
}

func (s *UploadPartChunkRequest) GetChunkNumber() *int64 {
	return s.ChunkNumber
}

func (s *UploadPartChunkRequest) GetPartUploadId() *string {
	return s.PartUploadId
}

func (s *UploadPartChunkRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UploadPartChunkRequest) SetBody(v io.Reader) *UploadPartChunkRequest {
	s.Body = v
	return s
}

func (s *UploadPartChunkRequest) SetChunkNumber(v int64) *UploadPartChunkRequest {
	s.ChunkNumber = &v
	return s
}

func (s *UploadPartChunkRequest) SetPartUploadId(v string) *UploadPartChunkRequest {
	s.PartUploadId = &v
	return s
}

func (s *UploadPartChunkRequest) SetUploadId(v string) *UploadPartChunkRequest {
	s.UploadId = &v
	return s
}

func (s *UploadPartChunkRequest) Validate() error {
	return dara.Validate(s)
}
