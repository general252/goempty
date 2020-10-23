package model

// MUser 用户
type MUser struct {
	Id       int64
	UserName string `gorm:"column:user_name;type:varchar(64);uniqueIndex;not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}
