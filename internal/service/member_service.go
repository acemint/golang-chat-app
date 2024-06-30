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

func (s *MemberServiceStruct) FindMemberById(id string) (*domain.Member, error) {
	return s.memberRepository.FindSingleActiveMemberByID(id)
}

// Say this is a dummy function where we want to make sure that the member has a certain transaction limit
func (s *MemberServiceStruct) IsTransactionOverLimit(sender *domain.Member) error {
	return nil
}

// Say this is also a dummy function to make sure to check whether the current transaction being made is suspicious
func (s *MemberServiceStruct) ValidateFraudActivity(member *domain.Member, receiver *domain.Member, transaction *domain.Transaction) error {
	return nil
}
