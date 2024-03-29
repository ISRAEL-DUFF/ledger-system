package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IBlockMetumRepository interface {
	types.IBaseRepository[IBlockMetumRepository]
	Create(input types.CreateBlockMetum) (*model.BlockMetum, error)
	FindById(id string) (*model.BlockMetum, error)
	Update(data *model.BlockMetum) error
	FindAllBlockMetaInBetween(startDate, endDate int64, accountId string) ([]*model.BlockMetum, error)
}

type BlockMetum struct {
	dbQuery *dao.Query
}

func NewBlockMetumRepository() *BlockMetum {
	dbInstance := config.DbInstance().GetDBQuery()

	return &BlockMetum{
		dbQuery: dbInstance,
	}
}

func (blockMetumRepo *BlockMetum) WithTransaction(queryTx types.IDBTransaction) IBlockMetumRepository {
	return &BlockMetum{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (blockMetumRepo *BlockMetum) BeginTransaction() types.IDBTransaction {
	return blockMetumRepo.dbQuery.Begin()
}

func (blockMetumRepo *BlockMetum) Create(input types.CreateBlockMetum) (*model.BlockMetum, error) {
	dbInstance := blockMetumRepo.dbQuery
	blockMetum := dbInstance.BlockMetum.WithContext(context.Background())

	createdBlockMetum := &model.BlockMetum{
		AccountID:      input.AccountId,
		BlockTxLimit:   int32(input.BlockTxLimit),
		TransitionTxID: input.TransitionTxId,
		OpeningDate:    input.OpeningDate,
		ClosingDate:    input.ClosingDate,
	}

	if err := blockMetum.Create(createdBlockMetum); err != nil {
		return nil, err
	}

	return createdBlockMetum, nil
}

func (blockMetumRepo *BlockMetum) FindById(id string) (*model.BlockMetum, error) {
	dbInstance := blockMetumRepo.dbQuery
	blockMetum := dbInstance.BlockMetum.WithContext(context.Background())

	blockMetumData, err := blockMetum.Where(dbInstance.BlockMetum.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return blockMetumData, nil
}

func (blockMetumRepo *BlockMetum) Update(data *model.BlockMetum) error {
	if data.ID == "" {
		return errors.New("can't update block meta without primary ID")
	}

	dbInstance := blockMetumRepo.dbQuery
	blockMetum := dbInstance.BlockMetum.WithContext(context.Background())

	if err := blockMetum.Save(data); err != nil {
		return err
	}

	return nil
}

func (blockMetumRepo *BlockMetum) FindAllBlockMetaInBetween(startDate, endDate int64, accountId string) ([]*model.BlockMetum, error) {
	dbInstance := blockMetumRepo.dbQuery
	blockMetum := dbInstance.BlockMetum.WithContext(context.Background())

	var err error

	// TODO: verify this query
	blockMeta, err := blockMetum.Where(dbInstance.BlockMetum.AccountID.Eq(accountId), dbInstance.BlockMetum.OpeningDate.Gt(time.UnixMilli(startDate)), dbInstance.BlockMetum.ClosingDate.Lt(time.UnixMilli(endDate))).Order(dbInstance.BlockMetum.CreatedAt).Find()

	if err != nil {
		return nil, err
	}

	return blockMeta, nil
}
