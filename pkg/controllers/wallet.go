package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/services"
	"github.com/israel-duff/ledger-system/pkg/types"
	httpUtil "github.com/israel-duff/ledger-system/pkg/utils/httpUtil"
)

type CreateWalletTypePayload struct {
	Name    string `json:"name"`
	OwnerId string `json:"ownerId"`
}

type UpdateWalletTypePayload struct {
	TypeId string                 `json:"id"`
	Rules  map[string]interface{} `json:"rules"`
}

type CreateWalletPayload struct {
	WalletType string `json:"walletType"`
	OwnerId    string `json:"ownerId"`
}

type WalletDto struct {
	Accounts   any     `json:"accounts"`
	WalletType any     `json:"type"`
	Balance    float32 `json:"balance"`
}

type WalletController struct {
	walletService      services.IAccountService
	transactionService services.TransactionService
	validatorInstance  *validator.Validate
}

func NewWalletController() *WalletController {
	coaRepo := repositories.NewChartOfAccountRepository()
	accountBlockRepo := repositories.NewAccountBlockRepository()
	ledgerAccountRepo := repositories.NewLedgerAccountRepository()
	journalEntry := repositories.NewJournalEntryRepository()
	walletRepo := repositories.NewWalletRepository()
	walletTypeRepo := repositories.NewWalletTypeRepository()
	transactionRepo := repositories.NewLedgerTransactionRepository()
	blockMetumRepo := repositories.NewBlockMetumRepository()
	transactionQService := services.NewTransactionQService()

	coaService := services.NewCoaService(coaRepo)
	walletService := services.NewAccountService(coaService, accountBlockRepo, ledgerAccountRepo, journalEntry, walletRepo, walletTypeRepo)
	transactionService := services.NewTransactionService(transactionRepo, journalEntry, accountBlockRepo, blockMetumRepo, transactionQService, walletService)

	return &WalletController{
		walletService:      walletService,
		transactionService: *transactionService,
		validatorInstance:  validator.New(),
	}
}

func (walletController *WalletController) CreateWallet(c *gin.Context) {
	var payload CreateWalletPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Invalid payload!!!")
		return
	}

	r := walletController.walletService.CreateWallet(payload.OwnerId, payload.WalletType)

	// if err != nil {
	// 	httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	httpUtil.SuccessResponseWithData(c, http.StatusCreated, WalletDto{
		Accounts:   r.GetAccounts(),
		WalletType: r.GetWalletType(),
	})
}

func (walletController *WalletController) CreateWalletType(c *gin.Context) {
	var payload CreateWalletTypePayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Invalid payload!!!")
		return
	}

	r, err := walletController.walletService.CreateWalletType(payload.OwnerId, payload.Name)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	httpUtil.SuccessResponseWithData(c, http.StatusCreated, r)
}

func (walletController *WalletController) UpdateWalletType(c *gin.Context) {
	var payload UpdateWalletTypePayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println(err)
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Invalid payload!!!")
		return
	}

	err := walletController.walletService.UpdateWalletType(payload.TypeId, payload.Rules)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	httpUtil.SuccessResponseWithMessage(c, http.StatusCreated, "Wallet Type Updated!!!")
}

func (walletController *WalletController) GetWalletByAccountNumber(c *gin.Context) {
	accountNumber, exists := c.Params.Get("accountNumber")

	if !exists {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "unknown account number")
		return
	}

	wallet, err := walletController.walletService.GetWalletByAccountNumber(accountNumber)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	httpUtil.SuccessResponseWithData(c, http.StatusAccepted, WalletDto{
		Accounts:   wallet.GetAccounts(),
		WalletType: wallet.GetWalletType(),
		Balance:    wallet.GetBalance(),
	})

}

func (walletController *WalletController) ListUserWalllets(c *gin.Context) {
	ownerId, exists := c.Params.Get("ownerId")

	if !exists {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "unknown ownerId")
		return
	}

	wallets, err := walletController.walletService.ListUserWallets(ownerId)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	walletDtoList := []WalletDto{}

	for _, w := range wallets {
		walletDtoList = append(walletDtoList, WalletDto{
			Accounts:   w.GetAccounts(),
			WalletType: w.GetWalletType(),
			Balance:    w.GetBalance(),
		})
	}

	httpUtil.SuccessResponseWithData(c, http.StatusAccepted, walletDtoList)

}

func (walletController *WalletController) GetLedgerAccount(c *gin.Context) {
	accountNumber, exists := c.Params.Get("accountNumber")

	if !exists {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "unknown account number")
		return
	}

	account := walletController.walletService.GetAccount(accountNumber)

	// if err != nil {
	// 	httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	httpUtil.SuccessResponseWithData(c, http.StatusAccepted, account)

}

func (walletController *WalletController) ListWalletTypes(c *gin.Context) {
	ownerId, exists := c.Params.Get("ownerId")

	if !exists {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "unknown owner ID")
		return
	}

	list, err := walletController.walletService.ListWalletTypes(ownerId)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	httpUtil.SuccessResponseWithData(c, http.StatusAccepted, list)
}

func (walletController *WalletController) PostTransaction(c *gin.Context) {
	var payload types.PostTransactionInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Invalid payload!!!")
		return
	}

	fmt.Println(payload)
	err := walletController.validatorInstance.Struct(payload)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err = walletController.transactionService.PostQueuedWalletTransaction(payload)

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	httpUtil.SuccessResponseWithMessage(c, http.StatusAccepted, "Transaction posted")
}

func (walletController *WalletController) GetWalletStatus(c *gin.Context) {
	accountNumber, exists := c.Params.Get("accountNumber")

	if !exists {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "unknown account number")
		return
	}

	account := walletController.walletService.WalletStatus(accountNumber)

	// if err != nil {
	// 	httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	httpUtil.SuccessResponseWithData(c, http.StatusAccepted, account)

}
