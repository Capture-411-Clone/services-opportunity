package endpoints

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"gorm.io/gorm"
)

/*
 * @apiDefine: OpportunityItems
 */
type OpportunityItems struct {
	ID                     int                     `json:"ID" openapi:"example:1;type:integer;"`
	Title                  string                  `json:"title" openapi:"example:Test Opportunity;type:string;"`
	Description            string                  `json:"description" openapi:"example:Test Description;type:string;"`
	UserKnowsContractValue bool                    `json:"user_knows_contract_value" openapi:"example:true;type:boolean;"`
	UserContractValue      string                  `json:"user_contract_value" openapi:"example:1000;type:string;"`
	CrawlerContractValue   string                  `json:"crawler_contract_value" openapi:"example:1000;type:string;"`
	SolicitationNumber     string                  `json:"solicitation_number" openapi:"example:548;type:string;"`
	MultiAward             bool                    `json:"multi_award" openapi:"example:true;type:boolean;"`
	UserID                 int                     `json:"user_id" openapi:"example:1;type:integer;"`
	StaffID                *int                    `json:"staff_id" openapi:"example:1;type:integer;"`
	ApprovedAt             *time.Time              `json:"approved_at" openapi:"example:2022-01-01T00:00:00Z;type:datetime;"`
	Requested              bool                    `json:"requested" openapi:"example:false;type:boolean;"`
	Duplicated             bool                    `json:"duplicated" openapi:"example:false;type:boolean;"`
	Depricated             bool                    `json:"depricated" openapi:"example:false;type:boolean;"`
	NeedToReview           bool                    `json:"need_to_review" openapi:"example:false;type:boolean;"`
	Market                 *models.Market          `json:"market" openapi:"$ref:Market;type:object;"`
	Department             *models.Department      `json:"department" openapi:"$ref:Department;type:object;"`
	Agency                 *models.Agency          `json:"agency" openapi:"$ref:Agency;type:object;"`
	Office                 *models.Office          `json:"office" openapi:"$ref:Office;type:object;"`
	Naics                  *models.Naics           `json:"naics" openapi:"$ref:Naics;type:object;"`
	FiscalYear             *models.FiscalYear      `json:"fiscal_year" openapi:"$ref:FiscalYear;type:object;"`
	SetAside               *models.SetAside        `json:"set_aside" openapi:"$ref:SetAside;type:object;"`
	ContractVehicle        *models.ContractVehicle `json:"contract_vehicle" openapi:"$ref:ContractVehicle;type:object;"`
	IsDraft                bool                    `json:"is_draft" openapi:"example:false;type:boolean;"`
	Documents              []*models.Document      `json:"documents" openapi:"$ref:Document;type:array;"`
	CreatedAt              time.Time               `json:"created_at" openapi:"example:2022-01-01T00:00:00Z;type:datetime;"`
	UpdatedAt              time.Time               `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z;type:datetime;"`
}

/*
 * @apiDefine: Data
 */
type Data struct {
	TotalRows        int                `json:"totalRows" openapi:"example:10;type:integer;"`
	OpportunityItems []OpportunityItems `json:"items" openapi:"$ref:OpportunityItems;type:array;"`
}

// Struct to represent the JSON request for multiple opportunities
/*
 * @apiDefine: ImportOpportunitiesRequest
 */
type ImportOpportunitiesRequest struct {
	Data Data `json:"data" openapi:"required;$ref:Data;type:object"`
}

func (c *ImportOpportunitiesRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"data": gocriteria.New("data").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"data": "Data",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

// Function to import opportunities
func (s *service) ImportOpportunities(ctx context.Context, importReq ImportOpportunitiesRequest) response.ErrorResponse {

	if !policy.CanImportOpportunity(ctx) {
		s.logger.With(ctx).Error("You do not have permission to import to opportunity")
		return response.ErrorForbidden("You do not have permission to import to opportunity")
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Iterate through each opportunity
		if importReq.Data.TotalRows <= 0 {
			return errors.New("no data provided")
		}

		// Map data from request to the new Opportunity model
		for _, oppReq := range importReq.Data.OpportunityItems {

			// Map data from request to the new Opportunity model
			opportunity := models.Opportunity{
				ID:                     uint(oppReq.ID),
				Title:                  oppReq.Title,
				Description:            oppReq.Description,
				UserKnowsContractValue: oppReq.UserKnowsContractValue,
				UserContractValue:      oppReq.UserContractValue,
				CrawlerContractValue:   oppReq.CrawlerContractValue, // Assuming this is populated later by some service
				SolicitationNumber:     oppReq.SolicitationNumber,
				MultiAward:             oppReq.MultiAward,
				UserID:                 oppReq.UserID,
				StaffID:                dtp.NullInt64{Int64: 0, Valid: false}, // Default to invalid
				ApprovedAt:             dtp.NullTime{Time: time.Time{}, Valid: false},
				Requested:              oppReq.Requested,
				Duplicated:             oppReq.Duplicated,
				Deprecated:             oppReq.Depricated,
				NeedToReview:           oppReq.NeedToReview,
				Market:                 exp.TerIf(oppReq.Market != nil, oppReq.Market.Name, ""),
				Department:             exp.TerIf(oppReq.Department != nil, oppReq.Department.Name, ""),
				Agency:                 exp.TerIf(oppReq.Agency != nil, oppReq.Agency.Name, ""),
				Office:                 exp.TerIf(oppReq.Office != nil, oppReq.Office.Name, ""),
				Naics:                  exp.TerIf(oppReq.Naics != nil, oppReq.Naics.Name, ""),
				FiscalYear:             exp.TerIf(oppReq.FiscalYear == nil, "", oppReq.FiscalYear.Year),
				SetAside:               exp.TerIf(oppReq.SetAside == nil, "", oppReq.SetAside.Name),
				ContractVehicle:        exp.TerIf(oppReq.ContractVehicle == nil, "", oppReq.ContractVehicle.Name),
				IsDraft:                oppReq.IsDraft,
				CreatedAt:              oppReq.CreatedAt,
				UpdatedAt:              oppReq.UpdatedAt,
			}

			// Check if StaffID is not nil before dereferencing
			if oppReq.StaffID != nil {
				opportunity.StaffID = dtp.NullInt64{Int64: int64(*oppReq.StaffID), Valid: true}
			}

			// Check if ApprovedAt is not nil before dereferencing
			if oppReq.ApprovedAt != nil {
				opportunity.ApprovedAt = dtp.NullTime{Time: *oppReq.ApprovedAt, Valid: true}
			}

			// Create the opportunity in the database
			if err := tx.Create(&opportunity).Error; err != nil {
				log.Printf("Failed to save opportunity ID %d: %v", opportunity.ID, err)
				return err
			}

			// Get the IDs of the documents
			var documents []models.Document
			var documentIDs []int
			for _, doc := range oppReq.Documents {
				documentIDs = append(documentIDs, int(doc.ID))
			}
			if len(documentIDs) > 0 {
				if err := tx.Where("ID IN ?", documentIDs).Find(&documents).Error; err != nil {
					log.Printf("Failed to find documents for opportunity ID %d: %v", opportunity.ID, err)
					return err
				}

				if err := tx.Model(&opportunity).Association("Documents").Append(&documents); err != nil {
					log.Printf("Failed to associate documents with opportunity ID %d: %v", opportunity.ID, err)
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		s.logger.With(ctx).Error("transaction failed: ", err)
		return response.GormErrorResponse(err, "Failed to import opportunities")
	}

	return response.ErrorResponse{}
}
