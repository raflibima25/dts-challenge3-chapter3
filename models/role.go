package models

type Role struct {
	ID       uint   `gorm:"primaryKey" json:"role_id"`
	RoleName string `gorm:"not null" json:"role_name"`
}
