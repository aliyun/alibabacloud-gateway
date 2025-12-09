// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *GetObjectMetaRequest
	GetVersionId() *string
}

type GetObjectMetaRequest struct {
	// The versionID of the object.
	//
	// example:
	//
	// CAEQNRiBgIDMh4mD0BYiIDUzNDA4OGNmZjBjYTQ0YmI4Y2I4ZmVlYzJlNGVk****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectMetaRequest) GoString() string {
	return s.String()
}

func (s *GetObjectMetaRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetObjectMetaRequest) SetVersionId(v string) *GetObjectMetaRequest {
	s.VersionId = &v
	return s
}

func (s *GetObjectMetaRequest) Validate() error {
	return dara.Validate(s)
}
