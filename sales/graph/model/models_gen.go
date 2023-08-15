// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Benefits struct {
	Data  []*Benefit `json:"data,omitempty"`
	Count *int       `json:"count,omitempty"`
}

type Companies struct {
	Data  []*Company `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Contacts struct {
	Count int        `json:"count"`
	Data  []*Contact `json:"data,omitempty"`
}

type Contracts struct {
	Data  []*Contract `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Coupons struct {
	Data  []*Coupon `json:"data,omitempty"`
	Count int       `json:"count"`
}

type DateRange struct {
	StartDate *string `json:"start_date,omitempty"`
	EndDate   *int    `json:"end_date,omitempty"`
}

type Email struct {
	Personal *string `json:"personal,omitempty"`
	Work     *string `json:"work,omitempty"`
	Other    *string `json:"other,omitempty"`
}

type EmailInput struct {
	Personal *string `json:"personal,omitempty"`
	Work     *string `json:"work,omitempty"`
	Other    *string `json:"other,omitempty"`
}

type Inventories struct {
	Data  []*Inventory `json:"data,omitempty"`
	Count *int         `json:"count,omitempty"`
}

type NewBenefit struct {
	Locale      string                 `json:"locale"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewCompany struct {
	User        *string                `json:"user,omitempty"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Parent      *string                `json:"parent,omitempty"`
	Industry    *string                `json:"industry,omitempty"`
	Employees   *int                   `json:"employees,omitempty"`
	Revenue     *float64               `json:"revenue,omitempty"`
	City        *string                `json:"city,omitempty"`
	State       *string                `json:"state,omitempty"`
	Zip         *string                `json:"zip,omitempty"`
	Country     *string                `json:"country,omitempty"`
	Timezone    *string                `json:"timezone,omitempty"`
	Phone       *string                `json:"phone,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
}

type NewContact struct {
	Company    *string                `json:"company,omitempty"`
	User       *string                `json:"user,omitempty"`
	Type       *string                `json:"type,omitempty"`
	FirstName  *string                `json:"first_name,omitempty"`
	LastName   *string                `json:"last_name,omitempty"`
	Picture    *string                `json:"picture,omitempty"`
	Email      *EmailInput            `json:"email,omitempty"`
	Phone      *PhoneInput            `json:"phone,omitempty"`
	Street     *string                `json:"street,omitempty"`
	City       *string                `json:"city,omitempty"`
	State      *string                `json:"state,omitempty"`
	Zip        *string                `json:"zip,omitempty"`
	Country    *string                `json:"country,omitempty"`
	Website    *string                `json:"website,omitempty"`
	Birthday   *string                `json:"birthday,omitempty"`
	JobTitle   *string                `json:"job_title,omitempty"`
	Gender     *Gender                `json:"gender,omitempty"`
	Timezone   *string                `json:"timezone,omitempty"`
	Language   *string                `json:"language,omitempty"`
	Source     *string                `json:"source,omitempty"`
	Revenue    *float64               `json:"revenue,omitempty"`
	Rating     *int                   `json:"rating,omitempty"`
	Notes      *string                `json:"notes,omitempty"`
	Stage      *string                `json:"stage,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Labels     []*string              `json:"labels,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Reviewable *bool                  `json:"reviewable,omitempty"`
}

type NewContract struct {
	User        string                 `json:"user"`
	Contact     string                 `json:"contact"`
	Reference   string                 `json:"reference"`
	Description string                 `json:"description"`
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	StartDate   string                 `json:"start_date"`
	EndDate     string                 `json:"end_date"`
	AutoRenew   bool                   `json:"auto_renew"`
	Categories  []string               `json:"categories,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewCoupon struct {
	Locale      string                 `json:"locale"`
	Code        string                 `json:"code"`
	Description string                 `json:"description"`
	Type        DiscountType           `json:"type"`
	Amount      float64                `json:"amount"`
	MaxDiscount *float64               `json:"max_discount,omitempty"`
	Currency    string                 `json:"currency"`
	StartDate   string                 `json:"start_date"`
	EndDate     int                    `json:"end_date"`
	Products    []string               `json:"products,omitempty"`
	MaxUses     *int                   `json:"max_uses,omitempty"`
	MinPurchase *float64               `json:"min_purchase,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewInventory struct {
	Date     int  `json:"date"`
	Quantity *int `json:"quantity,omitempty"`
}

type NewOrder struct {
	Reference   string   `json:"reference"`
	Data        []string `json:"data,omitempty"`
	Total       float64  `json:"total"`
	Currency    string   `json:"currency"`
	Cancellable bool     `json:"cancellable"`
	Status      string   `json:"status"`
}

type NewPoint struct {
	Points    int                    `json:"points"`
	Type      string                 `json:"type"`
	ExpiresAt *string                `json:"expires_at,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Status    string                 `json:"status"`
}

type NewPrice struct {
	StartDate *string                `json:"start_date,omitempty"`
	EndDate   *int                   `json:"end_date,omitempty"`
	Regular   float64                `json:"regular"`
	Sale      *float64               `json:"sale,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type NewProduct struct {
	Sku         string                 `json:"sku"`
	Locale      string                 `json:"locale"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Quantity    int                    `json:"quantity"`
	Currency    string                 `json:"currency"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewProgram struct {
	Type        string                 `json:"type"`
	Locale      string                 `json:"locale"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Points      int                    `json:"points"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewQuote struct {
	User        string                 `json:"user"`
	Contact     *string                `json:"contact,omitempty"`
	Locale      string                 `json:"locale"`
	Reference   string                 `json:"reference"`
	Purchase    string                 `json:"purchase"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Terms       *string                `json:"terms,omitempty"`
	Notes       *string                `json:"notes,omitempty"`
	Template    string                 `json:"template"`
	ValidUntil  int                    `json:"valid_until"`
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	Billing     map[string]interface{} `json:"billing"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type NewReward struct {
	Locale      string                 `json:"locale"`
	Tier        string                 `json:"tier"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Cost        int                    `json:"cost"`
	ExpiresAt   *string                `json:"expires_at,omitempty"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewTier struct {
	Locale      string                 `json:"locale"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Cost        int                    `json:"cost"`
	Rewards     []string               `json:"rewards"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewTransaction struct {
	User        string                 `json:"user"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Date        string                 `json:"date"`
	Description string                 `json:"description"`
}

type NewWishlist struct {
	Product  string                 `json:"product"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Orders struct {
	Data  []*Order `json:"data,omitempty"`
	Count int      `json:"count"`
}

type Phone struct {
	Mobile *string `json:"mobile,omitempty"`
	Work   *string `json:"work,omitempty"`
	Home   *string `json:"home,omitempty"`
	Other  *string `json:"other,omitempty"`
}

type PhoneInput struct {
	Mobile *string `json:"mobile,omitempty"`
	Work   *string `json:"work,omitempty"`
	Home   *string `json:"home,omitempty"`
	Other  *string `json:"other,omitempty"`
}

type Pipeline struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Parent      *Pipeline              `json:"parent,omitempty"`
}

type Points struct {
	Count int      `json:"count"`
	Data  []*Point `json:"data,omitempty"`
}

type Prices struct {
	Data  []*Price `json:"data,omitempty"`
	Count *int     `json:"count,omitempty"`
}

type Products struct {
	Data  []*Product `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Programs struct {
	Data  []*Program `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Quotes struct {
	Count int      `json:"count"`
	Data  []*Quote `json:"data,omitempty"`
}

type Rewards struct {
	Data  []*Reward `json:"data,omitempty"`
	Count int       `json:"count"`
}

type Tiers struct {
	Data  []*Tier `json:"data,omitempty"`
	Count int     `json:"count"`
}

type Transaction struct {
	ID          string                 `json:"id"`
	User        *User                  `json:"user"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Date        string                 `json:"date"`
	Description string                 `json:"description"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type Transactions struct {
	Data  []*Transaction `json:"data,omitempty"`
	Count int            `json:"count"`
}

type UpdateBenefit struct {
	Locale      *string                `json:"locale,omitempty"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateCompany struct {
	User        *string                `json:"user,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Parent      *string                `json:"parent,omitempty"`
	Industry    *string                `json:"industry,omitempty"`
	Employees   *int                   `json:"employees,omitempty"`
	Revenue     *float64               `json:"revenue,omitempty"`
	City        *string                `json:"city,omitempty"`
	Zip         *string                `json:"zip,omitempty"`
	State       *string                `json:"state,omitempty"`
	Country     *string                `json:"country,omitempty"`
	Timezone    *string                `json:"timezone,omitempty"`
	Phone       *string                `json:"phone,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type UpdateContact struct {
	Company    *string                `json:"company,omitempty"`
	User       *string                `json:"user,omitempty"`
	Reference  *string                `json:"reference,omitempty"`
	Type       *string                `json:"type,omitempty"`
	FirstName  *string                `json:"first_name,omitempty"`
	LastName   *string                `json:"last_name,omitempty"`
	Picture    *string                `json:"picture,omitempty"`
	Email      *EmailInput            `json:"email,omitempty"`
	Phone      *PhoneInput            `json:"phone,omitempty"`
	Street     *string                `json:"street,omitempty"`
	City       *string                `json:"city,omitempty"`
	State      *string                `json:"state,omitempty"`
	Zip        *string                `json:"zip,omitempty"`
	Country    *string                `json:"country,omitempty"`
	Website    *string                `json:"website,omitempty"`
	Birthday   *string                `json:"birthday,omitempty"`
	JobTitle   *string                `json:"job_title,omitempty"`
	Gender     *Gender                `json:"gender,omitempty"`
	Timezone   *string                `json:"timezone,omitempty"`
	Language   *string                `json:"language,omitempty"`
	Source     *string                `json:"source,omitempty"`
	Revenue    *float64               `json:"revenue,omitempty"`
	Rating     *int                   `json:"rating,omitempty"`
	Notes      *string                `json:"notes,omitempty"`
	Stage      *string                `json:"stage,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Labels     []*string              `json:"labels,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Reviewable *bool                  `json:"reviewable,omitempty"`
}

type UpdateContract struct {
	User        *string                `json:"user,omitempty"`
	Contact     *string                `json:"contact,omitempty"`
	Reference   *string                `json:"reference,omitempty"`
	Description *string                `json:"description,omitempty"`
	Amount      *float64               `json:"amount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	StartDate   *string                `json:"start_date,omitempty"`
	EndDate     *string                `json:"end_date,omitempty"`
	AutoRenew   *bool                  `json:"auto_renew,omitempty"`
	Categories  []string               `json:"categories,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateCoupon struct {
	Locale      string                 `json:"locale"`
	Code        *string                `json:"code,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *DiscountType          `json:"type,omitempty"`
	Amount      *float64               `json:"amount,omitempty"`
	MaxDiscount *float64               `json:"max_discount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	StartDate   *string                `json:"start_date,omitempty"`
	EndDate     *int                   `json:"end_date,omitempty"`
	Products    []string               `json:"products,omitempty"`
	MaxUses     *int                   `json:"max_uses,omitempty"`
	MinPurchase *float64               `json:"min_purchase,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateInventory struct {
	Date     *int `json:"date,omitempty"`
	Quantity *int `json:"quantity,omitempty"`
}

type UpdateOrder struct {
	Reference   *string  `json:"reference,omitempty"`
	Data        []string `json:"data,omitempty"`
	Total       *float64 `json:"total,omitempty"`
	Currency    *string  `json:"currency,omitempty"`
	Cancellable *bool    `json:"cancellable,omitempty"`
	Status      *string  `json:"status,omitempty"`
	CancelledAt *int     `json:"cancelled_at,omitempty"`
}

type UpdatePoint struct {
	Points    *int                   `json:"points,omitempty"`
	Type      *string                `json:"type,omitempty"`
	ExpiresAt *string                `json:"expires_at,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Status    *string                `json:"status,omitempty"`
}

type UpdatePrice struct {
	StartDate *string                `json:"start_date,omitempty"`
	EndDate   *int                   `json:"end_date,omitempty"`
	Regular   *float64               `json:"regular,omitempty"`
	Sale      *float64               `json:"sale,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateProduct struct {
	Sku         *string                `json:"sku,omitempty"`
	Locale      *string                `json:"locale,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Quantity    *int                   `json:"quantity,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateProgram struct {
	Locale      *string                `json:"locale,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Points      *int                   `json:"points,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateQuote struct {
	User        *string                `json:"user,omitempty"`
	Contact     *string                `json:"contact,omitempty"`
	Locale      *string                `json:"locale,omitempty"`
	Reference   *string                `json:"reference,omitempty"`
	Purchase    *string                `json:"purchase,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Terms       *string                `json:"terms,omitempty"`
	Notes       *string                `json:"notes,omitempty"`
	Template    *string                `json:"template,omitempty"`
	ValidUntil  *int                   `json:"valid_until,omitempty"`
	Amount      *float64               `json:"amount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Billing     map[string]interface{} `json:"billing,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type UpdateReward struct {
	Locale      string                 `json:"locale"`
	Tier        *string                `json:"tier,omitempty"`
	Cost        *string                `json:"cost,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Points      *int                   `json:"points,omitempty"`
	ExpiresAt   *string                `json:"expires_at,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateTier struct {
	Locale      string                 `json:"locale"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Cost        *int                   `json:"cost,omitempty"`
	Rewards     []string               `json:"rewards,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateTransaction struct {
	User        *string                `json:"user,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Amount      *float64               `json:"amount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Date        *string                `json:"date,omitempty"`
	Description *string                `json:"description,omitempty"`
}

type UpdateWishlist struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type Wishlist struct {
	ID        string                 `json:"id"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Status    string                 `json:"status"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	Product   string                 `json:"product"`
	UID       string                 `json:"uid"`
	CreatedBy *string                `json:"created_by,omitempty"`
	UpdatedBy *string                `json:"updated_by,omitempty"`
}

type Wishlists struct {
	Count int         `json:"count"`
	Data  []*Wishlist `json:"data"`
}

type ContactStage string

const (
	ContactStageNew        ContactStage = "NEW"
	ContactStageNurturing  ContactStage = "NURTURING"
	ContactStageQualified  ContactStage = "QUALIFIED"
	ContactStageClosedLost ContactStage = "CLOSED_LOST"
	ContactStageClosedWon  ContactStage = "CLOSED_WON"
)

var AllContactStage = []ContactStage{
	ContactStageNew,
	ContactStageNurturing,
	ContactStageQualified,
	ContactStageClosedLost,
	ContactStageClosedWon,
}

func (e ContactStage) IsValid() bool {
	switch e {
	case ContactStageNew, ContactStageNurturing, ContactStageQualified, ContactStageClosedLost, ContactStageClosedWon:
		return true
	}
	return false
}

func (e ContactStage) String() string {
	return string(e)
}

func (e *ContactStage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContactStage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactStage", str)
	}
	return nil
}

func (e ContactStage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContactStatus string

const (
	ContactStatusActive   ContactStatus = "ACTIVE"
	ContactStatusInactive ContactStatus = "INACTIVE"
	ContactStatusArchived ContactStatus = "ARCHIVED"
)

var AllContactStatus = []ContactStatus{
	ContactStatusActive,
	ContactStatusInactive,
	ContactStatusArchived,
}

func (e ContactStatus) IsValid() bool {
	switch e {
	case ContactStatusActive, ContactStatusInactive, ContactStatusArchived:
		return true
	}
	return false
}

func (e ContactStatus) String() string {
	return string(e)
}

func (e *ContactStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContactStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactStatus", str)
	}
	return nil
}

func (e ContactStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContractStatus string

const (
	ContractStatusDraft      ContractStatus = "DRAFT"
	ContractStatusPending    ContractStatus = "PENDING"
	ContractStatusActive     ContractStatus = "ACTIVE"
	ContractStatusExpired    ContractStatus = "EXPIRED"
	ContractStatusTerminated ContractStatus = "TERMINATED"
)

var AllContractStatus = []ContractStatus{
	ContractStatusDraft,
	ContractStatusPending,
	ContractStatusActive,
	ContractStatusExpired,
	ContractStatusTerminated,
}

func (e ContractStatus) IsValid() bool {
	switch e {
	case ContractStatusDraft, ContractStatusPending, ContractStatusActive, ContractStatusExpired, ContractStatusTerminated:
		return true
	}
	return false
}

func (e ContractStatus) String() string {
	return string(e)
}

func (e *ContractStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContractStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContractStatus", str)
	}
	return nil
}

func (e ContractStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DiscountType string

const (
	DiscountTypePercentageOff  DiscountType = "PERCENTAGE_OFF"
	DiscountTypeFixedAmountOff DiscountType = "FIXED_AMOUNT_OFF"
)

var AllDiscountType = []DiscountType{
	DiscountTypePercentageOff,
	DiscountTypeFixedAmountOff,
}

func (e DiscountType) IsValid() bool {
	switch e {
	case DiscountTypePercentageOff, DiscountTypeFixedAmountOff:
		return true
	}
	return false
}

func (e DiscountType) String() string {
	return string(e)
}

func (e *DiscountType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscountType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscountType", str)
	}
	return nil
}

func (e DiscountType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EmailType string

const (
	EmailTypePersonal EmailType = "PERSONAL"
	EmailTypeWork     EmailType = "WORK"
	EmailTypeOther    EmailType = "OTHER"
)

var AllEmailType = []EmailType{
	EmailTypePersonal,
	EmailTypeWork,
	EmailTypeOther,
}

func (e EmailType) IsValid() bool {
	switch e {
	case EmailTypePersonal, EmailTypeWork, EmailTypeOther:
		return true
	}
	return false
}

func (e EmailType) String() string {
	return string(e)
}

func (e *EmailType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmailType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmailType", str)
	}
	return nil
}

func (e EmailType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderOther,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOther:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PhoneType string

const (
	PhoneTypeHome   PhoneType = "HOME"
	PhoneTypeMobile PhoneType = "MOBILE"
	PhoneTypeWork   PhoneType = "WORK"
	PhoneTypeOther  PhoneType = "OTHER"
)

var AllPhoneType = []PhoneType{
	PhoneTypeHome,
	PhoneTypeMobile,
	PhoneTypeWork,
	PhoneTypeOther,
}

func (e PhoneType) IsValid() bool {
	switch e {
	case PhoneTypeHome, PhoneTypeMobile, PhoneTypeWork, PhoneTypeOther:
		return true
	}
	return false
}

func (e PhoneType) String() string {
	return string(e)
}

func (e *PhoneType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhoneType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PhoneType", str)
	}
	return nil
}

func (e PhoneType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PointType string

const (
	PointTypeEarn       PointType = "EARN"
	PointTypeRedeem     PointType = "REDEEM"
	PointTypeTransfer   PointType = "TRANSFER"
	PointTypeAdjustment PointType = "ADJUSTMENT"
)

var AllPointType = []PointType{
	PointTypeEarn,
	PointTypeRedeem,
	PointTypeTransfer,
	PointTypeAdjustment,
}

func (e PointType) IsValid() bool {
	switch e {
	case PointTypeEarn, PointTypeRedeem, PointTypeTransfer, PointTypeAdjustment:
		return true
	}
	return false
}

func (e PointType) String() string {
	return string(e)
}

func (e *PointType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PointType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PointType", str)
	}
	return nil
}

func (e PointType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProgramType string

const (
	ProgramTypePurchase     ProgramType = "PURCHASE"
	ProgramTypeReferral     ProgramType = "REFERRAL"
	ProgramTypeSignup       ProgramType = "SIGNUP"
	ProgramTypeReview       ProgramType = "REVIEW"
	ProgramTypeBirthday     ProgramType = "BIRTHDAY"
	ProgramTypeSubscription ProgramType = "SUBSCRIPTION"
)

var AllProgramType = []ProgramType{
	ProgramTypePurchase,
	ProgramTypeReferral,
	ProgramTypeSignup,
	ProgramTypeReview,
	ProgramTypeBirthday,
	ProgramTypeSubscription,
}

func (e ProgramType) IsValid() bool {
	switch e {
	case ProgramTypePurchase, ProgramTypeReferral, ProgramTypeSignup, ProgramTypeReview, ProgramTypeBirthday, ProgramTypeSubscription:
		return true
	}
	return false
}

func (e ProgramType) String() string {
	return string(e)
}

func (e *ProgramType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProgramType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProgramType", str)
	}
	return nil
}

func (e ProgramType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionMethod string

const (
	TransactionMethodCreditCard TransactionMethod = "CREDIT_CARD"
	TransactionMethodDebitCard  TransactionMethod = "DEBIT_CARD"
	TransactionMethodPaypal     TransactionMethod = "PAYPAL"
	TransactionMethodStripe     TransactionMethod = "STRIPE"
)

var AllTransactionMethod = []TransactionMethod{
	TransactionMethodCreditCard,
	TransactionMethodDebitCard,
	TransactionMethodPaypal,
	TransactionMethodStripe,
}

func (e TransactionMethod) IsValid() bool {
	switch e {
	case TransactionMethodCreditCard, TransactionMethodDebitCard, TransactionMethodPaypal, TransactionMethodStripe:
		return true
	}
	return false
}

func (e TransactionMethod) String() string {
	return string(e)
}

func (e *TransactionMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionMethod", str)
	}
	return nil
}

func (e TransactionMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
	TransactionStatusRefunded  TransactionStatus = "REFUNDED"
)

var AllTransactionStatus = []TransactionStatus{
	TransactionStatusPending,
	TransactionStatusCompleted,
	TransactionStatusFailed,
	TransactionStatusRefunded,
}

func (e TransactionStatus) IsValid() bool {
	switch e {
	case TransactionStatusPending, TransactionStatusCompleted, TransactionStatusFailed, TransactionStatusRefunded:
		return true
	}
	return false
}

func (e TransactionStatus) String() string {
	return string(e)
}

func (e *TransactionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionStatus", str)
	}
	return nil
}

func (e TransactionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WishlistStatus string

const (
	WishlistStatusActive   WishlistStatus = "ACTIVE"
	WishlistStatusArchived WishlistStatus = "ARCHIVED"
	WishlistStatusDeleted  WishlistStatus = "DELETED"
)

var AllWishlistStatus = []WishlistStatus{
	WishlistStatusActive,
	WishlistStatusArchived,
	WishlistStatusDeleted,
}

func (e WishlistStatus) IsValid() bool {
	switch e {
	case WishlistStatusActive, WishlistStatusArchived, WishlistStatusDeleted:
		return true
	}
	return false
}

func (e WishlistStatus) String() string {
	return string(e)
}

func (e *WishlistStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WishlistStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WishlistStatus", str)
	}
	return nil
}

func (e WishlistStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
