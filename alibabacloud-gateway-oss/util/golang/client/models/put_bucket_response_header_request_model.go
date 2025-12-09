// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResponseHeaderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderConfiguration(v *ResponseHeaderConfiguration) *PutBucketResponseHeaderRequest
	GetResponseHeaderConfiguration() *ResponseHeaderConfiguration
}

type PutBucketResponseHeaderRequest struct {
	ResponseHeaderConfiguration *ResponseHeaderConfiguration `json:"ResponseHeaderConfiguration,omitempty" xml:"ResponseHeaderConfiguration,omitempty"`
}

func (s PutBucketResponseHeaderRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResponseHeaderRequest) GoString() string {
	return s.String()
}

func (s *PutBucketResponseHeaderRequest) GetResponseHeaderConfiguration() *ResponseHeaderConfiguration {
	return s.ResponseHeaderConfiguration
}

func (s *PutBucketResponseHeaderRequest) SetResponseHeaderConfiguration(v *ResponseHeaderConfiguration) *PutBucketResponseHeaderRequest {
	s.ResponseHeaderConfiguration = v
	return s
}

func (s *PutBucketResponseHeaderRequest) Validate() error {
	if s.ResponseHeaderConfiguration != nil {
		if err := s.ResponseHeaderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
