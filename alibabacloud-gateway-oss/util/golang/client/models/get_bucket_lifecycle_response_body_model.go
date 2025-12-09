// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLifecycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleConfiguration(v *LifecycleConfiguration) *GetBucketLifecycleResponseBody
	GetLifecycleConfiguration() *LifecycleConfiguration
}

type GetBucketLifecycleResponseBody struct {
	// The container that stores the lifecycle rules configured for the bucket.
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitempty" xml:"LifecycleConfiguration,omitempty"`
}

func (s GetBucketLifecycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponseBody) GetLifecycleConfiguration() *LifecycleConfiguration {
	return s.LifecycleConfiguration
}

func (s *GetBucketLifecycleResponseBody) SetLifecycleConfiguration(v *LifecycleConfiguration) *GetBucketLifecycleResponseBody {
	s.LifecycleConfiguration = v
	return s
}

func (s *GetBucketLifecycleResponseBody) Validate() error {
	if s.LifecycleConfiguration != nil {
		if err := s.LifecycleConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
