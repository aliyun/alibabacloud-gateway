// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutProcessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketProcessConfiguration(v *BucketProcessConfiguration) *PutProcessConfigurationRequest
	GetBucketProcessConfiguration() *BucketProcessConfiguration
}

type PutProcessConfigurationRequest struct {
	// Bucket Image Processing Configuration
	//
	// This parameter is required.
	BucketProcessConfiguration *BucketProcessConfiguration `json:"BucketProcessConfiguration,omitempty" xml:"BucketProcessConfiguration,omitempty"`
}

func (s PutProcessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutProcessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PutProcessConfigurationRequest) GetBucketProcessConfiguration() *BucketProcessConfiguration {
	return s.BucketProcessConfiguration
}

func (s *PutProcessConfigurationRequest) SetBucketProcessConfiguration(v *BucketProcessConfiguration) *PutProcessConfigurationRequest {
	s.BucketProcessConfiguration = v
	return s
}

func (s *PutProcessConfigurationRequest) Validate() error {
	if s.BucketProcessConfiguration != nil {
		if err := s.BucketProcessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
