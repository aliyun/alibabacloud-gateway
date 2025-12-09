// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualBucketResult interface {
	dara.Model
	String() string
	GoString() string
	SetVirtualBucket(v []*VirtualBucket) *ListVirtualBucketResult
	GetVirtualBucket() []*VirtualBucket
}

type ListVirtualBucketResult struct {
	VirtualBucket []*VirtualBucket `json:"VirtualBucket,omitempty" xml:"VirtualBucket,omitempty" type:"Repeated"`
}

func (s ListVirtualBucketResult) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBucketResult) GoString() string {
	return s.String()
}

func (s *ListVirtualBucketResult) GetVirtualBucket() []*VirtualBucket {
	return s.VirtualBucket
}

func (s *ListVirtualBucketResult) SetVirtualBucket(v []*VirtualBucket) *ListVirtualBucketResult {
	s.VirtualBucket = v
	return s
}

func (s *ListVirtualBucketResult) Validate() error {
	if s.VirtualBucket != nil {
		for _, item := range s.VirtualBucket {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
