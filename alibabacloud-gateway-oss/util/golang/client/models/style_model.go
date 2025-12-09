// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStyle interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *Style
	GetContent() *string
}

type Style struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s Style) String() string {
	return dara.Prettify(s)
}

func (s Style) GoString() string {
	return s.String()
}

func (s *Style) GetContent() *string {
	return s.Content
}

func (s *Style) SetContent(v string) *Style {
	s.Content = &v
	return s
}

func (s *Style) Validate() error {
	return dara.Validate(s)
}
