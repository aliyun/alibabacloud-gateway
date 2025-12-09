// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *GetObjectTaggingRequest
	GetVersionId() *string
}

type GetObjectTaggingRequest struct {
	// The versionID of the object that you want to query.
	//
	// example:
	//
	// CAEQExiBgID.jImWlxciIDQ2ZjgwODIyNDk5MTRhNzBiYmQwYTZkMTYzZjM0****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectTaggingRequest) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetObjectTaggingRequest) SetVersionId(v string) *GetObjectTaggingRequest {
	s.VersionId = &v
	return s
}

func (s *GetObjectTaggingRequest) Validate() error {
	return dara.Validate(s)
}
