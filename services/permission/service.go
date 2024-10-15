package permission

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/policy"
	"gorm.io/gorm"
)

type Service interface {
	AccessList(ctx context.Context) (acl map[string]bool, err error)
}

type service struct {
	db     *gorm.DB
	logger log.Logger
}

func MakeService(db *gorm.DB, logger log.Logger) Service {
	return &service{db, logger}
}

func (s *service) AccessList(ctx context.Context) (map[string]bool, error) {
	acl := map[string]bool{
		"CanAccessAgency":          policy.CanAccessAgency(ctx),
		"CanCreateAgency":          policy.CanCreateAgency(ctx),
		"CanAccessCategory":        policy.CanAccessCategory(ctx),
		"CanCreateCategory":        policy.CanCreateCategory(ctx),
		"CanAccessDepartment":      policy.CanAccessDepartment(ctx),
		"CanCreateDepartment":      policy.CanCreateDepartment(ctx),
		"CanAccessFiscalYear":      policy.CanAccessFiscalYear(ctx),
		"CanCreateFiscalYear":      policy.CanCreateFiscalYear(ctx),
		"CanAccessOpportunity":     policy.CanAccessOpportunity(ctx),
		"CanCreateOpportunity":     policy.CanCreateOpportunity(ctx),
		"CanCreateContractVehicle": policy.CanCreateContractVehicle(ctx),
		"CanAccessInvite":          policy.CanAccessInvite(ctx),
		"CanAccessMarket":          policy.CanAccessMarket(ctx),
		"CanCreateMarket":          policy.CanCreateMarket(ctx),
		"CanAccessNaics":           policy.CanAccessNaics(ctx),
		"CanCreateNaics":           policy.CanCreateNaics(ctx),
		"CanAccessOffice":          policy.CanAccessOffice(ctx),
		"CanCreateOffice":          policy.CanCreateOffice(ctx),
		"CanAccessSetAside":        policy.CanAccessSetAside(ctx),
		"CanCreateSetAside":        policy.CanCreateSetAside(ctx),
		"CanAccessContractVehicle": policy.CanAccessContractVehicle(ctx),
		"CanAccessUser":            policy.CanAccessUser(ctx),
		"CanCreateUser":            policy.CanCreateUser(ctx),
	}
	return acl, nil
}
