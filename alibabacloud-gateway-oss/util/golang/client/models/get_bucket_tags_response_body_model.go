// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTagging(v *GetBucketTagsResponseBodyTagging) *GetBucketTagsResponseBody
	GetTagging() *GetBucketTagsResponseBodyTagging
}

type GetBucketTagsResponseBody struct {
	// The container that stores the returned tags of the bucket.
	//
	// > If no tags are configured for the bucket, an XML message body is returned in which the Tagging element is empty.
	Tagging *GetBucketTagsResponseBodyTagging `json:"Tagging,omitempty" xml:"Tagging,omitempty" type:"Struct"`
}

func (s GetBucketTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketTagsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketTagsResponseBody) GetTagging() *GetBucketTagsResponseBodyTagging {
	return s.Tagging
}

func (s *GetBucketTagsResponseBody) SetTagging(v *GetBucketTagsResponseBodyTagging) *GetBucketTagsResponseBody {
	s.Tagging = v
	return s
}

func (s *GetBucketTagsResponseBody) Validate() error {
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketTagsResponseBodyTagging struct {
	// The container that stores the returned tags of the bucket.
	TagSet *TagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty"`
}

func (s GetBucketTagsResponseBodyTagging) String() string {
	return dara.Prettify(s)
}

func (s GetBucketTagsResponseBodyTagging) GoString() string {
	return s.String()
}

func (s *GetBucketTagsResponseBodyTagging) GetTagSet() *TagSet {
	return s.TagSet
}

func (s *GetBucketTagsResponseBodyTagging) SetTagSet(v *TagSet) *GetBucketTagsResponseBodyTagging {
	s.TagSet = v
	return s
}

func (s *GetBucketTagsResponseBodyTagging) Validate() error {
	if s.TagSet != nil {
		if err := s.TagSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}
