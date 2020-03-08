package data

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User is test struct.
type User struct {
	gorm.Model
	UserID   string `json:"user_id" gorm:"unique; not null; primary_key; column:user_id"`
	Password string `json:"password" gorm:"not null; column:password"`
	Name     string
	Age      int
	Posts    []Post `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
}

func (user *User) Update(newUser User) (err error) {
	err = db.Model(&user).Updates(newUser).Error
	return
}

// CreateSession is function to create session and set cookie.
func (user *User) CreateSession(w http.ResponseWriter) error {
	sid := make([]byte, 32)
	io.ReadFull(rand.Reader, sid)
	session := Session{
		UUID:   base64.URLEncoding.EncodeToString(sid),
		UserID: user.UserID,
	}
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: session.UUID,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	return db.Create(&session).Error
}

// GetPosts returns all posts user has.
func (user *User) GetPosts() (posts []Post, err error) {
	query := db.Table("posts").
		Select("posts.*, users.name, users.age").
		Joins("inner join users on posts.user_id = users.user_id").
		Where("posts.user_id = ? and posts.deleted_at is NULL", user.UserID)
	rows, err := query.Rows()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var post Post
	var passUser User
	for rows.Next() {
		err = query.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = query.ScanRows(rows, &passUser)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		post.User = passUser
		posts = append(posts, post)
	}
	return
}

// Save is method to save user.
func (user *User) Save() error {
	return db.Create(&user).Error
}

// Session is session
type Session struct {
	gorm.Model
	UUID   string
	UserID string
}

// Delete is method to delete session from table.
func (session *Session) Delete() error {
	return db.Delete(&session).Error
}

// Post is post
type Post struct {
	gorm.Model
	User   User `gorm:"foreignkey:UserID"`
	UserID string
	Title  string
	Text   string
}

// Save is save posts
func (post *Post) Save() error {
	fmt.Printf("%+v\n", post)
	return db.Create(&post).Error
}

func (post *Post) Update(newPost Post) error {
	return db.Model(&post).Updates(newPost).Error
}

func (post *Post) Delete() error {
	return db.Delete(&post).Error
}
