package repository

import "gorm.io/gorm"

type MemberRepositoryStruct struct {
	db *gorm.DB
}

func (r *MemberRepositoryStruct) FindActiveMember(email string) (*Member, error) {
	var member Member
	result := r.db.Where(&Member{Email: email}).First(&member)
	return &member, result.Error
}

func (r *MemberRepositoryStruct) CreateMember(member *Member) (*Member, error) {
	result := r.db.Create(member)
	return member, result.Error
}

func (r *MemberRepositoryStruct) DeleteMember(email string) error {
	result := r.db.Delete(&Member{Email: email})
	return result.Error
}

func (r *MemberRepositoryStruct) UpdateMember(member *Member, updateableMemberData *UpdateableMember) (*Member, error) {
	member.Email = updateableMemberData.Email
	member.Name = updateableMemberData.Name
	member.Age = updateableMemberData.Age
	member.Gender = updateableMemberData.Gender
	member.Password = updateableMemberData.Password

	result := r.db.Save(member)
	return member, result.Error
}
