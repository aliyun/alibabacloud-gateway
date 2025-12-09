// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPolicyForObjectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PutAccessPointPolicyForObjectProcessRequest
	GetBody() *string
}

type PutAccessPointPolicyForObjectProcessRequest struct {
	// The json format permission policies for an Object FC Access Point.
	//
	// example:
	//
	// {
	//
	// 	            "Version": "1",
	//
	// 	            "Statement": [{
	//
	// 		            "Effect": "Allow",
	//
	// 		            "Action": [
	//
	// 			            "oss:GetObject"
	//
	// 		            ],
	//
	// 		            "Principal": [
	//
	// 			            "237210000000000xxxx"
	//
	// 		            ],
	//
	// 		            "Resource": [
	//
	// 			            "acs:oss:cn-qingdao:1030700000xxxx:accesspointforobjectprocess/fc-test/object/*"
	//
	// 		            ]
	//
	// 	            }]
	//
	//             }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutAccessPointPolicyForObjectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPolicyForObjectProcessRequest) GoString() string {
	return s.String()
}

func (s *PutAccessPointPolicyForObjectProcessRequest) GetBody() *string {
	return s.Body
}

func (s *PutAccessPointPolicyForObjectProcessRequest) SetBody(v string) *PutAccessPointPolicyForObjectProcessRequest {
	s.Body = &v
	return s
}

func (s *PutAccessPointPolicyForObjectProcessRequest) Validate() error {
	return dara.Validate(s)
}
