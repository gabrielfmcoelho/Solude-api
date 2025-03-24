package route

import (
	"time"

	"github.com/gabrielfmcoelho/platform-core/api/controller"
	"github.com/gabrielfmcoelho/platform-core/bootstrap"
	"github.com/gabrielfmcoelho/platform-core/repository"
	"github.com/gabrielfmcoelho/platform-core/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPublicWebsiteRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	sr := repository.NewServiceRepository(db)
	uslr := repository.NewUserServiceLogRepository(db)
	pwc := &controller.PublicWebsiteController{
		ServiceUsecase: usecase.NewServiceUsecase(sr, uslr, timeout),
	}

	group.GET("/services/marketing", pwc.GetMarketingServices)
}
