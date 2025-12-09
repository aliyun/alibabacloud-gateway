// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCORSConfiguration(v *GetBucketCorsResponseBodyCORSConfiguration) *GetBucketCorsResponseBody
	GetCORSConfiguration() *GetBucketCorsResponseBodyCORSConfiguration
}

type GetBucketCorsResponseBody struct {
	// The container that stores CORS configuration.
	CORSConfiguration *GetBucketCorsResponseBodyCORSConfiguration `json:"CORSConfiguration,omitempty" xml:"CORSConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketCorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketCorsResponseBody) GetCORSConfiguration() *GetBucketCorsResponseBodyCORSConfiguration {
	return s.CORSConfiguration
}

func (s *GetBucketCorsResponseBody) SetCORSConfiguration(v *GetBucketCorsResponseBodyCORSConfiguration) *GetBucketCorsResponseBody {
	s.CORSConfiguration = v
	return s
}

func (s *GetBucketCorsResponseBody) Validate() error {
	if s.CORSConfiguration != nil {
		if err := s.CORSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketCorsResponseBodyCORSConfiguration struct {
	// The container that stores CORS rules. Up to 10 rules can be configured for a bucket.
	CORSRule []*CORSRule `json:"CORSRule,omitempty" xml:"CORSRule,omitempty" type:"Repeated"`
	// Indicates whether the Vary: Origin header was returned. Default value: false.
	//
	// - true: The Vary: Origin header is returned regardless whether the request is a cross-origin request or whether the cross-origin request succeeds.
	//
	// - false: The Vary: Origin header is not returned.
	//
	// example:
	//
	// true
	ResponseVary *bool `json:"ResponseVary,omitempty" xml:"ResponseVary,omitempty"`
}

func (s GetBucketCorsResponseBodyCORSConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCorsResponseBodyCORSConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketCorsResponseBodyCORSConfiguration) GetCORSRule() []*CORSRule {
	return s.CORSRule
}

func (s *GetBucketCorsResponseBodyCORSConfiguration) GetResponseVary() *bool {
	return s.ResponseVary
}

func (s *GetBucketCorsResponseBodyCORSConfiguration) SetCORSRule(v []*CORSRule) *GetBucketCorsResponseBodyCORSConfiguration {
	s.CORSRule = v
	return s
}

func (s *GetBucketCorsResponseBodyCORSConfiguration) SetResponseVary(v bool) *GetBucketCorsResponseBodyCORSConfiguration {
	s.ResponseVary = &v
	return s
}

func (s *GetBucketCorsResponseBodyCORSConfiguration) Validate() error {
	if s.CORSRule != nil {
		for _, item := range s.CORSRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
