// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetResourcePoolInfoResponse(v *GetResourcePoolInfoResp) *GetResourcePoolInfoResponseBody
	GetGetResourcePoolInfoResponse() *GetResourcePoolInfoResp
}

type GetResourcePoolInfoResponseBody struct {
	GetResourcePoolInfoResponse *GetResourcePoolInfoResp `json:"GetResourcePoolInfoResponse,omitempty" xml:"GetResourcePoolInfoResponse,omitempty"`
}

func (s GetResourcePoolInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePoolInfoResponseBody) GetGetResourcePoolInfoResponse() *GetResourcePoolInfoResp {
	return s.GetResourcePoolInfoResponse
}

func (s *GetResourcePoolInfoResponseBody) SetGetResourcePoolInfoResponse(v *GetResourcePoolInfoResp) *GetResourcePoolInfoResponseBody {
	s.GetResourcePoolInfoResponse = v
	return s
}

func (s *GetResourcePoolInfoResponseBody) Validate() error {
	if s.GetResourcePoolInfoResponse != nil {
		if err := s.GetResourcePoolInfoResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}
