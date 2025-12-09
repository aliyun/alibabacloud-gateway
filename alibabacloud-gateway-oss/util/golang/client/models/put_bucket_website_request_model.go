// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketWebsiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebsiteConfiguration(v *WebsiteConfiguration) *PutBucketWebsiteRequest
	GetWebsiteConfiguration() *WebsiteConfiguration
}

type PutBucketWebsiteRequest struct {
	// The container that stores the website configuration.
	WebsiteConfiguration *WebsiteConfiguration `json:"WebsiteConfiguration,omitempty" xml:"WebsiteConfiguration,omitempty"`
}

func (s PutBucketWebsiteRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketWebsiteRequest) GoString() string {
	return s.String()
}

func (s *PutBucketWebsiteRequest) GetWebsiteConfiguration() *WebsiteConfiguration {
	return s.WebsiteConfiguration
}

func (s *PutBucketWebsiteRequest) SetWebsiteConfiguration(v *WebsiteConfiguration) *PutBucketWebsiteRequest {
	s.WebsiteConfiguration = v
	return s
}

func (s *PutBucketWebsiteRequest) Validate() error {
	if s.WebsiteConfiguration != nil {
		if err := s.WebsiteConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
