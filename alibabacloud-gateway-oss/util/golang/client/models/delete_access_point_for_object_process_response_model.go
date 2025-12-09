// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessPointForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessPointForObjectProcessResponse
	GetStatusCode() *int32
}

type DeleteAccessPointForObjectProcessResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAccessPointForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessPointForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessPointForObjectProcessResponse) SetHeaders(v map[string]*string) *DeleteAccessPointForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPointForObjectProcessResponse) SetStatusCode(v int32) *DeleteAccessPointForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPointForObjectProcessResponse) Validate() error {
	return dara.Validate(s)
}
