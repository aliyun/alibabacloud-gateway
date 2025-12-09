// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLoggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketLoggingStatus(v *BucketLoggingStatus) *GetBucketLoggingResponseBody
	GetBucketLoggingStatus() *BucketLoggingStatus
}

type GetBucketLoggingResponseBody struct {
	// Indicates the container used to store access logging configuration of a bucket.
	BucketLoggingStatus *BucketLoggingStatus `json:"BucketLoggingStatus,omitempty" xml:"BucketLoggingStatus,omitempty"`
}

func (s GetBucketLoggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLoggingResponseBody) GetBucketLoggingStatus() *BucketLoggingStatus {
	return s.BucketLoggingStatus
}

func (s *GetBucketLoggingResponseBody) SetBucketLoggingStatus(v *BucketLoggingStatus) *GetBucketLoggingResponseBody {
	s.BucketLoggingStatus = v
	return s
}

func (s *GetBucketLoggingResponseBody) Validate() error {
	if s.BucketLoggingStatus != nil {
		if err := s.BucketLoggingStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}
