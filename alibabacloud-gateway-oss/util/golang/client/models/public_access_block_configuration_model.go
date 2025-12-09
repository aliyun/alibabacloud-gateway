// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicAccessBlockConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetBlockPublicAccess(v bool) *PublicAccessBlockConfiguration
	GetBlockPublicAccess() *bool
}

type PublicAccessBlockConfiguration struct {
	// example:
	//
	// true
	BlockPublicAccess *bool `json:"BlockPublicAccess,omitempty" xml:"BlockPublicAccess,omitempty"`
}

func (s PublicAccessBlockConfiguration) String() string {
	return dara.Prettify(s)
}

func (s PublicAccessBlockConfiguration) GoString() string {
	return s.String()
}

func (s *PublicAccessBlockConfiguration) GetBlockPublicAccess() *bool {
	return s.BlockPublicAccess
}

func (s *PublicAccessBlockConfiguration) SetBlockPublicAccess(v bool) *PublicAccessBlockConfiguration {
	s.BlockPublicAccess = &v
	return s
}

func (s *PublicAccessBlockConfiguration) Validate() error {
	return dara.Validate(s)
}
