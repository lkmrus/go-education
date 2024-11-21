package auth

import (
	"github.com/gorilla/mux"
	"net/http"
)

type User struct{}

func Route() *mux.Router {
	user := User{}

	r := mux.NewRouter()
	router := r.PathPrefix("/user").Subrouter()
	router.HandleFunc("/register", user.register)
	router.HandleFunc("/login", user.login)
	return router
}

func (u *User) register(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Register!"))
}

func (u *User) login(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Login!"))
}
