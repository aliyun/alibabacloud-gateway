// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolRequesterQoSInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequesterQoSInfo(v *RequesterQoSInfo) *GetResourcePoolRequesterQoSInfoResponseBody
	GetRequesterQoSInfo() *RequesterQoSInfo
}

type GetResourcePoolRequesterQoSInfoResponseBody struct {
	RequesterQoSInfo *RequesterQoSInfo `json:"RequesterQoSInfo,omitempty" xml:"RequesterQoSInfo,omitempty"`
}

func (s GetResourcePoolRequesterQoSInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolRequesterQoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePoolRequesterQoSInfoResponseBody) GetRequesterQoSInfo() *RequesterQoSInfo {
	return s.RequesterQoSInfo
}

func (s *GetResourcePoolRequesterQoSInfoResponseBody) SetRequesterQoSInfo(v *RequesterQoSInfo) *GetResourcePoolRequesterQoSInfoResponseBody {
	s.RequesterQoSInfo = v
	return s
}

func (s *GetResourcePoolRequesterQoSInfoResponseBody) Validate() error {
	if s.RequesterQoSInfo != nil {
		if err := s.RequesterQoSInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
