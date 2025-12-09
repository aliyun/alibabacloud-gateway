// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CommentConfiguration
	GetComment() *string
}

type CommentConfiguration struct {
	// example:
	//
	// comments
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s CommentConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CommentConfiguration) GoString() string {
	return s.String()
}

func (s *CommentConfiguration) GetComment() *string {
	return s.Comment
}

func (s *CommentConfiguration) SetComment(v string) *CommentConfiguration {
	s.Comment = &v
	return s
}

func (s *CommentConfiguration) Validate() error {
	return dara.Validate(s)
}
