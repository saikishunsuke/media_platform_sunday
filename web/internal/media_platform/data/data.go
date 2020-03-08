package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

// Db is database.
var db *gorm.DB

// GetSession returns session struct from session id.
func GetSession(sessionID string) (session Session, err error) {
	err = db.First(&session, "uuid = ?", sessionID).Error
	return
}

// GetUser returns user from user id.
func GetUser(userID string) (user User, err error) {
	err = db.First(&user, "user_id = ?", userID).Error
	return
}

// GetAllPosts returns all posts.
func GetAllPosts() (posts []Post, err error) {
	query := db.Table("posts").
		Select("posts.*, users.name, users.age").
		Joins("inner join users on posts.user_id = users.user_id").
		Where("posts.deleted_at is NULL")
	rows, err := query.Rows()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var post Post
	var user User
	for rows.Next() {
		err = db.ScanRows(rows, &post)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = db.ScanRows(rows, &user)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		post.User = user
		posts = append(posts, post)
	}
	return
}

func GetPost(postID uint) (post Post, err error) {
	err = db.First(&post, postID).Error
	return
}

func init() {
	db = getDb()
	db.AutoMigrate(&User{}, &Session{}, &Post{})
	db.LogMode(true)
}

func getDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@tcp(mysql:3306)/data_base?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
