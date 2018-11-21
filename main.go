package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
  "test/auth"
  "test/middleware"
)

func Login(w http.ResponseWriter, r *http.Request) {
  type Credetial struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }
  var c Credetial

  err := json.NewDecoder(r.Body).Decode(&c)
  log.Println("username: ", c.Username)
  log.Println("password: ", c.Password)

	log.Println("Login called")
	validToken, err := auth.GenerateJWT(c.Username, "admin")
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	type LoginResponse struct {
		Token string `json:"token"`
	}
	lr := LoginResponse{validToken}
	json.NewEncoder(w).Encode(lr)
	return
}
func SelectAllArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("SelectAllArticle")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/articles", SelectAllArticle).Methods("GET")
  router.Use(middleware.ValidTokenMiddleware)
	log.Fatal(http.ListenAndServe(":8000", router))
}
