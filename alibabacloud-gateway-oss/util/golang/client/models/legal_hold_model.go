// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLegalHold interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *LegalHold
	GetStatus() *string
}

type LegalHold struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s LegalHold) String() string {
	return dara.Prettify(s)
}

func (s LegalHold) GoString() string {
	return s.String()
}

func (s *LegalHold) GetStatus() *string {
	return s.Status
}

func (s *LegalHold) SetStatus(v string) *LegalHold {
	s.Status = &v
	return s
}

func (s *LegalHold) Validate() error {
	return dara.Validate(s)
}
