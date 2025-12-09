// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagging(v *Tagging) *PutBucketTagsRequest
	GetTagging() *Tagging
}

type PutBucketTagsRequest struct {
	// The container used to store TagSet.
	Tagging *Tagging `json:"Tagging,omitempty" xml:"Tagging,omitempty"`
}

func (s PutBucketTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketTagsRequest) GoString() string {
	return s.String()
}

func (s *PutBucketTagsRequest) GetTagging() *Tagging {
	return s.Tagging
}

func (s *PutBucketTagsRequest) SetTagging(v *Tagging) *PutBucketTagsRequest {
	s.Tagging = v
	return s
}

func (s *PutBucketTagsRequest) Validate() error {
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			return err
		}
	}
	return nil
}
