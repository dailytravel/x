// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Address struct {
	Street  *string `json:"street,omitempty"`
	City    *string `json:"city,omitempty"`
	State   *string `json:"state,omitempty"`
	Zip     *string `json:"zip,omitempty"`
	Country *string `json:"country,omitempty"`
}

type Balances struct {
	Count int        `json:"count"`
	Data  []*Balance `json:"data,omitempty"`
}

type Benefits struct {
	Data  []*Benefit `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Companies struct {
	Data  []*Company `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Contacts struct {
	Count int        `json:"count"`
	Data  []*Contact `json:"data,omitempty"`
}

type Coupons struct {
	Data  []*Coupon `json:"data,omitempty"`
	Count int       `json:"count"`
}

type Items struct {
	Items []*Item `json:"items,omitempty"`
	Count int     `json:"count"`
}

type Memberships struct {
	Data  []*Membership `json:"data,omitempty"`
	Count int           `json:"count"`
}

type NewBalance struct {
	UID      string                 `json:"uid"`
	Type     string                 `json:"type"`
	Credits  int                    `json:"credits"`
	Notes    *string                `json:"notes,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   string                 `json:"status"`
}

type NewBenefit struct {
	Locale      string                 `json:"locale"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewCompany struct {
	UID         *string                `json:"uid,omitempty"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Industry    *string                `json:"industry,omitempty"`
	Employees   *int                   `json:"employees,omitempty"`
	Revenue     *float64               `json:"revenue,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Address     map[string]interface{} `json:"address,omitempty"`
	Email       *string                `json:"email,omitempty"`
	Phone       *string                `json:"phone,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
}

type NewContact struct {
	FirstName    *string                `json:"first_name,omitempty"`
	LastName     *string                `json:"last_name,omitempty"`
	Email        *string                `json:"email,omitempty"`
	Phone        *string                `json:"phone,omitempty"`
	Picture      *string                `json:"picture,omitempty"`
	Address      map[string]interface{} `json:"address,omitempty"`
	Birthday     *string                `json:"birthday,omitempty"`
	Company      *string                `json:"company,omitempty"`
	JobTitle     *string                `json:"job_title,omitempty"`
	Timezone     *string                `json:"timezone,omitempty"`
	Language     *string                `json:"language,omitempty"`
	Source       *string                `json:"source,omitempty"`
	Subscribed   *bool                  `json:"subscribed,omitempty"`
	Rating       *int                   `json:"rating,omitempty"`
	Notes        *string                `json:"notes,omitempty"`
	Status       *string                `json:"status,omitempty"`
	Labels       []*string              `json:"labels,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	LastActivity *string                `json:"last_activity,omitempty"`
	Created      string                 `json:"created"`
	Updated      string                 `json:"updated"`
	UID          *string                `json:"uid,omitempty"`
}

type NewCoupon struct {
	Locale      string                 `json:"locale"`
	Code        string                 `json:"code"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Amount      float64                `json:"amount"`
	MaxDiscount *float64               `json:"max_discount,omitempty"`
	MinPurchase *float64               `json:"min_purchase,omitempty"`
	Currency    string                 `json:"currency"`
	MaxUses     *int                   `json:"max_uses,omitempty"`
	Uses        *int                   `json:"uses,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Starts      *string                `json:"starts,omitempty"`
	Expires     *string                `json:"expires,omitempty"`
	UID         *string                `json:"uid,omitempty"`
	Products    []*string              `json:"products,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type NewItem struct {
	Package     string                 `json:"package"`
	Locale      string                 `json:"locale"`
	Type        string                 `json:"type"`
	Code        string                 `json:"code"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Quantity    int                    `json:"quantity"`
	Price       float64                `json:"price"`
	Discount    *float64               `json:"discount,omitempty"`
	Currency    string                 `json:"currency"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type NewMembership struct {
	Tier     string                 `json:"tier"`
	Number   string                 `json:"number"`
	Since    string                 `json:"since"`
	Until    string                 `json:"until"`
	Billing  map[string]interface{} `json:"billing,omitempty"`
	Payment  map[string]interface{} `json:"payment,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   string                 `json:"status"`
}

type NewOrder struct {
	Code        string                 `json:"code"`
	Cancellable bool                   `json:"cancellable"`
	Payment     string                 `json:"payment"`
	Coupon      *string                `json:"coupon,omitempty"`
	UID         *string                `json:"uid,omitempty"`
	Contact     *string                `json:"contact,omitempty"`
	Items       []string               `json:"items"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Shares      []*string              `json:"shares,omitempty"`
}

type NewPackage struct {
	Locale       string                 `json:"locale"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Includes     string                 `json:"includes"`
	Excludes     string                 `json:"excludes"`
	Redeem       string                 `json:"redeem"`
	Cancellation string                 `json:"cancellation"`
	Instant      bool                   `json:"instant"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Status       string                 `json:"status"`
	Product      string                 `json:"product"`
}

type NewProduct struct {
	UID         *string                `json:"uid,omitempty"`
	Locale      string                 `json:"locale"`
	Type        string                 `json:"type"`
	Slug        string                 `json:"slug"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Duration    string                 `json:"duration"`
	Reviewable  bool                   `json:"reviewable"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
	Terms       []*string              `json:"terms,omitempty"`
	Place       *string                `json:"place,omitempty"`
	Places      []*string              `json:"places,omitempty"`
}

type NewPromotion struct {
	Type        string                 `json:"type"`
	Locale      string                 `json:"locale"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Credits     int                    `json:"credits"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewQuote struct {
	UID         string                 `json:"uid"`
	Contact     *string                `json:"contact,omitempty"`
	Locale      string                 `json:"locale"`
	Code        string                 `json:"code"`
	Purchase    string                 `json:"purchase"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Terms       *string                `json:"terms,omitempty"`
	Notes       *string                `json:"notes,omitempty"`
	Template    string                 `json:"template"`
	ValidUntil  int                    `json:"validUntil"`
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	Billing     map[string]interface{} `json:"billing"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Shares      []*string              `json:"shares,omitempty"`
}

type NewReward struct {
	Locale      string                 `json:"locale"`
	Tier        string                 `json:"tier"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Cost        int                    `json:"cost"`
	Expires     *string                `json:"expires,omitempty"`
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

type NewVoucher struct {
	Locale      string                 `json:"locale"`
	Code        string                 `json:"code"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Start       string                 `json:"start"`
	End         string                 `json:"end"`
	Price       float64                `json:"price"`
	Discount    *float64               `json:"discount,omitempty"`
	Currency    string                 `json:"currency"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
	Package     string                 `json:"package"`
	Supplier    string                 `json:"supplier"`
}

type NewWishlist struct {
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Orders struct {
	Data  []*Order `json:"data,omitempty"`
	Count int      `json:"count"`
}

type Packages struct {
	Total int        `json:"total"`
	Data  []*Package `json:"data,omitempty"`
}

type Products struct {
	Count int        `json:"count"`
	Data  []*Product `json:"data,omitempty"`
}

type Promotions struct {
	Data  []*Promotion `json:"data,omitempty"`
	Count int          `json:"count"`
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

type UpdateBalance struct {
	Type     *string                `json:"type,omitempty"`
	Credits  *int                   `json:"credits,omitempty"`
	Notes    *string                `json:"notes,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type UpdateBenefit struct {
	Locale      *string                `json:"locale,omitempty"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateCompany struct {
	UID         *string                `json:"uid,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Industry    *string                `json:"industry,omitempty"`
	Employees   *int                   `json:"employees,omitempty"`
	Revenue     *float64               `json:"revenue,omitempty"`
	Address     map[string]interface{} `json:"address,omitempty"`
	Email       *string                `json:"email,omitempty"`
	Phone       *string                `json:"phone,omitempty"`
	Website     *string                `json:"website,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type UpdateContact struct {
	FirstName    *string                `json:"first_name,omitempty"`
	LastName     *string                `json:"last_name,omitempty"`
	Email        *string                `json:"email,omitempty"`
	Phone        *string                `json:"phone,omitempty"`
	Picture      *string                `json:"picture,omitempty"`
	Address      map[string]interface{} `json:"address,omitempty"`
	Birthday     *string                `json:"birthday,omitempty"`
	Company      *string                `json:"company,omitempty"`
	JobTitle     *string                `json:"job_title,omitempty"`
	Timezone     *string                `json:"timezone,omitempty"`
	Language     *string                `json:"language,omitempty"`
	Source       *string                `json:"source,omitempty"`
	Subscribed   *bool                  `json:"subscribed,omitempty"`
	Rating       *int                   `json:"rating,omitempty"`
	Notes        *string                `json:"notes,omitempty"`
	Status       *string                `json:"status,omitempty"`
	Labels       []*string              `json:"labels,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	LastActivity *string                `json:"last_activity,omitempty"`
	Created      string                 `json:"created"`
	Updated      string                 `json:"updated"`
	UID          string                 `json:"uid"`
}

type UpdateCoupon struct {
	Locale      *string                `json:"locale,omitempty"`
	Code        *string                `json:"code,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Amount      *float64               `json:"amount,omitempty"`
	MaxDiscount *float64               `json:"max_discount,omitempty"`
	MinPurchase *float64               `json:"min_purchase,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	MaxUses     *int                   `json:"max_uses,omitempty"`
	Uses        *int                   `json:"uses,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Starts      *string                `json:"starts,omitempty"`
	Expires     *string                `json:"expires,omitempty"`
	UID         *string                `json:"uid,omitempty"`
	Products    []*string              `json:"products,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type UpdateItem struct {
	Package     *string                `json:"package,omitempty"`
	Locale      *string                `json:"locale,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Code        *string                `json:"code,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Quantity    *int                   `json:"quantity,omitempty"`
	Price       *float64               `json:"price,omitempty"`
	Discount    *float64               `json:"discount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type UpdateMembership struct {
	Tier     *string                `json:"tier,omitempty"`
	Number   *string                `json:"number,omitempty"`
	Since    *string                `json:"since,omitempty"`
	Until    *string                `json:"until,omitempty"`
	Billing  map[string]interface{} `json:"billing,omitempty"`
	Payment  map[string]interface{} `json:"payment,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type UpdateOrder struct {
	Code        *string                `json:"code,omitempty"`
	Cancellable *bool                  `json:"cancellable,omitempty"`
	Payment     *string                `json:"payment,omitempty"`
	Coupon      *string                `json:"coupon,omitempty"`
	UID         *string                `json:"uid,omitempty"`
	Contact     *string                `json:"contact,omitempty"`
	Items       []string               `json:"items,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Shares      []*string              `json:"shares,omitempty"`
}

type UpdatePackage struct {
	Locale       *string                `json:"locale,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Includes     *string                `json:"includes,omitempty"`
	Excludes     *string                `json:"excludes,omitempty"`
	Redeem       *string                `json:"redeem,omitempty"`
	Cancellation *string                `json:"cancellation,omitempty"`
	Instant      *bool                  `json:"instant,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Status       *string                `json:"status,omitempty"`
	Product      *string                `json:"product,omitempty"`
}

type UpdateProduct struct {
	Locale      *string                `json:"locale,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Slug        *string                `json:"slug,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Duration    *string                `json:"duration,omitempty"`
	Reviewable  *bool                  `json:"reviewable,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Terms       []*string              `json:"terms,omitempty"`
	Place       *string                `json:"place,omitempty"`
	Places      []*string              `json:"places,omitempty"`
}

type UpdatePromotion struct {
	Locale      *string                `json:"locale,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Credits     *int                   `json:"credits,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateQuote struct {
	UID         *string                `json:"uid,omitempty"`
	Contact     *string                `json:"contact,omitempty"`
	Locale      *string                `json:"locale,omitempty"`
	Code        *string                `json:"code,omitempty"`
	Purchase    *string                `json:"purchase,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Terms       *string                `json:"terms,omitempty"`
	Notes       *string                `json:"notes,omitempty"`
	Template    *string                `json:"template,omitempty"`
	ValidUntil  *int                   `json:"validUntil,omitempty"`
	Amount      *float64               `json:"amount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Billing     map[string]interface{} `json:"billing,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Shares      []*string              `json:"shares,omitempty"`
}

type UpdateReward struct {
	Locale      string                 `json:"locale"`
	Tier        *string                `json:"tier,omitempty"`
	Cost        *string                `json:"cost,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Credits     *int                   `json:"credits,omitempty"`
	Expires     *string                `json:"expires,omitempty"`
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

type UpdateVoucher struct {
	Locale      *string                `json:"locale,omitempty"`
	Code        *string                `json:"code,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Start       *string                `json:"start,omitempty"`
	End         string                 `json:"end"`
	Price       *float64               `json:"price,omitempty"`
	Discount    *float64               `json:"discount,omitempty"`
	Currency    *string                `json:"currency,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Package     *string                `json:"package,omitempty"`
	Supplier    *string                `json:"supplier,omitempty"`
}

type UpdateWishlist struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type Vouchers struct {
	Count int        `json:"count"`
	Data  []*Voucher `json:"data,omitempty"`
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

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "PERCENTAGE"
	DiscountTypeFlatRate   DiscountType = "FLAT_RATE"
)

var AllDiscountType = []DiscountType{
	DiscountTypePercentage,
	DiscountTypeFlatRate,
}

func (e DiscountType) IsValid() bool {
	switch e {
	case DiscountTypePercentage, DiscountTypeFlatRate:
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

type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
	ProductStatusPending  ProductStatus = "PENDING"
	ProductStatusArchived ProductStatus = "ARCHIVED"
)

var AllProductStatus = []ProductStatus{
	ProductStatusActive,
	ProductStatusInactive,
	ProductStatusPending,
	ProductStatusArchived,
}

func (e ProductStatus) IsValid() bool {
	switch e {
	case ProductStatusActive, ProductStatusInactive, ProductStatusPending, ProductStatusArchived:
		return true
	}
	return false
}

func (e ProductStatus) String() string {
	return string(e)
}

func (e *ProductStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductStatus", str)
	}
	return nil
}

func (e ProductStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PromotionType string

const (
	PromotionTypePurchase     PromotionType = "PURCHASE"
	PromotionTypeReferral     PromotionType = "REFERRAL"
	PromotionTypeSignup       PromotionType = "SIGNUP"
	PromotionTypeReview       PromotionType = "REVIEW"
	PromotionTypeBirthday     PromotionType = "BIRTHDAY"
	PromotionTypeSubscription PromotionType = "SUBSCRIPTION"
)

var AllPromotionType = []PromotionType{
	PromotionTypePurchase,
	PromotionTypeReferral,
	PromotionTypeSignup,
	PromotionTypeReview,
	PromotionTypeBirthday,
	PromotionTypeSubscription,
}

func (e PromotionType) IsValid() bool {
	switch e {
	case PromotionTypePurchase, PromotionTypeReferral, PromotionTypeSignup, PromotionTypeReview, PromotionTypeBirthday, PromotionTypeSubscription:
		return true
	}
	return false
}

func (e PromotionType) String() string {
	return string(e)
}

func (e *PromotionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PromotionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PromotionType", str)
	}
	return nil
}

func (e PromotionType) MarshalGQL(w io.Writer) {
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
