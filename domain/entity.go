package domain

import (
	"time"

	"gorm.io/gorm"
)

func (Member) TableName() string {
	return "ca_member"
}

type Member struct {
	ID        string         `gorm:"column:id;type:uuid;primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	Email     string         `gorm:"column:email;not null"`
	Name      string         `gorm:"column:name;not null"`
	Age       int            `gorm:"column:age;not null"`
	Gender    string         `gorm:"column:gender;not null"`
	Password  string         `gorm:"column:password;not null"`
}

func (Transaction) TableName() string {
	return "ca_transaction"
}

type Transaction struct {
	ID               string         `gorm:"column:id;type:uuid;primaryKey"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	MemberIDSender   string         `gorm:"column:member_id_sender;type:uuid;not null"`
	MemberIDReceiver string         `gorm:"column:member_id_receiver;type:uuid;not null"`
	SentAt           time.Time      `gorm:"column:sent_at;not null"`
	Amount           int            `gorm:"column:amount;not null"`
}
