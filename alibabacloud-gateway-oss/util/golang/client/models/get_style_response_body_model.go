// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetStyle(v *StyleInfo) *GetStyleResponseBody
	GetStyle() *StyleInfo
}

type GetStyleResponseBody struct {
	// The container in which the queried image styles are stored.
	Style *StyleInfo `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GetStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GetStyleResponseBody) GetStyle() *StyleInfo {
	return s.Style
}

func (s *GetStyleResponseBody) SetStyle(v *StyleInfo) *GetStyleResponseBody {
	s.Style = v
	return s
}

func (s *GetStyleResponseBody) Validate() error {
	if s.Style != nil {
		if err := s.Style.Validate(); err != nil {
			return err
		}
	}
	return nil
}
