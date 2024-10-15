package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/role"
)

/*
 * @apiDefine: CanAccessDocumentRequest
 */
type CanAccessDocumentRequest struct {
	FilePath string `json:"file_path" openapi:"example:file.pdf;type:string"`
	UserID   int    `json:"user_id" openapi:"example:1;type:int"`
}

func (m *CanAccessDocumentRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"file_path": gocriteria.New("file_path").Required(),
		"user_id":   gocriteria.New("user_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"file_path": "file",
			"user_id":   "UserID",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, m)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) CanAccess(ctx context.Context, input CanAccessDocumentRequest) (string, response.ErrorResponse) {
	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", input.UserID).Preload("Roles").First(&user).Error
	if err != nil {
		s.logger.With(ctx).Error("Failed to download file", err)
		return "", response.ErrorBadRequest(nil, "Failed to download file")
	}

	var document models.Document
	err = s.db.WithContext(ctx).Where("file_path = ?", input.FilePath).First(&document).Error
	if err != nil {
		s.logger.With(ctx).Error("Error finding document", err)
		return "", response.GormErrorResponse(nil, "Error finding document")
	}

	canAccess, err := s.CanAccessDocument(ctx, document, user.ID)
	if err != nil {
		s.logger.With(ctx).Error("Error checking access to document", err)
		return "", response.GormErrorResponse(nil, "Error checking access to document")
	}

	var oppo models.Opportunity
	err = s.db.Model(&models.Opportunity{}).Where("id = ?", document.OwnerID).First(&oppo).Error
	if err != nil {
		s.logger.With(ctx).Error("Error finding opportunity", err)
		return "", response.GormErrorResponse(nil, "Error finding opportunity")
	}

	if !canAccess && !s.UserBoughtOpportunity(ctx, user, oppo) {
		s.logger.With(ctx).Error("you do not have permission to access this document")
		return "", response.ErrorForbidden(nil, "you do not have permission to access this document")
	}

	return "OK-HasAccess", response.ErrorResponse{}
}

func (s *service) CanAccessDocument(ctx context.Context, doc models.Document, userID uint) (bool, error) {
	// This method needs user roles preloaded

	isAdmin := false
	isReviewer := false

	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", userID).Preload("Roles").First(&user).Error
	if err != nil {
		return false, err
	}

	roles := user.Roles
	for _, r := range roles {
		if r.Title == role.Admin {
			isAdmin = true
		}
		if r.Title == role.Reviewer {
			isReviewer = true
		}
	}

	if isAdmin || isReviewer || user.ID == uint(doc.UserID) || user.ID == uint(doc.StaffID.Int64) {
		return true, nil
	}

	return false, nil
}

func (s *service) UserBoughtOpportunity(ctx context.Context, user models.User, opportunity models.Opportunity) bool {
	var orders []models.Order
	err := s.db.Model(models.Order{}).Where("user_id = ? AND opportunity_id = ? AND paid_at IS NOT NULL AND refunded_at IS NULL", user.ID, opportunity.ID).Find(&orders).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Checking Prev Orders Failed: %v", err)
		return false
	}

	return len(orders) > 0
}
