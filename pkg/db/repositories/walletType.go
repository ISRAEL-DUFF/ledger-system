package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/types/datastructure"
)

type IWalletTypeRepository interface {
	types.IBaseRepository[IWalletTypeRepository]
	Create(input types.CreateWalletType) (*model.WalletType, error)
	FindById(id string) (*model.WalletType, error)
	GetWalletRulesByTypeId(id string) *WalletType
	FindByOwnerId(id string) ([]*model.WalletType, error)
	UpdateWalletType(id string, typeObj map[string]interface{}) error
}

type WalletTypeRepository struct {
	dbQuery *dao.Query
}

// Wallet type datastructure
type WalletType struct {
	Name  string                 `json:"name" validate:"required"`
	Rules []types.WalletRuleType `json:"rules" validate:"required,dive,required"`
}

func NewWalletType(input map[string]any) *WalletType {
	jsonData, err := json.Marshal(input)

	if err != nil {
		panic("Unable to initialize wallet type")
	}

	var walletType WalletType
	jsonErr := json.Unmarshal(jsonData, &walletType)

	if jsonErr != nil {
		panic("unable to initialize wallet type from json")
	}

	return &walletType
}

func NewWalletTypeRepository() *WalletTypeRepository {
	dbInstance := config.DbInstance().GetDBQuery()
	return &WalletTypeRepository{
		dbQuery: dbInstance,
	}
}

func (walletTypeRepo *WalletTypeRepository) WithTransaction(queryTx types.IDBTransaction) IWalletTypeRepository {
	return &WalletTypeRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (repo *WalletTypeRepository) BeginTransaction() types.IDBTransaction {
	return repo.dbQuery.Begin()
}

func (walletTypeRepo *WalletTypeRepository) Create(input types.CreateWalletType) (*model.WalletType, error) {
	walletTypeData := NewWalletType(input.Rules)

	// TODO: Remove this...
	fmt.Println(walletTypeData)

	dbInstance := walletTypeRepo.dbQuery
	walletType := dbInstance.WalletType.WithContext(context.Background())

	jsonStr, _ := json.Marshal(input.Rules)

	createdWalletType := &model.WalletType{
		Name:    input.Name,
		OwnerID: input.OwnerId,
		Rules:   string(jsonStr),
	}

	if err := walletType.Create(createdWalletType); err != nil {
		return nil, err
	}

	return createdWalletType, nil
}

func (walletTypeRepo *WalletTypeRepository) FindById(id string) (*model.WalletType, error) {
	dbInstance := walletTypeRepo.dbQuery
	WalletType := dbInstance.WalletType.WithContext(context.Background())

	w, err := WalletType.Where(dbInstance.WalletType.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (walletTypeRepo *WalletTypeRepository) FindByOwnerId(id string) ([]*model.WalletType, error) {
	dbInstance := walletTypeRepo.dbQuery
	WalletType := dbInstance.WalletType.WithContext(context.Background())

	w, err := WalletType.Where(dbInstance.WalletType.OwnerID.Eq(id)).Find()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (walletTypeRepo *WalletTypeRepository) GetWalletRulesByTypeId(id string) *WalletType {
	walletTypeModel, err := walletTypeRepo.FindById(id)

	if err != nil {
		panic("unable to get wallet type")
	}

	var walletTypeMap map[string]interface{}

	convErr := json.Unmarshal([]byte(walletTypeModel.Rules), &walletTypeMap)

	if convErr != nil {
		panic("unable to read string rules")
	}

	walletType := NewWalletType(walletTypeMap)

	return walletType

}

func (walletTypeRepo *WalletTypeRepository) UpdateWalletType(id string, typeObj map[string]interface{}) error {
	walletType := NewWalletType(typeObj)
	validatorInstance := validator.New()

	fmt.Println(walletType)

	err := validatorInstance.Struct(walletType)

	if err != nil {
		return err
	}

	err = walletType.Validate()

	if err != nil {
		return err
	}

	walletTypeDbb := walletTypeRepo.dbQuery.WithContext(context.Background())
	walletTypeDb := walletTypeDbb.WalletType

	walletTypeModel, err := walletTypeRepo.FindById(id)

	if err != nil {
		return err
	}

	walletTypeStr, err := json.Marshal(typeObj)

	if err != nil {
		return err
	}

	walletTypeModel.Rules = string(walletTypeStr)
	walletTypeModel.Name = typeObj["name"].(string)

	if updateErr := walletTypeDb.Save(walletTypeModel); updateErr != nil {
		return updateErr
	}

	return nil
}

func (walletType *WalletType) HasEvent(eventName string) bool {
	for _, evt := range walletType.Rules {
		if evt.Event == eventName {
			return true
		}
	}

	return false
}

func (walletType *WalletType) AccountLabels() *datastructure.Set[string] {
	accountSet := datastructure.NewSet[string]()

	for _, evt := range walletType.Rules {
		accountSet.Add(evt.Rule.Credit)
		accountSet.Add(evt.Rule.Debit)
	}

	return accountSet
}

func (walletType *WalletType) Validate() error {
	for _, wt := range walletType.Rules {
		err := wt.Validate()

		if err != nil {
			return err
		}
	}

	return nil
}
