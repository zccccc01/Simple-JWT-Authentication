package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"` // 确保 ID 是主键
	Tel      string `json:"tel" gorm:"unique"`    // 电话号码唯一
	Password string `json:"password"`             // 可以打个-标签忽略掉,不返回,但是忽略了post传不进去
	Name     string `json:"name"`
}
