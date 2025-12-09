// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointPolicyForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessPointPolicyForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessPointPolicyForObjectProcessResponse
	GetStatusCode() *int32
}

type DeleteAccessPointPolicyForObjectProcessResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAccessPointPolicyForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointPolicyForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointPolicyForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessPointPolicyForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessPointPolicyForObjectProcessResponse) SetHeaders(v map[string]*string) *DeleteAccessPointPolicyForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPointPolicyForObjectProcessResponse) SetStatusCode(v int32) *DeleteAccessPointPolicyForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPointPolicyForObjectProcessResponse) Validate() error {
	return dara.Validate(s)
}
