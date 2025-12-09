// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaQueryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaQueryStatus(v *GetMetaQueryStatusResponseBodyMetaQueryStatus) *GetMetaQueryStatusResponseBody
	GetMetaQueryStatus() *GetMetaQueryStatusResponseBodyMetaQueryStatus
}

type GetMetaQueryStatusResponseBody struct {
	// The container that stores the metadata information.
	MetaQueryStatus *GetMetaQueryStatusResponseBodyMetaQueryStatus `json:"MetaQueryStatus,omitempty" xml:"MetaQueryStatus,omitempty" type:"Struct"`
}

func (s GetMetaQueryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaQueryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaQueryStatusResponseBody) GetMetaQueryStatus() *GetMetaQueryStatusResponseBodyMetaQueryStatus {
	return s.MetaQueryStatus
}

func (s *GetMetaQueryStatusResponseBody) SetMetaQueryStatus(v *GetMetaQueryStatusResponseBodyMetaQueryStatus) *GetMetaQueryStatusResponseBody {
	s.MetaQueryStatus = v
	return s
}

func (s *GetMetaQueryStatusResponseBody) Validate() error {
	if s.MetaQueryStatus != nil {
		if err := s.MetaQueryStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetaQueryStatusResponseBodyMetaQueryStatus struct {
	// The time when the metadata index library was created. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
	//
	// example:
	//
	// 2021-08-02T10:49:17.289372919+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The scan type. Valid values:
	//
	// - FullScanning: Full scanning is in progress.
	//
	// - IncrementalScanning: Incremental scanning is in progress.
	//
	// example:
	//
	// FullScanning
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The status of the metadata index library. Valid values:
	//
	// - Ready: The metadata index library is being prepared after it is created.
	//
	// In this case, the metadata index library cannot be used to query data.
	//
	// - Stop: The metadata index library is paused.
	//
	// - Running: The metadata index library is running.
	//
	// - Retrying: The metadata index library failed to be created and is being created again.
	//
	// - Failed: The metadata index library failed to be created.
	//
	// - Deleted: The metadata index library is deleted.
	//
	// example:
	//
	// Runing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the metadata index library was updated. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
	//
	// example:
	//
	// 2021-08-02T10:49:17.289372919+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetMetaQueryStatusResponseBodyMetaQueryStatus) String() string {
	return dara.Prettify(s)
}

func (s GetMetaQueryStatusResponseBodyMetaQueryStatus) GoString() string {
	return s.String()
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) GetPhase() *string {
	return s.Phase
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) GetState() *string {
	return s.State
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) SetCreateTime(v string) *GetMetaQueryStatusResponseBodyMetaQueryStatus {
	s.CreateTime = &v
	return s
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) SetPhase(v string) *GetMetaQueryStatusResponseBodyMetaQueryStatus {
	s.Phase = &v
	return s
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) SetState(v string) *GetMetaQueryStatusResponseBodyMetaQueryStatus {
	s.State = &v
	return s
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) SetUpdateTime(v string) *GetMetaQueryStatusResponseBodyMetaQueryStatus {
	s.UpdateTime = &v
	return s
}

func (s *GetMetaQueryStatusResponseBodyMetaQueryStatus) Validate() error {
	return dara.Validate(s)
}
