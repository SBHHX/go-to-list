package models

type List struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"` // 关联用户
	Tasks  []Task `gorm:"foreignKey:ListID"` // 关联任务
}
