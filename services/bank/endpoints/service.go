package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"
	"github.com/stripe/stripe-go/v76"
	"gorm.io/gorm"
)

type Service interface {
	CreateOpportunityOrderCheckoutSession(ctx context.Context, input CreateOpportunityOrderCheckoutSessionRequestBody) (string, string, response.ErrorResponse)
	CreateSubscriptionCheckoutSession(ctx context.Context, input CreateSubscriptionCheckoutSessionRequestBody) (string, string, response.ErrorResponse)
	CreateCreditCheckoutSession(ctx context.Context, input CreateCreditCheckoutSessionRequestBody) (string, string, response.ErrorResponse)

	GetAllProducts(ctx context.Context) ([]stripe.Product, response.ErrorResponse)
	GetAllPrices(ctx context.Context, params GetAllPricesRequestParams) ([]stripe.Price, response.ErrorResponse)
	GetSessionInfo(ctx context.Context, params GetSessInfoRequestParams) (*stripe.CheckoutSession, response.ErrorResponse)
	GetSubscriptionInfo(ctx context.Context, params GetSubscriptionInfoRequestParams) (*stripe.Subscription, response.ErrorResponse)
	CreatePortalSession(ctx context.Context) (string, response.ErrorResponse)
	ChangeSubscriptionForDowngrade(ctx context.Context, input ChangeSubscriptionForDowngradeRequestBody) (*stripe.Subscription, response.ErrorResponse)
	PlanDowngradeSubscription(ctx context.Context, input PlanDowngradeRequestBody) response.ErrorResponse
	CancelPendingDowngrade(ctx context.Context, params CancelPendingDowngradeRequestParams) response.ErrorResponse
	GetMyDowngrades(ctx context.Context) ([]models.Downgrade, response.ErrorResponse)
	UpgradeSubscription(ctx context.Context, input UpgradeSubscriptionRequestBody) (UpgradeSubscriptionResponseData, response.ErrorResponse)
	ApplyFreshDowngrades(ctx context.Context) error

	HandleOpportunityOrderPaid(ctx context.Context, cust *stripe.Customer, orderID int64, eventID string) response.ErrorResponse
	HandleSubscriptionPaid(ctx context.Context, customer *stripe.Customer, subscriptionID *stripe.Subscription, price *stripe.Price, eventID string) response.ErrorResponse
	HandleSubscriptionUpdated(ctx context.Context, cust *stripe.Customer, sub *stripe.Subscription) response.ErrorResponse
	HandleAddMoreCredits(ctx context.Context, customer *stripe.Customer, quantity int64, amountPaid int64, eventID string) response.ErrorResponse
}

type service struct {
	db       *gorm.DB
	logger   log.Logger
	notifier notificationv2.Notifier
}

func MakeService(db *gorm.DB, logger log.Logger, notifier notificationv2.Notifier) Service {
	return &service{
		db,
		logger,
		notifier,
	}
}
