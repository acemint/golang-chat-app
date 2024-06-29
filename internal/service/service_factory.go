package service

import (
	repository "chat-app/internal/repository/postgres"

	"gorm.io/gorm"
)

var MemberService *MemberServiceStruct

func InitializeService(db *gorm.DB, mr *repository.MemberRepositoryStruct) {
	MemberService = &MemberServiceStruct{
		db:               db,
		memberRepository: mr,
	}
}
