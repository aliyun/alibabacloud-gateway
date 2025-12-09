// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteBucketWormRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWormId(v string) *CompleteBucketWormRequest
	GetWormId() *string
}

type CompleteBucketWormRequest struct {
	// The ID of the retention policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1666E2CFB2B3418****
	WormId *string `json:"wormId,omitempty" xml:"wormId,omitempty"`
}

func (s CompleteBucketWormRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteBucketWormRequest) GoString() string {
	return s.String()
}

func (s *CompleteBucketWormRequest) GetWormId() *string {
	return s.WormId
}

func (s *CompleteBucketWormRequest) SetWormId(v string) *CompleteBucketWormRequest {
	s.WormId = &v
	return s
}

func (s *CompleteBucketWormRequest) Validate() error {
	return dara.Validate(s)
}
