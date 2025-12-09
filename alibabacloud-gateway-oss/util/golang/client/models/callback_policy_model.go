// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyItem(v []*CallbackPolicyPolicyItem) *CallbackPolicy
	GetPolicyItem() []*CallbackPolicyPolicyItem
}

type CallbackPolicy struct {
	PolicyItem []*CallbackPolicyPolicyItem `json:"PolicyItem,omitempty" xml:"PolicyItem,omitempty" type:"Repeated"`
}

func (s CallbackPolicy) String() string {
	return dara.Prettify(s)
}

func (s CallbackPolicy) GoString() string {
	return s.String()
}

func (s *CallbackPolicy) GetPolicyItem() []*CallbackPolicyPolicyItem {
	return s.PolicyItem
}

func (s *CallbackPolicy) SetPolicyItem(v []*CallbackPolicyPolicyItem) *CallbackPolicy {
	s.PolicyItem = v
	return s
}

func (s *CallbackPolicy) Validate() error {
	if s.PolicyItem != nil {
		for _, item := range s.PolicyItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CallbackPolicyPolicyItem struct {
	// example:
	//
	// e1wiY2Fsb...9keVwiOlwiYnVja2V0PSR7YnU=
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// Q2Fs...FcIiwgXCJ4OmJcIjpcImJcIn0=
	CallbackVar *string `json:"CallbackVar,omitempty" xml:"CallbackVar,omitempty"`
	// example:
	//
	// first
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s CallbackPolicyPolicyItem) String() string {
	return dara.Prettify(s)
}

func (s CallbackPolicyPolicyItem) GoString() string {
	return s.String()
}

func (s *CallbackPolicyPolicyItem) GetCallback() *string {
	return s.Callback
}

func (s *CallbackPolicyPolicyItem) GetCallbackVar() *string {
	return s.CallbackVar
}

func (s *CallbackPolicyPolicyItem) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CallbackPolicyPolicyItem) SetCallback(v string) *CallbackPolicyPolicyItem {
	s.Callback = &v
	return s
}

func (s *CallbackPolicyPolicyItem) SetCallbackVar(v string) *CallbackPolicyPolicyItem {
	s.CallbackVar = &v
	return s
}

func (s *CallbackPolicyPolicyItem) SetPolicyName(v string) *CallbackPolicyPolicyItem {
	s.PolicyName = &v
	return s
}

func (s *CallbackPolicyPolicyItem) Validate() error {
	return dara.Validate(s)
}
