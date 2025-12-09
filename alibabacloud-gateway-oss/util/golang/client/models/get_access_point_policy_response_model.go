// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessPointPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessPointPolicyResponse
	GetStatusCode() *int32
	SetBody(v string) *GetAccessPointPolicyResponse
	GetBody() *string
}

type GetAccessPointPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPointPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPointPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessPointPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessPointPolicyResponse) GetBody() *string {
	return s.Body
}

func (s *GetAccessPointPolicyResponse) SetHeaders(v map[string]*string) *GetAccessPointPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPointPolicyResponse) SetStatusCode(v int32) *GetAccessPointPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPointPolicyResponse) SetBody(v string) *GetAccessPointPolicyResponse {
	s.Body = &v
	return s
}

func (s *GetAccessPointPolicyResponse) Validate() error {
	return dara.Validate(s)
}
