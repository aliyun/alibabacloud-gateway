// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketWebsiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetWebsiteConfiguration(v *WebsiteConfiguration) *GetBucketWebsiteResponseBody
	GetWebsiteConfiguration() *WebsiteConfiguration
}

type GetBucketWebsiteResponseBody struct {
	// The containers of the website configuration.
	WebsiteConfiguration *WebsiteConfiguration `json:"WebsiteConfiguration,omitempty" xml:"WebsiteConfiguration,omitempty"`
}

func (s GetBucketWebsiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketWebsiteResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketWebsiteResponseBody) GetWebsiteConfiguration() *WebsiteConfiguration {
	return s.WebsiteConfiguration
}

func (s *GetBucketWebsiteResponseBody) SetWebsiteConfiguration(v *WebsiteConfiguration) *GetBucketWebsiteResponseBody {
	s.WebsiteConfiguration = v
	return s
}

func (s *GetBucketWebsiteResponseBody) Validate() error {
	if s.WebsiteConfiguration != nil {
		if err := s.WebsiteConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
