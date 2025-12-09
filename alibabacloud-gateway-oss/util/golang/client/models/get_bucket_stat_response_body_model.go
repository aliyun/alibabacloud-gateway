// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketStat(v *BucketStat) *GetBucketStatResponseBody
	GetBucketStat() *BucketStat
}

type GetBucketStatResponseBody struct {
	// The container that stores all information returned for the GetBucketStat request.
	BucketStat *BucketStat `json:"BucketStat,omitempty" xml:"BucketStat,omitempty"`
}

func (s GetBucketStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketStatResponseBody) GetBucketStat() *BucketStat {
	return s.BucketStat
}

func (s *GetBucketStatResponseBody) SetBucketStat(v *BucketStat) *GetBucketStatResponseBody {
	s.BucketStat = v
	return s
}

func (s *GetBucketStatResponseBody) Validate() error {
	if s.BucketStat != nil {
		if err := s.BucketStat.Validate(); err != nil {
			return err
		}
	}
	return nil
}
