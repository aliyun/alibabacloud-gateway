// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketResponseHeaderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderConfiguration(v *ResponseHeaderConfiguration) *GetBucketResponseHeaderResponseBody
	GetResponseHeaderConfiguration() *ResponseHeaderConfiguration
}

type GetBucketResponseHeaderResponseBody struct {
	ResponseHeaderConfiguration *ResponseHeaderConfiguration `json:"ResponseHeaderConfiguration,omitempty" xml:"ResponseHeaderConfiguration,omitempty"`
}

func (s GetBucketResponseHeaderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketResponseHeaderResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketResponseHeaderResponseBody) GetResponseHeaderConfiguration() *ResponseHeaderConfiguration {
	return s.ResponseHeaderConfiguration
}

func (s *GetBucketResponseHeaderResponseBody) SetResponseHeaderConfiguration(v *ResponseHeaderConfiguration) *GetBucketResponseHeaderResponseBody {
	s.ResponseHeaderConfiguration = v
	return s
}

func (s *GetBucketResponseHeaderResponseBody) Validate() error {
	if s.ResponseHeaderConfiguration != nil {
		if err := s.ResponseHeaderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
