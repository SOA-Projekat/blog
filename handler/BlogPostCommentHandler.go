package handler

import (
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPostCommentHandler struct {
	CommentService *service.BlogPostCommentService
}

func (handler *BlogPostCommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogID, err := strconv.Atoi(params["blogID"])
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	var commentData struct {
		Text   string `json:"text"`
		UserID int    `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&commentData); err != nil {
		http.Error(w, "Failed to decode comment data", http.StatusBadRequest)
		return
	}

	err = handler.CommentService.AddComment(blogID, commentData.Text, commentData.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *BlogPostCommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentID, err := strconv.Atoi(params["commentID"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var commentData struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&commentData); err != nil {
		http.Error(w, "Failed to decode comment data", http.StatusBadRequest)
		return
	}

	err = handler.CommentService.UpdateComment(commentID, commentData.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostCommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentID, err := strconv.Atoi(params["commentID"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	err = handler.CommentService.DeleteComment(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
