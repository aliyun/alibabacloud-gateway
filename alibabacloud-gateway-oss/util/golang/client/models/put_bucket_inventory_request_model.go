// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInventoryConfiguration(v *InventoryConfiguration) *PutBucketInventoryRequest
	GetInventoryConfiguration() *InventoryConfiguration
	SetInventoryId(v string) *PutBucketInventoryRequest
	GetInventoryId() *string
}

type PutBucketInventoryRequest struct {
	// The container that stores the Inventory configuration.
	InventoryConfiguration *InventoryConfiguration `json:"InventoryConfiguration,omitempty" xml:"InventoryConfiguration,omitempty"`
	// The name of the inventory.
	//
	// This parameter is required.
	//
	// example:
	//
	// report1
	InventoryId *string `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
}

func (s PutBucketInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *PutBucketInventoryRequest) GetInventoryConfiguration() *InventoryConfiguration {
	return s.InventoryConfiguration
}

func (s *PutBucketInventoryRequest) GetInventoryId() *string {
	return s.InventoryId
}

func (s *PutBucketInventoryRequest) SetInventoryConfiguration(v *InventoryConfiguration) *PutBucketInventoryRequest {
	s.InventoryConfiguration = v
	return s
}

func (s *PutBucketInventoryRequest) SetInventoryId(v string) *PutBucketInventoryRequest {
	s.InventoryId = &v
	return s
}

func (s *PutBucketInventoryRequest) Validate() error {
	if s.InventoryConfiguration != nil {
		if err := s.InventoryConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
