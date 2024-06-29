package controller

import (
	"net/http"

	"chat-app/domain"
	"chat-app/dto"
	"chat-app/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	baseMemberPath = "/member"
	rootMemberPath = "/"
)

type MemberRouterStruct struct {
	memberService *service.MemberServiceStruct
}

func (mr *MemberRouterStruct) addRoutesToGroup(rg *gin.RouterGroup) {
	memberRoutes := rg.Group(baseMemberPath)

	memberRoutes.POST(rootMemberPath, func(ctx *gin.Context) {
		var createMemberRequest dto.CreateMemberRequest
		if err := ctx.ShouldBind(&createMemberRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		member := convertToMember(&createMemberRequest)
		result, err := mr.memberService.CreateMember(member)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	})

}

func convertToMember(request *dto.CreateMemberRequest) *domain.Member {
	member := &domain.Member{
		Name:     request.Name,
		Email:    request.Email,
		Gender:   request.Gender,
		Password: request.Password,
	}
	return member
}
