// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualBucket interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *VirtualBucket
	GetName() *string
	SetRealBucket(v []*VirtualBucketRealBucket) *VirtualBucket
	GetRealBucket() []*VirtualBucketRealBucket
}

type VirtualBucket struct {
	// example:
	//
	// virtual-bucket
	Name       *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
	RealBucket []*VirtualBucketRealBucket `json:"RealBucket,omitempty" xml:"RealBucket,omitempty" type:"Repeated"`
}

func (s VirtualBucket) String() string {
	return dara.Prettify(s)
}

func (s VirtualBucket) GoString() string {
	return s.String()
}

func (s *VirtualBucket) GetName() *string {
	return s.Name
}

func (s *VirtualBucket) GetRealBucket() []*VirtualBucketRealBucket {
	return s.RealBucket
}

func (s *VirtualBucket) SetName(v string) *VirtualBucket {
	s.Name = &v
	return s
}

func (s *VirtualBucket) SetRealBucket(v []*VirtualBucketRealBucket) *VirtualBucket {
	s.RealBucket = v
	return s
}

func (s *VirtualBucket) Validate() error {
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

type VirtualBucketRealBucket struct {
	// example:
	//
	// real-bucket
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VirtualBucketRealBucket) String() string {
	return dara.Prettify(s)
}

func (s VirtualBucketRealBucket) GoString() string {
	return s.String()
}

func (s *VirtualBucketRealBucket) GetName() *string {
	return s.Name
}

func (s *VirtualBucketRealBucket) GetStatus() *string {
	return s.Status
}

func (s *VirtualBucketRealBucket) SetName(v string) *VirtualBucketRealBucket {
	s.Name = &v
	return s
}

func (s *VirtualBucketRealBucket) SetStatus(v string) *VirtualBucketRealBucket {
	s.Status = &v
	return s
}

func (s *VirtualBucketRealBucket) Validate() error {
	return dara.Validate(s)
}
