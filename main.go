package main

import (
	"context"
	"log"
	"time"
	"tutorial/db"
	"tutorial/logger"
	"tutorial/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// init logger
	logger.Init()
	defer logger.Log.Sync()

	// Connect Database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.ConnectDB(ctx)

	if err != nil {
		log.Fatal("Gagal Koneksi Ke database: ", err)
	}

	defer pool.Close()

	r := routers.SetupRouter(pool)
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Recovery())
	r.Run(":8081")

}
