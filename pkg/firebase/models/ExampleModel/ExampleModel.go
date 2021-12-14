package ExampleModel

import (
	"net/url"
	"context"
	"log"
	"time"

	"github.com/satori/go.uuid"
	"firebase.google.com/go/db"
)

//firebase model
type Example struct { 
	ID 			uuid.UUID 
	URI 		string 
	ImageURL 	string 
	Category 	string 
	Title 		string 
	Author 		string 
	Content 	string
	Key			string 
	TotalView 	int64 
	IsActive 	bool 
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	CreatedBy 	int
	UpdatedBy 	uuid.UUID 
}

var TableName = "example"

func GetAll(ctx context.Context, initDb *db.Client) (map[string]Example, error) { 
	db := initDb.NewRef(TableName)
	var Blogs map[string]Example
	if err := db.Get(ctx, &Blogs); err != nil { 
		log.Println(err.Error()) 
		return Blogs, err 
	} 
	return Blogs, nil 
}

func Insert(ctx context.Context, initDb *db.Client) (*Example, error) {
	db := initDb.NewRef(TableName)
    blog := &Example{
		ID:        uuid.NewV4(),
		URI:       url.QueryEscape("http://localhost:8080/"),
		ImageURL:  "url image here",
		Category:  "category",
		Title:     "title",
		Author:    "author",
		Content:   "Bla bla bla",
		TotalView: 100,
		IsActive:  false,
		CreatedBy: 1,
		CreatedAt: time.Now(),
	}
	
	if err := db.Child(blog.ID.String()).Set(ctx, blog); err != nil {
		log.Println(err.Error()) 
		return nil, err
	}
	return blog, nil
}

func GetOneBlog(ctx context.Context, initDb *db.Client, param string) (Example, error) {
	var Blog Example
	db := initDb.NewRef(TableName)
	if err := db.Child(param).Get(ctx, &Blog); err != nil {
		log.Println(err.Error()) 
		return Blog, nil
	}
	Blog.Key = param

	return Blog, nil
}


func (blog *Example) Update(ctx context.Context, initDb *db.Client) error {
	db := initDb.NewRef(TableName)
	err := db.Child(blog.ID.String()).Set(ctx, blog)
	if err != nil {
		log.Println(err.Error()) 
		return err 
	}
	return nil
}

func BlogDelete(ctx context.Context, initDb *db.Client, id string) error {
	db := initDb.NewRef(TableName)
	err := db.Child(id).Delete(ctx)
	if err != nil {
		log.Println(err.Error()) 
	 	return err
	}
	return nil
}
