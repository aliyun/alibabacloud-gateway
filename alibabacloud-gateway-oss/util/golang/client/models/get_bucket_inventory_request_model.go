// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInventoryId(v string) *GetBucketInventoryRequest
	GetInventoryId() *string
}

type GetBucketInventoryRequest struct {
	// The name of the inventory to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// list1
	InventoryId *string `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
}

func (s GetBucketInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryRequest) GetInventoryId() *string {
	return s.InventoryId
}

func (s *GetBucketInventoryRequest) SetInventoryId(v string) *GetBucketInventoryRequest {
	s.InventoryId = &v
	return s
}

func (s *GetBucketInventoryRequest) Validate() error {
	return dara.Validate(s)
}
