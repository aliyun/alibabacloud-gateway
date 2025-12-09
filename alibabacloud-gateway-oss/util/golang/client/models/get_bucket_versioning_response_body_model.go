// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketVersioningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVersioningConfiguration(v *GetBucketVersioningResponseBodyVersioningConfiguration) *GetBucketVersioningResponseBody
	GetVersioningConfiguration() *GetBucketVersioningResponseBodyVersioningConfiguration
}

type GetBucketVersioningResponseBody struct {
	// The container that stores the versioning state of the bucket.
	VersioningConfiguration *GetBucketVersioningResponseBodyVersioningConfiguration `json:"VersioningConfiguration,omitempty" xml:"VersioningConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketVersioningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketVersioningResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketVersioningResponseBody) GetVersioningConfiguration() *GetBucketVersioningResponseBodyVersioningConfiguration {
	return s.VersioningConfiguration
}

func (s *GetBucketVersioningResponseBody) SetVersioningConfiguration(v *GetBucketVersioningResponseBodyVersioningConfiguration) *GetBucketVersioningResponseBody {
	s.VersioningConfiguration = v
	return s
}

func (s *GetBucketVersioningResponseBody) Validate() error {
	if s.VersioningConfiguration != nil {
		if err := s.VersioningConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketVersioningResponseBodyVersioningConfiguration struct {
	// The versioning state of the bucket.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBucketVersioningResponseBodyVersioningConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketVersioningResponseBodyVersioningConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketVersioningResponseBodyVersioningConfiguration) GetStatus() *string {
	return s.Status
}

func (s *GetBucketVersioningResponseBodyVersioningConfiguration) SetStatus(v string) *GetBucketVersioningResponseBodyVersioningConfiguration {
	s.Status = &v
	return s
}

func (s *GetBucketVersioningResponseBodyVersioningConfiguration) Validate() error {
	return dara.Validate(s)
}
