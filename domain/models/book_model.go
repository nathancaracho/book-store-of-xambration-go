package models

type Book struct {
	ID               uint   `gorm:"primaryKey,autoIncrement"`
	Name             string `gorm:"size:100"`
	Description      string `gorm:"size:500"`
	BriefDescription string `gorm:"size:250"`
	Value            float32
}
