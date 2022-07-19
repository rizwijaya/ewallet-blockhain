package product

import (
	api "ewallet-blockhain/app/contracts"
	"ewallet-blockhain/modules/v1/utilities/wallet/repository"
	"ewallet-blockhain/modules/v1/utilities/wallet/service"
	"net/http"

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

	response := apiResponse.APIResponse("Successfully get Balance", http.StatusOK, "success", balance)
	c.JSON(http.StatusOK, response)
}
