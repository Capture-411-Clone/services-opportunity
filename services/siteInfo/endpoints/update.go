package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateSiteInfoRequest
 */
type UpdateSiteInfoRequest struct {
	CompanyName      string `json:"company_name" openapi:"example:My Company;type:string"`
	MissionStatement string `json:"mission_statement" openapi:"example:To be the best;type:string"`
	History          string `json:"history" openapi:"example:To be the best;type:string"`
	Goal             string `json:"goal" openapi:"example:To be the best;type:string"`
	Value            string `json:"value" openapi:"example:To be the best;type:string"`
	Achievement      string `json:"achievement" openapi:"example:To be the best;type:string"`
	Member           string `json:"member" openapi:"example:Ayman;type:string"`
	GeneralContact   string `json:"general_contact" openapi:"example:mey@gmail.com;type:string"`
	Address          string `json:"address" openapi:"example:My Address;type:string"`
	SocialMedia      string `json:"social_media" openapi:"example:@Christina;type:string"`
	PhoneNumber      string `json:"phone_number" openapi:"example:+1234567890;type:string"`
	EmailAddress     string `json:"email_address" openapi:"example:CM@gmail.com;type:string"`
	OfficeHours      string `json:"office_hours" openapi:"example:07:00 AM - 10:00 PM;type:string"`
	PhysicalAddress  string `json:"physical_address" openapi:"example:Right Here;type:string"`
	MapOrDirections  string `json:"map_or_directions" openapi:"example:98.76.54;type:string"`
	EmergencyContact string `json:"emergency_contact" openapi:"example:+18787654321;type:string"`
	FeedbackLink     string `json:"feedback_link" openapi:"example:My Feedback Link;type:string"`
	SupportLink      string `json:"support_link" openapi:"example:My Support Link;type:string"`
}

func (c *UpdateSiteInfoRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"company_name":      gocriteria.New("company_name").Required(),
		"mission_statement": gocriteria.New("mission_statement").Optional(),
		"history":           gocriteria.New("history").Optional(),
		"goals":             gocriteria.New("goals").Optional(),
		"values":            gocriteria.New("values").Optional(),
		"achievements":      gocriteria.New("achievements").Optional(),
		"members":           gocriteria.New("members").Optional(),
		"general_contact":   gocriteria.New("general_contact").Optional(),
		"address":           gocriteria.New("address").Optional(),
		"social_media":      gocriteria.New("social_media").Optional(),
		"phone_number":      gocriteria.New("phone_number").Optional(),
		"email_address":     gocriteria.New("email_address").Optional(),
		"office_hours":      gocriteria.New("office_hours").Optional(),
		"physical_address":  gocriteria.New("physical_address").Optional(),
		"map_or_directions": gocriteria.New("map_or_directions").Optional(),
		"emergency_contact": gocriteria.New("emergency_contact").Optional(),
		"feedback_link":     gocriteria.New("feedback_link").Optional(),
		"support_link":      gocriteria.New("support_link").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"company_name":      "Company Name",
			"mission_statement": "Mission Statement",
			"history":           "History",
			"goal":              "Goal",
			"value":             "Value",
			"achievement":       "Achievement",
			"member":            "Member",
			"general_contact":   "General Contact",
			"address":           "Address",
			"social_media":      "Social Media",
			"phone_number":      "Phone Number",
			"email_address":     "Email Address",
			"office_hours":      "Office Hours",
			"physical_address":  "Physical Address",
			"map_or_directions": "Map or Directions",
			"emergency_contact": "Emergency Contact",
			"feedback_link":     "Feedback Link",
			"support_link":      "Support Link",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, input UpdateSiteInfoRequest) (
	models.SiteInfo, response.ErrorResponse,
) {
	var site models.SiteInfo

	if !policy.CanUpdateSiteInfo(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a site")
		return models.SiteInfo{}, response.ErrorForbidden(nil, "You do not have permission to update a site")
	}

	err := s.db.WithContext(ctx).First(&site, "id", 1).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.SiteInfo{}, response.GormErrorResponse(err, "An error occurred while finding the site")
	}

	site.CompanyName = input.CompanyName
	site.MissionStatement = input.MissionStatement
	site.History = input.History
	site.Goal = input.Goal
	site.Value = input.Value
	site.Achievement = input.Achievement
	site.Member = input.Member
	site.GeneralContact = input.GeneralContact
	site.Address = input.Address
	site.SocialMedia = input.SocialMedia
	site.PhoneNumber = input.PhoneNumber
	site.EmailAddress = input.EmailAddress
	site.OfficeHours = input.OfficeHours
	site.PhysicalAddress = input.PhysicalAddress
	site.MapOrDirections = input.MapOrDirections
	site.EmergencyContact = input.EmergencyContact
	site.FeedbackLink = input.FeedbackLink

	err = s.db.WithContext(ctx).Save(&site).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.SiteInfo{}, response.GormErrorResponse(err, "An error occurred while updating the site")
	}
	return site, response.ErrorResponse{}
}
