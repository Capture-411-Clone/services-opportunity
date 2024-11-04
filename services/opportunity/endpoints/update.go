package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
 * @apiDefine: UpdateOpportunityRequest
 */
type UpdateOpportunityRequest struct {
	Title                  string               `json:"title" openapi:"example:opportunity1;type:string;"`
	Description            string               `json:"description" openapi:"example:opportunity1;type:string;"`
	UserKnowsContractValue bool                 `json:"user_knows_contract_value" openapi:"example:true;type:boolean"`
	UserContractValue      string               `json:"user_contract_value" openapi:"example:1000;type:string"`
	SolicitationNumber     string               `json:"solicitation_number" openapi:"example:548;type:string"`
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
	Deprecated             bool                 `json:"deprecated" openapi:"example:false;type:boolean"`
	Documents              []AddDocumentRequest `json:"documents" openapi:"$ref:AddDocumentRequest;example:;type:array"`
}

func (c *UpdateOpportunityRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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
		"deprecated":                gocriteria.New("deprecated").Optional(),
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
			"deprecated":                "Deprecated",
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

/*
 * @apiDefine: UpdateOpportunityRequestParam
 */
type UpdateOpportunityRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateOpportunityRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, id string, input UpdateOpportunityRequest) (
	models.Opportunity, response.ErrorResponse,
) {
	Id := policy.ExtractIdClaim(ctx)
	userID, _ := strconv.Atoi(Id)

	var user models.User
	err := s.db.WithContext(ctx).First(&user, "id", userID).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
	}

	var opportunity models.Opportunity

	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		s.logger.With(ctx).Error(tx.Error)
		return models.Opportunity{}, response.GormErrorResponse(tx.Error, "An error occurred while creating the opportunity")
	}

	txx := tx.Preload("Documents")
	err = txx.First(&opportunity, "id = ?", id).Error

	if err != nil {
		tx.Rollback()
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while finding the opportunity")
	}

	// This is the other user trying to add dcouments :if not owner and not requested add the documents with titles(exept OTHER) which does not exists in opportunity
	// then commit the contribution and exit
	if int(opportunity.StaffID.Int64) != userID && opportunity.UserID != userID && !opportunity.Requested && !policy.CanApproveOpportunity(ctx) {
		approved := true
		for _, doc := range input.Documents {

			// make an array of titles already exisits except title called Other
			var titles []string
			for _, doc := range opportunity.Documents {
				if doc.Title != "Other" {
					titles = append(titles, doc.Title)
				}
			}

			// if the title of the document is not in the titles array add it to the opportunity and save it

			if !InArray(doc.Title, titles) {
				var documentUserID int
				if user.ContributorID.Valid {
					documentUserID = int(user.ContributorID.Int64)
				} else {
					documentUserID = userID
				}

				document := models.Document{
					Title:    doc.Title,
					FilePath: doc.FilePath,
					Priority: doc.Priority,
					LinkType: exp.TerIf(doc.LinkType == "", "private", doc.LinkType),
					UserID:   documentUserID,
				}

				err = tx.Model(&opportunity).Association("Documents").Append(&document)
				if err != nil {
					tx.Rollback()
					s.logger.With(ctx).Error(err)
					return models.Opportunity{}, response.GormErrorResponse(err, "error to add document")
				}

				approved = false
			}
		}

		if !approved {
			err = tx.Model(&opportunity).Update("approved_at", nil).Error
			if err != nil {
				tx.Rollback()
				s.logger.With(ctx).Error(err)
				return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while updating the opportunity")
			}
		}

		fmt.Println("Requested:", opportunity.Requested)

		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(err)
			return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while updating the opportunity")
		}

		return opportunity, response.ErrorResponse{}
	}

	var duplicated bool
	var sameOpportunitiesCount int64
	var sameOpportunities []models.Opportunity

	if input.SolicitationNumber != "" {
		err = s.db.WithContext(ctx).Model(&models.Opportunity{}).Where("solicitation_number = ? AND deprecated = ? AND id != ?", input.SolicitationNumber, false, opportunity.ID).Find(&sameOpportunities).Count(&sameOpportunitiesCount).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while creating the opportunity")
		}
		if sameOpportunitiesCount > 0 {
			duplicated = true
		}
	}

	opportunity.Duplicated = duplicated

	if input.Deprecated {
		// Assuming 'sameOpportunities' already contains opportunities with the same solicitation number,
		// excluding the current one. You may need to adjust the query if not.
		if len(sameOpportunities) == 1 {
			// If there's exactly one other opportunity with the same solicitation number and the current
			// opportunity is being set to deprecated, update the other opportunity to not be marked as duplicated.
			otherOpportunity := &sameOpportunities[0]
			otherOpportunity.Duplicated = false
			if err := tx.Save(&otherOpportunity).Error; err != nil {
				tx.Rollback()
				s.logger.With(ctx).Error(err)
				return models.Opportunity{}, response.GormErrorResponse(err, "Failed to update duplicate status of other opportunity")
			}
		}
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

	if policy.CanApproveOpportunity(ctx) {
		opportunity.ApprovedAt = dtp.NullTime{
			Time:  time.Now(),
			Valid: input.Approved,
		}
	}

	opportunity.Title = input.Title

	opportunity.Description = input.Description

	opportunity.SolicitationNumber = input.SolicitationNumber

	opportunity.Market = input.Market

	opportunity.Department = input.Department

	opportunity.Agency = input.Agency

	opportunity.Office = input.Office

	opportunity.Naics = input.Naics

	opportunity.FiscalYear = input.FiscalYear

	opportunity.SetAside = input.SetAside

	opportunity.ContractVehicle = input.ContractVehicle

	opportunity.IsDraft = input.IsDraft

	if policy.IsAdmin(ctx) || policy.IsReviewer(ctx) {
		opportunity.UserContractValue = input.UserContractValue
		opportunity.UserKnowsContractValue = input.UserKnowsContractValue
	}

	if input.SolicitationNumber != "" && opportunity.CrawlerContractValue == "" {
		contractValue, multiAward, err := collector.ContractValue(input.SolicitationNumber)
		if err != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(err)

			opportunity.CrawlerContractValue = ""
			opportunity.MultiAward = false
		} else {
			opportunity.CrawlerContractValue = contractValue
			opportunity.MultiAward = multiAward
		}
	}

	opportunity.Deprecated = input.Deprecated

	err = tx.Save(&opportunity).Error
	if err != nil {
		tx.Rollback()
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while updating the opportunity")
	}

	if len(input.Documents) != 0 {
		// Remove all documents
		err = tx.Model(&opportunity).Association("Documents").Clear()
		if err != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(err)
			return models.Opportunity{}, response.GormErrorResponse(err, "error to remove documents")
		}

		for _, doc := range input.Documents {

			// If user is admin and doc.UserID exists we pick it. if not we use the toke userID
			var docOwnerID int
			if doc.UserID != 0 && policy.CanDeleteUser(ctx) {
				docOwnerID = doc.UserID
			} else {
				if user.ContributorID.Valid {
					docOwnerID = int(user.ContributorID.Int64)
				} else {
					docOwnerID = userID
				}
			}

			document := models.Document{
				Title:    doc.Title,
				FilePath: doc.FilePath,
				Priority: doc.Priority,
				LinkType: exp.TerIf(doc.LinkType == "", "private", doc.LinkType),
				UserID:   docOwnerID,
				StaffID: dtp.NullInt64{
					Int64: int64(userID),
					Valid: user.ContributorID.Valid,
				},
			}

			err = tx.Model(&opportunity).Association("Documents").Append(&document)
			if err != nil {
				tx.Rollback()
				s.logger.With(ctx).Error(err)
				return models.Opportunity{}, response.GormErrorResponse(err, "error to add document")
			}
		}
	}

	if len(opportunity.Documents) > 0 && opportunity.SolicitationNumber != "" {
		err = tx.Model(&opportunity).Update("requested", input.Requested).Error
		if err != nil {
			tx.Rollback()
			s.logger.With(ctx).Error(err)
			return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while updating the opportunity")
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "An error occurred while updating the opportunity")
	}

	return opportunity, response.ErrorResponse{}
}

func InArray(val string, array []string) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}
