package models

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"not null"`
	Priority  string `gorm:"default:'p2'"` // 默认优先级p2
	DueDate   string
	Completed bool `gorm:"default:false"` // 默认未完成
	ListID    uint `gorm:"not null"`
	List      List `gorm:"foreignKey:ListID"` // 关联清单
}
