// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketReplicationLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketReplicationLocationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketReplicationLocationResponseBody) *GetBucketReplicationLocationResponse
	GetBody() *GetBucketReplicationLocationResponseBody
}

type GetBucketReplicationLocationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketReplicationLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketReplicationLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationLocationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketReplicationLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketReplicationLocationResponse) GetBody() *GetBucketReplicationLocationResponseBody {
	return s.Body
}

func (s *GetBucketReplicationLocationResponse) SetHeaders(v map[string]*string) *GetBucketReplicationLocationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketReplicationLocationResponse) SetStatusCode(v int32) *GetBucketReplicationLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketReplicationLocationResponse) SetBody(v *GetBucketReplicationLocationResponseBody) *GetBucketReplicationLocationResponse {
	s.Body = v
	return s
}

func (s *GetBucketReplicationLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
