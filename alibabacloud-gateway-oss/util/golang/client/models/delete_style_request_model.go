// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStyleName(v string) *DeleteStyleRequest
	GetStyleName() *string
}

type DeleteStyleRequest struct {
	// The name of the image style.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagestyle
	StyleName *string `json:"styleName,omitempty" xml:"styleName,omitempty"`
}

func (s DeleteStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStyleRequest) GoString() string {
	return s.String()
}

func (s *DeleteStyleRequest) GetStyleName() *string {
	return s.StyleName
}

func (s *DeleteStyleRequest) SetStyleName(v string) *DeleteStyleRequest {
	s.StyleName = &v
	return s
}

func (s *DeleteStyleRequest) Validate() error {
	return dara.Validate(s)
}
