// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListCnameResult(v *ListCnameResponseBodyListCnameResult) *ListCnameResponseBody
	GetListCnameResult() *ListCnameResponseBodyListCnameResult
}

type ListCnameResponseBody struct {
	// The container that is used to query information about all CNAME records.
	ListCnameResult *ListCnameResponseBodyListCnameResult `json:"ListCnameResult,omitempty" xml:"ListCnameResult,omitempty" type:"Struct"`
}

func (s ListCnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCnameResponseBody) GoString() string {
	return s.String()
}

func (s *ListCnameResponseBody) GetListCnameResult() *ListCnameResponseBodyListCnameResult {
	return s.ListCnameResult
}

func (s *ListCnameResponseBody) SetListCnameResult(v *ListCnameResponseBodyListCnameResult) *ListCnameResponseBody {
	s.ListCnameResult = v
	return s
}

func (s *ListCnameResponseBody) Validate() error {
	if s.ListCnameResult != nil {
		if err := s.ListCnameResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCnameResponseBodyListCnameResult struct {
	// The name of the bucket to which the CNAME records you want to query are mapped.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The container that is used to store the information about all CNAME records.
	Cname []*CnameInfo `json:"Cname,omitempty" xml:"Cname,omitempty" type:"Repeated"`
	// The name of the bucket owner.
	//
	// example:
	//
	// 133413***273506
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s ListCnameResponseBodyListCnameResult) String() string {
	return dara.Prettify(s)
}

func (s ListCnameResponseBodyListCnameResult) GoString() string {
	return s.String()
}

func (s *ListCnameResponseBodyListCnameResult) GetBucket() *string {
	return s.Bucket
}

func (s *ListCnameResponseBodyListCnameResult) GetCname() []*CnameInfo {
	return s.Cname
}

func (s *ListCnameResponseBodyListCnameResult) GetOwner() *string {
	return s.Owner
}

func (s *ListCnameResponseBodyListCnameResult) SetBucket(v string) *ListCnameResponseBodyListCnameResult {
	s.Bucket = &v
	return s
}

func (s *ListCnameResponseBodyListCnameResult) SetCname(v []*CnameInfo) *ListCnameResponseBodyListCnameResult {
	s.Cname = v
	return s
}

func (s *ListCnameResponseBodyListCnameResult) SetOwner(v string) *ListCnameResponseBodyListCnameResult {
	s.Owner = &v
	return s
}

func (s *ListCnameResponseBodyListCnameResult) Validate() error {
	if s.Cname != nil {
		for _, item := range s.Cname {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
