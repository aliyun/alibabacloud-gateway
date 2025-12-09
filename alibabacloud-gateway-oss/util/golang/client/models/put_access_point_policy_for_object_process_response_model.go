// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPolicyForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAccessPointPolicyForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAccessPointPolicyForObjectProcessResponse
	GetStatusCode() *int32
}

type PutAccessPointPolicyForObjectProcessResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutAccessPointPolicyForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPolicyForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *PutAccessPointPolicyForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAccessPointPolicyForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAccessPointPolicyForObjectProcessResponse) SetHeaders(v map[string]*string) *PutAccessPointPolicyForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *PutAccessPointPolicyForObjectProcessResponse) SetStatusCode(v int32) *PutAccessPointPolicyForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAccessPointPolicyForObjectProcessResponse) Validate() error {
	return dara.Validate(s)
}
