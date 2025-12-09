// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketInfo(v *BucketInfo) *GetBucketInfoResponseBody
	GetBucketInfo() *BucketInfo
}

type GetBucketInfoResponseBody struct {
	// The container that stores the information about the bucket.
	BucketInfo *BucketInfo `json:"BucketInfo,omitempty" xml:"BucketInfo,omitempty"`
}

func (s GetBucketInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBody) GetBucketInfo() *BucketInfo {
	return s.BucketInfo
}

func (s *GetBucketInfoResponseBody) SetBucketInfo(v *BucketInfo) *GetBucketInfoResponseBody {
	s.BucketInfo = v
	return s
}

func (s *GetBucketInfoResponseBody) Validate() error {
	if s.BucketInfo != nil {
		if err := s.BucketInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
