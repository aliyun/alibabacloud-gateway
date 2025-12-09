// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelPublishUrls interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *LiveChannelPublishUrls
	GetUrl() *string
}

type LiveChannelPublishUrls struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s LiveChannelPublishUrls) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelPublishUrls) GoString() string {
	return s.String()
}

func (s *LiveChannelPublishUrls) GetUrl() *string {
	return s.Url
}

func (s *LiveChannelPublishUrls) SetUrl(v string) *LiveChannelPublishUrls {
	s.Url = &v
	return s
}

func (s *LiveChannelPublishUrls) Validate() error {
	return dara.Validate(s)
}
