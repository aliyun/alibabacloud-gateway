// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLifecycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleConfiguration(v *LifecycleConfiguration) *PutBucketLifecycleRequest
	GetLifecycleConfiguration() *LifecycleConfiguration
}

type PutBucketLifecycleRequest struct {
	// The container that stores the lifecycle configuration.
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitempty" xml:"LifecycleConfiguration,omitempty"`
}

func (s PutBucketLifecycleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLifecycleRequest) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleRequest) GetLifecycleConfiguration() *LifecycleConfiguration {
	return s.LifecycleConfiguration
}

func (s *PutBucketLifecycleRequest) SetLifecycleConfiguration(v *LifecycleConfiguration) *PutBucketLifecycleRequest {
	s.LifecycleConfiguration = v
	return s
}

func (s *PutBucketLifecycleRequest) Validate() error {
	if s.LifecycleConfiguration != nil {
		if err := s.LifecycleConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
