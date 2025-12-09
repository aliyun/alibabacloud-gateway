// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonPrefix interface {
	dara.Model
	String() string
	GoString() string
	SetPrefix(v string) *CommonPrefix
	GetPrefix() *string
}

type CommonPrefix struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s CommonPrefix) String() string {
	return dara.Prettify(s)
}

func (s CommonPrefix) GoString() string {
	return s.String()
}

func (s *CommonPrefix) GetPrefix() *string {
	return s.Prefix
}

func (s *CommonPrefix) SetPrefix(v string) *CommonPrefix {
	s.Prefix = &v
	return s
}

func (s *CommonPrefix) Validate() error {
	return dara.Validate(s)
}
