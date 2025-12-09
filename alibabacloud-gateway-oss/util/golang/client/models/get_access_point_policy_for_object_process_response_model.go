// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPolicyForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessPointPolicyForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessPointPolicyForObjectProcessResponse
	GetStatusCode() *int32
	SetBody(v string) *GetAccessPointPolicyForObjectProcessResponse
	GetBody() *string
}

type GetAccessPointPolicyForObjectProcessResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPointPolicyForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPolicyForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPointPolicyForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessPointPolicyForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessPointPolicyForObjectProcessResponse) GetBody() *string {
	return s.Body
}

func (s *GetAccessPointPolicyForObjectProcessResponse) SetHeaders(v map[string]*string) *GetAccessPointPolicyForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPointPolicyForObjectProcessResponse) SetStatusCode(v int32) *GetAccessPointPolicyForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPointPolicyForObjectProcessResponse) SetBody(v string) *GetAccessPointPolicyForObjectProcessResponse {
	s.Body = &v
	return s
}

func (s *GetAccessPointPolicyForObjectProcessResponse) Validate() error {
	return dara.Validate(s)
}
