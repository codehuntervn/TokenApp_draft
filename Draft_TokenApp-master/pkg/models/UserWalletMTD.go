package models

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var dbUserWalletMTD *gorm.DB

type UserWalletMTD struct {
	gorm.Model
	UserAddress  string `gorm:""json:"userAddress"`
	Type         string `json:"type"`
	Nonce        string `json: "nonce"` //mấy cái state con nữa
	TokenExpired string `json:"tokenExpired"`
}

func init() {
	config.Connect()
	dbUserWalletMTD = config.GetDB()
	dbUserWalletMTD.AutoMigrate(&UserWalletMTD{})
}

func (b *UserWalletMTD) CreateUserWalletMTD() *UserWalletMTD {
	dbUserWalletMTD.NewRecord(b)
	dbUserWalletMTD.Create(&b)
	return b
}

func GetAllUserWalletMTDs() []UserWalletMTD {
	var UserWalletMTDs []UserWalletMTD
	dbUserWalletMTD.Find(&UserWalletMTDs)
	return UserWalletMTDs
}

func GetUserWalletMTDById(Id int64) (*UserWalletMTD, *gorm.DB) {
	var getUserWalletMTD UserWalletMTD
	dbUserWalletMTD := dbUserWalletMTD.Where("ID=?", Id).Find(&getUserWalletMTD)
	return &getUserWalletMTD, dbUserWalletMTD
}

func DeleteUserWalletMTD(ID int64) UserWalletMTD {
	var userWalletMTD UserWalletMTD
	dbUserWalletMTD.Where("ID=?", ID).Delete(userWalletMTD)
	return userWalletMTD
}
