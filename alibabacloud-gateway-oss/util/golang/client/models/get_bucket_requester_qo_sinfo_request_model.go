// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRequesterQoSInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQosRequester(v string) *GetBucketRequesterQoSInfoRequest
	GetQosRequester() *string
}

type GetBucketRequesterQoSInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 26753xxxxxxxx14340
	QosRequester *string `json:"qosRequester,omitempty" xml:"qosRequester,omitempty"`
}

func (s GetBucketRequesterQoSInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRequesterQoSInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBucketRequesterQoSInfoRequest) GetQosRequester() *string {
	return s.QosRequester
}

func (s *GetBucketRequesterQoSInfoRequest) SetQosRequester(v string) *GetBucketRequesterQoSInfoRequest {
	s.QosRequester = &v
	return s
}

func (s *GetBucketRequesterQoSInfoRequest) Validate() error {
	return dara.Validate(s)
}
