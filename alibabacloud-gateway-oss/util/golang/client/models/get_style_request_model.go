// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStyleName(v string) *GetStyleRequest
	GetStyleName() *string
}

type GetStyleRequest struct {
	// The name of the image style.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagestyle
	StyleName *string `json:"styleName,omitempty" xml:"styleName,omitempty"`
}

func (s GetStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStyleRequest) GoString() string {
	return s.String()
}

func (s *GetStyleRequest) GetStyleName() *string {
	return s.StyleName
}

func (s *GetStyleRequest) SetStyleName(v string) *GetStyleRequest {
	s.StyleName = &v
	return s
}

func (s *GetStyleRequest) Validate() error {
	return dara.Validate(s)
}
