// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveRecord interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *LiveRecord
	GetEndTime() *string
	SetRemoteAddr(v string) *LiveRecord
	GetRemoteAddr() *string
	SetStartTime(v string) *LiveRecord
	GetStartTime() *string
}

type LiveRecord struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RemoteAddr *string `json:"RemoteAddr,omitempty" xml:"RemoteAddr,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LiveRecord) String() string {
	return dara.Prettify(s)
}

func (s LiveRecord) GoString() string {
	return s.String()
}

func (s *LiveRecord) GetEndTime() *string {
	return s.EndTime
}

func (s *LiveRecord) GetRemoteAddr() *string {
	return s.RemoteAddr
}

func (s *LiveRecord) GetStartTime() *string {
	return s.StartTime
}

func (s *LiveRecord) SetEndTime(v string) *LiveRecord {
	s.EndTime = &v
	return s
}

func (s *LiveRecord) SetRemoteAddr(v string) *LiveRecord {
	s.RemoteAddr = &v
	return s
}

func (s *LiveRecord) SetStartTime(v string) *LiveRecord {
	s.StartTime = &v
	return s
}

func (s *LiveRecord) Validate() error {
	return dara.Validate(s)
}
