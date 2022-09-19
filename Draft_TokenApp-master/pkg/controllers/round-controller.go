package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akhil/go-bookstore/pkg/models"
	"github.com/akhil/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewRound models.Round

func GetRound(w http.ResponseWriter, r *http.Request) {
	newRounds := models.GetAllRounds()
	res, _ := json.Marshal(newRounds)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoundById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roundId := vars["roundId"]
	ID, err := strconv.ParseInt(roundId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	roundDetails, _ := models.GetRoundById(ID)
	res, _ := json.Marshal(roundDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateRound(w http.ResponseWriter, r *http.Request) {
	CreateRound := &models.Round{}
	utils.ParseBody(r, CreateRound)
	b := CreateRound.CreateRound()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roundId := vars["roundId"]
	ID, err := strconv.ParseInt(roundId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	round := models.DeleteRound(ID)
	res, _ := json.Marshal(round)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRound(w http.ResponseWriter, r *http.Request) {
	var updateRound = &models.Round{}
	utils.ParseBody(r, updateRound)
	vars := mux.Vars(r)
	roundId := vars["roundId"]
	ID, err := strconv.ParseInt(roundId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	roundDetails, db := models.GetRoundById(ID)
	if updateRound.RoundID != "" {
		roundDetails.RoundID = updateRound.RoundID
	}
	if updateRound.RoundName != "" {
		roundDetails.RoundName = updateRound.RoundName
	}
	if updateRound.DefaultPrice != "" {
		roundDetails.DefaultPrice = updateRound.DefaultPrice
	}
	if updateRound.Start != "" {
		roundDetails.Start = updateRound.Start
	}
	if updateRound.End != "" {
		roundDetails.End = updateRound.End
	}
	if updateRound.State != "" {
		roundDetails.State = updateRound.State
	}
	if updateRound.Info != "" {
		roundDetails.Info = updateRound.Info
	}
	db.Save(&roundDetails)
	res, _ := json.Marshal(roundDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
