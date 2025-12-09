// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelPlayUrls interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *LiveChannelPlayUrls
	GetUrl() *string
}

type LiveChannelPlayUrls struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s LiveChannelPlayUrls) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelPlayUrls) GoString() string {
	return s.String()
}

func (s *LiveChannelPlayUrls) GetUrl() *string {
	return s.Url
}

func (s *LiveChannelPlayUrls) SetUrl(v string) *LiveChannelPlayUrls {
	s.Url = &v
	return s
}

func (s *LiveChannelPlayUrls) Validate() error {
	return dara.Validate(s)
}
