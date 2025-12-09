// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketPublicAccessBlockResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketPublicAccessBlockResponseBody) *GetBucketPublicAccessBlockResponse
	GetBody() *GetBucketPublicAccessBlockResponseBody
}

type GetBucketPublicAccessBlockResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketPublicAccessBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *GetBucketPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketPublicAccessBlockResponse) GetBody() *GetBucketPublicAccessBlockResponseBody {
	return s.Body
}

func (s *GetBucketPublicAccessBlockResponse) SetHeaders(v map[string]*string) *GetBucketPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *GetBucketPublicAccessBlockResponse) SetStatusCode(v int32) *GetBucketPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketPublicAccessBlockResponse) SetBody(v *GetBucketPublicAccessBlockResponseBody) *GetBucketPublicAccessBlockResponse {
	s.Body = v
	return s
}

func (s *GetBucketPublicAccessBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
