package user

import (
	"demo/app/pkg/db"
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

func UserRoute(dbConnection *db.Db) *mux.Router {
	r := mux.NewRouter()
	router := r.PathPrefix("/user").Subrouter()

	userRepo := UserRepository{
		Database: dbConnection,
	}

	router.HandleFunc("/register", userRepo.register).Methods("POST")
	router.HandleFunc("/login", userRepo.login).Methods("POST")
	return router
}

func (uRepo *UserRepository) register(writer http.ResponseWriter, request *http.Request) {
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

	result, createErr := uRepo.CreateUser(&User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hash,
	})
	fmt.Println(result)

	if createErr != nil {
		Json(writer, createErr.Error(), 500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err.Error())
	}
	writer.Write(jsonBytes)
}

func (uRepo *UserRepository) login(writer http.ResponseWriter, request *http.Request) {
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

	if !CheckPassword(payload.Password, "hash") {
		Json(writer, "not allowed", 400)
		return
	}

	passwordHash, err := HashPassword(payload.Password)
	if err != nil {
		Json(writer, err.Error(), 500)
		return
	}

	if payload.Password != passwordHash {
		Json(writer, "password was wrong", 400)
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
