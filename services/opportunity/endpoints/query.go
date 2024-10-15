package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/filter"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: OpportunityQueryFilterType
 */
type OpportunityQueryFilterType struct {
	Title                filter.FilterValue[string] `json:"title,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Description          filter.FilterValue[string] `json:"description" openapi:"nullable;$ref:FilterValueString;"`
	UserContractValue    filter.FilterValue[string] `json:"user_contract_value" openapi:"nullable;$ref:FilterValueString;"`
	GPTContractValue     filter.FilterValue[string] `json:"gpt_contract_value" openapi:"nullable;$ref:FilterValueString;"`
	CrawlerContractValue filter.FilterValue[string] `json:"crawler_contract_value" openapi:"nullable;$ref:FilterValueString;"`
	SolicitationNumber   filter.FilterValue[string] `json:"solicitation_number" openapi:"nullable;$ref:FilterValueString;"`
	ApprovedAt           filter.FilterValue[string] `json:"approved_at" openapi:"nullable;$ref:FilterValueString;"`
	Requested            filter.FilterValue[string] `json:"requested" openapi:"nullable;$ref:FilterValueString;"`
	Duplicated           filter.FilterValue[string] `json:"duplicated" openapi:"nullable;$ref:FilterValueString;"`
	Deprecated           filter.FilterValue[string] `json:"deprecated" openapi:"nullable;$ref:FilterValueString;"`
	Market               filter.FilterValue[string] `json:"market" openapi:"nullable;$ref:FilterValueString;"`
	Department           filter.FilterValue[string] `json:"department" openapi:"nullable;$ref:FilterValueString;"`
	Agency               filter.FilterValue[string] `json:"agency" openapi:"nullable;$ref:FilterValueString;"`
	Office               filter.FilterValue[string] `json:"office" openapi:"nullable;$ref:FilterValueString;"`
	Naics                filter.FilterValue[string] `json:"naics" openapi:"nullable;$ref:FilterValueString;"`
	FiscalYear           filter.FilterValue[string] `json:"fiscal_year" openapi:"nullable;$ref:FilterValueString;"`
	SetAside             filter.FilterValue[string] `json:"set_aside" openapi:"nullable;$ref:FilterValueString;"`
	ContractVehicle      filter.FilterValue[string] `json:"contract_vehicle" openapi:"nullable;$ref:FilterValueString;"`
	UserID               filter.FilterValue[string] `json:"user_id" openapi:"nullable;$ref:FilterValueInt;"`
	StaffID              filter.FilterValue[string] `json:"staff_id" openapi:"nullable;$ref:FilterValueInt;"`
	CreatedAt            filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: OpportunityQueryRequestParams
 */
type OpportunityQueryRequestParams struct {
	ID         string                     `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page       string                     `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage    string                     `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order      string                     `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy    string                     `json:"order_by,omitempty" openapi:"example:title;in:query"`
	Query      string                     `json:"query,omitempty" openapi:"in:query"`
	MineOnly   string                     `json:"mine_only,omitempty" openapi:"in:query"`
	MyPipeline string                     `json:"my_pipeline,omitempty" openapi:"example:true;in:query"`
	Filters    OpportunityQueryFilterType `json:"filters,omitempty" openapi:"$ref:OpportunityQueryFilterType;in:query"`
}

func (data *OpportunityQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":        gocriteria.New("id").Optional(),
		"page":      gocriteria.New("page").Optional(),
		"per_page":  gocriteria.New("per_page").Optional(),
		"query":     gocriteria.New("query").Optional(),
		"order":     gocriteria.New("order").Optional(),
		"order_by":  gocriteria.New("order_by").Optional(),
		"mine_only": gocriteria.New("mine_only").Optional(),
		"filters": gocriteria.Schema{
			"title": gocriteria.Schema{
				"op":    gocriteria.New("filter.title.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.title.value").Optional(),
			},
			"description": gocriteria.Schema{
				"op":    gocriteria.New("filter.description.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.description.value").Optional(),
			},
			"user_contract_value": gocriteria.Schema{
				"op":    gocriteria.New("filter.user_contract_value.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.user_contract_value.value").Optional(),
			},
			"gpt_contract_value": gocriteria.Schema{
				"op":    gocriteria.New("filter.gpt_contract_value.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.gpt_contract_value.value").Optional(),
			},
			"crawler_contract_value": gocriteria.Schema{
				"op":    gocriteria.New("filter.crawler_contract_value.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.crawler_contract_value.value").Optional(),
			},
			"solicitation_number": gocriteria.Schema{
				"op":    gocriteria.New("filter.solicitation_number.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.solicitation_number.value").Optional(),
			},
			"approved_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.approved_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.approved_at.value").Optional(),
			},
			"market": gocriteria.Schema{
				"op":    gocriteria.New("filter.market.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.market.value").Optional(),
			},
			"department": gocriteria.Schema{
				"op":    gocriteria.New("filter.department.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.department.value").Optional(),
			},
			"agency": gocriteria.Schema{
				"op":    gocriteria.New("filter.agency.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.agency.value").Optional(),
			},
			"office": gocriteria.Schema{
				"op":    gocriteria.New("filter.office.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.office.value").Optional(),
			},
			"naics": gocriteria.Schema{
				"op":    gocriteria.New("filter.naics.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.naics.value").Optional(),
			},
			"fiscal_year": gocriteria.Schema{
				"op":    gocriteria.New("filter.fiscal_year.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.fiscal_year.value").Optional(),
			},
			"set_aside": gocriteria.Schema{
				"op":    gocriteria.New("filter.set_aside.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.set_aside.value").Optional(),
			},
			"user_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.user_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.user_id.value").Optional(),
			},
			"staff_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.staff_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.staff_id.value").Optional(),
			},
			"contract_vehicle": gocriteria.Schema{
				"op":    gocriteria.New("filter.contract_vehicle.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.contract_vehicle.value").Optional(),
			},
			"requested": gocriteria.Schema{
				"op":    gocriteria.New("filter.requested.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.requested.value").Optional(),
			},
			"duplicated": gocriteria.Schema{
				"op":    gocriteria.New("filter.duplicated.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.duplicated.value").Optional(),
			},
			"deprecated": gocriteria.Schema{
				"op":    gocriteria.New("filter.deprecated.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.deprecated.value").Optional(),
			},
			"created_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.created_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.created_at.value").Optional(),
			},
		},
	}
	errs := gocriteria.ValidateQueries(r, schema, data)
	if len(errs) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errs)
		return dumpedErrors
	}

	return nil
}

func makeFilters(params OpportunityQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	if params.Filters.Title.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Title.Op, params.Filters.Title.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("title %s %s", opValue.Operator, val))
	}

	if params.Filters.Description.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Description.Op, params.Filters.Description.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("descriptions %s %s", opValue.Operator, val))
	}

	if params.Filters.UserContractValue.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.UserContractValue.Op, params.Filters.UserContractValue.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("contract_value %s %s", opValue.Operator, val))
	}

	if params.Filters.GPTContractValue.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.GPTContractValue.Op, params.Filters.GPTContractValue.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("gpt_contract_value %s %s", opValue.Operator, val))
	}

	if params.Filters.CrawlerContractValue.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CrawlerContractValue.Op, params.Filters.CrawlerContractValue.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("crawler_contract_value %s %s", opValue.Operator, val))
	}

	if params.Filters.SolicitationNumber.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.SolicitationNumber.Op, params.Filters.SolicitationNumber.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("solicitation_number %s %s", opValue.Operator, val))
	}

	if params.Filters.ApprovedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.ApprovedAt.Op, params.Filters.ApprovedAt.Value)
		// transform boolean to to null or not null date time
		if opValue.Value != "" {
			transformedValue := exp.TerIf(opValue.Value == "true", "IS NOT NULL", "IS NULL")
			where = append(where, fmt.Sprintf("approved_at %s", transformedValue))
		}
	}

	if params.Filters.Duplicated.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Duplicated.Op, params.Filters.Duplicated.Value)
		// Assuming the value is 'true', 'false', '1', or '0'
		boolValue := strings.ToLower(opValue.Value) == "true" || opValue.Value == "1"
		where = append(where, fmt.Sprintf("duplicated %s %t", opValue.Operator, boolValue))
	}

	if params.Filters.Deprecated.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Deprecated.Op, params.Filters.Deprecated.Value)
		// Assuming the value is 'true', 'false', '1', or '0'
		boolValue := strings.ToLower(opValue.Value) == "true" || opValue.Value == "1"
		where = append(where, fmt.Sprintf("deprecated %s %t", opValue.Operator, boolValue)) // Corrected typo here
	}

	if params.Filters.Requested.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Requested.Op, params.Filters.Requested.Value)
		// Assuming the value is 'true', 'false', or could be numeric '1', '0'
		boolValue := strings.ToLower(opValue.Value) == "true" || opValue.Value == "1"
		// Format the condition without quotes, using TRUE or FALSE directly
		where = append(where, fmt.Sprintf("requested %s %t", opValue.Operator, boolValue))
	}

	if params.Filters.Market.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Market.Op, params.Filters.Market.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("market %s %s", opValue.Operator, val))
	}

	if params.Filters.Department.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Department.Op, params.Filters.Department.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("department %s %s", opValue.Operator, val))
	}

	if params.Filters.Agency.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Agency.Op, params.Filters.Agency.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("agency %s %s", opValue.Operator, val))
	}

	if params.Filters.Office.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Office.Op, params.Filters.Office.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("office %s %s", opValue.Operator, val))
	}

	if params.Filters.Naics.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Naics.Op, params.Filters.Naics.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("naics %s %s", opValue.Operator, val))
	}

	if params.Filters.FiscalYear.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.FiscalYear.Op, params.Filters.FiscalYear.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("fiscal_year %s %s", opValue.Operator, val))
	}

	if params.Filters.SetAside.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.SetAside.Op, params.Filters.SetAside.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("set_aside %s %s", opValue.Operator, val))
	}

	if params.Filters.UserID.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.UserID.Op, params.Filters.UserID.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("user_id %s %s", opValue.Operator, val))
	}

	if params.Filters.StaffID.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.StaffID.Op, params.Filters.StaffID.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("staff_id %s %s", opValue.Operator, val))
	}

	if params.Filters.ContractVehicle.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.ContractVehicle.Op, params.Filters.ContractVehicle.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("contract_vehicle %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	return where
}

func (s *service) Query(
	ctx context.Context, offset, limit int, params OpportunityQueryRequestParams,
) ([]models.Opportunity, response.ErrorResponse) {
	// we do not exclude bought orders from buy list referring to app store page buying apps and purchased ones are in the list

	Id := policy.ExtractIdClaim(ctx)
	userID, _ := strconv.Atoi(Id)

	if !policy.CanQueryOpportunity(ctx) {
		s.logger.With(ctx).Error("unauthorized to query opportunity")
		return []models.Opportunity{}, response.ErrorUnAuthorized(nil, "unauthorized to query opportunity")
	}

	var allCount int64
	err := s.db.Model(&models.Opportunity{}).Where("requested = ?", false).Count(&allCount).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Opportunity{}, response.ErrorInternalServerError(err, "cant process the request. contact support")
	}

	var opportunities []models.Opportunity

	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User").
		Preload("Staff").
		Preload("Documents.User")

	if policy.CanDeleteOrder(ctx) {
		tx = tx.Preload("Orders")
	}

	where := makeFilters(params)
	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	// if allCount < int64(config.AppConfig.ContributionThreshold) && !policy.CanApproveOpportunity(ctx) {
	// 	tx = tx.Where("opportunities.user_id = ? OR opportunities.requested = true", userID)
	// }

	if params.MineOnly == "true" {
		tx = tx.Where("opportunities.user_id = ? OR opportunities.staff_id = ?", userID, userID)
	}

	// if cant approve means user and user can only see approved opportunities or not approved but requested OR owned ones
	// if !policy.CanApproveOpportunity(ctx) {
	// 	tx = tx.Where("opportunities.approved_at IS NOT NULL OR opportunities.requested = true")
	// }

	// if only my pipeline: orders.user_id = the user id requesting
	if params.MyPipeline == "true" {
		subQuery := s.db.Model(&models.Order{}).
			Select("opportunity_id").
			Where("orders.user_id = ?", userID).
			Group("opportunity_id")
		tx = tx.Joins("JOIN (?) AS valid_opportunities ON valid_opportunities.opportunity_id = opportunities.id", subQuery)
	}

	err = tx.Find(&opportunities).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Opportunity{}, response.GormErrorResponse(err, "error in finding opportunities")
	}

	return opportunities, response.ErrorResponse{}
}
