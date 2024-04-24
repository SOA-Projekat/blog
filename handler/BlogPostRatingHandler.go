package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPostRatingHandler struct {
	BlogPostRatingService *service.BlogPostRatingService
}

func (handler *BlogPostRatingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	blogPostRatings, err := handler.BlogPostRatingService.GetAll(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogPostRatings)
}

func (handler *BlogPostRatingHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogPostRatingId := params["blogPostRatingId"]

	blogPostRating, err := handler.BlogPostRatingService.GetById(blogPostRatingId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogPostRating)
}

func (handler *BlogPostRatingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var rating model.BlogPostRating

	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogPostRatingService.CreateRating(&rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostRatingHandler) Update(w http.ResponseWriter, r *http.Request) {
	var rating model.BlogPostRating

	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogPostRatingService.Update(&rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *BlogPostRatingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogPostRatingIdStr := params["blogPostRatingId"]
	blogPostRatingId, err := strconv.ParseUint(blogPostRatingIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid blogPostRatingId", http.StatusBadRequest)
		return
	}

	err = handler.BlogPostRatingService.Delete(uint(blogPostRatingId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
