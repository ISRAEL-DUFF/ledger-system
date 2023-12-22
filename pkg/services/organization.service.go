package services

import (
	"errors"
	"fmt"

	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

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

func (orgService *OrganizationService) GenerateAPIKey(organizationId string) (*types.CreateAPIKEYResponse, error) {
	record, _ := orgService.ApikRepository.FindByOrgId(organizationId)
	if record != nil {
		return nil, errors.New("api key already exists for the organization")
	}

	testKeys, err := orgService.generateKeys(types.TEST_KEY)
	if err != nil {
		return nil, err
	}

	liveKeys, err := orgService.generateKeys(types.LIVE_KEY)
	if err != nil {
		return nil, err
	}

	_, err = orgService.ApikRepository.Create(types.CreateAPIKEY{
		OrganizationID: organizationId,
		TestSecretKey:  testKeys.Secretkey,
		TestPublicKey:  testKeys.PublicKey,
		LivePublicKey:  liveKeys.PublicKey,
		LiveSecretKey:  record.LiveSecretKey,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateAPIKEYResponse{
		TestKeys: *testKeys,
		LiveKeys: *liveKeys,
	}, nil
}

func (orgService *OrganizationService) ReGenerateAPIKey(organizationId string, keyType types.API_KEY_TYPE) (*types.APIKEY, error) {
	record, err := orgService.ApikRepository.FindByOrgId(organizationId)
	if err != nil {
		return nil, errors.New("api key record does not exist for the organization")
	}

	var keys *types.APIKEY

	if keyType == types.LIVE_KEY {
		keys, err = orgService.generateKeys(types.LIVE_KEY)
		record.LivePublicKey = keys.PublicKey
		record.LiveSecretKey = keys.Secretkey
	} else {
		keys, err = orgService.generateKeys(types.TEST_KEY)
		record.TestPublicKey = keys.PublicKey
		record.TestPublicKey = keys.Secretkey
	}

	if err = orgService.ApikRepository.Update(record); err != nil {
		return nil, err
	}

	return keys, nil
}

func (orgService *OrganizationService) GetAPIKeys(organizationId string) (*types.CreateAPIKEYResponse, error) {
	record, err := orgService.ApikRepository.FindByOrgId(organizationId)
	if err != nil {
		return nil, errors.New("no api key available!!!")
	}

	return &types.CreateAPIKEYResponse{
		TestKeys: types.APIKEY{
			PublicKey: record.TestPublicKey,
			Secretkey: record.TestSecretKey,
		},
		LiveKeys: types.APIKEY{
			PublicKey: record.LivePublicKey,
			Secretkey: record.LiveSecretKey,
		},
	}, nil
}

func (orgService *OrganizationService) CreateOrganization(input types.CreateOrganizationDto) (string, error) {
	organizationOpts := types.CreateOrganization{
		Name:    input.Name,
		Address: input.Address,
	}

	if input.EmailAddress != "" && input.PhoneNumber != "" {
		user, err := orgService.UserRepository.FindByEmailAndPhone(input.EmailAddress, input.PhoneNumber)
		if err != nil {
			return "", err
		}

		organizationOpts.EmailAddress = input.EmailAddress
		organizationOpts.PhoneNumber = input.PhoneNumber
		organizationOpts.OwnerID = user.ID
	} else if input.EmailAddress != "" {
		user, err := orgService.UserRepository.FindByEmail(input.EmailAddress)
		if err != nil {
			return "", err
		}

		organizationOpts.EmailAddress = input.EmailAddress
		organizationOpts.OwnerID = user.ID
	} else if input.PhoneNumber != "" {
		user, err := orgService.UserRepository.FindByPhoneNumber(input.PhoneNumber)
		if err != nil {
			return "", err
		}

		organizationOpts.PhoneNumber = input.PhoneNumber
		organizationOpts.OwnerID = user.ID
	} else {
		return "", errors.New("invalid email or phone number")
	}

	organization, err := orgService.OrganizationRepository.Create(organizationOpts)
	if err != nil {
		return "", err
	}

	return organization.ID, nil
}

func (orgService *OrganizationService) GetOrganizationByOwnerIdentfier(identifier string) (*model.Organization, error) {
	user, err := orgService.UserRepository.FindByEmail(identifier)
	if err != nil {
		user, err = orgService.UserRepository.FindByPhoneNumber(identifier)
		if err != nil {
			return nil, err
		}
	}

	organization, err := orgService.OrganizationRepository.FindByOwnerId(user.ID)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (orgService *OrganizationService) GetOrganizationById(id string) (*model.Organization, error) {
	organization, err := orgService.OrganizationRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (orgService *OrganizationService) generateKeys(keyType types.API_KEY_TYPE) (*types.APIKEY, error) {
	publicApiKeyGenerator := utils.NewApiKeyGenerator()
	secretApiKeyGenerator := utils.NewApiKeyGenerator()

	pKey, err := publicApiKeyGenerator()
	if err != nil {
		return nil, err
	}

	sKey, err := secretApiKeyGenerator()
	if err != nil {
		return nil, err
	}

	publicPrefix := "pk_test_"
	privatePrefix := "sk_test_"

	if keyType == types.LIVE_KEY {
		publicPrefix = "pk_live_"
		privatePrefix = "sk_live_"
	}

	return &types.APIKEY{
		PublicKey: fmt.Sprintf("%s%s", publicPrefix, pKey),
		Secretkey: fmt.Sprintf("%s%s", privatePrefix, sKey),
	}, nil
}
