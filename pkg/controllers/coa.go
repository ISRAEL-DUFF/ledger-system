package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/services"
	"github.com/israel-duff/ledger-system/pkg/types"
	httpUtil "github.com/israel-duff/ledger-system/pkg/utils/httpUtil"
)

type CreateCoaAccountPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCoaAccountResponse struct {
	ID            string `json:"id"`
	AccountNumber string `json:"accountNumber"`
}

func CreateAccount(c *gin.Context) {
	coaRepo := repositories.NewChartOfAccountRepository()
	coaService := services.NewCoaService(coaRepo)
	var payload CreateCoaAccountPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Invalid payload!!!")
	}

	id, accountNumber := coaService.CreateAccount(types.ASSET, payload.Name, payload.Description)

	if id == "" {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Unable to create account!!!")
	}

	httpUtil.SuccessResponseWithData(c, http.StatusCreated, &CreateCoaAccountResponse{
		ID:            id,
		AccountNumber: accountNumber,
	})
}

func ListAccounts(c *gin.Context) {
	coaRepo := repositories.NewChartOfAccountRepository()
	coaService := services.NewCoaService(coaRepo)

	accounts, err := coaService.ListAll()

	if err != nil {
		httpUtil.ErrorResponseWithMessage(c, http.StatusBadRequest, "Unable to List accounts!!!")
		return
	}

	httpUtil.SuccessResponseWithData(c, http.StatusCreated, accounts)
}
