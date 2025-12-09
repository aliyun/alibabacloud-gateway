// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitPartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPartUploadId(v string) *CommitPartRequest
	GetPartUploadId() *string
	SetUploadId(v string) *CommitPartRequest
	GetUploadId() *string
}

type CommitPartRequest struct {
	// This parameter is required.
	PartUploadId *string `json:"partUploadId,omitempty" xml:"partUploadId,omitempty"`
	// This parameter is required.
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s CommitPartRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitPartRequest) GoString() string {
	return s.String()
}

func (s *CommitPartRequest) GetPartUploadId() *string {
	return s.PartUploadId
}

func (s *CommitPartRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CommitPartRequest) SetPartUploadId(v string) *CommitPartRequest {
	s.PartUploadId = &v
	return s
}

func (s *CommitPartRequest) SetUploadId(v string) *CommitPartRequest {
	s.UploadId = &v
	return s
}

func (s *CommitPartRequest) Validate() error {
	return dara.Validate(s)
}
