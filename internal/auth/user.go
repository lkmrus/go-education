package auth

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct{}

func Route() *mux.Router {
	user := User{}

	r := mux.NewRouter()
	router := r.PathPrefix("/user").Subrouter()
	router.HandleFunc("/register", user.register).Methods("POST")
	router.HandleFunc("/login", user.login).Methods("POST")
	return router
}

func (u *User) register(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Register!"))
}

func (u *User) login(writer http.ResponseWriter, request *http.Request) {
	var payload LoginRequest
	response := LoginResponse{
		Token: "token",
	}

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		Json(writer, err.Error(), 402)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		Json(writer, "Email and password are required", 400)
		return
	}

	Json(writer, response, 201)
}

func Json(writer http.ResponseWriter, response any, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	writer.Write(jsonBytes)
}
