// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetObjectsRequest(v *GetObjectsReq) *GetObjectsRequest
	GetGetObjectsRequest() *GetObjectsReq
}

type GetObjectsRequest struct {
	GetObjectsRequest *GetObjectsReq `json:"GetObjectsRequest,omitempty" xml:"GetObjectsRequest,omitempty"`
}

func (s GetObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetObjectsRequest) GetGetObjectsRequest() *GetObjectsReq {
	return s.GetObjectsRequest
}

func (s *GetObjectsRequest) SetGetObjectsRequest(v *GetObjectsReq) *GetObjectsRequest {
	s.GetObjectsRequest = v
	return s
}

func (s *GetObjectsRequest) Validate() error {
	if s.GetObjectsRequest != nil {
		if err := s.GetObjectsRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
