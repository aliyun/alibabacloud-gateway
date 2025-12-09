// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketHashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetObjectHashConfiguration(v *ObjectHashConfiguration) *GetBucketHashResponseBody
	GetObjectHashConfiguration() *ObjectHashConfiguration
}

type GetBucketHashResponseBody struct {
	ObjectHashConfiguration *ObjectHashConfiguration `json:"ObjectHashConfiguration,omitempty" xml:"ObjectHashConfiguration,omitempty"`
}

func (s GetBucketHashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketHashResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketHashResponseBody) GetObjectHashConfiguration() *ObjectHashConfiguration {
	return s.ObjectHashConfiguration
}

func (s *GetBucketHashResponseBody) SetObjectHashConfiguration(v *ObjectHashConfiguration) *GetBucketHashResponseBody {
	s.ObjectHashConfiguration = v
	return s
}

func (s *GetBucketHashResponseBody) Validate() error {
	if s.ObjectHashConfiguration != nil {
		if err := s.ObjectHashConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
