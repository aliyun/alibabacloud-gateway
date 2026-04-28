// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetObjectRetentionRequest struct {
}

func (s GetObjectRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectRetentionRequest) GoString() string {
	return s.String()
}

func (s *GetObjectRetentionRequest) Validate() error {
	return dara.Validate(s)
}
