package models

import "time"

type User struct {
	ID        int       `gorm:"column:id; primary_key; auto_increment" json:"id"`
	Username  string    `gorm:"column:username; type:varchar(200); not null" json:"username"`
	Password  string    `gorm:"column:password; type:varchar(200); not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at; default null" json:"created_at"`
}

func (m *User) TableName() string {
	return "user"
}
