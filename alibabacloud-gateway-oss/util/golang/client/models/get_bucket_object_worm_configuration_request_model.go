// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketObjectWormConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetBucketObjectWormConfigurationRequest struct {
}

func (s GetBucketObjectWormConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketObjectWormConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetBucketObjectWormConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
