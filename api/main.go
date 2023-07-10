package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
	"urlshorter/controllers"
	"urlshorter/migrations"
	"urlshorter/services"
)

var (
	g errgroup.Group
)

func routerApi() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.POST("/add", controllers.AddLink)
	e.POST("/auth", controllers.Auth)

	return e
}

func main() {
	migrations.ApplyMigrations()

	go services.ListenQueue()

	serverApi := &http.Server{
		Addr:         ":9999",
		Handler:      routerApi(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return serverApi.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
