// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PutAccessPointPolicyRequest
	GetBody() *string
}

type PutAccessPointPolicyRequest struct {
	// The configurations of the access point policy.
	//
	// example:
	//
	// {
	//
	//    "Version":"1",
	//
	//    "Statement":[
	//
	//    {
	//
	//      "Action":[
	//
	//        "oss:PutObject",
	//
	//        "oss:GetObject"
	//
	//     ],
	//
	//     "Effect":"Deny",
	//
	//     "Principal":["27737962156157xxxx"],
	//
	//     "Resource":[
	//
	//        "acs:oss:ap-southeast-2:111933544165xxxx:accesspoint/$ap-01",
	//
	//        "acs:oss:ap-southeast-2:111933544165xxxx:accesspoint/$ap-01/object/*"
	//
	//      ]
	//
	//    }
	//
	//   ]
	//
	//  }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutAccessPointPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutAccessPointPolicyRequest) GetBody() *string {
	return s.Body
}

func (s *PutAccessPointPolicyRequest) SetBody(v string) *PutAccessPointPolicyRequest {
	s.Body = &v
	return s
}

func (s *PutAccessPointPolicyRequest) Validate() error {
	return dara.Validate(s)
}
