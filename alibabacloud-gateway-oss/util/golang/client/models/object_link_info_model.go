// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectLinkInfo interface {
	dara.Model
	String() string
	GoString() string
	SetPart(v []*ObjectLinkInfoPart) *ObjectLinkInfo
	GetPart() []*ObjectLinkInfoPart
}

type ObjectLinkInfo struct {
	Part []*ObjectLinkInfoPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s ObjectLinkInfo) String() string {
	return dara.Prettify(s)
}

func (s ObjectLinkInfo) GoString() string {
	return s.String()
}

func (s *ObjectLinkInfo) GetPart() []*ObjectLinkInfoPart {
	return s.Part
}

func (s *ObjectLinkInfo) SetPart(v []*ObjectLinkInfoPart) *ObjectLinkInfo {
	s.Part = v
	return s
}

func (s *ObjectLinkInfo) Validate() error {
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

type ObjectLinkInfoPart struct {
	// example:
	//
	// test.txt
	PartName *string `json:"PartName,omitempty" xml:"PartName,omitempty"`
	// example:
	//
	// 3
	PartNumber *int64 `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
}

func (s ObjectLinkInfoPart) String() string {
	return dara.Prettify(s)
}

func (s ObjectLinkInfoPart) GoString() string {
	return s.String()
}

func (s *ObjectLinkInfoPart) GetPartName() *string {
	return s.PartName
}

func (s *ObjectLinkInfoPart) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *ObjectLinkInfoPart) SetPartName(v string) *ObjectLinkInfoPart {
	s.PartName = &v
	return s
}

func (s *ObjectLinkInfoPart) SetPartNumber(v int64) *ObjectLinkInfoPart {
	s.PartNumber = &v
	return s
}

func (s *ObjectLinkInfoPart) Validate() error {
	return dara.Validate(s)
}
