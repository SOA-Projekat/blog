package handler

import (
	"database-example/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPostRatingHandler struct {
	RatingService *service.BlogPostRatingService
}

// AddRating dodaje novu ocenu za dati blog post.
func (handler *BlogPostRatingHandler) AddRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogID, err := strconv.Atoi(params["blogID"])
	if err != nil {
		http.Error(w, "Invalid blogID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	isPositive, err := strconv.ParseBool(params["isPositive"])
	if err != nil {
		http.Error(w, "Invalid isPositive value", http.StatusBadRequest)
		return
	}

	// Dodajte ocenu
	err = handler.RatingService.AddRating(uint(blogID), userID, isPositive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteRating briše ocenu za dati blog post na osnovu korisničkog ID-ja.
func (handler *BlogPostRatingHandler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blogID, err := strconv.Atoi(params["blogID"])
	if err != nil {
		http.Error(w, "Invalid blogID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	// Obrišite ocenu
	err = handler.RatingService.DeleteRating(uint(blogID), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
