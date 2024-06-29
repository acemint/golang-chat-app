package service

import (
	"chat-app/domain"
	repository "chat-app/internal/repository/postgres"
	"errors"

	"gorm.io/gorm"
)

type MemberServiceStruct struct {
	db               *gorm.DB
	memberRepository *repository.MemberRepositoryStruct
}

func (s *MemberServiceStruct) CreateMember(member *domain.Member) (*domain.Member, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		existingMember, err := s.memberRepository.FindSingleActiveMember(member.Email)
		if err != nil {
			return err
		}
		if existingMember != nil {
			return errors.New("email has been taken")
		}

		if _, err := s.memberRepository.CreateMember(member); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.memberRepository.FindSingleActiveMember(member.Email)
}
