// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddressInput struct {
	Street  *string `json:"street,omitempty"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Zip     string  `json:"zip"`
	Country string  `json:"country"`
}

type Apis struct {
	Data  []*Api `json:"data,omitempty"`
	Count int    `json:"count"`
}

type Clients struct {
	Data  []*Client `json:"data,omitempty"`
	Count int       `json:"count"`
}

type Connections struct {
	Data  []*Connection `json:"data,omitempty"`
	Count int           `json:"count"`
}

type Identity struct {
	ID          string                 `json:"id"`
	User        *User                  `json:"user"`
	Provider    string                 `json:"provider"`
	AccessToken string                 `json:"access_token"`
	ExpiresIn   int                    `json:"expires_in"`
	Connection  string                 `json:"connection"`
	IsSocial    bool                   `json:"is_social"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type Integrations struct {
	Data  []*Integration `json:"data,omitempty"`
	Count int            `json:"count"`
}

type Invitations struct {
	Data  []*Invitation `json:"data,omitempty"`
	Count int           `json:"count"`
}

type Keys struct {
	Data  []*Key `json:"data,omitempty"`
	Count int    `json:"count"`
}

type Mfa struct {
	Enabled bool   `json:"enabled"`
	Code    string `json:"code"`
}

type MFAInput struct {
	Enabled bool   `json:"enabled"`
	Code    string `json:"code"`
}

type Membership struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Membership) IsEntity() {}

type NewAPI struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Identifier  string                 `json:"identifier"`
	Algorithm   Algorithm              `json:"algorithm"`
	Expiration  int                    `json:"expiration"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewClient struct {
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Domain      *string                `json:"domain,omitempty"`
	Redirect    *string                `json:"redirect,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewConnection struct {
	Client      string                 `json:"client"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewIntegration struct {
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewInvitation struct {
	Email    string                 `json:"email"`
	Roles    []string               `json:"roles"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type NewKey struct {
	Name string `json:"name"`
}

type NewPermission struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewRole struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Permissions []*string `json:"permissions,omitempty"`
}

type NewToken struct {
	Name      string    `json:"name"`
	Abilities []*string `json:"abilities,omitempty"`
}

type Order struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Order) IsEntity() {}

type Organization struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Created   *User  `json:"created,omitempty"`
	Updated   *User  `json:"updated,omitempty"`
}

func (Organization) IsEntity() {}

type PasswordInput struct {
	CurrentPassword      string `json:"currentPassword"`
	NewPassword          string `json:"newPassword"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type Payload struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
}

type Permissions struct {
	Data  []*Permission `json:"data,omitempty"`
	Count int           `json:"count"`
}

type Point struct {
	ID string `json:"id"`
}

func (Point) IsEntity() {}

type ProfileInput struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	JobTitle  *string `json:"jobTitle,omitempty"`
	Birthday  *string `json:"birthday,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	Bio       *string `json:"bio,omitempty"`
	Company   *string `json:"company,omitempty"`
	Website   *string `json:"website,omitempty"`
}

type Quote struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Quote) IsEntity() {}

type Roles struct {
	Data  []*Role `json:"data,omitempty"`
	Count int     `json:"count"`
}

type Tokens struct {
	Data  []*Token `json:"data,omitempty"`
	Count int      `json:"count"`
}

type UpdateAPI struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Identifier  *string                `json:"identifier,omitempty"`
	Algorithm   *Algorithm             `json:"algorithm,omitempty"`
	Expiration  *int                   `json:"expiration,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateClient struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Domain      *string                `json:"domain,omitempty"`
	Redirect    *string                `json:"redirect,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateConnection struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateIntegration struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateInvitation struct {
	Roles  []string `json:"roles"`
	Status *string  `json:"status,omitempty"`
}

type UpdatePermission struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateRole struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Permissions []*string `json:"permissions,omitempty"`
}

type UpdateToken struct {
	Name      *string   `json:"name,omitempty"`
	Abilities []*string `json:"abilities,omitempty"`
}

type UserInput struct {
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Phone    *string       `json:"phone,omitempty"`
	Password string        `json:"password"`
	Roles    []*string     `json:"roles,omitempty"`
	Timezone *string       `json:"timezone,omitempty"`
	Locale   *string       `json:"locale,omitempty"`
	Picture  *string       `json:"picture,omitempty"`
	Profile  *ProfileInput `json:"profile,omitempty"`
	Address  *AddressInput `json:"address,omitempty"`
	Status   *string       `json:"status,omitempty"`
}

type Users struct {
	Count int     `json:"count"`
	Data  []*User `json:"data,omitempty"`
}

type Wishlist struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Wishlist) IsEntity() {}

type Algorithm string

const (
	AlgorithmHs256 Algorithm = "HS256"
	AlgorithmHs384 Algorithm = "HS384"
	AlgorithmHs512 Algorithm = "HS512"
	AlgorithmRs256 Algorithm = "RS256"
	AlgorithmRs384 Algorithm = "RS384"
	AlgorithmRs512 Algorithm = "RS512"
	AlgorithmEs256 Algorithm = "ES256"
	AlgorithmEs384 Algorithm = "ES384"
	AlgorithmEs512 Algorithm = "ES512"
	AlgorithmPs256 Algorithm = "PS256"
	AlgorithmPs384 Algorithm = "PS384"
	AlgorithmPs512 Algorithm = "PS512"
)

var AllAlgorithm = []Algorithm{
	AlgorithmHs256,
	AlgorithmHs384,
	AlgorithmHs512,
	AlgorithmRs256,
	AlgorithmRs384,
	AlgorithmRs512,
	AlgorithmEs256,
	AlgorithmEs384,
	AlgorithmEs512,
	AlgorithmPs256,
	AlgorithmPs384,
	AlgorithmPs512,
}

func (e Algorithm) IsValid() bool {
	switch e {
	case AlgorithmHs256, AlgorithmHs384, AlgorithmHs512, AlgorithmRs256, AlgorithmRs384, AlgorithmRs512, AlgorithmEs256, AlgorithmEs384, AlgorithmEs512, AlgorithmPs256, AlgorithmPs384, AlgorithmPs512:
		return true
	}
	return false
}

func (e Algorithm) String() string {
	return string(e)
}

func (e *Algorithm) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Algorithm(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Algorithm", str)
	}
	return nil
}

func (e Algorithm) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InvitationStatus string

const (
	InvitationStatusPending  InvitationStatus = "PENDING"
	InvitationStatusAccepted InvitationStatus = "ACCEPTED"
	InvitationStatusDeclined InvitationStatus = "DECLINED"
)

var AllInvitationStatus = []InvitationStatus{
	InvitationStatusPending,
	InvitationStatusAccepted,
	InvitationStatusDeclined,
}

func (e InvitationStatus) IsValid() bool {
	switch e {
	case InvitationStatusPending, InvitationStatusAccepted, InvitationStatusDeclined:
		return true
	}
	return false
}

func (e InvitationStatus) String() string {
	return string(e)
}

func (e *InvitationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InvitationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InvitationStatus", str)
	}
	return nil
}

func (e InvitationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SocialProvider string

const (
	SocialProviderFacebook SocialProvider = "FACEBOOK"
	SocialProviderGoogle   SocialProvider = "GOOGLE"
	SocialProviderTwitter  SocialProvider = "TWITTER"
	SocialProviderGithub   SocialProvider = "GITHUB"
)

var AllSocialProvider = []SocialProvider{
	SocialProviderFacebook,
	SocialProviderGoogle,
	SocialProviderTwitter,
	SocialProviderGithub,
}

func (e SocialProvider) IsValid() bool {
	switch e {
	case SocialProviderFacebook, SocialProviderGoogle, SocialProviderTwitter, SocialProviderGithub:
		return true
	}
	return false
}

func (e SocialProvider) String() string {
	return string(e)
}

func (e *SocialProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SocialProvider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SocialProvider", str)
	}
	return nil
}

func (e SocialProvider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusInactive  UserStatus = "INACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
	UserStatusPending   UserStatus = "PENDING"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusInactive,
	UserStatusSuspended,
	UserStatusPending,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusInactive, UserStatusSuspended, UserStatusPending:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
