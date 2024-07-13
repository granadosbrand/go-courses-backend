package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirsName string `json:"first_name"`
		LastName string `json:"last_name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: MakeCreateEnpoint(),
		Get:    MakeGetEnpoint(),
		GetAll: MakeGetAllEnpoint(),
		Update: MakeUpdateEnpoint(),
		Delete: MakeDeleteEnpoint(),
	}
}

func MakeCreateEnpoint() Controller {

	
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Create User")
		json.NewEncoder(w).Encode(req)
	}
}

func MakeUpdateEnpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update User")
		json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
	}
}

func MakeGetAllEnpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All Users")
		json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
	}
}

func MakeGetEnpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get User")
		json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
	}
}

func MakeDeleteEnpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete User")
		json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
	}
}
