// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationPrefixSet interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixs(v []*string) *ReplicationPrefixSet
	GetPrefixs() []*string
}

type ReplicationPrefixSet struct {
	Prefixs []*string `json:"Prefix,omitempty" xml:"Prefix,omitempty" type:"Repeated"`
}

func (s ReplicationPrefixSet) String() string {
	return dara.Prettify(s)
}

func (s ReplicationPrefixSet) GoString() string {
	return s.String()
}

func (s *ReplicationPrefixSet) GetPrefixs() []*string {
	return s.Prefixs
}

func (s *ReplicationPrefixSet) SetPrefixs(v []*string) *ReplicationPrefixSet {
	s.Prefixs = v
	return s
}

func (s *ReplicationPrefixSet) Validate() error {
	return dara.Validate(s)
}
