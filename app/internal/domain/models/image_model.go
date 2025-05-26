package models

type Image struct {
	Id        int
	ProductId int
	Product   *Product `gorm:"foreignKey:ProductId"`
	Thumbnail string
	Mobile    string
	Tablet    string
	Desktop   string
}

func (i *Image) TableName() string {
	return "images"
}
