// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileGroup interface {
	dara.Model
	String() string
	GoString() string
	SetPart(v []*CreateFileGroupPart) *CreateFileGroup
	GetPart() []*CreateFileGroupPart
}

type CreateFileGroup struct {
	Part []*CreateFileGroupPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s CreateFileGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateFileGroup) GoString() string {
	return s.String()
}

func (s *CreateFileGroup) GetPart() []*CreateFileGroupPart {
	return s.Part
}

func (s *CreateFileGroup) SetPart(v []*CreateFileGroupPart) *CreateFileGroup {
	s.Part = v
	return s
}

func (s *CreateFileGroup) Validate() error {
	if s.Part != nil {
		for _, item := range s.Part {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFileGroupPart struct {
	// example:
	//
	// "EB327B57BB20D17C293A966115FE27BD"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// test.txt
	PartName *string `json:"PartName,omitempty" xml:"PartName,omitempty"`
	// example:
	//
	// 3
	PartNumber *int64 `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
}

func (s CreateFileGroupPart) String() string {
	return dara.Prettify(s)
}

func (s CreateFileGroupPart) GoString() string {
	return s.String()
}

func (s *CreateFileGroupPart) GetETag() *string {
	return s.ETag
}

func (s *CreateFileGroupPart) GetPartName() *string {
	return s.PartName
}

func (s *CreateFileGroupPart) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *CreateFileGroupPart) SetETag(v string) *CreateFileGroupPart {
	s.ETag = &v
	return s
}

func (s *CreateFileGroupPart) SetPartName(v string) *CreateFileGroupPart {
	s.PartName = &v
	return s
}

func (s *CreateFileGroupPart) SetPartNumber(v int64) *CreateFileGroupPart {
	s.PartNumber = &v
	return s
}

func (s *CreateFileGroupPart) Validate() error {
	return dara.Validate(s)
}
