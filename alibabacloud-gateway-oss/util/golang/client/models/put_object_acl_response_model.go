// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutObjectAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutObjectAclResponse
	GetStatusCode() *int32
}

type PutObjectAclResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutObjectAclResponse) String() string {
	return dara.Prettify(s)
}

func (s PutObjectAclResponse) GoString() string {
	return s.String()
}

func (s *PutObjectAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutObjectAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutObjectAclResponse) SetHeaders(v map[string]*string) *PutObjectAclResponse {
	s.Headers = v
	return s
}

func (s *PutObjectAclResponse) SetStatusCode(v int32) *PutObjectAclResponse {
	s.StatusCode = &v
	return s
}

func (s *PutObjectAclResponse) Validate() error {
	return dara.Validate(s)
}
