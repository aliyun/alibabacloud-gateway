// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectLegalHoldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutObjectLegalHoldRequestBody) *PutObjectLegalHoldRequest
	GetBody() *PutObjectLegalHoldRequestBody
}

type PutObjectLegalHoldRequest struct {
	Body *PutObjectLegalHoldRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PutObjectLegalHoldRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLegalHoldRequest) GoString() string {
	return s.String()
}

func (s *PutObjectLegalHoldRequest) GetBody() *PutObjectLegalHoldRequestBody {
	return s.Body
}

func (s *PutObjectLegalHoldRequest) SetBody(v *PutObjectLegalHoldRequestBody) *PutObjectLegalHoldRequest {
	s.Body = v
	return s
}

func (s *PutObjectLegalHoldRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutObjectLegalHoldRequestBody struct {
	LegalHold *LegalHold `json:"LegalHold,omitempty" xml:"LegalHold,omitempty"`
}

func (s PutObjectLegalHoldRequestBody) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLegalHoldRequestBody) GoString() string {
	return s.String()
}

func (s *PutObjectLegalHoldRequestBody) GetLegalHold() *LegalHold {
	return s.LegalHold
}

func (s *PutObjectLegalHoldRequestBody) SetLegalHold(v *LegalHold) *PutObjectLegalHoldRequestBody {
	s.LegalHold = v
	return s
}

func (s *PutObjectLegalHoldRequestBody) Validate() error {
	if s.LegalHold != nil {
		if err := s.LegalHold.Validate(); err != nil {
			return err
		}
	}
	return nil
}
