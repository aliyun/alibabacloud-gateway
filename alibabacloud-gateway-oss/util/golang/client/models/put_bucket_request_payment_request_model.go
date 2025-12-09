// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRequestPaymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestPaymentConfiguration(v *RequestPaymentConfiguration) *PutBucketRequestPaymentRequest
	GetRequestPaymentConfiguration() *RequestPaymentConfiguration
}

type PutBucketRequestPaymentRequest struct {
	// The container that stores pay-by-requester configurations.
	RequestPaymentConfiguration *RequestPaymentConfiguration `json:"RequestPaymentConfiguration,omitempty" xml:"RequestPaymentConfiguration,omitempty"`
}

func (s PutBucketRequestPaymentRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRequestPaymentRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequestPaymentRequest) GetRequestPaymentConfiguration() *RequestPaymentConfiguration {
	return s.RequestPaymentConfiguration
}

func (s *PutBucketRequestPaymentRequest) SetRequestPaymentConfiguration(v *RequestPaymentConfiguration) *PutBucketRequestPaymentRequest {
	s.RequestPaymentConfiguration = v
	return s
}

func (s *PutBucketRequestPaymentRequest) Validate() error {
	if s.RequestPaymentConfiguration != nil {
		if err := s.RequestPaymentConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
