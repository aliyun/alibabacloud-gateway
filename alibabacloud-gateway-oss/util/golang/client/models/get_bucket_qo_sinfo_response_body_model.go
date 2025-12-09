// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketQoSInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQoSConfiguration(v *BucketQoSConfiguration) *GetBucketQoSInfoResponseBody
	GetQoSConfiguration() *BucketQoSConfiguration
}

type GetBucketQoSInfoResponseBody struct {
	QoSConfiguration *BucketQoSConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
}

func (s GetBucketQoSInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketQoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketQoSInfoResponseBody) GetQoSConfiguration() *BucketQoSConfiguration {
	return s.QoSConfiguration
}

func (s *GetBucketQoSInfoResponseBody) SetQoSConfiguration(v *BucketQoSConfiguration) *GetBucketQoSInfoResponseBody {
	s.QoSConfiguration = v
	return s
}

func (s *GetBucketQoSInfoResponseBody) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
