// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTagging(v *GetObjectTaggingResponseBodyTagging) *GetObjectTaggingResponseBody
	GetTagging() *GetObjectTaggingResponseBodyTagging
}

type GetObjectTaggingResponseBody struct {
	// The container that stores the returned tag of the bucket.
	Tagging *GetObjectTaggingResponseBodyTagging `json:"Tagging,omitempty" xml:"Tagging,omitempty" type:"Struct"`
}

func (s GetObjectTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingResponseBody) GetTagging() *GetObjectTaggingResponseBodyTagging {
	return s.Tagging
}

func (s *GetObjectTaggingResponseBody) SetTagging(v *GetObjectTaggingResponseBodyTagging) *GetObjectTaggingResponseBody {
	s.Tagging = v
	return s
}

func (s *GetObjectTaggingResponseBody) Validate() error {
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetObjectTaggingResponseBodyTagging struct {
	// The tag set of the target object.
	TagSet *TagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty"`
}

func (s GetObjectTaggingResponseBodyTagging) String() string {
	return dara.Prettify(s)
}

func (s GetObjectTaggingResponseBodyTagging) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingResponseBodyTagging) GetTagSet() *TagSet {
	return s.TagSet
}

func (s *GetObjectTaggingResponseBodyTagging) SetTagSet(v *TagSet) *GetObjectTaggingResponseBodyTagging {
	s.TagSet = v
	return s
}

func (s *GetObjectTaggingResponseBodyTagging) Validate() error {
	if s.TagSet != nil {
		if err := s.TagSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}
