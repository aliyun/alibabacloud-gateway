// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationFailureReasons interface {
	dara.Model
	String() string
	GoString() string
	SetFailureCode(v string) *BatchOperationFailureReasons
	GetFailureCode() *string
	SetFailureReason(v string) *BatchOperationFailureReasons
	GetFailureReason() *string
}

type BatchOperationFailureReasons struct {
	FailureCode   *string `json:"FailureCode,omitempty" xml:"FailureCode,omitempty"`
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
}

func (s BatchOperationFailureReasons) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationFailureReasons) GoString() string {
	return s.String()
}

func (s *BatchOperationFailureReasons) GetFailureCode() *string {
	return s.FailureCode
}

func (s *BatchOperationFailureReasons) GetFailureReason() *string {
	return s.FailureReason
}

func (s *BatchOperationFailureReasons) SetFailureCode(v string) *BatchOperationFailureReasons {
	s.FailureCode = &v
	return s
}

func (s *BatchOperationFailureReasons) SetFailureReason(v string) *BatchOperationFailureReasons {
	s.FailureReason = &v
	return s
}

func (s *BatchOperationFailureReasons) Validate() error {
	return dara.Validate(s)
}
