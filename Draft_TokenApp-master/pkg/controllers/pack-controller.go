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

var NewPack models.Pack

func GetPack(w http.ResponseWriter, r *http.Request) {
	newPacks := models.GetAllPacks()
	res, _ := json.Marshal(newPacks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPackById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packId := vars["packId"]
	ID, err := strconv.ParseInt(packId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	packDetails, _ := models.GetPackById(ID)
	res, _ := json.Marshal(packDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePack(w http.ResponseWriter, r *http.Request) {
	CreatePack := &models.Pack{}
	utils.ParseBody(r, CreatePack)
	b := CreatePack.CreatePack()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packId := vars["packId"]
	ID, err := strconv.ParseInt(packId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pack := models.DeletePack(ID)
	res, _ := json.Marshal(pack)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePack(w http.ResponseWriter, r *http.Request) {
	var updatePack = &models.Pack{}
	utils.ParseBody(r, updatePack)
	vars := mux.Vars(r)
	packId := vars["packId"]
	ID, err := strconv.ParseInt(packId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	packDetails, dbPack := models.GetPackById(ID)
	if updatePack.PackID != "" {
		packDetails.PackID = updatePack.PackID
	}
	if updatePack.Volume != "" {
		packDetails.Volume = updatePack.Volume
	}
	if updatePack.State != "" {
		packDetails.State = updatePack.State
	}
	if updatePack.RoundID != "" {
		packDetails.RoundID = updatePack.RoundID
	}
	dbPack.Save(&packDetails)
	res, _ := json.Marshal(packDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
