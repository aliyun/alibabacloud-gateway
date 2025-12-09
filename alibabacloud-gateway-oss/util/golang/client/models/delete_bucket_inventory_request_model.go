// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInventoryId(v string) *DeleteBucketInventoryRequest
	GetInventoryId() *string
}

type DeleteBucketInventoryRequest struct {
	// The name of the inventory that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// list1
	InventoryId *string `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
}

func (s DeleteBucketInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketInventoryRequest) GetInventoryId() *string {
	return s.InventoryId
}

func (s *DeleteBucketInventoryRequest) SetInventoryId(v string) *DeleteBucketInventoryRequest {
	s.InventoryId = &v
	return s
}

func (s *DeleteBucketInventoryRequest) Validate() error {
	return dara.Validate(s)
}
