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

var NewUserWalletMTD models.UserWalletMTD

func GetUserWalletMTD(w http.ResponseWriter, r *http.Request) {
	newUserWalletMTDs := models.GetAllUserWalletMTDs()
	res, _ := json.Marshal(newUserWalletMTDs)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserWalletMTDById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userAddress := vars["userAddress"]
	ID, err := strconv.ParseInt(userAddress, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userWalletMTDDetails, _ := models.GetUserWalletMTDById(ID)
	res, _ := json.Marshal(userWalletMTDDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUserWalletMTD(w http.ResponseWriter, r *http.Request) {
	CreateUserWalletMTD := &models.UserWalletMTD{}
	utils.ParseBody(r, CreateUserWalletMTD)
	b := CreateUserWalletMTD.CreateUserWalletMTD()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUserWalletMTD(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userAddress := vars["userAddress"]
	ID, err := strconv.ParseInt(userAddress, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userWalletMTD := models.DeleteUserWalletMTD(ID)
	res, _ := json.Marshal(userWalletMTD)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserWalletMTD(w http.ResponseWriter, r *http.Request) {
	var updateUserWalletMTD = &models.UserWalletMTD{}
	utils.ParseBody(r, updateUserWalletMTD)
	vars := mux.Vars(r)
	userAddress := vars["userAddress"]
	ID, err := strconv.ParseInt(userAddress, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userWalletMTDDetails, dbUserWalletMTD := models.GetUserWalletMTDById(ID)
	if updateUserWalletMTD.UserAddress != "" {
		userWalletMTDDetails.UserAddress = updateUserWalletMTD.UserAddress
	}
	if updateUserWalletMTD.Type != "" {
		userWalletMTDDetails.Type = updateUserWalletMTD.Type
	}
	if updateUserWalletMTD.Nonce != "" {
		userWalletMTDDetails.Nonce = updateUserWalletMTD.Nonce
	}
	if updateUserWalletMTD.TokenExpired != "" {
		userWalletMTDDetails.TokenExpired = updateUserWalletMTD.TokenExpired
	}
	dbUserWalletMTD.Save(&userWalletMTDDetails)
	res, _ := json.Marshal(userWalletMTDDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
