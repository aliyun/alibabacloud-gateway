// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectAclResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectAclResponseBody) *GetObjectAclResponse
	GetBody() *GetObjectAclResponseBody
}

type GetObjectAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectAclResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectAclResponse) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectAclResponse) GetBody() *GetObjectAclResponseBody {
	return s.Body
}

func (s *GetObjectAclResponse) SetHeaders(v map[string]*string) *GetObjectAclResponse {
	s.Headers = v
	return s
}

func (s *GetObjectAclResponse) SetStatusCode(v int32) *GetObjectAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectAclResponse) SetBody(v *GetObjectAclResponseBody) *GetObjectAclResponse {
	s.Body = v
	return s
}

func (s *GetObjectAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
