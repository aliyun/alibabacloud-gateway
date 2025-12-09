// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetStyleList(v *ListStyleResponseBodyStyleList) *ListStyleResponseBody
	GetStyleList() *ListStyleResponseBodyStyleList
}

type ListStyleResponseBody struct {
	// The container that was used to query the information about image styles.
	StyleList *ListStyleResponseBodyStyleList `json:"StyleList,omitempty" xml:"StyleList,omitempty" type:"Struct"`
}

func (s ListStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ListStyleResponseBody) GetStyleList() *ListStyleResponseBodyStyleList {
	return s.StyleList
}

func (s *ListStyleResponseBody) SetStyleList(v *ListStyleResponseBodyStyleList) *ListStyleResponseBody {
	s.StyleList = v
	return s
}

func (s *ListStyleResponseBody) Validate() error {
	if s.StyleList != nil {
		if err := s.StyleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStyleResponseBodyStyleList struct {
	// The list of styles.
	Style []*StyleInfo `json:"Style,omitempty" xml:"Style,omitempty" type:"Repeated"`
}

func (s ListStyleResponseBodyStyleList) String() string {
	return dara.Prettify(s)
}

func (s ListStyleResponseBodyStyleList) GoString() string {
	return s.String()
}

func (s *ListStyleResponseBodyStyleList) GetStyle() []*StyleInfo {
	return s.Style
}

func (s *ListStyleResponseBodyStyleList) SetStyle(v []*StyleInfo) *ListStyleResponseBodyStyleList {
	s.Style = v
	return s
}

func (s *ListStyleResponseBodyStyleList) Validate() error {
	if s.Style != nil {
		for _, item := range s.Style {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
