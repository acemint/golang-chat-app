package service

import (
	repository "chat-app/internal/repository/postgres"
)

func InitializeService() {
	MemberService = &MemberServiceStruct{
		db:               repository.DB,
		memberRepository: repository.MemberRepository,
	}
}
