// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSymlinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *GetSymlinkRequest
	GetVersionId() *string
}

type GetSymlinkRequest struct {
	// The version of the object to which the symbolic link points.
	//
	// example:
	//
	// CAEQNRiBgMClj7qD0BYiIDQ5Y2QyMjc3NGZkODRlMTU5M2VkY2U3MWRiNGRh****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetSymlinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSymlinkRequest) GoString() string {
	return s.String()
}

func (s *GetSymlinkRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetSymlinkRequest) SetVersionId(v string) *GetSymlinkRequest {
	s.VersionId = &v
	return s
}

func (s *GetSymlinkRequest) Validate() error {
	return dara.Validate(s)
}
