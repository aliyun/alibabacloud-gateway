// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipleObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteResult(v *DeleteMultipleObjectsResponseBodyDeleteResult) *DeleteMultipleObjectsResponseBody
	GetDeleteResult() *DeleteMultipleObjectsResponseBodyDeleteResult
}

type DeleteMultipleObjectsResponseBody struct {
	DeleteResult *DeleteMultipleObjectsResponseBodyDeleteResult `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty" type:"Struct"`
}

func (s DeleteMultipleObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipleObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsResponseBody) GetDeleteResult() *DeleteMultipleObjectsResponseBodyDeleteResult {
	return s.DeleteResult
}

func (s *DeleteMultipleObjectsResponseBody) SetDeleteResult(v *DeleteMultipleObjectsResponseBodyDeleteResult) *DeleteMultipleObjectsResponseBody {
	s.DeleteResult = v
	return s
}

func (s *DeleteMultipleObjectsResponseBody) Validate() error {
	if s.DeleteResult != nil {
		if err := s.DeleteResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMultipleObjectsResponseBodyDeleteResult struct {
	Deleted      []*DeletedObject `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Repeated"`
	EncodingType *string          `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
}

func (s DeleteMultipleObjectsResponseBodyDeleteResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipleObjectsResponseBodyDeleteResult) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsResponseBodyDeleteResult) GetDeleted() []*DeletedObject {
	return s.Deleted
}

func (s *DeleteMultipleObjectsResponseBodyDeleteResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *DeleteMultipleObjectsResponseBodyDeleteResult) SetDeleted(v []*DeletedObject) *DeleteMultipleObjectsResponseBodyDeleteResult {
	s.Deleted = v
	return s
}

func (s *DeleteMultipleObjectsResponseBodyDeleteResult) SetEncodingType(v string) *DeleteMultipleObjectsResponseBodyDeleteResult {
	s.EncodingType = &v
	return s
}

func (s *DeleteMultipleObjectsResponseBodyDeleteResult) Validate() error {
	if s.Deleted != nil {
		for _, item := range s.Deleted {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
