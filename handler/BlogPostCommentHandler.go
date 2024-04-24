package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPostCommentHandler struct {
	BlogPostCommentService *service.BlogPostCommentService
}

func (handler *BlogPostCommentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	blogComments, err := handler.BlogPostCommentService.GetAll(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogComments)
}

func (handler *BlogPostCommentHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogPostCommentId := params["blogPostCommentId"]

	blogPostComment, err := handler.BlogPostCommentService.GetById(blogPostCommentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogPostComment)
}

func (handler *BlogPostCommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var comment model.BlogPostComment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogPostCommentService.CreateComment(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostCommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	var comment model.BlogPostComment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogPostCommentService.Update(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostCommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogPostCommentIdStr := params["blogPostCommentId"]
	blogPostCommentId, err := strconv.ParseUint(blogPostCommentIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid blogPostId", http.StatusBadRequest)
		return
	}

	err = handler.BlogPostCommentService.Delete(uint(blogPostCommentId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
