// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectLegalHoldRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetObjectLegalHoldRequest struct {
}

func (s GetObjectLegalHoldRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectLegalHoldRequest) GoString() string {
	return s.String()
}

func (s *GetObjectLegalHoldRequest) Validate() error {
	return dara.Validate(s)
}
