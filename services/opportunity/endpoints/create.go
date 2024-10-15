package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/Capture-411/services-opportunity/services/opportunity/collector"
)

/*
 * @apiDefine: CreateOpportunityRequest
 */
type CreateOpportunityRequest struct {
	Title                  string               `json:"title" openapi:"example:Test Opportunity;type:string;"`
	Description            string               `json:"description" openapi:"example:Test Description;type:string;"`
	SolicitationNumber     string               `json:"solicitation_number" openapi:"example:548;type:string"`
	UserKnowsContractValue bool                 `json:"user_knows_contract_value" openapi:"example:true;type:boolean"`
	UserContractValue      string               `json:"user_contract_value" openapi:"example:1000;type:string"`
	Approved               bool                 `json:"approved" openapi:"example:true;type:boolean"`
	Market                 string               `json:"market" openapi:"example:Market Name;type:string;"`
	Department             string               `json:"department" openapi:"example:Department Name;type:string;"`
	Agency                 string               `json:"agency" openapi:"example:Agency Name;type:string;"`
	Office                 string               `json:"office" openapi:"example:Office Name;type:string;"`
	Naics                  string               `json:"naics" openapi:"example:NAICS Name;type:string;"`
	FiscalYear             string               `json:"fiscal_year" openapi:"example:2022;type:string;"`
	SetAside               string               `json:"set_aside" openapi:"example:Set Aside Name;type:string;"`
	ContractVehicle        string               `json:"contract_vehicle" openapi:"example:Contract Vehicle Name;type:string;"`
	IsDraft                bool                 `json:"is_draft" openapi:"example:false;type:boolean"`
	Requested              bool                 `json:"requested" openapi:"example:false;type:boolean"`
	Documents              []AddDocumentRequest `json:"documents" openapi:"$ref:AddDocumentRequest;example:;type:array"`
}

func (c *CreateOpportunityRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"title":                     gocriteria.New("title").Optional(),
		"description":               gocriteria.New("description").Optional(),
		"user_contract_value":       gocriteria.New("user_contract_value").Optional(),
		"user_knows_contract_value": gocriteria.New("user_knows_contract_value").Optional(),
		"solicitation_number":       gocriteria.New("solicitation_number").Optional(),
		"approved":                  gocriteria.New("approved").Optional(),
		"market":                    gocriteria.New("market").Optional(),
		"department":                gocriteria.New("department").Optional(),
		"agency":                    gocriteria.New("agency").Optional(),
		"office":                    gocriteria.New("office").Optional(),
		"naics":                     gocriteria.New("naics").Optional(),
		"fiscal_year":               gocriteria.New("fiscal_year").Optional(),
		"set_aside":                 gocriteria.New("set_aside").Optional(),
		"contract_vehicle":          gocriteria.New("contract_vehicle").Optional(),
		"is_draft":                  gocriteria.New("is_draft").Optional(),
		"requested":                 gocriteria.New("requested").Optional(),
		"documents":                 gocriteria.New("documents").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"title":                     "Title",
			"description":               "Description",
			"user_contract_value":       "User Contract Value",
			"user_knows_contract_value": "User Knows Contract Value",
			"solicitation_number":       "Solicitation Number",
			"approved":                  "Approved",
			"market":                    "Market",
			"department":                "Department",
			"agency":                    "Agency",
			"office":                    "Office",
			"naics":                     "NAICS",
			"fiscal_year":               "Fiscal Year",
			"set_aside":                 "Set Aside",
			"contract_vehicle":          "Contract Vehicle",
			"is_draft":                  "Is Draft",
			"requested":                 "Requested",
			"documents":                 "Documents",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

type ResponseErrorType map[string][]string

func (s *service) Create(ctx context.Context, input CreateOpportunityRequest, duplicateAllowed bool) (models.Opportunity, response.ErrorResponse) {
	if !policy.CanCreateOpportunity(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a opportunity")
		return models.Opportunity{}, response.ErrorForbidden("You do not have permission to create a opportunity")
	}

	// if is not draft and no solicitation number return error
	if !input.IsDraft && input.SolicitationNumber == "" {
		reserr := make(ResponseErrorType)
		reserr["solicitation_number"] = []string{"Solicitation number is required"}
		s.logger.With(ctx).Error("Solicitation number is required")
		return models.Opportunity{}, response.ErrorBadRequest(reserr, "Solicitation number is required. You need to mark it as draft if you don't have a solicitation number")
	}

	var duplicated bool
	var sameSolicitationCount int64

	if input.SolicitationNumber != "" {
		err := s.db.WithContext(ctx).
			Model(&models.Opportunity{}).
			Where("solicitation_number = ? AND deprecated = ?", input.SolicitationNumber, false).
			Count(&sameSolicitationCount).Error

		if err != nil {
			s.logger.With(ctx).Error(err)
			return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
		}

		if sameSolicitationCount > 0 {
			if duplicateAllowed {
				duplicated = true
			} else {
				return models.Opportunity{}, response.ErrorBadRequest("Opportunity with the same solicitation number already exists")
			}
		}
	}

	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		s.logger.With(ctx).Error(tx.Error)
		return models.Opportunity{}, response.GormErrorResponse(tx.Error, "An error occurred while creating the opportunity")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	var user models.User
	err := tx.WithContext(ctx).First(&user, "id", id).Error
	if err != nil {
		tx.Rollback()
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
	}

	contributionRequestTimeLimitHours := config.AppConfig.ContributionRequestTimeLimitHours
	contributionRequestSizeLimit := config.AppConfig.ContributionRequestSizeLimit

	// check if input requested true and last 24 hours more than 20 opportunities created return error
	if input.Requested {
		var count int64
		err := tx.Model(&models.Opportunity{}).Where("created_at > ? AND user_id = ?", time.Now().Add(-time.Hour*time.Duration(contributionRequestTimeLimitHours)), id).Count(&count).Error
		if err != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(err)
			return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
		}
		if count >= int64(contributionRequestSizeLimit) {
			tx.Rollback()
			return models.Opportunity{}, response.ErrorBadRequest(fmt.Sprintf("You can only create %d opportunities in %d hours", contributionRequestSizeLimit, contributionRequestTimeLimitHours))
		}
	}

	var contractValue string
	var multiAward bool
	var er error
	if input.SolicitationNumber != "" {
		contractValue, multiAward, er = collector.ContractValue(input.SolicitationNumber)
		if er != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(er)
		}
	}

	var opportunityUserID int
	if user.ContributorID.Valid {
		opportunityUserID = int(user.ContributorID.Int64)
	} else {
		opportunityUserID = id
	}

	opportunity := models.Opportunity{
		ApprovedAt: dtp.NullTime{
			Time:  time.Now(),
			Valid: input.Approved && policy.CanApproveOpportunity(ctx),
		},
		CrawlerContractValue: contractValue,
		MultiAward:           multiAward,
		UserID:               opportunityUserID,
		StaffID: dtp.NullInt64{
			Int64: int64(id),
			Valid: user.ContributorID.Valid,
		},
		Title:                  input.Title,
		Description:            input.Description,
		SolicitationNumber:     input.SolicitationNumber,
		Requested:              input.Requested,
		Market:                 input.Market,
		Department:             input.Department,
		Agency:                 input.Agency,
		Office:                 input.Office,
		Naics:                  input.Naics,
		FiscalYear:             input.FiscalYear,
		SetAside:               input.SetAside,
		ContractVehicle:        input.ContractVehicle,
		UserKnowsContractValue: input.UserKnowsContractValue,
		UserContractValue:      input.UserContractValue,
		Duplicated:             duplicated,
		IsDraft:                input.IsDraft,
	}

	if input.SolicitationNumber != "" && opportunity.CrawlerContractValue == "" {
		go func() {
			contractValue, multiAward, err := collector.ContractValue(input.SolicitationNumber)
			if err != nil {
				tx.Rollback()
				s.logger.With(ctx).Error(err)
			} else {
				opportunityData := map[string]interface{}{
					"crawler_contract_value": contractValue,
					"multi_award":            multiAward,
				}
				s.db.WithContext(ctx).Model(&models.Opportunity{}).Where("solicitation_number = ?", input.SolicitationNumber).Updates(opportunityData)
			}
		}()
	}

	err = tx.Create(&opportunity).Error
	if err != nil {
		tx.Rollback()
		s.logger.With(ctx).Error(err)
		return opportunity, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
	}

	// check if it is more than one RFP/RFQ document title return error
	rfpCount := 0
	for _, document := range input.Documents {
		if document.Title == "RFP/RFQ" {
			rfpCount++
		}
	}
	if rfpCount > 1 {
		tx.Rollback()
		return opportunity, response.ErrorBadRequest("Only one RFP/RFQ document is allowed")
	}

	// If user is admin and doc.UserID exists we pick it. if not we use the toke userID

	for _, document := range input.Documents {
		var docOwnerID int
		if document.UserID != 0 && policy.CanDeleteUser(ctx) {
			docOwnerID = document.UserID
		} else {
			if user.ContributorID.Valid {
				docOwnerID = int(user.ContributorID.Int64)
			} else {
				docOwnerID = id
			}
		}

		document := models.Document{
			Title:    document.Title,
			FilePath: document.FilePath,
			Priority: document.Priority,
			LinkType: exp.TerIf(document.LinkType == "", "private", document.LinkType),
			UserID:   docOwnerID,
			StaffID: dtp.NullInt64{
				Int64: int64(id),
				Valid: user.ContributorID.Valid,
			},
		}

		err := tx.Model(&opportunity).Association("Documents").Append(&document)
		if err != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(err)
			return opportunity, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		s.logger.With(ctx).Error(err)
		return opportunity, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
	}

	return opportunity, response.ErrorResponse{}
}
