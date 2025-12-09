// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInventoryConfiguration(v *InventoryConfiguration) *GetBucketInventoryResponseBody
	GetInventoryConfiguration() *InventoryConfiguration
}

type GetBucketInventoryResponseBody struct {
	// The inventory task configured for a bucket.
	InventoryConfiguration *InventoryConfiguration `json:"InventoryConfiguration,omitempty" xml:"InventoryConfiguration,omitempty"`
}

func (s GetBucketInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryResponseBody) GetInventoryConfiguration() *InventoryConfiguration {
	return s.InventoryConfiguration
}

func (s *GetBucketInventoryResponseBody) SetInventoryConfiguration(v *InventoryConfiguration) *GetBucketInventoryResponseBody {
	s.InventoryConfiguration = v
	return s
}

func (s *GetBucketInventoryResponseBody) Validate() error {
	if s.InventoryConfiguration != nil {
		if err := s.InventoryConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
