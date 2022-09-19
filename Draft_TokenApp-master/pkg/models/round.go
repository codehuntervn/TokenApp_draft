package models

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Round struct {
	gorm.Model
	RoundID      string `gorm:""json:"roundID"`
	RoundName    string `json:"roundName"`
	DefaultPrice string `json:"defaultPrice"`
	Start        string `json: "start"`
	End          string `json: "end"`
	State        string `json: "state"` //mấy cái state con nữa
	Info         string `json: "info"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Round{})
}

func (b *Round) CreateRound() *Round {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllRounds() []Round {
	var Rounds []Round
	db.Find(&Rounds)
	return Rounds
}

func GetRoundById(Id int64) (*Round, *gorm.DB) {
	var getRound Round
	db := db.Where("ID=?", Id).Find(&getRound)
	return &getRound, db
}

func DeleteRound(ID int64) Round {
	var round Round
	db.Where("ID=?", ID).Delete(round)
	return round
}
