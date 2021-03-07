package main

import (
	"flag"
	"fmt"
	product_delivery "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/product/delivery"
	product_repo "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/product/repository"
	product_usecase "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/product/usecase"
	session_repo "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/session/repository"
	session_usecase "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/session/usecase"
	user_delivery "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/user/delivery"
	user_repo "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/user/repository"
	user_usecase "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/user/usecase"
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/server/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	flag.Parse()

	sessionRepo := session_repo.NewSessionLocalRepository()
	sessionManager := session_usecase.NewUseCase(sessionRepo)

	userRepo := user_repo.NewSessionLocalRepository()
	userUCase := user_usecase.NewUseCase()
	userHandler := &user_delivery.UserHandler{
		UserUCase:      userUCase,
		SessionManager: sessionManager,
		UserRepo: userRepo,
	}

	productRepo := product_repo.NewSessionLocalRepository()
	productUCase := product_usecase.NewUseCase()
	productHandler := &product_delivery.ProductHandler{
		ProductRepo: productRepo,
		ProductUCase: productUCase,
	}



	mainMux := mux.NewRouter()
	mainMux.Use(middleware.Panic)
	mainMux.Use(middleware.Cors)
	mainMux.HandleFunc("/api/v1/user/signup", userHandler.Signup).Methods("POST", "OPTIONS")
	mainMux.HandleFunc("/api/v1/user/login", userHandler.Login).Methods("POST", "OPTIONS")
	mainMux.HandleFunc("/api/v1/product/{id:[0-9]+}", productHandler.GetProduct).Methods("GET", "OPTIONS")
	mainMux.HandleFunc("/api/v1/product", productHandler.GetRangeProducts).Methods("POST", "OPTIONS")


	// Handlers with Auth middleware
	authMux := mainMux.PathPrefix("/").Subrouter()
	middlewareAuth := middleware.Auth(sessionManager)
	authMux.Use(middlewareAuth)
	authMux.HandleFunc("/api/v1/user/profile", userHandler.GetProfile).Methods("GET", "OPTIONS")
	authMux.HandleFunc("/api/v1/user/logout", userHandler.Logout).Methods("DELETE", "OPTIONS")
	authMux.HandleFunc("/api/v1/user/profile", userHandler.UpdateProfile).Methods("PUT", "OPTIONS")
	authMux.HandleFunc("/api/v1/user/profile/avatar", userHandler.GetProfileAvatar).Methods("GET", "OPTIONS")
	authMux.HandleFunc("/api/v1/user/profile/avatar", userHandler.UpdateProfileAvatar).Methods("PUT", "OPTIONS")



	//log.Fatal(http.ListenAndServe(":"+*port, mainMux))

	server := &http.Server{
		Addr:         ":"+*port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: mainMux, // Pass our instance of gorilla/mux in.
	}

	fmt.Println("starting server")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
