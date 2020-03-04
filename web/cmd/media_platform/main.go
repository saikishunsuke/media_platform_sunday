package main

import (
	"fmt"
	"net/http"

	"work/internal/media_platform/auth"
	"work/internal/media_platform/post"
)

func main() {
	fmt.Println("Listening on http://localhost:8088/")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my API."))
	})
	http.HandleFunc("/users/signup", auth.SignUpHandler)
	http.HandleFunc("/users/signin", auth.SignInHandler)
	http.HandleFunc("/users/signout", auth.SignOutHandler)
	http.HandleFunc("/users/update", auth.SignInRequired(auth.UpdateUserDataHandler))
	http.HandleFunc("/post/new", auth.SignInRequired(post.CreatePostHandler))
	http.HandleFunc("/post/index", post.ReadAllPostsHandler)
	http.HandleFunc("/post/mine", auth.SignInRequired(post.ReadOwnPostsHandler))
	http.ListenAndServe(":80", nil)
}
