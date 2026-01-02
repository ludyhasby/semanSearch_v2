package entity

type DoaCollection struct {
	ID       uint   `gorm:"primary_key;column:id"`
	Name     string `gorm:"column:name"`
	Doa      string `gorm:"column:doa"`
	Meaning  string `gorm:"column:meaning"`
	Guidence string `gorm:"column:guidence"`
	Virtue   string `gorm:"column:virtue"`
	Source   string `gorm:"column:source"`
}

func (t DoaCollection) TableName() string {
	return "doa_collection"
}
