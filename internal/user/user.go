package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func Route() *mux.Router {
	user := User{}

	r := mux.NewRouter()
	router := r.PathPrefix("/user").Subrouter()
	router.HandleFunc("/register", user.register).Methods("POST")
	router.HandleFunc("/login", user.login).Methods("POST")
	return router
}

func (u *User) register(writer http.ResponseWriter, request *http.Request) {
	var payload RegisterRequest

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		Json(writer, err.Error(), 402)
		return
	}

	err = validator.New().Struct(payload)
	if err != nil {
		Json(writer, err.Error(), 402)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		Json(writer, "Email and password are required", 400)
		return
	}

	hash, errHash := HashPassword(payload.Password)
	if errHash != nil {
		Json(writer, errHash.Error(), 500)
		return
	}

	// TODO create user
	fmt.Println(hash)

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

func HashPassword(password string) (string, error) {
	// GenerateFromPassword возвращает bcrypt хеш пароля с заданной сложностью
	// DefaultCost = 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(bytes), nil
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
