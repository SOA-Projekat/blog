package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPostHandler struct {
	BlogPostService *service.BlogPostService
}

func (handler *BlogPostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	blogPosts, err := handler.BlogPostService.GetAll(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogPosts)
}

func (handler *BlogPostHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogPostId := params["blogPostId"]

	blogPost, err := handler.BlogPostService.GetById(blogPostId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogPost)
}

func (handler *BlogPostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var blog model.BlogPost

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogPostService.CreateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostHandler) Update(w http.ResponseWriter, r *http.Request) {
	var blog model.BlogPost

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogPostService.Update(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogPostId := params["blogPostId"]

	err := handler.BlogPostService.Delete(blogPostId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
