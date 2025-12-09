// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReservedCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReservedCapacityRecord(v *ReservedCapacityRecord) *GetReservedCapacityResponseBody
	GetReservedCapacityRecord() *ReservedCapacityRecord
}

type GetReservedCapacityResponseBody struct {
	ReservedCapacityRecord *ReservedCapacityRecord `json:"ReservedCapacityRecord,omitempty" xml:"ReservedCapacityRecord,omitempty"`
}

func (s GetReservedCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReservedCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *GetReservedCapacityResponseBody) GetReservedCapacityRecord() *ReservedCapacityRecord {
	return s.ReservedCapacityRecord
}

func (s *GetReservedCapacityResponseBody) SetReservedCapacityRecord(v *ReservedCapacityRecord) *GetReservedCapacityResponseBody {
	s.ReservedCapacityRecord = v
	return s
}

func (s *GetReservedCapacityResponseBody) Validate() error {
	if s.ReservedCapacityRecord != nil {
		if err := s.ReservedCapacityRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}
