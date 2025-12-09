// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteObjectTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteObjectTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteObjectTaggingResponse
	GetStatusCode() *int32
}

type DeleteObjectTaggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteObjectTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteObjectTaggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteObjectTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteObjectTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteObjectTaggingResponse) SetHeaders(v map[string]*string) *DeleteObjectTaggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteObjectTaggingResponse) SetStatusCode(v int32) *DeleteObjectTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteObjectTaggingResponse) Validate() error {
	return dara.Validate(s)
}
