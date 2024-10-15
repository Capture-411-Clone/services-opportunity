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
 * @apiDefine: NotificationQueryFilterType
 */
type NotificationQueryFilterType struct {
	Body         filter.FilterValue[string] `json:"body,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"contains\",\"value\":\"hello\"}"`
	Recipient    filter.FilterValue[string] `json:"recipient,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"contains\",\"value\":\"michael\"}\"}"`
	Drivers      filter.FilterValue[string] `json:"drivers,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"contains\",\"value\":\"email,database\"}\"}"`
	SenderUserID filter.FilterValue[string] `json:"sender_user_id,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"equals\",\"value\":1\"}"`
	TargetUserID filter.FilterValue[string] `json:"target_user_id,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"equals\",\"value\":1\"}\"}"`
	Seen         filter.FilterValue[string] `json:"seen,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"equals\",\"value\":true\"}\"}"`
	Subject      filter.FilterValue[string] `json:"subject,omitempty" openapi:"$ref:FilterValueString;example:{\"name\":{\"op\":\"contains\",\"value\":\"hello\"}\"}"`
	CreatedAt    filter.FilterValue[string] `json:"created_at,omitempty" openapi:"$ref:FilterValueString;example:{\"created_at\":{\"op\":\"equals\",\"value\":\"2021-01-01\"}"`
}

/*
 * @apiDefine: NotificationQueryRequestParams
 */
type NotificationQueryRequestParams struct {
	ID      string                      `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                      `json:"page,omitempty" openapi:"required;pattern:^[0-9]+$;in:query;example:1"`
	PerPage string                      `json:"per_page,omitempty" openapi:"required;pattern:^[0-9]+$;in:query;example:10"`
	Order   string                      `json:"order,omitempty" openapi:"in:query;example:desc;"`
	OrderBy string                      `json:"order_by,omitempty" openapi:"in:query;example:recipient"`
	Filters NotificationQueryFilterType `json:"filters,omitempty" openapi:"$ref:NotificationQueryFilterType;type:object;in;query;"`
}

func (r *NotificationQueryRequestParams) ValidateQueries(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"page":     gocriteria.New("page").Required(),
		"per_page": gocriteria.New("per_page").Required(),
		"order":    gocriteria.New("order"),
		"order_by": gocriteria.New("order_by"),
		"filters": gocriteria.Schema{
			"body": gocriteria.Schema{
				"op":    gocriteria.New("filter.body.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.body.value").Optional(),
			},
			"recipient": gocriteria.Schema{
				"op":    gocriteria.New("filter.recipient.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.recipient.value").Optional(),
			},
			"drivers": gocriteria.Schema{
				"op":    gocriteria.New("filter.drivers.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.drivers.value").Optional(),
			},
			"sender_user_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.sender_user_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.sender_user_id.value").Optional(),
			},
			"target_user_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.target_user_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.target_user_id.value").Optional(),
			},
			"seen": gocriteria.Schema{
				"op":    gocriteria.New("filter.seen.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.seen.value").Optional(),
			},
			"created_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.created_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.created_at.value").Optional(),
			},
		},
	}

	errr := gocriteria.ValidateQueries(req, schema, r)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func makeFilters(params NotificationQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	if params.Filters.Body.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Body.Op, params.Filters.Body.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("body %s %s", opValue.Operator, val))
	}

	if params.Filters.Recipient.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Recipient.Op, params.Filters.Recipient.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("recipient %s %s", opValue.Operator, val))
	}

	if params.Filters.Drivers.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Drivers.Op, params.Filters.Drivers.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("driver %s %s", opValue.Operator, val))
	}

	if params.Filters.SenderUserID.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.SenderUserID.Op, params.Filters.SenderUserID.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("sender_user_id %s %s", opValue.Operator, val))
	}

	if params.Filters.TargetUserID.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.TargetUserID.Op, params.Filters.TargetUserID.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("target_user_id %s %s", opValue.Operator, val))
	}

	if params.Filters.Seen.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Seen.Op, params.Filters.Seen.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("seen %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	if params.Filters.Subject.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Subject.Op, params.Filters.Subject.Value)
		val := exp.TerIf(opValue.Value == "", "", opValue.Value)
		where = append(where, fmt.Sprintf("subject %s %s", opValue.Operator, val))
	}

	return where
}

func (s *service) Query(
	ctx context.Context, offset, limit int, requestParams NotificationQueryRequestParams,
) ([]models.Notification, response.ErrorResponse) {

	var notifications []models.Notification
	var user models.User

	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	err := s.db.WithContext(ctx).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Notification{}, response.GormErrorResponse(err, "Error checking user role")
	}

	tx := s.db.WithContext(ctx).
		Order(fmt.Sprintf("%s %s", requestParams.OrderBy, requestParams.Order)).
		Offset(offset).Limit(limit).Preload("SenderUser").Preload("TargetUser")

	where := makeFilters(requestParams)
	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	onlyOwner := !policy.CanQueryAllNotifications(ctx)

	if onlyOwner {
		tx = tx.Where("target_user_id = ? OR sender_user_id = ?", theID, theID)
		// TODO: Change this in a way checking seen at one week ago no created_at. No Seen at field Yet
		tx = tx.Where("seen = false OR (seen = true AND created_at > now() - interval '1 week')")
	}

	err = tx.Find(&notifications).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Notification{}, response.GormErrorResponse(err, "Error finding notifications")
	}

	return notifications, response.ErrorResponse{}
}
