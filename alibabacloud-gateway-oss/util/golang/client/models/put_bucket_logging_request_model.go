// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLoggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketLoggingStatus(v *BucketLoggingStatus) *PutBucketLoggingRequest
	GetBucketLoggingStatus() *BucketLoggingStatus
}

type PutBucketLoggingRequest struct {
	// The container that stores the logging status information.
	BucketLoggingStatus *BucketLoggingStatus `json:"BucketLoggingStatus,omitempty" xml:"BucketLoggingStatus,omitempty"`
}

func (s PutBucketLoggingRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLoggingRequest) GoString() string {
	return s.String()
}

func (s *PutBucketLoggingRequest) GetBucketLoggingStatus() *BucketLoggingStatus {
	return s.BucketLoggingStatus
}

func (s *PutBucketLoggingRequest) SetBucketLoggingStatus(v *BucketLoggingStatus) *PutBucketLoggingRequest {
	s.BucketLoggingStatus = v
	return s
}

func (s *PutBucketLoggingRequest) Validate() error {
	if s.BucketLoggingStatus != nil {
		if err := s.BucketLoggingStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}
