package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/models"
	notificationv2 "github.com/Capture-411/services-opportunity/notification-v2"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx context.Context, input LoginRequest) (response LoginResponseData, err response.ErrorResponse)
	Register(ctx context.Context, input RegisterRequest) (response LoginResponseData, err response.ErrorResponse)
	ResetPassword(ctx context.Context, input ResetPasswordRequest) (response LoginResponseData, err response.ErrorResponse)
	ChangePassword(ctx context.Context, input ChangePasswordRequest) (
		response LoginResponseData, err response.ErrorResponse,
	)
	UserInfo(ctx context.Context, accessToken string) (user models.User, err response.ErrorResponse)
	ApprovePhone(ctx context.Context, code string) (ok string, err response.ErrorResponse)
	ApproveEmail(ctx context.Context, token string) (ok string, err response.ErrorResponse)
	authenticate(ctx context.Context, username, password string) Identity
	generateAccessTokenJWT(identity Identity) (string, error)
	checkAndDeleteVerificationBySessionCodeAndPhone(ctx context.Context, sessionCode string, phone string) (phn string, err response.ErrorResponse)
	checkAndDeleteVerificationBySessionCodeAndEmail(ctx context.Context, sessionCode string, email string) (string, response.ErrorResponse)
	checkAndDeleteVerificationByCode(ctx context.Context, code string) (phone string, err response.ErrorResponse)
}

type Identity interface {
	GetID() string
	GetFullName() string
	GetPhone() string
	GetRoles() []string
}

type service struct {
	db                     *gorm.DB
	logger                 log.Logger
	notifierv2             notificationv2.Notifier
	accessTokenSigningKey  string
	refreshTokenSigningKey string
	tokenExpiration        int
	refreshTokenExpiration int
}

func MakeService(
	db *gorm.DB, logger log.Logger, notifierv2 notificationv2.Notifier,
	accessTokenSigningKey, refreshTokenSigningKey string, accessTokenExpiration, refreshTokenExpiration int,
) Service {

	return &service{
		db,
		logger,
		notifierv2,
		accessTokenSigningKey,
		refreshTokenSigningKey,
		accessTokenExpiration,
		refreshTokenExpiration,
	}
}
