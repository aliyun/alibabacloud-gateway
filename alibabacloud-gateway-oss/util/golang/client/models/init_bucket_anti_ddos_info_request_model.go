// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitBucketAntiDDosInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAntiDDOSConfiguration(v *BucketAntiDDOSConfiguration) *InitBucketAntiDDosInfoRequest
	GetAntiDDOSConfiguration() *BucketAntiDDOSConfiguration
}

type InitBucketAntiDDosInfoRequest struct {
	// The container that stores the configurations of Anti-DDoS instances.
	AntiDDOSConfiguration *BucketAntiDDOSConfiguration `json:"AntiDDOSConfiguration,omitempty" xml:"AntiDDOSConfiguration,omitempty"`
}

func (s InitBucketAntiDDosInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s InitBucketAntiDDosInfoRequest) GoString() string {
	return s.String()
}

func (s *InitBucketAntiDDosInfoRequest) GetAntiDDOSConfiguration() *BucketAntiDDOSConfiguration {
	return s.AntiDDOSConfiguration
}

func (s *InitBucketAntiDDosInfoRequest) SetAntiDDOSConfiguration(v *BucketAntiDDOSConfiguration) *InitBucketAntiDDosInfoRequest {
	s.AntiDDOSConfiguration = v
	return s
}

func (s *InitBucketAntiDDosInfoRequest) Validate() error {
	if s.AntiDDOSConfiguration != nil {
		if err := s.AntiDDOSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
