package controller

import (
	"errors"

	"chat-app/domain"
	repository "chat-app/internal/repository/postgres"

	"gorm.io/gorm"
)

var MemberService *MemberServiceStruct

type MemberServiceInterface interface {
	CreateMember(*domain.Member)
}

type MemberServiceStruct struct {
	db               *gorm.DB
	memberRepository *repository.MemberRepositoryStruct
}

func (s *MemberServiceStruct) CreateMember(member *domain.Member) (*domain.Member, error) {
	s.db.Transaction(func(tx *gorm.DB) error {
		_, err := s.memberRepository.FindActiveMember(member.Email)
		if err == nil {
			return errors.New("email already taken")
		}
		s.memberRepository.CreateMember(member)
		return nil
	})
	return s.memberRepository.FindActiveMember(member.Email)
}
