package transports

import (
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/mid"
	"github.com/Capture-411/services-opportunity/services/bank/endpoints"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Echo, service endpoints.Service, logger log.Logger, prefix string) {
	res := resource{service, logger}

	g := r.Group(prefix)

	// /api/v1/bank/webhook
	g.POST("/stripe/bank/webhooks", res.stripeWebhook)

	rg := g.Group("")
	rg.Use(mid.EchoJWTHandler(), mid.BindUserToContext)

	rg.GET("/bank/products", res.getAllProducts)
	rg.GET("/bank/products/:product_id/prices", res.getAllPrices)
	rg.POST("/bank/create-subscription-checkout-session", res.createSubscriptionCheckoutSession)
	rg.POST("/bank/create-credits-checkout-session", res.createCreditCheckoutSession)
	rg.POST("/bank/create-opportunity-order-checkout-session", res.createOpportunityOrderCheckoutSession)

	rg.GET("/bank/session/:session_id", res.getSessInfo)
	rg.GET("/bank/subscription/:subscription_id", res.getSubscriptionInfo)
	rg.GET("/bank/create-portal-session", res.createPortalSession)
	rg.POST("/bank/subscription/plan-downgrade", res.planDowngradeSubscription)
	rg.DELETE("/bank/subscription/cancel-pending-downgrade/:downgrade_id", res.cancelPendingDowngrade)
	rg.GET("/bank/subscription/downgrades/my-downgrades", res.getMyDowngrades)
	rg.POST("/bank/subscription/upgrade", res.upgradeSubscription)

}

type resource struct {
	service endpoints.Service
	logger  log.Logger
}
