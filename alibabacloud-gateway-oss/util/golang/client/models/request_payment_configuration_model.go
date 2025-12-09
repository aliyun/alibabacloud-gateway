// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestPaymentConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetPayer(v string) *RequestPaymentConfiguration
	GetPayer() *string
}

type RequestPaymentConfiguration struct {
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
}

func (s RequestPaymentConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RequestPaymentConfiguration) GoString() string {
	return s.String()
}

func (s *RequestPaymentConfiguration) GetPayer() *string {
	return s.Payer
}

func (s *RequestPaymentConfiguration) SetPayer(v string) *RequestPaymentConfiguration {
	s.Payer = &v
	return s
}

func (s *RequestPaymentConfiguration) Validate() error {
	return dara.Validate(s)
}
