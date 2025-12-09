// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagging(v *Tagging) *PutObjectTaggingRequest
	GetTagging() *Tagging
	SetVersionId(v string) *PutObjectTaggingRequest
	GetVersionId() *string
}

type PutObjectTaggingRequest struct {
	// The container of the tag set.
	Tagging *Tagging `json:"Tagging,omitempty" xml:"Tagging,omitempty"`
	// The version id of the target object.
	//
	// example:
	//
	// CAEQNRiBgMClj7qD0BYiIDQ5Y2QyMjc3NGZkODRlMTU5M2VkY2U3MWRiNGRh****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PutObjectTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectTaggingRequest) GoString() string {
	return s.String()
}

func (s *PutObjectTaggingRequest) GetTagging() *Tagging {
	return s.Tagging
}

func (s *PutObjectTaggingRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PutObjectTaggingRequest) SetTagging(v *Tagging) *PutObjectTaggingRequest {
	s.Tagging = v
	return s
}

func (s *PutObjectTaggingRequest) SetVersionId(v string) *PutObjectTaggingRequest {
	s.VersionId = &v
	return s
}

func (s *PutObjectTaggingRequest) Validate() error {
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			return err
		}
	}
	return nil
}
