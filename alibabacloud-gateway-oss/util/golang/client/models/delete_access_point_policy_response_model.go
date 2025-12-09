// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessPointPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessPointPolicyResponse
	GetStatusCode() *int32
}

type DeleteAccessPointPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAccessPointPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessPointPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessPointPolicyResponse) SetHeaders(v map[string]*string) *DeleteAccessPointPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPointPolicyResponse) SetStatusCode(v int32) *DeleteAccessPointPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPointPolicyResponse) Validate() error {
	return dara.Validate(s)
}
