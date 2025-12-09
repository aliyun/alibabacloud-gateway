// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVectorsResponse
	GetStatusCode() *int32
}

type DeleteVectorsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteVectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteVectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVectorsResponse) SetHeaders(v map[string]*string) *DeleteVectorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteVectorsResponse) SetStatusCode(v int32) *DeleteVectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVectorsResponse) Validate() error {
	return dara.Validate(s)
}
