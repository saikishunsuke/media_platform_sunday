package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"work/internal/media_platform/auth"
	"work/internal/media_platform/data"
)

// CreatePostHandler is handler which create post.
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Only Post."))
		return
	}
	var post data.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user, err := auth.GetSignInUser(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	post.UserID = user.UserID
	if err != nil {
		fmt.Println(err)
		return
	}
	err = post.Save()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write([]byte("Success."))
	return
}

// ReadAllPostsHandler is handler which get all posts.
func ReadAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Only Get."))
		return
	}
	posts, err := data.GetAllPosts()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytes, _ := json.Marshal(&posts)
	w.Write(bytes)
	return
}

// ReadOwnPostsHandler is handler which get sign in user's posts.
func ReadOwnPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Only Get."))
		return
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Please Sign in."))
		return
	}
	session, err := data.GetSession(cookie.Value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user, err := data.GetUser(session.UserID)
	if err != nil {
		w.Write([]byte("Please Sign in."))
		return
	}
	posts, err := user.GetPosts()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytes, err := json.Marshal(&posts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write(bytes)
	return
}
