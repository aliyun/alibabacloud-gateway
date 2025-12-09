// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketLoggingStatus interface {
	dara.Model
	String() string
	GoString() string
	SetLoggingEnabled(v *LoggingEnabled) *BucketLoggingStatus
	GetLoggingEnabled() *LoggingEnabled
}

type BucketLoggingStatus struct {
	// This parameter is required.
	LoggingEnabled *LoggingEnabled `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
}

func (s BucketLoggingStatus) String() string {
	return dara.Prettify(s)
}

func (s BucketLoggingStatus) GoString() string {
	return s.String()
}

func (s *BucketLoggingStatus) GetLoggingEnabled() *LoggingEnabled {
	return s.LoggingEnabled
}

func (s *BucketLoggingStatus) SetLoggingEnabled(v *LoggingEnabled) *BucketLoggingStatus {
	s.LoggingEnabled = v
	return s
}

func (s *BucketLoggingStatus) Validate() error {
	if s.LoggingEnabled != nil {
		if err := s.LoggingEnabled.Validate(); err != nil {
			return err
		}
	}
	return nil
}
