// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagging interface {
	dara.Model
	String() string
	GoString() string
	SetTagSet(v *TagSet) *Tagging
	GetTagSet() *TagSet
}

type Tagging struct {
	TagSet *TagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty"`
}

func (s Tagging) String() string {
	return dara.Prettify(s)
}

func (s Tagging) GoString() string {
	return s.String()
}

func (s *Tagging) GetTagSet() *TagSet {
	return s.TagSet
}

func (s *Tagging) SetTagSet(v *TagSet) *Tagging {
	s.TagSet = v
	return s
}

func (s *Tagging) Validate() error {
	if s.TagSet != nil {
		if err := s.TagSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}
