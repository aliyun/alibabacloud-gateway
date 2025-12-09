// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDataRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssRedundancyTransitionTaskid(v string) *GetBucketDataRedundancyTransitionRequest
	GetXOssRedundancyTransitionTaskid() *string
}

type GetBucketDataRedundancyTransitionRequest struct {
	// The ID of the redundancy change task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 751f5243f8ac4ae89f34726534d1****
	XOssRedundancyTransitionTaskid *string `json:"x-oss-redundancy-transition-taskid,omitempty" xml:"x-oss-redundancy-transition-taskid,omitempty"`
}

func (s GetBucketDataRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDataRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *GetBucketDataRedundancyTransitionRequest) GetXOssRedundancyTransitionTaskid() *string {
	return s.XOssRedundancyTransitionTaskid
}

func (s *GetBucketDataRedundancyTransitionRequest) SetXOssRedundancyTransitionTaskid(v string) *GetBucketDataRedundancyTransitionRequest {
	s.XOssRedundancyTransitionTaskid = &v
	return s
}

func (s *GetBucketDataRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
