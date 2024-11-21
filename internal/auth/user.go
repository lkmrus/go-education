package auth

import (
	"github.com/gorilla/mux"
	"net/http"
)

type User struct{}

func (u *User) Route() *mux.Router {
	r := mux.NewRouter()
	router := r.PathPrefix("/user").Subrouter()
	router.HandleFunc("/register", register)
	router.HandleFunc("/login", login)
	return router
}

func register(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Register!"))
}

func login(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Login!"))
}
