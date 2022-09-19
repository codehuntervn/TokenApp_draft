package models

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var dbPack *gorm.DB

type Pack struct {
	gorm.Model
	PackID  string `gorm:""json:"packID"`
	Volume  string `json:"volume"`
	State   string `json: "state"` //mấy cái state con nữa
	RoundID string `json:"roundID"`
}

func init() {
	config.Connect()
	dbPack = config.GetDB()
	dbPack.AutoMigrate(&Pack{})
}

func (b *Pack) CreatePack() *Pack {
	dbPack.NewRecord(b)
	dbPack.Create(&b)
	return b
}

func GetAllPacks() []Pack {
	var Packs []Pack
	dbPack.Find(&Packs)
	return Packs
}

func GetPackById(Id int64) (*Pack, *gorm.DB) {
	var getPack Pack
	dbPack := dbPack.Where("ID=?", Id).Find(&getPack)
	return &getPack, dbPack
}

func DeletePack(ID int64) Pack {
	var pack Pack
	dbPack.Where("ID=?", ID).Delete(pack)
	return pack
}
