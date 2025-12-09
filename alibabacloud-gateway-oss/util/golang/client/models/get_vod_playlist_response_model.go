// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGetVodPlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVodPlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVodPlaylistResponse
	GetStatusCode() *int32
	SetBody(v io.Reader) *GetVodPlaylistResponse
	GetBody() io.Reader
}

type GetVodPlaylistResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVodPlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVodPlaylistResponse) GoString() string {
	return s.String()
}

func (s *GetVodPlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVodPlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVodPlaylistResponse) GetBody() io.Reader {
	return s.Body
}

func (s *GetVodPlaylistResponse) SetHeaders(v map[string]*string) *GetVodPlaylistResponse {
	s.Headers = v
	return s
}

func (s *GetVodPlaylistResponse) SetStatusCode(v int32) *GetVodPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodPlaylistResponse) SetBody(v io.Reader) *GetVodPlaylistResponse {
	s.Body = v
	return s
}

func (s *GetVodPlaylistResponse) Validate() error {
	return dara.Validate(s)
}
