package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=blog port=5432 sslmode=disable search_path=public"

	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.BlogPost{})
	database.AutoMigrate(&model.BlogPostComment{})
	database.AutoMigrate(&model.BlogPostRating{})

	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	// blogPost
	blogPostRepo := &repo.BlogPostRepository{DatabaseConnection: database}
	blogPostService := &service.BlogPostService{BlogRepo: blogPostRepo}
	blogPostHandler := &handler.BlogPostHandler{BlogPostService: blogPostService}

	// blogPostComment
	blogPostCommentRepo := &repo.BlogPostCommentRepository{DatabaseConnection: database}
	blogPostCommentService := &service.BlogPostCommentService{CommentRepo: blogPostCommentRepo}
	blogPostCommentHandler := &handler.BlogPostCommentHandler{CommentService: blogPostCommentService}

	// blogPostRating
	blogPostRatingRepo := &repo.BlogPostRatingRepository{DatabaseConnection: database}
	blogPostRatingService := &service.BlogPostRatingService{RatingRepo: blogPostRatingRepo}
	blogPostRatingHandler := &handler.BlogPostRatingHandler{RatingService: blogPostRatingService}

	router := mux.NewRouter().StrictSlash(true)

	//routes for blogPost
	router.HandleFunc("/api/blog/blogpost", blogPostHandler.GetAll).Methods("GET")
	router.HandleFunc("/api/blog/blogpost/{blogPostId}", blogPostHandler.GetById).Methods("GET")
	router.HandleFunc("/api/blog/blogpost", blogPostHandler.Create).Methods("POST")
	router.HandleFunc("/api/blog/blogpost/{id}", blogPostHandler.Update).Methods("PUT")
	router.HandleFunc("/api/blog/blogpost/{id}", blogPostHandler.Delete).Methods("DELETE")

	//routes for blogPostComment
	router.HandleFunc("/api/blog/blogpost/{blogPostId:[0-9]+}/comments/{userId:[0-9]+}/{creationTime}", blogPostCommentHandler.DeleteComment).Methods("DELETE")
	router.HandleFunc("/api/blog/blogpost/{blogPostId:[0-9]+}/comments", blogPostCommentHandler.UpdateComment).Methods("PUT")
	router.HandleFunc("/api/blog/blogpost/{blogPostId:[0-9]+}/comments", blogPostCommentHandler.AddComment).Methods("POST")

	//routes for blogPostRating
	router.HandleFunc("/api/blog/blogpost/{blogPostId:[0-9]+}/ratings/{userId:[0-9]+}", blogPostRatingHandler.DeleteRating).Methods("DELETE")
	router.HandleFunc("/api/blog/blogpost/{blogPostId:[0-9]+}/ratings", blogPostRatingHandler.AddRating).Methods("POST")

	permitedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permitedOrigins := handlers.AllowedOrigins([]string{"*"})
	permitedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permitedHeaders, permitedOrigins, permitedMethods)(router)))
}
