package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/thanhngodinh/mysql-simple/controller/util"
	"github.com/thanhngodinh/mysql-simple/controller/model"
)

var NewUser = model.User

func Get(w http.ResponseWriter, r *http.Request) {
	users := model.GetAll()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Load(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user, _:= model.GetById(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	CreateUser := &model.User{}
	util.ParseBody (r, CreateUser)
	u := CreateUser.Create()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while passing")
	}
	user := model.Delete(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Update(w http.ResponseWriter, r *http.Request){
	var updateUser = &model.User{}
	util.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, db := model.GetById(ID)
	if updateUser.UserName != "" {
		userDetails.UserName = updateUser.UserName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}