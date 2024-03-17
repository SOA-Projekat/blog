package model

import (
	"errors"
	"time"
)

type BlogPostStatus int

const (
	DRAFT BlogPostStatus = iota
	PUBLISHED
	CLOSED
	ACTIVE
	FAMOUS
)

type BlogPost struct {
	AuthorID     int
	TourID       int
	Title        string
	Description  string
	CreationDate time.Time
	//ImageURLs    []string
	//Comments     []BlogPostComment
	//Ratings      []BlogPostRating
	Status BlogPostStatus
}

func NewBlogPost(authorID, tourID int, title, description string, creationDate time.Time, status BlogPostStatus) (*BlogPost, error) {
	if authorID == 0 {
		return nil, errors.New("field required: AuthorID")
	}
	if title == "" {
		return nil, errors.New("invalid Title")
	}
	if description == "" {
		return nil, errors.New("invalid Description")
	}
	if creationDate.IsZero() {
		return nil, errors.New("invalid Creation Date")
	}
	if status != DRAFT && status != PUBLISHED && status != CLOSED && status != ACTIVE && status != FAMOUS {
		return nil, errors.New("invalid Post Status")
	}

	return &BlogPost{
		AuthorID:     authorID,
		TourID:       tourID,
		Title:        title,
		Description:  description,
		CreationDate: creationDate,
		//ImageURLs:    imageURLs,
		//Comments:     comments,
		//Ratings:      ratings,
		Status: status,
	}, nil
}
