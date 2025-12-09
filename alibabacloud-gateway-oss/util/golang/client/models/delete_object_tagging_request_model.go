// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteObjectTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *DeleteObjectTaggingRequest
	GetVersionId() *string
}

type DeleteObjectTaggingRequest struct {
	// The version ID of the object that you want to delete.
	//
	// example:
	//
	// CAEQExiBgID.jImWlxciIDQ2ZjgwODIyNDk5MTRhNzBiYmQwYTZkMTYzZjM0****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s DeleteObjectTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteObjectTaggingRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectTaggingRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeleteObjectTaggingRequest) SetVersionId(v string) *DeleteObjectTaggingRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteObjectTaggingRequest) Validate() error {
	return dara.Validate(s)
}
