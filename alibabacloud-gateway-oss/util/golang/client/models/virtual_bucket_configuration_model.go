// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualBucketConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetRealBucket(v []*VirtualBucketConfigurationRealBucket) *VirtualBucketConfiguration
	GetRealBucket() []*VirtualBucketConfigurationRealBucket
}

type VirtualBucketConfiguration struct {
	RealBucket []*VirtualBucketConfigurationRealBucket `json:"RealBucket,omitempty" xml:"RealBucket,omitempty" type:"Repeated"`
}

func (s VirtualBucketConfiguration) String() string {
	return dara.Prettify(s)
}

func (s VirtualBucketConfiguration) GoString() string {
	return s.String()
}

func (s *VirtualBucketConfiguration) GetRealBucket() []*VirtualBucketConfigurationRealBucket {
	return s.RealBucket
}

func (s *VirtualBucketConfiguration) SetRealBucket(v []*VirtualBucketConfigurationRealBucket) *VirtualBucketConfiguration {
	s.RealBucket = v
	return s
}

func (s *VirtualBucketConfiguration) Validate() error {
	if s.RealBucket != nil {
		for _, item := range s.RealBucket {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VirtualBucketConfigurationRealBucket struct {
	// example:
	//
	// realbucket-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s VirtualBucketConfigurationRealBucket) String() string {
	return dara.Prettify(s)
}

func (s VirtualBucketConfigurationRealBucket) GoString() string {
	return s.String()
}

func (s *VirtualBucketConfigurationRealBucket) GetName() *string {
	return s.Name
}

func (s *VirtualBucketConfigurationRealBucket) SetName(v string) *VirtualBucketConfigurationRealBucket {
	s.Name = &v
	return s
}

func (s *VirtualBucketConfigurationRealBucket) Validate() error {
	return dara.Validate(s)
}
