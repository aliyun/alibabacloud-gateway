// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostAsyncFetchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncFetchTaskConfiguration(v *AsyncFetchTaskConfiguration) *PostAsyncFetchTaskRequest
	GetAsyncFetchTaskConfiguration() *AsyncFetchTaskConfiguration
}

type PostAsyncFetchTaskRequest struct {
	AsyncFetchTaskConfiguration *AsyncFetchTaskConfiguration `json:"AsyncFetchTaskConfiguration,omitempty" xml:"AsyncFetchTaskConfiguration,omitempty"`
}

func (s PostAsyncFetchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PostAsyncFetchTaskRequest) GoString() string {
	return s.String()
}

func (s *PostAsyncFetchTaskRequest) GetAsyncFetchTaskConfiguration() *AsyncFetchTaskConfiguration {
	return s.AsyncFetchTaskConfiguration
}

func (s *PostAsyncFetchTaskRequest) SetAsyncFetchTaskConfiguration(v *AsyncFetchTaskConfiguration) *PostAsyncFetchTaskRequest {
	s.AsyncFetchTaskConfiguration = v
	return s
}

func (s *PostAsyncFetchTaskRequest) Validate() error {
	if s.AsyncFetchTaskConfiguration != nil {
		if err := s.AsyncFetchTaskConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
