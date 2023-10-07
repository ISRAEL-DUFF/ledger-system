package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	coaService := services.NewCoaService()

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
	httpUtil.SuccessResponseWithMessage(c, http.StatusCreated, "No Accounts Yet")
}
