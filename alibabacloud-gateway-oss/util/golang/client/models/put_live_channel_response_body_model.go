// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateLiveChannelResult(v *PutLiveChannelResponseBodyCreateLiveChannelResult) *PutLiveChannelResponseBody
	GetCreateLiveChannelResult() *PutLiveChannelResponseBodyCreateLiveChannelResult
}

type PutLiveChannelResponseBody struct {
	// The container that stores the result of the CreateLiveChannel request.
	CreateLiveChannelResult *PutLiveChannelResponseBodyCreateLiveChannelResult `json:"CreateLiveChannelResult,omitempty" xml:"CreateLiveChannelResult,omitempty" type:"Struct"`
}

func (s PutLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *PutLiveChannelResponseBody) GetCreateLiveChannelResult() *PutLiveChannelResponseBodyCreateLiveChannelResult {
	return s.CreateLiveChannelResult
}

func (s *PutLiveChannelResponseBody) SetCreateLiveChannelResult(v *PutLiveChannelResponseBodyCreateLiveChannelResult) *PutLiveChannelResponseBody {
	s.CreateLiveChannelResult = v
	return s
}

func (s *PutLiveChannelResponseBody) Validate() error {
	if s.CreateLiveChannelResult != nil {
		if err := s.CreateLiveChannelResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutLiveChannelResponseBodyCreateLiveChannelResult struct {
	// 保存播放地址的容器。
	PlayUrls *LiveChannelPlayUrls `json:"PlayUrls,omitempty" xml:"PlayUrls,omitempty"`
	// 保存推流地址的容器。
	PublishUrls *LiveChannelPublishUrls `json:"PublishUrls,omitempty" xml:"PublishUrls,omitempty"`
}

func (s PutLiveChannelResponseBodyCreateLiveChannelResult) String() string {
	return dara.Prettify(s)
}

func (s PutLiveChannelResponseBodyCreateLiveChannelResult) GoString() string {
	return s.String()
}

func (s *PutLiveChannelResponseBodyCreateLiveChannelResult) GetPlayUrls() *LiveChannelPlayUrls {
	return s.PlayUrls
}

func (s *PutLiveChannelResponseBodyCreateLiveChannelResult) GetPublishUrls() *LiveChannelPublishUrls {
	return s.PublishUrls
}

func (s *PutLiveChannelResponseBodyCreateLiveChannelResult) SetPlayUrls(v *LiveChannelPlayUrls) *PutLiveChannelResponseBodyCreateLiveChannelResult {
	s.PlayUrls = v
	return s
}

func (s *PutLiveChannelResponseBodyCreateLiveChannelResult) SetPublishUrls(v *LiveChannelPublishUrls) *PutLiveChannelResponseBodyCreateLiveChannelResult {
	s.PublishUrls = v
	return s
}

func (s *PutLiveChannelResponseBodyCreateLiveChannelResult) Validate() error {
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
