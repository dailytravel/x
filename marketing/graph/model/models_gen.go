// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Audience struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Segments    []*Segment             `json:"segments,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type Audiences struct {
	Data  []*Audience `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Campaign struct {
	ID        string                 `json:"id"`
	Owner     *User                  `json:"owner"`
	Audience  *Audience              `json:"audience"`
	Type      string                 `json:"type"`
	Name      string                 `json:"name"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Status    string                 `json:"status"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	CreatedBy *User                  `json:"created_by,omitempty"`
	UpdatedBy *User                  `json:"updated_by,omitempty"`
	Responses []*Response            `json:"responses,omitempty"`
}

type Campaigns struct {
	Count int         `json:"count"`
	Data  []*Campaign `json:"data,omitempty"`
}

type Category struct {
	ID string `json:"id"`
}

func (Category) IsEntity() {}

type CreateSegmentInput struct {
	Name        string       `json:"name"`
	Description *string      `json:"description,omitempty"`
	Rules       []*RuleInput `json:"rules"`
}

type Follow struct {
	ID string `json:"id"`
}

func (Follow) IsEntity() {}

type Link struct {
	ID          string                 `json:"id"`
	Owner       *User                  `json:"owner"`
	Domain      string                 `json:"domain"`
	Reference   string                 `json:"reference"`
	Title       *string                `json:"title,omitempty"`
	Destination string                 `json:"destination"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Categories  []*Category            `json:"categories,omitempty"`
}

type Links struct {
	Data  []*Link `json:"data,omitempty"`
	Count int     `json:"count"`
}

type NewAudience struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Segment     *SegmentInput          `json:"segment"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewCampaign struct {
	Audience string                 `json:"audience"`
	Slug     string                 `json:"slug"`
	Type     CampaignType           `json:"type"`
	Name     string                 `json:"name"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type NewLink struct {
	Owner       *string                `json:"owner,omitempty"`
	Title       *string                `json:"title,omitempty"`
	Domain      string                 `json:"domain"`
	Reference   string                 `json:"reference"`
	Destination string                 `json:"destination"`
	Categories  []string               `json:"categories,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type Response struct {
	ID string `json:"id"`
}

func (Response) IsEntity() {}

type Rule struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RuleInput struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Schedule struct {
	ID        string                 `json:"id"`
	Campaign  string                 `json:"campaign"`
	StartDate string                 `json:"start_date"`
	EndDate   string                 `json:"end_date"`
	Recurring bool                   `json:"recurring"`
	Frequency *string                `json:"frequency,omitempty"`
	Interval  *int                   `json:"interval,omitempty"`
	Days      []*string              `json:"days,omitempty"`
	Time      string                 `json:"time"`
	Timezone  string                 `json:"timezone"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

type Segment struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Rules       []*Rule `json:"rules,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	CreatedBy   *User   `json:"created_by,omitempty"`
	UpdatedBy   *User   `json:"updated_by,omitempty"`
}

type SegmentInput struct {
	ID          *string      `json:"id,omitempty"`
	Name        string       `json:"name"`
	Description *string      `json:"description,omitempty"`
	Type        string       `json:"type"`
	Rules       []*RuleInput `json:"rules,omitempty"`
}

type Segments struct {
	Count int        `json:"count"`
	Data  []*Segment `json:"data,omitempty"`
}

type UpdateAudience struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Segment     *SegmentInput          `json:"segment,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateCampaign struct {
	Audience *string                `json:"audience,omitempty"`
	Slug     *string                `json:"slug,omitempty"`
	Type     *CampaignType          `json:"type,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateLink struct {
	Owner      *string                `json:"owner,omitempty"`
	Title      *string                `json:"title,omitempty"`
	Domain     *string                `json:"domain,omitempty"`
	Reference  *string                `json:"reference,omitempty"`
	Categories []string               `json:"categories,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateSegmentInput struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Rules       []*RuleInput `json:"rules,omitempty"`
}

type User struct {
	ID string `json:"id"`
}

func (User) IsEntity() {}

type CampaignStatus string

const (
	CampaignStatusDraft      CampaignStatus = "DRAFT"
	CampaignStatusScheduled  CampaignStatus = "SCHEDULED"
	CampaignStatusInProgress CampaignStatus = "IN_PROGRESS"
	CampaignStatusCompleted  CampaignStatus = "COMPLETED"
)

var AllCampaignStatus = []CampaignStatus{
	CampaignStatusDraft,
	CampaignStatusScheduled,
	CampaignStatusInProgress,
	CampaignStatusCompleted,
}

func (e CampaignStatus) IsValid() bool {
	switch e {
	case CampaignStatusDraft, CampaignStatusScheduled, CampaignStatusInProgress, CampaignStatusCompleted:
		return true
	}
	return false
}

func (e CampaignStatus) String() string {
	return string(e)
}

func (e *CampaignStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CampaignStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CampaignStatus", str)
	}
	return nil
}

func (e CampaignStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CampaignType string

const (
	CampaignTypeEmail            CampaignType = "EMAIL"
	CampaignTypeSms              CampaignType = "SMS"
	CampaignTypePushNotification CampaignType = "PUSH_NOTIFICATION"
)

var AllCampaignType = []CampaignType{
	CampaignTypeEmail,
	CampaignTypeSms,
	CampaignTypePushNotification,
}

func (e CampaignType) IsValid() bool {
	switch e {
	case CampaignTypeEmail, CampaignTypeSms, CampaignTypePushNotification:
		return true
	}
	return false
}

func (e CampaignType) String() string {
	return string(e)
}

func (e *CampaignType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CampaignType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CampaignType", str)
	}
	return nil
}

func (e CampaignType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
