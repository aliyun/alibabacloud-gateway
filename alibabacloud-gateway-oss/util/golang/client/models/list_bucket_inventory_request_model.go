// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListBucketInventoryRequest
	GetContinuationToken() *string
}

type ListBucketInventoryRequest struct {
	// Specify the start position of the list operation. You can obtain this token from the NextContinuationToken field of last ListBucketInventory\\"s result.
	//
	// example:
	//
	// test1.txt
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
}

func (s ListBucketInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListBucketInventoryRequest) SetContinuationToken(v string) *ListBucketInventoryRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListBucketInventoryRequest) Validate() error {
	return dara.Validate(s)
}
