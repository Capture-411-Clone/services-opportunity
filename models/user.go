package models

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/plugin/soft_delete"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"golang.org/x/crypto/bcrypt"
)

/*
 * @apiDefine: User
 */
type User struct {
	ID                   uint                  `gorm:"primaryKey;uniqueIndex:udx_users" openapi:"example:1;type:integer;"`
	Title                string                `json:"title" gorm:"size:100" openapi:"example:User Title"`
	Name                 string                `json:"name" gorm:"size:100" openapi:"example:John Doe"`
	Email                dtp.NullString        `json:"email" gorm:"size:256;index;unique" openapi:"example:YnV0c@example.com"`
	Phone                dtp.NullString        `json:"phone" gorm:"size:20;index;unique" openapi:"example:1234567890"`
	IDCode               dtp.NullString        `json:"id_code" gorm:"size:100;index;unique;" openapi:"example:1234567890"`
	Password             string                `json:"password" gorm:"size:256" openapi:"example:7c222fb2927d828af22f592134e8932480637c0d"`
	LastName             string                `json:"last_name" gorm:"size:100" openapi:"example:Doe"`
	Username             dtp.NullString        `json:"username" gorm:"size:100;index;unique;" openapi:"example:john"`
	DateOfBirth          dtp.NullTime          `json:"date_of_birth" openapi:"example:1980-01-01"`
	Gender               string                `json:"gender" openapi:"example:Male"`
	Address              string                `json:"address" openapi:"example:Ave;type:string;"`
	TownCity             string                `json:"town_city" openapi:"example:New York City;type:string;"`
	State                string                `json:"state" openapi:"example:NY;type:string;"`
	ZipCode              string                `json:"zip_code" openapi:"example:123875;type:string;"`
	Country              string                `json:"country" openapi:"example:USA;type:string;"`
	CompanyName          string                `json:"company_name" openapi:"example:Beta;type:string;"`
	SuspendedAt          dtp.NullTime          `json:"suspended_at" openapi:"example:2021-01-01"`
	PhoneVerifiedAt      dtp.NullTime          `json:"phone_verified_at" openapi:"example:2021-01-01"`
	EmailVerifiedAt      dtp.NullTime          `json:"email_verified_at" openapi:"example:2021-01-01"`
	ProfileCompletedAt   dtp.NullTime          `json:"profile_completed_at" openapi:"example:2021-01-01"` // In update and create and updateProfile a user check if the data is enough set the time to make it true.
	PolicyApprovedAt     dtp.NullTime          `json:"policy_approved_at"`
	Invite               *Invite               `json:"invite"`
	Roles                []*Role               `json:"roles" gorm:"many2many:user_role;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"example:[1,2,3];$ref:Role"`
	ContributorID        dtp.NullInt64         `json:"contributor_id" openapi:"example:1;type:integer;"`
	Contributor          *User                 `json:"contributor" gorm:"foreignKey:ContributorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	ReferralCode         string                `json:"referral_code" gorm:"size:100" openapi:"example:BOSSLADY"`
	ReferredWithCode     string                `json:"referred_with_code" gorm:"size:100" openapi:"example:BOSSLADY"`
	ShouldChangePassword bool                  `json:"should_change_password" gorm:"default:false" openapi:"example:false"`
	StripeCustomerID     string                `json:"stripe_customer_id" gorm:"size:100" openapi:"example:cus_1234567890"`
	SubscriptionID       string                `json:"subscription_id" gorm:"size:100" openapi:"example:sub_1234567890"`
	Subscription         []byte                `json:"subscription" gorm:"type:jsonb"`
	Credits              int                   `json:"credits" gorm:"default:0" openapi:"example:0"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
	DeletedAt            soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_users"`
}

/*
 * @apiDefine: UserData
 */
type UserData struct {
	ID                 uint                  `gorm:"primaryKey;uniqueIndex:udx_users" openapi:"example:1;type:integer;"`
	Title              string                `json:"title" gorm:"size:100" openapi:"example:User Title"`
	Name               string                `json:"name" gorm:"size:100" openapi:"example:John Doe"`
	Email              dtp.NullString        `json:"email" gorm:"size:256;index;unique" openapi:"example:YnV0c@example.com"`
	Phone              dtp.NullString        `json:"phone" gorm:"size:20;index;unique" openapi:"example:1234567890"`
	IDCode             dtp.NullString        `json:"id_code" gorm:"size:100;index;unique;" openapi:"example:1234567890"`
	LastName           string                `json:"last_name" gorm:"size:100" openapi:"example:Doe"`
	DateOfBirth        dtp.NullTime          `json:"date_of_birth" openapi:"example:1980-01-01"`
	Gender             string                `json:"gender" openapi:"example:Male"`
	Address            string                `json:"address" openapi:"example:Ave;type:string;"`
	TownCity           string                `json:"town_city" openapi:"example:New York City;type:string;"`
	State              string                `json:"state" openapi:"example:NY;type:string;"`
	ZipCode            string                `json:"zip_code" openapi:"example:123875;type:string;"`
	Country            string                `json:"country" openapi:"example:USA;type:string;"`
	CompanyName        string                `json:"company_name" openapi:"example:Beta;type:string;"`
	SuspendedAt        dtp.NullTime          `json:"suspended_at" openapi:"example:2021-01-01"`
	PhoneVerifiedAt    dtp.NullTime          `json:"phone_verified_at" openapi:"example:2021-01-01"`
	EmailVerifiedAt    dtp.NullTime          `json:"email_verified_at" openapi:"example:2021-01-01"`
	ProfileCompletedAt dtp.NullTime          `json:"profile_completed_at" openapi:"example:2021-01-01"` // In update and create and updateProfile a user check if the data is enough set the time to make it true.
	PolicyApprovedAt   dtp.NullTime          `json:"policy_approved_at"`
	Invite             *Invite               `json:"invite"`
	Roles              []*Role               `json:"roles" gorm:"many2many:user_role;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"example:[1,2,3];$ref:Role"`
	ContributorID      dtp.NullInt64         `json:"contributor_id" openapi:"example:1;type:integer;"`
	Contributor        *User                 `json:"contributor" gorm:"foreignKey:ContributorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	ReferralCode       string                `json:"invite_code" gorm:"size:100" openapi:"example:BOSSLADY"`
	ReferredWithCode   string                `json:"referred_with_code" gorm:"size:100" openapi:"example:BOSSLADY"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
	DeletedAt          soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_users"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u User) StarsAvg() int {
	return 5
}

func (u *User) CheckProfileCompleted() bool {
	return false
}

func (u User) GetID() string {
	return strconv.FormatUint(uint64(u.ID), 10)
}
func (u User) GetFullName() string {
	return strings.TrimSpace(u.Name + " " + u.LastName)
}
func (u User) GetPhone() string {
	return u.Phone.String
}
func (u User) GetRoles() []string {
	rls := []string{}
	for _, v := range u.Roles {
		rls = append(rls, v.Title)
	}
	return rls
}
