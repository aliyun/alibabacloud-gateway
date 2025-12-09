// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostVodPlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostVodPlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostVodPlaylistResponse
	GetStatusCode() *int32
}

type PostVodPlaylistResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PostVodPlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s PostVodPlaylistResponse) GoString() string {
	return s.String()
}

func (s *PostVodPlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostVodPlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostVodPlaylistResponse) SetHeaders(v map[string]*string) *PostVodPlaylistResponse {
	s.Headers = v
	return s
}

func (s *PostVodPlaylistResponse) SetStatusCode(v int32) *PostVodPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *PostVodPlaylistResponse) Validate() error {
	return dara.Validate(s)
}
