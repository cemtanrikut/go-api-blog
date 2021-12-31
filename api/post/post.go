package post

import "time"

type Post struct {
	Author      string    `json:"author"`
	AuthorImage string    `json:"author_image"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
	IsActive    bool      `json:"is_active"`
	Tag         Tag       `json:"tag"`
}

type Tag struct {
	Name string `json:"tag"`
}

func AddPost() {

}

func GetPost(Id int) {

}

func GetPostsUnderTag(tagName Tag) {

}

func GetPosts() {

}

func UpdatePost() {

}

func DeletePost() {

}
