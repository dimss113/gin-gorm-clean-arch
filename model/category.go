package model

type Category struct {
	ID   uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	BaseModel
}
