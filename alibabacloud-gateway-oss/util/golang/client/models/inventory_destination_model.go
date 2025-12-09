// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInventoryDestination interface {
	dara.Model
	String() string
	GoString() string
	SetOSSBucketDestination(v *InventoryOSSBucketDestination) *InventoryDestination
	GetOSSBucketDestination() *InventoryOSSBucketDestination
}

type InventoryDestination struct {
	OSSBucketDestination *InventoryOSSBucketDestination `json:"OSSBucketDestination,omitempty" xml:"OSSBucketDestination,omitempty"`
}

func (s InventoryDestination) String() string {
	return dara.Prettify(s)
}

func (s InventoryDestination) GoString() string {
	return s.String()
}

func (s *InventoryDestination) GetOSSBucketDestination() *InventoryOSSBucketDestination {
	return s.OSSBucketDestination
}

func (s *InventoryDestination) SetOSSBucketDestination(v *InventoryOSSBucketDestination) *InventoryDestination {
	s.OSSBucketDestination = v
	return s
}

func (s *InventoryDestination) Validate() error {
	if s.OSSBucketDestination != nil {
		if err := s.OSSBucketDestination.Validate(); err != nil {
			return err
		}
	}
	return nil
}
