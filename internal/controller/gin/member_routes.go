package controller

import (
	"errors"
	"log"
	"net/http"

	"chat-app/domain"
	"chat-app/dto"
	"chat-app/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	baseMemberPath        = "/member"
	rootMemberPath        = ""
	createTransactionPath = "/transaction"
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"value": result})
	})

	memberRoutes.POST(createTransactionPath, func(ctx *gin.Context) {
		ctxCopy := ctx.Copy()
		var createTransactionRequest dto.CreateTransactionRequest
		if err := ctx.ShouldBind(&createTransactionRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sender, receiver, transactionData, err := convertToTransactionComponents(&createTransactionRequest, mr.memberService)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		channelErrorCheck := make(chan error, 2)
		go func() {
			err := mr.memberService.IsTransactionOverLimit(sender)
			channelErrorCheck <- err
		}()
		go func() {
			err := mr.memberService.ValidateFraudActivity(sender, receiver, transactionData)
			channelErrorCheck <- err
		}()

		for err := range channelErrorCheck {
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		log.Println("Done! in path " + ctxCopy.Request.URL.Path)
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

// The first return member is sender while the second is receiver
func convertToTransactionComponents(request *dto.CreateTransactionRequest, service *service.MemberServiceStruct) (*domain.Member, *domain.Member, *domain.Transaction, error) {
	sender, err := service.FindMemberById(request.MemberIDSender)
	if err != nil {
		return nil, nil, nil, errors.New("sender not found")
	}
	receiver, err := service.FindMemberById(request.MemberIDReceiver)
	if err != nil {
		return nil, nil, nil, errors.New("receiver not found")
	}
	transaction := &domain.Transaction{
		MemberIDSender:   sender.ID,
		MemberIDReceiver: receiver.ID,
		Amount:           request.Amount,
	}
	return sender, receiver, transaction, nil
}
