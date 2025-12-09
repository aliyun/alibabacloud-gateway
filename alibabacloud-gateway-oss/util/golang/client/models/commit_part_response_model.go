// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitPartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommitPartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommitPartResponse
	GetStatusCode() *int32
}

type CommitPartResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CommitPartResponse) String() string {
	return dara.Prettify(s)
}

func (s CommitPartResponse) GoString() string {
	return s.String()
}

func (s *CommitPartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommitPartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommitPartResponse) SetHeaders(v map[string]*string) *CommitPartResponse {
	s.Headers = v
	return s
}

func (s *CommitPartResponse) SetStatusCode(v int32) *CommitPartResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitPartResponse) Validate() error {
	return dara.Validate(s)
}
