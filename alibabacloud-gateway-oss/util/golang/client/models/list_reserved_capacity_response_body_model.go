// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReservedCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReservedCapacityRecordList(v *ReservedCapacityRecordList) *ListReservedCapacityResponseBody
	GetReservedCapacityRecordList() *ReservedCapacityRecordList
}

type ListReservedCapacityResponseBody struct {
	ReservedCapacityRecordList *ReservedCapacityRecordList `json:"ReservedCapacityRecordList,omitempty" xml:"ReservedCapacityRecordList,omitempty"`
}

func (s ListReservedCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReservedCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *ListReservedCapacityResponseBody) GetReservedCapacityRecordList() *ReservedCapacityRecordList {
	return s.ReservedCapacityRecordList
}

func (s *ListReservedCapacityResponseBody) SetReservedCapacityRecordList(v *ReservedCapacityRecordList) *ListReservedCapacityResponseBody {
	s.ReservedCapacityRecordList = v
	return s
}

func (s *ListReservedCapacityResponseBody) Validate() error {
	if s.ReservedCapacityRecordList != nil {
		if err := s.ReservedCapacityRecordList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
