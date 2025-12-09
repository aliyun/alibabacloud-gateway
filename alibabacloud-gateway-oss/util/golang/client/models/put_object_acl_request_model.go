// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *PutObjectAclRequest
	GetVersionId() *string
}

type PutObjectAclRequest struct {
	// The version id of the object.
	//
	// example:
	//
	// CAEQMhiBgIC3rpSD0BYiIDBjYTk5MmIzN2JlNjQxZTFiNGIzM2E3OTliODA0****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PutObjectAclRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectAclRequest) GoString() string {
	return s.String()
}

func (s *PutObjectAclRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PutObjectAclRequest) SetVersionId(v string) *PutObjectAclRequest {
	s.VersionId = &v
	return s
}

func (s *PutObjectAclRequest) Validate() error {
	return dara.Validate(s)
}
