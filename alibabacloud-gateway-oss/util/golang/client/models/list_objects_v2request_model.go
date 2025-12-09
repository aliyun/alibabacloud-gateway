// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListObjectsV2Request
	GetContinuationToken() *string
	SetDelimiter(v string) *ListObjectsV2Request
	GetDelimiter() *string
	SetEncodingType(v string) *ListObjectsV2Request
	GetEncodingType() *string
	SetFetchOwner(v bool) *ListObjectsV2Request
	GetFetchOwner() *bool
	SetMaxKeys(v int64) *ListObjectsV2Request
	GetMaxKeys() *int64
	SetPrefix(v string) *ListObjectsV2Request
	GetPrefix() *string
	SetStartAfter(v string) *ListObjectsV2Request
	GetStartAfter() *string
}

type ListObjectsV2Request struct {
	// The token from which the list operation starts. You can obtain the token from NextContinuationToken in the response of the ListObjectsV2 request.
	//
	// example:
	//
	// test1.txt
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// The character that is used to group objects by name. If you specify delimiter in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
	//
	// example:
	//
	// /
	Delimiter *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	// The encoding format of the returned objects in the response.
	//
	// >  The values of Delimiter, StartAfter, Prefix, NextContinuationToken, and Key are UTF-8 encoded. If the value of Delimiter, StartAfter, Prefix, NextContinuationToken, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	// Specifies whether to include the information about the bucket owner in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	FetchOwner *bool `json:"fetch-owner,omitempty" xml:"fetch-owner,omitempty"`
	// The maximum number of objects to be returned.\\
	//
	// Valid values: 1 to 999.\\
	//
	// Default value: 100.
	//
	// >  If the number of returned objects exceeds the value of max-keys, the response contains NextContinuationToken.Use the value of NextContinuationToken as the value of continuation-token in the next request.
	//
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// The prefix that must be contained in names of the returned objects.\\
	//
	//
	// 	- The value of prefix can be up to 1,024 bytes in length.
	//
	// 	- If you specify prefix, the names of the returned objects contain the prefix.
	//
	// If you set prefix to a directory name, the objects whose names start with this prefix are listed. The objects consist of all objects and subdirectories in this directory.\\
	//
	// If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.\\
	//
	// For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The name of the object after which the list operation begins. If this parameter is specified, objects whose names are alphabetically after the value of start-after are returned.\\
	//
	// The objects are returned by page based on start-after. The value of start-after can be up to 1,024 bytes in length.\\
	//
	// If the value of start-after does not exist when you perform a conditional query, the list starts from the object whose name is alphabetically after the value of start-after.
	//
	// example:
	//
	// b
	StartAfter *string `json:"start-after,omitempty" xml:"start-after,omitempty"`
}

func (s ListObjectsV2Request) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsV2Request) GoString() string {
	return s.String()
}

func (s *ListObjectsV2Request) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListObjectsV2Request) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListObjectsV2Request) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectsV2Request) GetFetchOwner() *bool {
	return s.FetchOwner
}

func (s *ListObjectsV2Request) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListObjectsV2Request) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectsV2Request) GetStartAfter() *string {
	return s.StartAfter
}

func (s *ListObjectsV2Request) SetContinuationToken(v string) *ListObjectsV2Request {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsV2Request) SetDelimiter(v string) *ListObjectsV2Request {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsV2Request) SetEncodingType(v string) *ListObjectsV2Request {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsV2Request) SetFetchOwner(v bool) *ListObjectsV2Request {
	s.FetchOwner = &v
	return s
}

func (s *ListObjectsV2Request) SetMaxKeys(v int64) *ListObjectsV2Request {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsV2Request) SetPrefix(v string) *ListObjectsV2Request {
	s.Prefix = &v
	return s
}

func (s *ListObjectsV2Request) SetStartAfter(v string) *ListObjectsV2Request {
	s.StartAfter = &v
	return s
}

func (s *ListObjectsV2Request) Validate() error {
	return dara.Validate(s)
}
