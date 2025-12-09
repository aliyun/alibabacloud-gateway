// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLiveChannelStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *PutLiveChannelStatusRequest
	GetStatus() *string
}

type PutLiveChannelStatusRequest struct {
	// The status of the LiveChannel.
	//
	// Valid values:
	//
	// - enabled: enables the LiveChannel.
	//
	// - disabled: disables the LiveChannel.
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutLiveChannelStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s PutLiveChannelStatusRequest) GoString() string {
	return s.String()
}

func (s *PutLiveChannelStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *PutLiveChannelStatusRequest) SetStatus(v string) *PutLiveChannelStatusRequest {
	s.Status = &v
	return s
}

func (s *PutLiveChannelStatusRequest) Validate() error {
	return dara.Validate(s)
}
