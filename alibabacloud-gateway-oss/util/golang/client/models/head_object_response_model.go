// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHeadObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HeadObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HeadObjectResponse
	GetStatusCode() *int32
}

type HeadObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s HeadObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s HeadObjectResponse) GoString() string {
	return s.String()
}

func (s *HeadObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HeadObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HeadObjectResponse) SetHeaders(v map[string]*string) *HeadObjectResponse {
	s.Headers = v
	return s
}

func (s *HeadObjectResponse) SetStatusCode(v int32) *HeadObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *HeadObjectResponse) Validate() error {
	return dara.Validate(s)
}
