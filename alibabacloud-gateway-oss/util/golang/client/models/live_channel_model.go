// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannel interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *LiveChannel
	GetDescription() *string
	SetLastModified(v string) *LiveChannel
	GetLastModified() *string
	SetName(v string) *LiveChannel
	GetName() *string
	SetPlayUrls(v *LiveChannelPlayUrls) *LiveChannel
	GetPlayUrls() *LiveChannelPlayUrls
	SetPublishUrls(v *LiveChannelPublishUrls) *LiveChannel
	GetPublishUrls() *LiveChannelPublishUrls
	SetStatus(v string) *LiveChannel
	GetStatus() *string
}

type LiveChannel struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	LastModified *string                 `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Name         *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PlayUrls     *LiveChannelPlayUrls    `json:"PlayUrls,omitempty" xml:"PlayUrls,omitempty"`
	PublishUrls  *LiveChannelPublishUrls `json:"PublishUrls,omitempty" xml:"PublishUrls,omitempty"`
	Status       *string                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s LiveChannel) String() string {
	return dara.Prettify(s)
}

func (s LiveChannel) GoString() string {
	return s.String()
}

func (s *LiveChannel) GetDescription() *string {
	return s.Description
}

func (s *LiveChannel) GetLastModified() *string {
	return s.LastModified
}

func (s *LiveChannel) GetName() *string {
	return s.Name
}

func (s *LiveChannel) GetPlayUrls() *LiveChannelPlayUrls {
	return s.PlayUrls
}

func (s *LiveChannel) GetPublishUrls() *LiveChannelPublishUrls {
	return s.PublishUrls
}

func (s *LiveChannel) GetStatus() *string {
	return s.Status
}

func (s *LiveChannel) SetDescription(v string) *LiveChannel {
	s.Description = &v
	return s
}

func (s *LiveChannel) SetLastModified(v string) *LiveChannel {
	s.LastModified = &v
	return s
}

func (s *LiveChannel) SetName(v string) *LiveChannel {
	s.Name = &v
	return s
}

func (s *LiveChannel) SetPlayUrls(v *LiveChannelPlayUrls) *LiveChannel {
	s.PlayUrls = v
	return s
}

func (s *LiveChannel) SetPublishUrls(v *LiveChannelPublishUrls) *LiveChannel {
	s.PublishUrls = v
	return s
}

func (s *LiveChannel) SetStatus(v string) *LiveChannel {
	s.Status = &v
	return s
}

func (s *LiveChannel) Validate() error {
	if s.PlayUrls != nil {
		if err := s.PlayUrls.Validate(); err != nil {
			return err
		}
	}
	if s.PublishUrls != nil {
		if err := s.PublishUrls.Validate(); err != nil {
			return err
		}
	}
	return nil
}
