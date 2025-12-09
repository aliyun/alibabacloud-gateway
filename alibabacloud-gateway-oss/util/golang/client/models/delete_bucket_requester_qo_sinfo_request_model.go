// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketRequesterQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQosRequester(v string) *DeleteBucketRequesterQoSInfoRequest
	GetQosRequester() *string
}

type DeleteBucketRequesterQoSInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 26753xxxxxxxx14340
	QosRequester *string `json:"qosRequester,omitempty" xml:"qosRequester,omitempty"`
}

func (s DeleteBucketRequesterQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketRequesterQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketRequesterQoSInfoRequest) GetQosRequester() *string {
	return s.QosRequester
}

func (s *DeleteBucketRequesterQoSInfoRequest) SetQosRequester(v string) *DeleteBucketRequesterQoSInfoRequest {
	s.QosRequester = &v
	return s
}

func (s *DeleteBucketRequesterQoSInfoRequest) Validate() error {
	return dara.Validate(s)
}
