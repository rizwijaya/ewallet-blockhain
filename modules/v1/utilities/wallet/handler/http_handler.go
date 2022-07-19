package product

import (
	"encoding/json"
	"ewallet-blockhain/app/blockhain"
	api "ewallet-blockhain/app/contracts"
	"ewallet-blockhain/modules/v1/utilities/wallet/repository"
	"ewallet-blockhain/modules/v1/utilities/wallet/service"
	"net/http"
	"strconv"

	apiResponse "ewallet-blockhain/pkg/api_response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WalletHandler interface {
	GetBalance(c *gin.Context)
}

type walletHandler struct {
	walletService service.Service
}

func NewWalletHandler(walletService service.Service) *walletHandler {
	return &walletHandler{walletService}
}

func Handler(db *gorm.DB, blockhain *api.Api) *walletHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository, blockhain)
	Handler := NewWalletHandler(Service)
	return Handler
}

func (h *walletHandler) GetBalance(c *gin.Context) {
	balance, err := h.walletService.GetBalance()
	if err != nil {
		response := apiResponse.APIResponse("Failed to get Balance", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := apiResponse.APIResponse("Successfully get Balance", http.StatusOK, "success", `{"balance": "`+balance.String()+`"}`)
	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) GetMyWallet(c *gin.Context) {
	mywallet, err := h.walletService.GetMyWallet()
	if err != nil {
		response := apiResponse.APIResponse("Failed to get wallet address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := apiResponse.APIResponse("Successfully get wallet address", http.StatusOK, "success", `{"address": "`+mywallet.String()+`"}`)
	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) Deposite(c *gin.Context) {
	amount := c.Param("amount")
	am, _ := strconv.Atoi(amount)
	var v map[string]interface{}
	//privateKey := c.Request.Header
	err := json.NewDecoder(c.Request.Body).Decode(&v)
	if err != nil {
		response := apiResponse.APIResponse("Failed wallet address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	auth := blockhain.GetAccountAuth(blockhain.Connect(), v["privateKey"].(string))
	deposite, err := h.walletService.Deposite(am, auth)
	if err != nil {
		response := apiResponse.APIResponse("Failed to deposit", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := apiResponse.APIResponse("Successfully to deposit", http.StatusOK, "success", deposite)
	c.JSON(http.StatusOK, response)
}

func (h *walletHandler) Withdraw(c *gin.Context) {
	amount := c.Param("amount")
	am, _ := strconv.Atoi(amount)
	var v map[string]interface{}
	//privateKey := c.Request.Header
	err := json.NewDecoder(c.Request.Body).Decode(&v)
	if err != nil {
		response := apiResponse.APIResponse("Failed wallet address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	auth := blockhain.GetAccountAuth(blockhain.Connect(), v["privateKey"].(string))
	withdraw, err := h.walletService.Withdraw(am, auth)
	if err != nil {
		response := apiResponse.APIResponse("Failed to Withdraw", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := apiResponse.APIResponse("Successfully to Withdraw", http.StatusOK, "success", withdraw)
	c.JSON(http.StatusOK, response)
}
