// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourcePoolRequesterQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourcePoolRequesterQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourcePoolRequesterQoSInfoResponse
	GetStatusCode() *int32
}

type DeleteResourcePoolRequesterQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteResourcePoolRequesterQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourcePoolRequesterQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourcePoolRequesterQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourcePoolRequesterQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourcePoolRequesterQoSInfoResponse) SetHeaders(v map[string]*string) *DeleteResourcePoolRequesterQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourcePoolRequesterQoSInfoResponse) SetStatusCode(v int32) *DeleteResourcePoolRequesterQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourcePoolRequesterQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
