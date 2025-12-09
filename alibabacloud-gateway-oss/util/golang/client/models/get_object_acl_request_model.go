// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *GetObjectAclRequest
	GetVersionId() *string
}

type GetObjectAclRequest struct {
	// The verison id of the target object.
	//
	// example:
	//
	// list1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectAclRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectAclRequest) GoString() string {
	return s.String()
}

func (s *GetObjectAclRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetObjectAclRequest) SetVersionId(v string) *GetObjectAclRequest {
	s.VersionId = &v
	return s
}

func (s *GetObjectAclRequest) Validate() error {
	return dara.Validate(s)
}
