// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateBucketWormRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInitiateWormConfiguration(v *InitiateWormConfiguration) *InitiateBucketWormRequest
	GetInitiateWormConfiguration() *InitiateWormConfiguration
}

type InitiateBucketWormRequest struct {
	// The parameters for WORM initialization.
	InitiateWormConfiguration *InitiateWormConfiguration `json:"InitiateWormConfiguration,omitempty" xml:"InitiateWormConfiguration,omitempty"`
}

func (s InitiateBucketWormRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiateBucketWormRequest) GoString() string {
	return s.String()
}

func (s *InitiateBucketWormRequest) GetInitiateWormConfiguration() *InitiateWormConfiguration {
	return s.InitiateWormConfiguration
}

func (s *InitiateBucketWormRequest) SetInitiateWormConfiguration(v *InitiateWormConfiguration) *InitiateBucketWormRequest {
	s.InitiateWormConfiguration = v
	return s
}

func (s *InitiateBucketWormRequest) Validate() error {
	if s.InitiateWormConfiguration != nil {
		if err := s.InitiateWormConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
