// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketInfo(v *BucketInfo) *GetVectorBucketResponseBody
	GetBucketInfo() *BucketInfo
}

type GetVectorBucketResponseBody struct {
	BucketInfo *BucketInfo `json:"BucketInfo,omitempty" xml:"BucketInfo,omitempty"`
}

func (s GetVectorBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVectorBucketResponseBody) GoString() string {
	return s.String()
}

func (s *GetVectorBucketResponseBody) GetBucketInfo() *BucketInfo {
	return s.BucketInfo
}

func (s *GetVectorBucketResponseBody) SetBucketInfo(v *BucketInfo) *GetVectorBucketResponseBody {
	s.BucketInfo = v
	return s
}

func (s *GetVectorBucketResponseBody) Validate() error {
	if s.BucketInfo != nil {
		if err := s.BucketInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
