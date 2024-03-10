package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"` // gorm 默认以 ID 作为主键
	Name      string `json:"name" gorm:"type:varchar(50);not null; unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
}
