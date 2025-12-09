// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBucketDataRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssTargetRedundancyType(v string) *CreateBucketDataRedundancyTransitionRequest
	GetXOssTargetRedundancyType() *string
}

type CreateBucketDataRedundancyTransitionRequest struct {
	// The redundancy type to which you want to convert the bucket. You can only convert the redundancy type of a bucket from LRS to ZRS.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZRS
	XOssTargetRedundancyType *string `json:"x-oss-target-redundancy-type,omitempty" xml:"x-oss-target-redundancy-type,omitempty"`
}

func (s CreateBucketDataRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBucketDataRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *CreateBucketDataRedundancyTransitionRequest) GetXOssTargetRedundancyType() *string {
	return s.XOssTargetRedundancyType
}

func (s *CreateBucketDataRedundancyTransitionRequest) SetXOssTargetRedundancyType(v string) *CreateBucketDataRedundancyTransitionRequest {
	s.XOssTargetRedundancyType = &v
	return s
}

func (s *CreateBucketDataRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
