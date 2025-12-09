// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHeadObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionId(v string) *HeadObjectRequest
	GetVersionId() *string
}

type HeadObjectRequest struct {
	// The version ID of the object for which you want to query metadata.
	//
	// example:
	//
	// CAEQMxiBgMCZov2D0BYiIDY4MDllOTc2YmY5MjQxMzdiOGI3OTlhNTU0ODIx****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s HeadObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s HeadObjectRequest) GoString() string {
	return s.String()
}

func (s *HeadObjectRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *HeadObjectRequest) SetVersionId(v string) *HeadObjectRequest {
	s.VersionId = &v
	return s
}

func (s *HeadObjectRequest) Validate() error {
	return dara.Validate(s)
}
