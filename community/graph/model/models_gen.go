// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Comments struct {
	Count int        `json:"count"`
	Data  []*Comment `json:"data,omitempty"`
}

type Contact struct {
	ID       string     `json:"id"`
	Comments []*Comment `json:"comments,omitempty"`
	Shares   []*Share   `json:"shares,omitempty"`
}

func (Contact) IsEntity() {}

type Conversations struct {
	Data  []*Conversation `json:"data,omitempty"`
	Count int             `json:"count"`
}

type Expense struct {
	ID       string     `json:"id"`
	Comments []*Comment `json:"comments,omitempty"`
	Shares   []*Share   `json:"shares,omitempty"`
}

func (Expense) IsEntity() {}

type File struct {
	ID       string     `json:"id"`
	Comments []*Comment `json:"comments,omitempty"`
	Shares   []*Share   `json:"shares,omitempty"`
}

func (File) IsEntity() {}

type Messages struct {
	Data  []*Message `json:"data,omitempty"`
	Count int        `json:"count"`
}

type NewComment struct {
	Parent      *string                `json:"parent,omitempty"`
	Object      map[string]interface{} `json:"object"`
	Locale      string                 `json:"locale"`
	UID         *string                `json:"uid,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Email       *string                `json:"email,omitempty"`
	Body        *string                `json:"body,omitempty"`
	Rating      *int                   `json:"rating,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Attachments []*string              `json:"attachments,omitempty"`
}

type NewConversation struct {
	Type        string                 `json:"type"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewMessage struct {
	Conversation string                 `json:"conversation"`
	Recipients   []string               `json:"recipients"`
	Subject      string                 `json:"subject"`
	Body         map[string]interface{} `json:"body"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type NewReaction struct {
	Object map[string]interface{} `json:"object"`
	Action string                 `json:"action"`
}

type NewRecipient struct {
	UID     string `json:"uid"`
	Message string `json:"message"`
}

type NewTemplate struct {
	Locale   string                 `json:"locale"`
	Name     string                 `json:"name"`
	Subject  string                 `json:"subject"`
	Body     string                 `json:"body"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type Notifications struct {
	Data  []*Notification `json:"data,omitempty"`
	Count int             `json:"count"`
}

type Place struct {
	ID      string     `json:"id"`
	Reviews []*Comment `json:"reviews,omitempty"`
}

func (Place) IsEntity() {}

type Quote struct {
	ID       string     `json:"id"`
	Comments []*Comment `json:"comments,omitempty"`
	Shares   []*Share   `json:"shares,omitempty"`
}

func (Quote) IsEntity() {}

type Reactions struct {
	Data  []*Reaction `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Recipients struct {
	Count int          `json:"count"`
	Data  []*Recipient `json:"data,omitempty"`
}

type ShareInput struct {
	UID        string                 `json:"uid"`
	Object     map[string]interface{} `json:"object"`
	Permission string                 `json:"permission"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Status     string                 `json:"status"`
	Created    string                 `json:"created"`
	Updated    string                 `json:"updated"`
}

type ShareUpdateInput struct {
	UID        *string                `json:"uid,omitempty"`
	Object     map[string]interface{} `json:"object,omitempty"`
	Permission *string                `json:"permission,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Created    string                 `json:"created"`
	Updated    string                 `json:"updated"`
}

type Shares struct {
	Data  []*Share `json:"data,omitempty"`
	Count int      `json:"count"`
}

type Templates struct {
	Count int         `json:"count"`
	Data  []*Template `json:"data,omitempty"`
}

type UpdateComment struct {
	Parent      *string                `json:"parent,omitempty"`
	Locale      string                 `json:"locale"`
	UID         *string                `json:"uid,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Email       *string                `json:"email,omitempty"`
	Body        *string                `json:"body,omitempty"`
	Rating      *int                   `json:"rating,omitempty"`
	Recommended *bool                  `json:"recommended,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Attachments []*string              `json:"attachments,omitempty"`
}

type UpdateConversation struct {
	Type        *string                `json:"type,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateMessage struct {
	Subject  *string                `json:"subject,omitempty"`
	Body     map[string]interface{} `json:"body,omitempty"`
	Status   *string                `json:"status,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateReaction struct {
	Action string `json:"action"`
}

type UpdateRecipient struct {
	Read *string `json:"read,omitempty"`
}

type UpdateTemplate struct {
	Locale   *string                `json:"locale,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Subject  *string                `json:"subject,omitempty"`
	Body     *string                `json:"body,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type ConversationStatus string

const (
	ConversationStatusActive   ConversationStatus = "ACTIVE"
	ConversationStatusArchived ConversationStatus = "ARCHIVED"
	ConversationStatusDeleted  ConversationStatus = "DELETED"
)

var AllConversationStatus = []ConversationStatus{
	ConversationStatusActive,
	ConversationStatusArchived,
	ConversationStatusDeleted,
}

func (e ConversationStatus) IsValid() bool {
	switch e {
	case ConversationStatusActive, ConversationStatusArchived, ConversationStatusDeleted:
		return true
	}
	return false
}

func (e ConversationStatus) String() string {
	return string(e)
}

func (e *ConversationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConversationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConversationStatus", str)
	}
	return nil
}

func (e ConversationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MessageStatus string

const (
	MessageStatusDraft     MessageStatus = "DRAFT"
	MessageStatusSent      MessageStatus = "SENT"
	MessageStatusArchived  MessageStatus = "ARCHIVED"
	MessageStatusSpam      MessageStatus = "SPAM"
	MessageStatusScheduled MessageStatus = "SCHEDULED"
	MessageStatusDelivered MessageStatus = "DELIVERED"
	MessageStatusRead      MessageStatus = "READ"
	MessageStatusFailed    MessageStatus = "FAILED"
)

var AllMessageStatus = []MessageStatus{
	MessageStatusDraft,
	MessageStatusSent,
	MessageStatusArchived,
	MessageStatusSpam,
	MessageStatusScheduled,
	MessageStatusDelivered,
	MessageStatusRead,
	MessageStatusFailed,
}

func (e MessageStatus) IsValid() bool {
	switch e {
	case MessageStatusDraft, MessageStatusSent, MessageStatusArchived, MessageStatusSpam, MessageStatusScheduled, MessageStatusDelivered, MessageStatusRead, MessageStatusFailed:
		return true
	}
	return false
}

func (e MessageStatus) String() string {
	return string(e)
}

func (e *MessageStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MessageStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MessageStatus", str)
	}
	return nil
}

func (e MessageStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReactionType string

const (
	ReactionTypeNone  ReactionType = "NONE"
	ReactionTypeLike  ReactionType = "LIKE"
	ReactionTypeLove  ReactionType = "LOVE"
	ReactionTypeWow   ReactionType = "WOW"
	ReactionTypeHaha  ReactionType = "HAHA"
	ReactionTypeSorry ReactionType = "SORRY"
	ReactionTypeAngry ReactionType = "ANGRY"
)

var AllReactionType = []ReactionType{
	ReactionTypeNone,
	ReactionTypeLike,
	ReactionTypeLove,
	ReactionTypeWow,
	ReactionTypeHaha,
	ReactionTypeSorry,
	ReactionTypeAngry,
}

func (e ReactionType) IsValid() bool {
	switch e {
	case ReactionTypeNone, ReactionTypeLike, ReactionTypeLove, ReactionTypeWow, ReactionTypeHaha, ReactionTypeSorry, ReactionTypeAngry:
		return true
	}
	return false
}

func (e ReactionType) String() string {
	return string(e)
}

func (e *ReactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReactionType", str)
	}
	return nil
}

func (e ReactionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShareStatus string

const (
	ShareStatusPending  ShareStatus = "PENDING"
	ShareStatusAccepted ShareStatus = "ACCEPTED"
	ShareStatusRejected ShareStatus = "REJECTED"
	ShareStatusRevoked  ShareStatus = "REVOKED"
)

var AllShareStatus = []ShareStatus{
	ShareStatusPending,
	ShareStatusAccepted,
	ShareStatusRejected,
	ShareStatusRevoked,
}

func (e ShareStatus) IsValid() bool {
	switch e {
	case ShareStatusPending, ShareStatusAccepted, ShareStatusRejected, ShareStatusRevoked:
		return true
	}
	return false
}

func (e ShareStatus) String() string {
	return string(e)
}

func (e *ShareStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShareStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShareStatus", str)
	}
	return nil
}

func (e ShareStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
