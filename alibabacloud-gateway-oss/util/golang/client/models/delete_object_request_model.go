// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *DeleteObjectRequest
	GetVersionId() *string
}

type DeleteObjectRequest struct {
	// The version ID of the object.
	//
	// example:
	//
	// CAEQMxiBgIDh3ZCB0BYiIGE4YjIyMjExZDhhYjQxNzZiNGUyZTI4ZjljZDcz****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s DeleteObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteObjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeleteObjectRequest) SetVersionId(v string) *DeleteObjectRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteObjectRequest) Validate() error {
	return dara.Validate(s)
}
