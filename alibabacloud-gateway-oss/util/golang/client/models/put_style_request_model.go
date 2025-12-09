// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStyle(v *Style) *PutStyleRequest
	GetStyle() *Style
	SetCategory(v string) *PutStyleRequest
	GetCategory() *string
	SetStyleName(v string) *PutStyleRequest
	GetStyleName() *string
}

type PutStyleRequest struct {
	// The style content.
	Style *Style `json:"Style,omitempty" xml:"Style,omitempty"`
	// The category of the style.
	//
	// example:
	//
	// image
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The name of the image style.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagestyle
	StyleName *string `json:"styleName,omitempty" xml:"styleName,omitempty"`
}

func (s PutStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutStyleRequest) GoString() string {
	return s.String()
}

func (s *PutStyleRequest) GetStyle() *Style {
	return s.Style
}

func (s *PutStyleRequest) GetCategory() *string {
	return s.Category
}

func (s *PutStyleRequest) GetStyleName() *string {
	return s.StyleName
}

func (s *PutStyleRequest) SetStyle(v *Style) *PutStyleRequest {
	s.Style = v
	return s
}

func (s *PutStyleRequest) SetCategory(v string) *PutStyleRequest {
	s.Category = &v
	return s
}

func (s *PutStyleRequest) SetStyleName(v string) *PutStyleRequest {
	s.StyleName = &v
	return s
}

func (s *PutStyleRequest) Validate() error {
	if s.Style != nil {
		if err := s.Style.Validate(); err != nil {
			return err
		}
	}
	return nil
}
