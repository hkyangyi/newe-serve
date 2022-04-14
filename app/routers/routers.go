package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	api := r.Group("api")

	admin := api.Group("admin")

	AdminRouter(admin)
	return
}
