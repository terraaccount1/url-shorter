package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/terraaccount1/url-shorter/internal/controller"
	"github.com/terraaccount1/url-shorter/internal/repository"
	"github.com/terraaccount1/url-shorter/internal/service"
)

var db *redis.Client

func main() {
	repo := repository.NewUrlRepository(db)
	svc := service.NewUrlService(repo)
	ctrl := controller.NewUrlController(svc)

	r := gin.Default()
	r.GET("/:uri", ctrl.GetUrl)
	r.POST("/", ctrl.SetUrl)
	slog.Info("Starting server ...")
	if err := r.Run(); err != nil {
		panic(fmt.Errorf("failed to start server : %w", err))
	}
}

func init() {
	db = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})
}
