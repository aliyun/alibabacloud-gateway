// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCORSConfiguration(v *CORSConfiguration) *PutBucketCorsRequest
	GetCORSConfiguration() *CORSConfiguration
}

type PutBucketCorsRequest struct {
	// The container that stores CORS rules.
	//
	// You can configure up to 10 CORS rules for a bucket. The XML message body in a request can be up to 16 KB in size.
	CORSConfiguration *CORSConfiguration `json:"CORSConfiguration,omitempty" xml:"CORSConfiguration,omitempty"`
}

func (s PutBucketCorsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCorsRequest) GoString() string {
	return s.String()
}

func (s *PutBucketCorsRequest) GetCORSConfiguration() *CORSConfiguration {
	return s.CORSConfiguration
}

func (s *PutBucketCorsRequest) SetCORSConfiguration(v *CORSConfiguration) *PutBucketCorsRequest {
	s.CORSConfiguration = v
	return s
}

func (s *PutBucketCorsRequest) Validate() error {
	if s.CORSConfiguration != nil {
		if err := s.CORSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
