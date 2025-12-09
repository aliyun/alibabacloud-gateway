// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketDataRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssRedundancyTransitionTaskid(v string) *DeleteBucketDataRedundancyTransitionRequest
	GetXOssRedundancyTransitionTaskid() *string
}

type DeleteBucketDataRedundancyTransitionRequest struct {
	// The ID of the redundancy type change task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4be5beb0f74f490186311b268bf6****
	XOssRedundancyTransitionTaskid *string `json:"x-oss-redundancy-transition-taskid,omitempty" xml:"x-oss-redundancy-transition-taskid,omitempty"`
}

func (s DeleteBucketDataRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketDataRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketDataRedundancyTransitionRequest) GetXOssRedundancyTransitionTaskid() *string {
	return s.XOssRedundancyTransitionTaskid
}

func (s *DeleteBucketDataRedundancyTransitionRequest) SetXOssRedundancyTransitionTaskid(v string) *DeleteBucketDataRedundancyTransitionRequest {
	s.XOssRedundancyTransitionTaskid = &v
	return s
}

func (s *DeleteBucketDataRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
