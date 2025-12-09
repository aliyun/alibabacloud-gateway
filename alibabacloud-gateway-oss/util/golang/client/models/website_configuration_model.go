// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsiteConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetErrorDocument(v *ErrorDocument) *WebsiteConfiguration
	GetErrorDocument() *ErrorDocument
	SetIndexDocument(v *IndexDocument) *WebsiteConfiguration
	GetIndexDocument() *IndexDocument
	SetRoutingRules(v *WebsiteConfigurationRoutingRules) *WebsiteConfiguration
	GetRoutingRules() *WebsiteConfigurationRoutingRules
}

type WebsiteConfiguration struct {
	ErrorDocument *ErrorDocument                    `json:"ErrorDocument,omitempty" xml:"ErrorDocument,omitempty"`
	IndexDocument *IndexDocument                    `json:"IndexDocument,omitempty" xml:"IndexDocument,omitempty"`
	RoutingRules  *WebsiteConfigurationRoutingRules `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty" type:"Struct"`
}

func (s WebsiteConfiguration) String() string {
	return dara.Prettify(s)
}

func (s WebsiteConfiguration) GoString() string {
	return s.String()
}

func (s *WebsiteConfiguration) GetErrorDocument() *ErrorDocument {
	return s.ErrorDocument
}

func (s *WebsiteConfiguration) GetIndexDocument() *IndexDocument {
	return s.IndexDocument
}

func (s *WebsiteConfiguration) GetRoutingRules() *WebsiteConfigurationRoutingRules {
	return s.RoutingRules
}

func (s *WebsiteConfiguration) SetErrorDocument(v *ErrorDocument) *WebsiteConfiguration {
	s.ErrorDocument = v
	return s
}

func (s *WebsiteConfiguration) SetIndexDocument(v *IndexDocument) *WebsiteConfiguration {
	s.IndexDocument = v
	return s
}

func (s *WebsiteConfiguration) SetRoutingRules(v *WebsiteConfigurationRoutingRules) *WebsiteConfiguration {
	s.RoutingRules = v
	return s
}

func (s *WebsiteConfiguration) Validate() error {
	if s.ErrorDocument != nil {
		if err := s.ErrorDocument.Validate(); err != nil {
			return err
		}
	}
	if s.IndexDocument != nil {
		if err := s.IndexDocument.Validate(); err != nil {
			return err
		}
	}
	if s.RoutingRules != nil {
		if err := s.RoutingRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WebsiteConfigurationRoutingRules struct {
	RoutingRule []*RoutingRule `json:"RoutingRule,omitempty" xml:"RoutingRule,omitempty" type:"Repeated"`
}

func (s WebsiteConfigurationRoutingRules) String() string {
	return dara.Prettify(s)
}

func (s WebsiteConfigurationRoutingRules) GoString() string {
	return s.String()
}

func (s *WebsiteConfigurationRoutingRules) GetRoutingRule() []*RoutingRule {
	return s.RoutingRule
}

func (s *WebsiteConfigurationRoutingRules) SetRoutingRule(v []*RoutingRule) *WebsiteConfigurationRoutingRules {
	s.RoutingRule = v
	return s
}

func (s *WebsiteConfigurationRoutingRules) Validate() error {
	if s.RoutingRule != nil {
		for _, item := range s.RoutingRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
