package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"` // 确保 ID 是主键
	Tel      string `json:"tel" gorm:"unique"`    // 电话号码唯一
	Password string `json:"password"`             
	Name     string `json:"name"`
}
