package transports

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/product"
	"github.com/stripe/stripe-go/v76/subscription"
	"github.com/stripe/stripe-go/v76/webhook"
)

// TODO: if user renew from portal i need to update the subscription id in user table
// TODO: if user downgrade i should update releted data

func (r resource) stripeWebhook(ctx echo.Context) error {
	req := ctx.Request()

	// Read the body of the request (needed for verification)
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error reading request body")
	}

	// Verify the event by checking its signature
	signature := req.Header.Get("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, signature, config.AppConfig.StripeWebhookSecret)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error verifying webhook signature")
	}

	// Handle the event
	switch event.Type {
	case "checkout.session.completed":
		var theSess stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &theSess)
		if err != nil {
			r.logger.Error("checkout.session.completed: Error unmarshalling session", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}

		orderID, ok := theSess.Metadata["order_id"]

		orderIDInt64, _ := strconv.ParseInt(orderID, 10, 64)

		if ok && orderIDInt64 != 0 {
			r.service.HandleOpportunityOrderPaid(ctx.Request().Context(), theSess.Customer, orderIDInt64, event.ID)
		} else if theSess.Subscription != nil {
			// Assuming you've stored the subscription ID in the session when creating it
			subscriptionID := theSess.Subscription.ID
			if subscriptionID != "" {
				// Fetch the subscription
				sub, err := subscription.Get(subscriptionID, nil)
				if err != nil {
					r.logger.Error("checkout.session.completed -> subscription: Error getting subscription", err)
				}

				// Fetch the customer
				customerID := sub.Customer.ID
				cust, err := customer.Get(customerID, nil)
				if err != nil {
					r.logger.Error("checkout.session.completed -> subscription: Error getting customer", err)
				}

				// Assuming the first item's price's product is what you need
				price := sub.Items.Data[0].Price
				// prod, err := product.Get(price.Product.ID, nil)
				// if err != nil {
				// 	r.logger.Error("checkout.session.completed -> subscription: Error getting product", err)
				// 	return ctx.NoContent(http.StatusInternalServerError)
				// }

				er := r.service.HandleSubscriptionPaid(ctx.Request().Context(), cust, sub, price, event.ID)
				if er.StatusCode != 0 {
					r.logger.Error("checkout.session.completed -> subscription: Error handling subscription paid", err)
				}
			}

			// if the buy was for credits (one-time purchase with quantity)
		} else {
			params := &stripe.CheckoutSessionListLineItemsParams{
				Session: stripe.String(theSess.ID),
			}
			i := session.ListLineItems(params)

			// Iterate through the line items
			for i.Next() {
				lineItem := i.LineItem()
				fmt.Printf("Line item: %v\n", lineItem)
				// Ensure lineItem and lineItem.Price are not nil to prevent dereference errors
				if lineItem == nil {
					continue
				}

				// Fetch the product associated with the line item
				prod, err := product.Get(lineItem.Price.Product.ID, nil)
				if err != nil {
					r.logger.Error("checkout.session.completed -> one-time purchase: Error getting product", err)
					return ctx.NoContent(http.StatusInternalServerError)
				}

				// Here you'd check if the product represents "credits"
				if prod.Metadata["type"] == "credit" {
					quantity := lineItem.Quantity
					amountPaid := lineItem.AmountTotal
					r.service.HandleAddMoreCredits(ctx.Request().Context(), theSess.Customer, quantity, amountPaid, event.ID)
				}
			}
		}

	case "invoice.payment_succeeded":
		var invoice stripe.Invoice
		err := json.Unmarshal(event.Data.Raw, &invoice)
		if err != nil {
			r.logger.Error("invoice.payment_succeeded: Error unmarshalling invoice", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
		if invoice.BillingReason == "subscription_cycle" {
			customerID := invoice.Customer.ID
			cust, err := customer.Get(customerID, nil)
			if err != nil {
				r.logger.Error("invoice.payment_succeeded: Error getting customer", err)
			}

			subscriptionID := invoice.Subscription.ID
			sub, err := subscription.Get(subscriptionID, nil)
			if err != nil {
				r.logger.Error("invoice.payment_succeeded: Error getting subscription", err)
			}

			price := sub.Items.Data[0].Price
			// Assuming the first item's price's product is what you need
			// productID := sub.Items.Data[0].Price.Product.ID
			// prod, err := product.Get(productID, nil)
			// if err != nil {
			// 	r.logger.Error("invoice.payment_succeeded: Error getting product", err)
			// }

			er := r.service.HandleSubscriptionPaid(ctx.Request().Context(), cust, sub, price, event.ID)
			if er.StatusCode != 0 {
				r.logger.Error("checkout.session.completed -> subscription: Error handling subscription paid", err)
			}
		}

	case "customer.subscription.updated":
		// Cancel, Downgrade
		var sub stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &sub)
		if err != nil {
			r.logger.Error("customer.subscription.updated: Error unmarshalling subscription", err)
		}

		// Fetch the customer
		customerID := sub.Customer.ID
		cust, err := customer.Get(customerID, nil)
		if err != nil {
			r.logger.Error("customer.subscription.updated: Error getting customer", err)
		}

		// Assuming the first item's price's product is what you need
		// productID := sub.Items.Data[0].Price.Product.ID
		// prod, err := product.Get(productID, nil)
		// if err != nil {
		// 	r.logger.Error("customer.subscription.updated: Error getting product", err)
		// }

		er := r.service.HandleSubscriptionUpdated(ctx.Request().Context(), cust, &sub)
		if er.StatusCode != 0 {
			r.logger.Error("customer.subscription.updated: Error handling subscription updated", err)
		}

	}

	return ctx.NoContent(http.StatusOK)
}
