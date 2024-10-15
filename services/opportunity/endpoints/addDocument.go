package endpoints

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: AddDocumentRequest
 */
type AddDocumentRequest struct {
	Title    string `json:"title" openapi:"example:document1;type:string;"`
	Priority int    `json:"priority" openapi:"example:1;type:integer;"`
	FilePath string `json:"file_path" openapi:"example:/file/path;type:string;"`
	LinkType string `json:"link_type" openapi:"example:;type:string;"`
	UserID   int    `json:"user_id" openapi:"example:1;type:integer;"`
}

func (m *AddDocumentRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"title":     gocriteria.New("title").Required(),
		"priority":  gocriteria.New("priority").Optional(),
		"file_path": gocriteria.New("file_path").Optional(),
		"link_type": gocriteria.New("link_type").Optional(),
		"user_id":   gocriteria.New("user_id").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"title":     "Title",
			"priority":  "Priority",
			"file_path": "File",
			"link_type": "Link Type",
			"user_id":   "User ID",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, m)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

/*
 * @apiDefine: AddDocumentRequestParam
 */
type AddDocumentRequestParam struct {
	OpportunityID int `json:"opportunity_id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *AddDocumentRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"opportunity_id": gocriteria.New("opportunity_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"opportunity_id": "Opportunity ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) AddDocument(ctx context.Context, opportunityID int, input AddDocumentRequest) (
	models.Opportunity, response.ErrorResponse,
) {
	Id := policy.ExtractIdClaim(ctx)
	userID, _ := strconv.Atoi(Id)

	var opportunity models.Opportunity

	if !policy.CanAddOpportunityDocument(ctx, opportunity) {
		s.logger.With(ctx).Error("you do not have permission to add document to opportunity")
		return models.Opportunity{}, response.ErrorForbidden("you do not have permission to add document to opportunity")
	}

	err := s.db.WithContext(ctx).First(&opportunity, "id", opportunityID).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "opportunity not found")
	}

	// If user is admin and doc.UserID exists we pick it. if not we use the toke userID
	var docOwnerID int
	if input.UserID != 0 && policy.CanDeleteUser(ctx) {
		docOwnerID = input.UserID
	} else {
		docOwnerID = userID
	}

	document := models.Document{
		Title:    input.Title,
		FilePath: input.FilePath,
		Priority: input.Priority,
		LinkType: exp.TerIf(input.LinkType == "", "private", input.LinkType),
		UserID:   docOwnerID,
	}

	err = s.db.WithContext(ctx).Model(&opportunity).Association("Documents").Append(&document)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Opportunity{}, response.GormErrorResponse(err, "error to add document")
	}

	return opportunity, response.ErrorResponse{}
}
