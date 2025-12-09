// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncFetchTaskConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *AsyncFetchTaskConfiguration
	GetCallback() *string
	SetContentMD5(v string) *AsyncFetchTaskConfiguration
	GetContentMD5() *string
	SetHost(v string) *AsyncFetchTaskConfiguration
	GetHost() *string
	SetIgnoreSameKey(v bool) *AsyncFetchTaskConfiguration
	GetIgnoreSameKey() *bool
	SetObject(v string) *AsyncFetchTaskConfiguration
	GetObject() *string
	SetStorageClass(v string) *AsyncFetchTaskConfiguration
	GetStorageClass() *string
	SetUrl(v string) *AsyncFetchTaskConfiguration
	GetUrl() *string
}

type AsyncFetchTaskConfiguration struct {
	// example:
	//
	// eyJjYWxsYmFja1VybCI6Ind3dy5hYmMuY29tL2NhbGxiYWNrIiwiY2FsbGJhY2tCb2R5IjoiJHtldGFnfSJ9
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// v23MlMRM/EgJczOs2yHTcA==
	ContentMD5 *string `json:"ContentMD5,omitempty" xml:"ContentMD5,omitempty"`
	// example:
	//
	// www.test.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// True
	IgnoreSameKey *bool `json:"IgnoreSameKey,omitempty" xml:"IgnoreSameKey,omitempty"`
	// example:
	//
	// abc.txt
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// example:
	//
	// www.test.com/abc.txt
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AsyncFetchTaskConfiguration) String() string {
	return dara.Prettify(s)
}

func (s AsyncFetchTaskConfiguration) GoString() string {
	return s.String()
}

func (s *AsyncFetchTaskConfiguration) GetCallback() *string {
	return s.Callback
}

func (s *AsyncFetchTaskConfiguration) GetContentMD5() *string {
	return s.ContentMD5
}

func (s *AsyncFetchTaskConfiguration) GetHost() *string {
	return s.Host
}

func (s *AsyncFetchTaskConfiguration) GetIgnoreSameKey() *bool {
	return s.IgnoreSameKey
}

func (s *AsyncFetchTaskConfiguration) GetObject() *string {
	return s.Object
}

func (s *AsyncFetchTaskConfiguration) GetStorageClass() *string {
	return s.StorageClass
}

func (s *AsyncFetchTaskConfiguration) GetUrl() *string {
	return s.Url
}

func (s *AsyncFetchTaskConfiguration) SetCallback(v string) *AsyncFetchTaskConfiguration {
	s.Callback = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) SetContentMD5(v string) *AsyncFetchTaskConfiguration {
	s.ContentMD5 = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) SetHost(v string) *AsyncFetchTaskConfiguration {
	s.Host = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) SetIgnoreSameKey(v bool) *AsyncFetchTaskConfiguration {
	s.IgnoreSameKey = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) SetObject(v string) *AsyncFetchTaskConfiguration {
	s.Object = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) SetStorageClass(v string) *AsyncFetchTaskConfiguration {
	s.StorageClass = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) SetUrl(v string) *AsyncFetchTaskConfiguration {
	s.Url = &v
	return s
}

func (s *AsyncFetchTaskConfiguration) Validate() error {
	return dara.Validate(s)
}
