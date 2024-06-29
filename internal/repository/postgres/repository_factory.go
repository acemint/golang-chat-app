package repository

import (
	"gorm.io/gorm"
)

var MemberRepository *MemberRepositoryStruct

func InitializeRepositories(db *gorm.DB) {
	MemberRepository = &MemberRepositoryStruct{
		db: db,
	}
}
