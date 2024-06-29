package domain

type UpdateableMember struct {
	Email    string `gorm:"column:email;not null"`
	Name     string `gorm:"column:name;not null"`
	Age      int    `gorm:"column:age;not null"`
	Gender   string `gorm:"column:gender;not null"`
	Password string `gorm:"column:password;not null"`
}
