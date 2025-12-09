// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetObjectInfoResult(v *GetObjectInfoResult) *GetObjectInfoResponseBody
	GetGetObjectInfoResult() *GetObjectInfoResult
}

type GetObjectInfoResponseBody struct {
	GetObjectInfoResult *GetObjectInfoResult `json:"GetObjectInfoResult,omitempty" xml:"GetObjectInfoResult,omitempty"`
}

func (s GetObjectInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectInfoResponseBody) GetGetObjectInfoResult() *GetObjectInfoResult {
	return s.GetObjectInfoResult
}

func (s *GetObjectInfoResponseBody) SetGetObjectInfoResult(v *GetObjectInfoResult) *GetObjectInfoResponseBody {
	s.GetObjectInfoResult = v
	return s
}

func (s *GetObjectInfoResponseBody) Validate() error {
	if s.GetObjectInfoResult != nil {
		if err := s.GetObjectInfoResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
