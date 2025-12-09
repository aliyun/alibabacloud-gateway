// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRequesterQoSInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequesterQoSInfo(v *RequesterQoSInfo) *GetBucketRequesterQoSInfoResponseBody
	GetRequesterQoSInfo() *RequesterQoSInfo
}

type GetBucketRequesterQoSInfoResponseBody struct {
	RequesterQoSInfo *RequesterQoSInfo `json:"RequesterQoSInfo,omitempty" xml:"RequesterQoSInfo,omitempty"`
}

func (s GetBucketRequesterQoSInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRequesterQoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketRequesterQoSInfoResponseBody) GetRequesterQoSInfo() *RequesterQoSInfo {
	return s.RequesterQoSInfo
}

func (s *GetBucketRequesterQoSInfoResponseBody) SetRequesterQoSInfo(v *RequesterQoSInfo) *GetBucketRequesterQoSInfoResponseBody {
	s.RequesterQoSInfo = v
	return s
}

func (s *GetBucketRequesterQoSInfoResponseBody) Validate() error {
	if s.RequesterQoSInfo != nil {
		if err := s.RequesterQoSInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
