// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagSet interface {
	dara.Model
	String() string
	GoString() string
	SetTags(v []*Tag) *TagSet
	GetTags() []*Tag
}

type TagSet struct {
	Tags []*Tag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagSet) String() string {
	return dara.Prettify(s)
}

func (s TagSet) GoString() string {
	return s.String()
}

func (s *TagSet) GetTags() []*Tag {
	return s.Tags
}

func (s *TagSet) SetTags(v []*Tag) *TagSet {
	s.Tags = v
	return s
}

func (s *TagSet) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
