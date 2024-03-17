package model

import (
	"fmt"
	"time"
)

// BlogStatus represents the status of a blog.
type BlogStatus int

const (
	BlogDraft BlogStatus = iota
	BlogPublished
	BlogArchived
	// BlogReady
)

// Blog represents the data structure for a blog
type Blog struct {
	BlogID       int        `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	CreationDate time.Time  `json:"creation_date"`
	Images       []string   `json:"images,omitempty"`
	Status       BlogStatus `json:"status"`
}

// NewBlog creates a new instance of Blog and returns it
func NewBlog(title, description string, images []string, status BlogStatus) *Blog {
	return &Blog{
		Title:        title,
		Description:  description,
		CreationDate: time.Now(),
		Images:       images,
		Status:       status,
	}
}

// Example function using NewBlog to create and use a new blog
func ExampleFunction() {
	// Creating a new blog
	newBlog := NewBlog("Example Blog", "This is an example blog that uses Markdown language.", []string{"image1.jpg", "image2.jpg"}, BlogDraft)

	// Using the new blog
	fmt.Println("New blog created:", newBlog.Title)
	fmt.Println("Description:", newBlog.Description)
	fmt.Println("Creation date:", newBlog.CreationDate)
	fmt.Println("Status:", newBlog.Status)
}
