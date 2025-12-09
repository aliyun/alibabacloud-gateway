// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourcePoolRequesterQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutResourcePoolRequesterQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutResourcePoolRequesterQoSInfoResponse
	GetStatusCode() *int32
}

type PutResourcePoolRequesterQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutResourcePoolRequesterQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PutResourcePoolRequesterQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *PutResourcePoolRequesterQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutResourcePoolRequesterQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutResourcePoolRequesterQoSInfoResponse) SetHeaders(v map[string]*string) *PutResourcePoolRequesterQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *PutResourcePoolRequesterQoSInfoResponse) SetStatusCode(v int32) *PutResourcePoolRequesterQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PutResourcePoolRequesterQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
