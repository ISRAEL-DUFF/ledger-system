package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpUtil "github.com/israel-duff/ledger-system/pkg/utils/httpUtil"
)

func RegisterController(c *gin.Context) {
	httpUtil.SuccessResponseWithMessage(c, http.StatusCreated, "User Created")
}

func LoginController(c *gin.Context) {
	httpUtil.SuccessResponseWithMessage(c, http.StatusCreated, "User Signed In")
}
