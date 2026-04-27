// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectLegalHoldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLegalHold(v *LegalHold) *GetObjectLegalHoldResponseBody
	GetLegalHold() *LegalHold
}

type GetObjectLegalHoldResponseBody struct {
	LegalHold *LegalHold `json:"LegalHold,omitempty" xml:"LegalHold,omitempty"`
}

func (s GetObjectLegalHoldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectLegalHoldResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectLegalHoldResponseBody) GetLegalHold() *LegalHold {
	return s.LegalHold
}

func (s *GetObjectLegalHoldResponseBody) SetLegalHold(v *LegalHold) *GetObjectLegalHoldResponseBody {
	s.LegalHold = v
	return s
}

func (s *GetObjectLegalHoldResponseBody) Validate() error {
	if s.LegalHold != nil {
		if err := s.LegalHold.Validate(); err != nil {
			return err
		}
	}
	return nil
}
