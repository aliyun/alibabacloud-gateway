// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanRestoredObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CleanRestoredObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CleanRestoredObjectResponse
	GetStatusCode() *int32
}

type CleanRestoredObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CleanRestoredObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CleanRestoredObjectResponse) GoString() string {
	return s.String()
}

func (s *CleanRestoredObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CleanRestoredObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CleanRestoredObjectResponse) SetHeaders(v map[string]*string) *CleanRestoredObjectResponse {
	s.Headers = v
	return s
}

func (s *CleanRestoredObjectResponse) SetStatusCode(v int32) *CleanRestoredObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CleanRestoredObjectResponse) Validate() error {
	return dara.Validate(s)
}
