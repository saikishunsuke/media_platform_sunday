package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"work/internal/media_platform/data"

	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Your request method is GET. This endpoint allows only POST."))
		return
	}
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.UserID == "" || user.Password == "" {
		w.Write([]byte("Invalid input."))
		return
	}
	user.Password = string(hash)
	if user.Save() != nil {
		w.Write([]byte("User id is already used."))
		return
	}
	if user.CreateSession(w) != nil {
		w.Write([]byte("Cannot create Session."))
		return
	}
	w.Write([]byte("OK."))
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Your request method is GET. This endpoint allows only POST."))
		return
	}
	var passedUser data.User
	err := json.NewDecoder(r.Body).Decode(&passedUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := data.GetUser(passedUser.UserID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Please sign up."))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passedUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Password is wrong."))
		return
	}
	if user.CreateSession(w) != nil {
		w.Write([]byte("Cannot create Session."))
		return
	}
	w.Write([]byte("Signed In."))
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		return
	}
	newCookie := &http.Cookie{Name: "session_id", Value: "", Path: "/"}
	http.SetCookie(w, newCookie)
	session, err := data.GetSession(cookie.Value)
	if err != nil {
		fmt.Println("Session Record Not Found.")
		return
	}
	session.Delete()
	w.Write([]byte("Signed out."))
}

func UpdateUserDataHandler(w http.ResponseWriter, r *http.Request) {
	loginUser, err := GetSignInUser(r)
	if err != nil {
		w.Write([]byte("Please Sign up."))
		return
	}
	var newUser data.User
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	if newUser.UserID != "" || newUser.Password != "" {
		w.Write([]byte("User id and password cannot be changed."))
		return
	}
	err = loginUser.Update(newUser)
	if err != nil {
		w.Write([]byte("Update Error."))
		return
	}
	return
}

func SignInRequired(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isSingedIn(r) {
			hf.ServeHTTP(w, r)
		} else {
			w.Write([]byte("Please Sign in."))
		}
	}
}

func isSingedIn(r *http.Request) bool {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = data.GetSession(cookie.Value)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func GetSignInUser(r *http.Request) (data.User, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println(err.Error())
		return data.User{}, err
	}
	session, err := data.GetSession(cookie.Value)
	if err != nil {
		fmt.Println(err.Error())
		return data.User{}, err
	}
	user, err := data.GetUser(session.UserID)
	if err != nil {
		fmt.Println(err.Error())
		return data.User{}, err
	}
	return user, err
}
