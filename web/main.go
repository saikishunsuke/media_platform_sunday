package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// AuthData is test struct.
type AuthData struct {
	gorm.Model
	UserID   string `json:"user_id" gorm:"unique; not null; primary_key; column:user_id"`
	Password string `json:"password" gorm:"unique; not null; column:password"`
}

// UserInformation is struct
type UserInformation struct {
	gorm.Model
	AuthData AuthData
	Name     string
	Age      int
}

var db *gorm.DB

func init() {
	db = getDb()
	db.AutoMigrate(&AuthData{}, &UserInformation{})
	db.LogMode(true)
}

func getDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@tcp(mysql:3306)/data_base?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	fmt.Println("Listening on http://localhost:8088/")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my API."))
	})
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World"))
	})
	http.HandleFunc("/users/signup", signUpHandler)
	http.HandleFunc("/users/signin", signInHandler)
	http.ListenAndServe(":80", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Your request method is GET. This endpoint allows only POST."))
		return
	}
	var auth AuthData
	err := json.NewDecoder(r.Body).Decode(&auth)
	fmt.Printf("%+v\n", auth)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	auth.Password = string(hash)
	if db.Create(&auth).Error != nil {
		w.Write([]byte("User id is already used."))
		return
	}
	w.Write([]byte("OK."))
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Your request method is GET. This endpoint allows only POST."))
		return
	}
	var passedAuth AuthData
	err := json.NewDecoder(r.Body).Decode(&passedAuth)
	if err != nil {
		fmt.Println(err)
		return
	}
	var auth AuthData
	if db.First(&auth, &AuthData{UserID: passedAuth.UserID}).RecordNotFound() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Please sign up."))
		return
	}
	fmt.Printf("%+v\n", auth)
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(passedAuth.Password))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Password is wrong."))
		return
	}
	w.Write([]byte("Signed In."))
}
