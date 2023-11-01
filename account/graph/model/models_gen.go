// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Apis struct {
	Data  []*Api `json:"data,omitempty"`
	Count int    `json:"count"`
}

type AuthPayload struct {
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	ExpiresIn    *int    `json:"expires_in,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
	User         *User   `json:"user,omitempty"`
}

type Clients struct {
	Data  []*Client `json:"data,omitempty"`
	Count int       `json:"count"`
}

type Connections struct {
	Data  []*Connection `json:"data,omitempty"`
	Count int           `json:"count"`
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

type LoginInput struct {
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	GrantType    *string `json:"grant_type,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
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

type NewCredential struct {
	Type     string                 `json:"type"`
	Secret   string                 `json:"secret"`
	Expires  string                 `json:"expires"`
	Status   string                 `json:"status"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type NewIdentity struct {
	UserID     string                 `json:"user_id"`
	Provider   string                 `json:"provider"`
	Connection string                 `json:"connection"`
	IsSocial   bool                   `json:"is_social"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Status     string                 `json:"status"`
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
	Roles    []*string              `json:"roles,omitempty"`
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

type NewUserInput struct {
	Name       string    `json:"name"`
	GivenName  *string   `json:"given_name,omitempty"`
	FamilyName *string   `json:"family_name,omitempty"`
	Email      string    `json:"email"`
	Phone      *string   `json:"phone,omitempty"`
	Password   string    `json:"password"`
	Roles      []*string `json:"roles,omitempty"`
	Timezone   *string   `json:"timezone,omitempty"`
	Locale     *string   `json:"locale,omitempty"`
	Picture    *string   `json:"picture,omitempty"`
	Status     *string   `json:"status,omitempty"`
}

type Order struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Order) IsEntity() {}

type Organization struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Organization) IsEntity() {}

type Payload struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	User         *User  `json:"user"`
}

type Permissions struct {
	Data  []*Permission `json:"data,omitempty"`
	Count int           `json:"count"`
}

type Quote struct {
	UID  string `json:"uid"`
	User *User  `json:"user,omitempty"`
}

func (Quote) IsEntity() {}

type RegisterInput struct {
	Name                 string  `json:"name"`
	Email                string  `json:"email"`
	Password             string  `json:"password"`
	PasswordConfirmation string  `json:"password_confirmation"`
	ClientID             *string `json:"client_id,omitempty"`
}

type ResetPasswordInput struct {
	Token                string `json:"token"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type Roles struct {
	Data  []*Role `json:"data,omitempty"`
	Count int     `json:"count"`
}

type SocialLoginInput struct {
	ClientID     string  `json:"client_id"`
	ClientSecret *string `json:"client_secret,omitempty"`
	Provider     string  `json:"provider"`
	Token        string  `json:"token"`
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

type UpdateCredential struct {
	Type     *string                `json:"type,omitempty"`
	Secret   *string                `json:"secret,omitempty"`
	Expires  *string                `json:"expires,omitempty"`
	Revoked  *bool                  `json:"revoked,omitempty"`
	Status   *string                `json:"status,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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

type UpdatePasswordInput struct {
	OldPassword          string  `json:"old_password"`
	Password             string  `json:"password"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty"`
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

type UpdateUserInput struct {
	Name       *string                `json:"name,omitempty"`
	GivenName  *string                `json:"given_name,omitempty"`
	FamilyName *string                `json:"family_name,omitempty"`
	Password   *string                `json:"password,omitempty"`
	Roles      []*string              `json:"roles,omitempty"`
	Timezone   *string                `json:"timezone,omitempty"`
	Locale     *string                `json:"locale,omitempty"`
	Picture    *string                `json:"picture,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type Users struct {
	Count int     `json:"count"`
	Data  []*User `json:"data,omitempty"`
}

type VerifyInput struct {
	Code string `json:"code"`
	Type string `json:"type"`
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
