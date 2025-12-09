// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAccessPointPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAccessPointPolicyResponse
	GetStatusCode() *int32
}

type PutAccessPointPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutAccessPointPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutAccessPointPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAccessPointPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAccessPointPolicyResponse) SetHeaders(v map[string]*string) *PutAccessPointPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutAccessPointPolicyResponse) SetStatusCode(v int32) *PutAccessPointPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAccessPointPolicyResponse) Validate() error {
	return dara.Validate(s)
}
