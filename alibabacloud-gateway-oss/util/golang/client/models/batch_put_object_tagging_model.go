// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutObjectTagging interface {
	dara.Model
	String() string
	GoString() string
	SetTagging(v *TagSet) *BatchPutObjectTagging
	GetTagging() *TagSet
}

type BatchPutObjectTagging struct {
	Tagging *TagSet `json:"tagging,omitempty" xml:"tagging,omitempty"`
}

func (s BatchPutObjectTagging) String() string {
	return dara.Prettify(s)
}

func (s BatchPutObjectTagging) GoString() string {
	return s.String()
}

func (s *BatchPutObjectTagging) GetTagging() *TagSet {
	return s.Tagging
}

func (s *BatchPutObjectTagging) SetTagging(v *TagSet) *BatchPutObjectTagging {
	s.Tagging = v
	return s
}

func (s *BatchPutObjectTagging) Validate() error {
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			return err
		}
	}
	return nil
}
