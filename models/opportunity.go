package models

import (
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Opportunity
 */
type Opportunity struct {
	ID                     uint                  `gorm:"primaryKey;uniqueIndex:udx_Opportunities" openapi:"example:1;type:integer;"`
	Title                  string                `json:"title" gorm:"size:2000;" openapi:"example:Opportunity Title;type:string;"`
	Description            string                `json:"description" openapi:"example:Opportunity Description;type:string;"`
	UserKnowsContractValue bool                  `json:"user_knows_contract_value" openapi:"example:true;type:boolean;"`
	UserContractValue      string                `json:"user_contract_value" openapi:"example:Contract Value;type:string;"`
	CrawlerContractValue   string                `json:"crawler_contract_value" openapi:"example:Contract Value;type:string;"`
	SolicitationNumber     string                `json:"solicitation_number" openapi:"example:Solicitation Number;type:string;"`
	MultiAward             bool                  `json:"multi_award" openapi:"example:true;type:boolean;"`
	User                   *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID                 int                   `json:"user_id" openapi:"example:1;type:integer;"`
	Staff                  *User                 `json:"staff" gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	StaffID                dtp.NullInt64         `json:"staff_id" openapi:"example:1;type:integer;"`
	ApprovedAt             dtp.NullTime          `json:"approved_at" openapi:"example:2022-01-01T00:00:00Z;type:datetime;"`
	Requested              bool                  `json:"requested" openapi:"example:false;type:boolean;"`
	Duplicated             bool                  `json:"duplicated" openapi:"example:false;type:boolean;"`
	Deprecated             bool                  `json:"deprecated" openapi:"example:false;type:boolean;"`
	NeedToReview           bool                  `json:"need_to_review" openapi:"example:false;type:boolean;"`
	Market                 string                `json:"market" openapi:"example:Market Name;type:string;"`
	Department             string                `json:"department" openapi:"example:Department Name;type:string;"`
	Agency                 string                `json:"agency" openapi:"example:Agency Name;type:string;"`
	Office                 string                `json:"office" openapi:"example:Office Name;type:string;"`
	Naics                  string                `json:"naics" openapi:"example:Naics Name;type:string;"`
	FiscalYear             string                `json:"fiscal_year" openapi:"example:2022;type:string;"`
	SetAside               string                `json:"set_aside" openapi:"example:Set Aside Name;type:string;"`
	ContractVehicle        string                `json:"contract_vehicle" openapi:"example:Contract Vehicle Name;type:string;"`
	IsDraft                bool                  `json:"is_draft" openapi:"example:false;type:boolean;"`
	Orders                 []*Order              `json:"orders" openapi:"$ref:Order;type:array;"`
	Documents              []*Document           `json:"documents" gorm:"polymorphic:Owner;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:Document;type:array;"`
	CreatedAt              time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt              time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt              soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_Opportunities"`
}
