// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoggingEnabled interface {
	dara.Model
	String() string
	GoString() string
	SetLoggingRole(v string) *LoggingEnabled
	GetLoggingRole() *string
	SetTargetBucket(v string) *LoggingEnabled
	GetTargetBucket() *string
	SetTargetPrefix(v string) *LoggingEnabled
	GetTargetPrefix() *string
}

type LoggingEnabled struct {
	// example:
	//
	// AliyunOSSLoggingDefaultRole
	LoggingRole *string `json:"LoggingRole,omitempty" xml:"LoggingRole,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// examplebucket
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// example:
	//
	// MyLog-
	TargetPrefix *string `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
}

func (s LoggingEnabled) String() string {
	return dara.Prettify(s)
}

func (s LoggingEnabled) GoString() string {
	return s.String()
}

func (s *LoggingEnabled) GetLoggingRole() *string {
	return s.LoggingRole
}

func (s *LoggingEnabled) GetTargetBucket() *string {
	return s.TargetBucket
}

func (s *LoggingEnabled) GetTargetPrefix() *string {
	return s.TargetPrefix
}

func (s *LoggingEnabled) SetLoggingRole(v string) *LoggingEnabled {
	s.LoggingRole = &v
	return s
}

func (s *LoggingEnabled) SetTargetBucket(v string) *LoggingEnabled {
	s.TargetBucket = &v
	return s
}

func (s *LoggingEnabled) SetTargetPrefix(v string) *LoggingEnabled {
	s.TargetPrefix = &v
	return s
}

func (s *LoggingEnabled) Validate() error {
	return dara.Validate(s)
}
