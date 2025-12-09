// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRequestPaymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestPaymentConfiguration(v *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) *GetBucketRequestPaymentResponseBody
	GetRequestPaymentConfiguration() *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration
}

type GetBucketRequestPaymentResponseBody struct {
	// Indicates the container for the payer.
	RequestPaymentConfiguration *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration `json:"RequestPaymentConfiguration,omitempty" xml:"RequestPaymentConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketRequestPaymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRequestPaymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketRequestPaymentResponseBody) GetRequestPaymentConfiguration() *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration {
	return s.RequestPaymentConfiguration
}

func (s *GetBucketRequestPaymentResponseBody) SetRequestPaymentConfiguration(v *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) *GetBucketRequestPaymentResponseBody {
	s.RequestPaymentConfiguration = v
	return s
}

func (s *GetBucketRequestPaymentResponseBody) Validate() error {
	if s.RequestPaymentConfiguration != nil {
		if err := s.RequestPaymentConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration struct {
	// Indicates who pays the download and request fees.
	//
	// example:
	//
	// Requester
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
}

func (s GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) GetPayer() *string {
	return s.Payer
}

func (s *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) SetPayer(v string) *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration {
	s.Payer = &v
	return s
}

func (s *GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration) Validate() error {
	return dara.Validate(s)
}
