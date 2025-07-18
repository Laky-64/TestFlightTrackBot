package models

type Link struct {
	ID     uint       `gorm:"primaryKey, autoIncrement"`
	URL    string     `gorm:"type:varchar(255);not null;unique;index"`
	AppID  *uint      `gorm:"index"`
	Status LinkStatus `gorm:"type:link_status_enum;index;"`
	BaseModel

	App App `gorm:"foreignKey:AppID;references:ID;constraint:OnDelete:CASCADE"`
}
