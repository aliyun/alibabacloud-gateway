// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteDataLakeCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPromoteDataLakeCacheRequest(v *PromoteDataLakeCacheReq) *PromoteDataLakeCacheRequest
	GetPromoteDataLakeCacheRequest() *PromoteDataLakeCacheReq
}

type PromoteDataLakeCacheRequest struct {
	PromoteDataLakeCacheRequest *PromoteDataLakeCacheReq `json:"PromoteDataLakeCacheRequest,omitempty" xml:"PromoteDataLakeCacheRequest,omitempty"`
}

func (s PromoteDataLakeCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s PromoteDataLakeCacheRequest) GoString() string {
	return s.String()
}

func (s *PromoteDataLakeCacheRequest) GetPromoteDataLakeCacheRequest() *PromoteDataLakeCacheReq {
	return s.PromoteDataLakeCacheRequest
}

func (s *PromoteDataLakeCacheRequest) SetPromoteDataLakeCacheRequest(v *PromoteDataLakeCacheReq) *PromoteDataLakeCacheRequest {
	s.PromoteDataLakeCacheRequest = v
	return s
}

func (s *PromoteDataLakeCacheRequest) Validate() error {
	if s.PromoteDataLakeCacheRequest != nil {
		if err := s.PromoteDataLakeCacheRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
