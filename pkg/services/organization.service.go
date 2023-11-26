package services

import "github.com/israel-duff/ledger-system/pkg/db/repositories"

type IOrganizationService interface {
}

type OrganizationService struct {
	OrganizationRepository repositories.IOrganizationRepository
	UserRepository         repositories.IUserRepository
	ApikRepository         repositories.IApikeyRepository
}

func NewOrganizationService(orgRepo repositories.IOrganizationRepository, userRepo repositories.IUserRepository, apiRepo repositories.IApikeyRepository) *OrganizationService {
	return &OrganizationService{
		OrganizationRepository: orgRepo,
		UserRepository:         userRepo,
		ApikRepository:         apiRepo,
	}
}

func (orgService *OrganizationService) GenerateAPIKey(organizationId string) {

}
