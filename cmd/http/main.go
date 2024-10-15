package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Capture-411/services-opportunity/database"
	notifierV2Module "github.com/Capture-411/services-opportunity/notification-v2"
	"github.com/Capture-411/services-opportunity/utils"
	"github.com/Capture-411/services-opportunity/validity"

	"github.com/Capture-411/services-opportunity/config"
	echohandlers "github.com/Capture-411/services-opportunity/kit/echo"
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/services/account"
	"github.com/Capture-411/services-opportunity/services/agency"
	"github.com/Capture-411/services-opportunity/services/bank"
	"github.com/Capture-411/services-opportunity/services/calculators"
	"github.com/Capture-411/services-opportunity/services/contractVehicle"
	"github.com/Capture-411/services-opportunity/services/department"
	"github.com/Capture-411/services-opportunity/services/document"
	"github.com/Capture-411/services-opportunity/services/entityHunt"
	"github.com/Capture-411/services-opportunity/services/fiscalYear"
	"github.com/Capture-411/services-opportunity/services/healthcheck"
	"github.com/Capture-411/services-opportunity/services/invite"
	"github.com/Capture-411/services-opportunity/services/market"
	"github.com/Capture-411/services-opportunity/services/naics"
	"github.com/Capture-411/services-opportunity/services/notification"
	"github.com/Capture-411/services-opportunity/services/office"
	"github.com/Capture-411/services-opportunity/services/opportunity"
	"github.com/Capture-411/services-opportunity/services/order"
	"github.com/Capture-411/services-opportunity/services/permission"
	"github.com/Capture-411/services-opportunity/services/role"
	"github.com/Capture-411/services-opportunity/services/setAside"
	"github.com/Capture-411/services-opportunity/services/siteInfo"
	"github.com/Capture-411/services-opportunity/services/user"
	"github.com/Capture-411/services-opportunity/services/verify"
	"github.com/Capture-411/services-opportunity/services/welcome"
	"github.com/Capture-411/services-opportunity/services/wishlist"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stripe/stripe-go/v76"
)

func main() {

	flag.Parse()

	config.Load()

	ctx := context.Background()

	validity.ApplyTranslations()

	stripe.Key = config.AppConfig.StripeSecretKey

	// ------------------------- Checkings -------------------------

	if pdftotextExists := utils.IsPdftotextInstalled(); !pdftotextExists {
		fmt.Println("pdftotext is not installed. Please install it and try again.")
		os.Exit(1)
	}

	// -------------------------------------------------------------

	var (
		// APP
		HostUri      = config.AppConfig.HostUri
		Protocol     = config.AppConfig.Protocol
		HostAddress  = config.AppConfig.HostAddress
		ExternalPort = strconv.Itoa(config.AppConfig.ExternalPort)
		InternalPort = strconv.Itoa(config.AppConfig.InternalPort)
		VERSION      = config.AppConfig.Version
		// JWT
		AccessTokenSigningKey      = config.AppConfig.AccessTokenSigningKey
		AccessTokenTokenExpiration = config.AppConfig.AccessTokenTokenExpiration
		RefreshTokenSigningKey     = config.AppConfig.RefreshTokenSigningKey
		RefreshTokenExpiration     = config.AppConfig.RefreshTokenExpiration
	)

	apiVersion1 := config.AppConfig.ApiVersion1

	logger := log.New().With(ctx, "version", VERSION)

	db := database.Connect()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: `[${time_rfc3339}] [${latency_human}] [${remote_ip}:${method}] [${status}]   err=[${error}]` + "\n",
			},
		),
	)
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.BodyLimit("20M"))

	e.HTTPErrorHandler = echohandlers.CustomHandler

	e.Static("/", config.AppConfig.StaticDirPath)

	healthcheck.RegisterHandlers(e, config.AppConfig.Version)

	welcome.RegisterHandlers(e)

	notifierV2 := notifierV2Module.MakeNotifier(db)

	account.Register(e, db, logger, notifierV2, AccessTokenSigningKey, RefreshTokenSigningKey, AccessTokenTokenExpiration, RefreshTokenExpiration, apiVersion1)

	user.Register(e, db, logger, apiVersion1)

	entityHunt.Register(e, db, logger, apiVersion1)

	verify.Register(e, db, notifierV2, logger, apiVersion1)

	role.Register(e, db, logger, apiVersion1)

	permission.Register(e, db, logger, apiVersion1)

	market.Register(e, db, logger, apiVersion1)

	invite.Register(e, db, logger, apiVersion1)

	department.Register(e, db, logger, apiVersion1)

	naics.Register(e, db, logger, apiVersion1)

	fiscalYear.Register(e, db, logger, apiVersion1)

	setAside.Register(e, db, logger, apiVersion1)

	contractVehicle.Register(e, db, logger, apiVersion1)

	agency.Register(e, db, logger, apiVersion1)

	office.Register(e, db, logger, apiVersion1)

	document.Register(e, db, logger, apiVersion1)

	opportunity.Register(e, db, logger, apiVersion1)

	siteInfo.Register(e, db, logger, apiVersion1)

	wishlist.Register(e, db, logger, apiVersion1)

	bank.Register(e, db, logger, notifierV2, apiVersion1)

	order.Register(e, db, logger, apiVersion1)

	calculators.Register(e, db, logger, apiVersion1)

	notification.Register(e, db, logger, notifierV2, apiVersion1)

	SWAGGER_UI_PATH := "apidocs"
	msg := make(chan error)

	if Protocol == "https" {
		ExternalPort = "443"
	}

	go func() {
		fmt.Printf(
			"listening on %s://%s:%s\n",
			Protocol, HostUri, ExternalPort,
		)
		fmt.Printf(
			"Swagger UI accessible at %s://%s:%s/%s\n", Protocol, HostUri, ExternalPort, SWAGGER_UI_PATH,
		)
		msg <- e.Start(HostAddress + ":" + InternalPort)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		msg <- fmt.Errorf("%s", <-c)
	}()
	logger.Error("chan err : ", <-msg)
	os.Exit(1)
}
