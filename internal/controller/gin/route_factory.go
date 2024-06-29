package controller

import (
	"chat-app/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	baseApiPath = "/v1"
)

var routerGroupV1 *gin.RouterGroup

func InitializeRoutes(r *gin.Engine, ms *service.MemberServiceStruct) {
	routerGroupV1 = r.Group(baseApiPath)
	memberRouter := &MemberRouterStruct{
		memberService: ms,
	}
	memberRouter.addRoutesToGroup(routerGroupV1)
}
