// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirtualBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVirtualBucketConfiguration(v *VirtualBucket) *GetVirtualBucketResponseBody
	GetVirtualBucketConfiguration() *VirtualBucket
}

type GetVirtualBucketResponseBody struct {
	VirtualBucketConfiguration *VirtualBucket `json:"VirtualBucketConfiguration,omitempty" xml:"VirtualBucketConfiguration,omitempty"`
}

func (s GetVirtualBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVirtualBucketResponseBody) GoString() string {
	return s.String()
}

func (s *GetVirtualBucketResponseBody) GetVirtualBucketConfiguration() *VirtualBucket {
	return s.VirtualBucketConfiguration
}

func (s *GetVirtualBucketResponseBody) SetVirtualBucketConfiguration(v *VirtualBucket) *GetVirtualBucketResponseBody {
	s.VirtualBucketConfiguration = v
	return s
}

func (s *GetVirtualBucketResponseBody) Validate() error {
	if s.VirtualBucketConfiguration != nil {
		if err := s.VirtualBucketConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
