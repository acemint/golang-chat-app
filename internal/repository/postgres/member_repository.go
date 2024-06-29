package repository

import (
	"chat-app/domain"

	"gorm.io/gorm"
)

type MemberRepositoryStruct struct {
	db *gorm.DB
}

func (r *MemberRepositoryStruct) FindSingleActiveMember(email string) (*domain.Member, error) {
	var members []domain.Member
	result := r.db.Where(&domain.Member{Email: email}).Limit(1).Find(&members)
	if len(members) == 0 {
		return nil, nil
	}
	return &members[0], result.Error
}

func (r *MemberRepositoryStruct) CreateMember(member *domain.Member) (*domain.Member, error) {
	result := r.db.Create(member)
	return member, result.Error
}

func (r *MemberRepositoryStruct) DeleteMember(email string) error {
	result := r.db.Delete(&domain.Member{Email: email})
	return result.Error
}

func (r *MemberRepositoryStruct) UpdateMember(member *domain.Member, updateableMemberData *domain.UpdateableMember) (*domain.Member, error) {
	member.Email = updateableMemberData.Email
	member.Name = updateableMemberData.Name
	member.Age = updateableMemberData.Age
	member.Gender = updateableMemberData.Gender
	member.Password = updateableMemberData.Password

	result := r.db.Save(member)
	return member, result.Error
}
