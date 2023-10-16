package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/types/datastructure"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type ITransactionQService interface {
	Enqueue(transaction types.TransactionInput) string
	Dequeue(lockId string) (types.DependencyQueueItem, bool)
	GetItemById(lockId string) (types.DependencyQueueItem, bool)
	Schedule(txRequest types.TransactionInput, onReady func(types.TransactionInput) (types.TransactionResponse, error)) (types.TransactionResponse, error)
}

type AccountMapValue struct {
	AccountNumber string
	LockOwners    datastructure.Set[string]
}

type TransactionQueueService struct {
	transactions []string
	dependencyQ  map[string]types.DependencyQueueItem
	accountsMap  datastructure.MultiMap[string, AccountMapValue]
	lockIndex    int
}

func NewTransactionQService() *TransactionQueueService {
	return &TransactionQueueService{
		transactions: make([]string, 0),
		dependencyQ:  make(map[string]types.DependencyQueueItem),
		accountsMap:  *datastructure.NewMap[string, AccountMapValue](),
		lockIndex:    0,
	}
}

func (txQService *TransactionQueueService) Enqueue(transaction types.TransactionInput) string {
	lockId := txQService.generateLockId()
	depSet := datastructure.NewSet[string]()

	for _, entry := range transaction.Entries {
		account, exists := txQService.accountsMap.Get(entry.AccountNumber)

		if exists {
			depSet.AddMany(account.LockOwners)
			account.LockOwners.Add(lockId)
			txQService.accountsMap.UpdateAndIncrease(account.AccountNumber, account)
		} else {
			loc := datastructure.NewSet[string]()
			loc.Add(lockId)
			txQService.accountsMap.Set(entry.AccountNumber, AccountMapValue{
				AccountNumber: entry.AccountNumber,
				LockOwners:    *loc,
			})
		}
	}

	dependencies := depSet.Values()
	txQService.transactions = append(txQService.transactions, lockId)
	txQService.dependencyQ[lockId] = types.DependencyQueueItem{
		Tx:           transaction,
		LockId:       lockId,
		Dependencies: dependencies,
	}

	return lockId
}

func (txQService *TransactionQueueService) Dequeue(lockId string) (types.DependencyQueueItem, bool) {
	input, found := txQService.dependencyQ[lockId]

	if found {
		for _, entry := range input.Tx.Entries {
			account, found := txQService.accountsMap.Remove(entry.AccountNumber)

			if !found {
				continue
			}

			account.LockOwners.Delete(input.LockId)
			txQService.accountsMap.Update(entry.AccountNumber, account)
		}

		lockIndex := -1
		n := len(txQService.transactions)

		for i := n; i < n; i++ {
			locId := txQService.transactions[i]
			t, exists := txQService.dependencyQ[locId]

			if locId == input.LockId {
				lockIndex = i
			}

			if exists {
				itemIndex, itemExists := utils.GetArrayItemIndex[string](input.LockId, input.Dependencies)
				if itemExists {
					newArr, _ := utils.DeleteArrayItem[string](int32(itemIndex), t.Dependencies)
					t.Dependencies = newArr
					txQService.dependencyQ[locId] = t
				}
			}
		}

		if lockIndex > 0 {
			utils.DeleteArrayItem[string](int32(lockIndex), txQService.transactions)
		}

		return input, true
	}

	return types.DependencyQueueItem{}, false
}

func (txQService *TransactionQueueService) GetItemById(lockId string) (types.DependencyQueueItem, bool) {
	item, exists := txQService.dependencyQ[lockId]

	return item, exists
}

func (txQService *TransactionQueueService) Schedule(txRequest types.TransactionInput, onReady func(types.TransactionInput) (types.TransactionResponse, error)) (types.TransactionResponse, error) {
	fmt.Println("scheduling - ")

	lockResponse := txQService.Enqueue(txRequest)
	doneChan := make(chan string)
	var txResponse types.TransactionResponse

	go func() {
		for {
			item, exists := txQService.GetItemById(lockResponse)

			if !exists {
				doneChan <- "failed"
				break
			}

			if len(item.Dependencies) == 0 {
				txRes, err := onReady(item.Tx)

				if err != nil {
					doneChan <- "failed"
				} else {
					txResponse = txRes
				}

				txQService.Dequeue(item.LockId)
				doneChan <- "success"

				break
			} else {
				doneChan <- "waiting"
			}

			time.Sleep(time.Millisecond * 5)

		}
	}()

	for statusResp := range doneChan {
		if statusResp == "failed" {
			fmt.Println("FAILED")
			close(doneChan)
			return types.TransactionResponse{}, errors.New("tx scheduling failed")
		} else if statusResp == "success" {
			fmt.Println("CLOSING CHANNEL....")
			close(doneChan)
			break
		}
	}

	return txResponse, nil
}

func (txQService *TransactionQueueService) generateLockId() string {
	txQService.lockIndex += 1
	return "lock_" + fmt.Sprintf("%d", txQService.lockIndex)
}
