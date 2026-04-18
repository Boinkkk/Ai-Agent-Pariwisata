package main

import (
	"context"
	"log"
	"time"
	"tutorial/db"
	"tutorial/routers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Connect Database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.ConnectDB(ctx)

	if err != nil {
		log.Fatal("Gagal Koneksi Ke database: ", err)
	}

	defer pool.Close()

	r := routers.SetupRouter(pool)

	r.Run(":8081")

}
