package service

import (
	"chat-app/domain"
	repository "chat-app/internal/repository/postgres"

	"gorm.io/gorm"
)

type MemberServiceStruct struct {
	db               *gorm.DB
	memberRepository *repository.MemberRepositoryStruct
}

func (s *MemberServiceStruct) CreateMember(member *domain.Member) (*domain.Member, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		_, err := s.memberRepository.FindActiveMember(member.Email)
		if err != nil {
			return err
		}
		s.memberRepository.CreateMember(member)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.memberRepository.FindActiveMember(member.Email)
}
