// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutObjectRetentionRequestBody) *PutObjectRetentionRequest
	GetBody() *PutObjectRetentionRequestBody
}

type PutObjectRetentionRequest struct {
	Body *PutObjectRetentionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PutObjectRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectRetentionRequest) GoString() string {
	return s.String()
}

func (s *PutObjectRetentionRequest) GetBody() *PutObjectRetentionRequestBody {
	return s.Body
}

func (s *PutObjectRetentionRequest) SetBody(v *PutObjectRetentionRequestBody) *PutObjectRetentionRequest {
	s.Body = v
	return s
}

func (s *PutObjectRetentionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutObjectRetentionRequestBody struct {
	Retention *Retention `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s PutObjectRetentionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s PutObjectRetentionRequestBody) GoString() string {
	return s.String()
}

func (s *PutObjectRetentionRequestBody) GetRetention() *Retention {
	return s.Retention
}

func (s *PutObjectRetentionRequestBody) SetRetention(v *Retention) *PutObjectRetentionRequestBody {
	s.Retention = v
	return s
}

func (s *PutObjectRetentionRequestBody) Validate() error {
	if s.Retention != nil {
		if err := s.Retention.Validate(); err != nil {
			return err
		}
	}
	return nil
}
