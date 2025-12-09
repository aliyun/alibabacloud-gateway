// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCommentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommentConfiguration(v *CommentConfiguration) *PutBucketCommentRequest
	GetCommentConfiguration() *CommentConfiguration
}

type PutBucketCommentRequest struct {
	CommentConfiguration *CommentConfiguration `json:"CommentConfiguration,omitempty" xml:"CommentConfiguration,omitempty"`
}

func (s PutBucketCommentRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCommentRequest) GoString() string {
	return s.String()
}

func (s *PutBucketCommentRequest) GetCommentConfiguration() *CommentConfiguration {
	return s.CommentConfiguration
}

func (s *PutBucketCommentRequest) SetCommentConfiguration(v *CommentConfiguration) *PutBucketCommentRequest {
	s.CommentConfiguration = v
	return s
}

func (s *PutBucketCommentRequest) Validate() error {
	if s.CommentConfiguration != nil {
		if err := s.CommentConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
