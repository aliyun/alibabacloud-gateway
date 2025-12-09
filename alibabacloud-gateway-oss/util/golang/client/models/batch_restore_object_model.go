// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRestoreObject interface {
	dara.Model
	String() string
	GoString() string
	SetDays(v int32) *BatchRestoreObject
	GetDays() *int32
	SetTier(v string) *BatchRestoreObject
	GetTier() *string
}

type BatchRestoreObject struct {
	// example:
	//
	// 5
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// example:
	//
	// Standard
	Tier *string `json:"Tier,omitempty" xml:"Tier,omitempty"`
}

func (s BatchRestoreObject) String() string {
	return dara.Prettify(s)
}

func (s BatchRestoreObject) GoString() string {
	return s.String()
}

func (s *BatchRestoreObject) GetDays() *int32 {
	return s.Days
}

func (s *BatchRestoreObject) GetTier() *string {
	return s.Tier
}

func (s *BatchRestoreObject) SetDays(v int32) *BatchRestoreObject {
	s.Days = &v
	return s
}

func (s *BatchRestoreObject) SetTier(v string) *BatchRestoreObject {
	s.Tier = &v
	return s
}

func (s *BatchRestoreObject) Validate() error {
	return dara.Validate(s)
}
