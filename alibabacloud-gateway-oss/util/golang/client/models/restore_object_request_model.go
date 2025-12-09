// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRestoreRequest(v *RestoreRequest) *RestoreObjectRequest
	GetRestoreRequest() *RestoreRequest
	SetVersionId(v string) *RestoreObjectRequest
	GetVersionId() *string
}

type RestoreObjectRequest struct {
	// The container that stores information about the RestoreObject request.
	RestoreRequest *RestoreRequest `json:"RestoreRequest,omitempty" xml:"RestoreRequest,omitempty"`
	// The version number of the object that you want to restore.
	//
	// example:
	//
	// CAEQNRiBgMClj7qD0BYiIDQ5Y2QyMjc3NGZkODRlMTU5M2VkY2U3MWRiNGRh
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s RestoreObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreObjectRequest) GoString() string {
	return s.String()
}

func (s *RestoreObjectRequest) GetRestoreRequest() *RestoreRequest {
	return s.RestoreRequest
}

func (s *RestoreObjectRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *RestoreObjectRequest) SetRestoreRequest(v *RestoreRequest) *RestoreObjectRequest {
	s.RestoreRequest = v
	return s
}

func (s *RestoreObjectRequest) SetVersionId(v string) *RestoreObjectRequest {
	s.VersionId = &v
	return s
}

func (s *RestoreObjectRequest) Validate() error {
	if s.RestoreRequest != nil {
		if err := s.RestoreRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
