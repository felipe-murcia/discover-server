package models

type Feature struct {
	ID    int32  `gorm:"primaryKey;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Icon  string `gorm:"column:icon" json:"icon"`
	State bool   `gorm:"column:state" json:"state"`
}

func (Feature) TableName() string {
	return "features"
}
