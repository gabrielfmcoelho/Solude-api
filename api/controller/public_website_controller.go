package controller

import (
	"net/http"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gin-gonic/gin"
)

type PublicWebsiteController struct {
	ServiceUsecase domain.ServiceUsecase
}

// GetMarketingServices retorna todos os serviços de marketing
// @Summary Get Marketing Services
// @Description Gets all marketing services
// @Tags Website Service
// @Produce json
// @Success 200 {array} domain.MarketingService
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/marketing [get]
func (sc *PublicWebsiteController) GetMarketingServices(c *gin.Context) {
	services, err := sc.ServiceUsecase.GetMarketing(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}
