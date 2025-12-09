// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationSourceSelectionCriteria interface {
	dara.Model
	String() string
	GoString() string
	SetSseKmsEncryptedObjects(v *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) *ReplicationSourceSelectionCriteria
	GetSseKmsEncryptedObjects() *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects
}

type ReplicationSourceSelectionCriteria struct {
	SseKmsEncryptedObjects *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects `json:"SseKmsEncryptedObjects,omitempty" xml:"SseKmsEncryptedObjects,omitempty" type:"Struct"`
}

func (s ReplicationSourceSelectionCriteria) String() string {
	return dara.Prettify(s)
}

func (s ReplicationSourceSelectionCriteria) GoString() string {
	return s.String()
}

func (s *ReplicationSourceSelectionCriteria) GetSseKmsEncryptedObjects() *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects {
	return s.SseKmsEncryptedObjects
}

func (s *ReplicationSourceSelectionCriteria) SetSseKmsEncryptedObjects(v *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) *ReplicationSourceSelectionCriteria {
	s.SseKmsEncryptedObjects = v
	return s
}

func (s *ReplicationSourceSelectionCriteria) Validate() error {
	if s.SseKmsEncryptedObjects != nil {
		if err := s.SseKmsEncryptedObjects.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) String() string {
	return dara.Prettify(s)
}

func (s ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) GoString() string {
	return s.String()
}

func (s *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) GetStatus() *string {
	return s.Status
}

func (s *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) SetStatus(v string) *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects {
	s.Status = &v
	return s
}

func (s *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) Validate() error {
	return dara.Validate(s)
}
