// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBucketAntiDDosInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAntiDDOSConfiguration(v *BucketAntiDDOSConfiguration) *UpdateBucketAntiDDosInfoRequest
	GetAntiDDOSConfiguration() *BucketAntiDDOSConfiguration
}

type UpdateBucketAntiDDosInfoRequest struct {
	// The container that stores the configurations of Anti-DDoS instances.
	AntiDDOSConfiguration *BucketAntiDDOSConfiguration `json:"AntiDDOSConfiguration,omitempty" xml:"AntiDDOSConfiguration,omitempty"`
}

func (s UpdateBucketAntiDDosInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBucketAntiDDosInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateBucketAntiDDosInfoRequest) GetAntiDDOSConfiguration() *BucketAntiDDOSConfiguration {
	return s.AntiDDOSConfiguration
}

func (s *UpdateBucketAntiDDosInfoRequest) SetAntiDDOSConfiguration(v *BucketAntiDDOSConfiguration) *UpdateBucketAntiDDosInfoRequest {
	s.AntiDDOSConfiguration = v
	return s
}

func (s *UpdateBucketAntiDDosInfoRequest) Validate() error {
	if s.AntiDDOSConfiguration != nil {
		if err := s.AntiDDOSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
