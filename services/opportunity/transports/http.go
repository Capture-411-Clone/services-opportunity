package transports

import (
	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/opportunity/collector"
	"github.com/Capture-411/services-opportunity/services/opportunity/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	if config.AppConfig.Environment == "development" {
		g.GET("/opportunities/contract-value/:solnum", func(c echo.Context) error {
			solNumm := c.Param("solnum")
			value, multiAwarded, err := collector.ContractValue(solNumm)
			if err != nil {
				return c.JSON(500, err)
			}
			return c.JSON(200, map[string]interface{}{
				"value":         value,
				"multi_awarded": multiAwarded,
			})
		})
	}

	// ------------------------------------------------

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.POST("/opportunities/pdf-doc-values", res.pdfDocValues)
	rg.POST("/opportunities/ai-doc-values", res.AIDocValues)
	rg.POST("/opportunities/ai-all-doc-values", res.AIAllDocValues)

	rg.GET("/opportunities", res.query)
	rg.POST("/opportunities", res.create)
	rg.Match([]string{"PUT", "PATCH"}, "/opportunities/:id", res.update)
	rg.DELETE("/opportunities", res.delete)
	rg.POST("/opportunities/:opportunity_id/documents", res.addDocument)

	rg.GET("/opportunities/:solicitation_number/check-duplicate", res.checkDuplicate)

	rg.GET("/opportunities/:id/price", res.calculatePrice)

	rg.GET("/opportunities/:id/download-all-documents", res.downloadAllDocuments)
	rg.GET("/opportunities/contributions-state", res.contributionsState)
	rg.POST("/opportunities/mass-import", res.importerOpportunities)

}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
