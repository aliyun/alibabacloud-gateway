// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSealAppendObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SealAppendObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SealAppendObjectResponse
	GetStatusCode() *int32
}

type SealAppendObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SealAppendObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s SealAppendObjectResponse) GoString() string {
	return s.String()
}

func (s *SealAppendObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SealAppendObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SealAppendObjectResponse) SetHeaders(v map[string]*string) *SealAppendObjectResponse {
	s.Headers = v
	return s
}

func (s *SealAppendObjectResponse) SetStatusCode(v int32) *SealAppendObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SealAppendObjectResponse) Validate() error {
	return dara.Validate(s)
}
