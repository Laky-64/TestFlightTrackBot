package models

type App struct {
	ID      uint   `gorm:"primaryKey;not null"`
	AppName string `gorm:"type:varchar(255);not null;unique;index"`
	BaseModel
}
